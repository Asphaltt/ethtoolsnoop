/**
 * Copyright 2024 Leon Hwang.
 * SPDX-License-Identifier: GPL-2.0
 */

#include "vmlinux.h"

#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_compiler.h>
#include <bpf/bpf_map_helpers.h>

#define IFNAMSIZ 16

// From include/uapi/linux/ethtool.h
#define ETHTOOL_PERQUEUE	0x0000004b /* Set per queue options */

#define EVENT_TYPE_IOCTL 1
#define EVENT_TYPE_GENL  2

struct event {
    u8 type;
    u8 genlhdr_cmd;
    u16 ethcmd;
    u32 pid;
    char ifname[IFNAMSIZ];
    char comm[TASK_COMM_LEN];

    struct ethnl_req_info *req;
} __attribute__((packed));

#define SIZEOF_EVENT (offsetof(struct event, req))

struct {
    __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
} events SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_PERCPU_ARRAY);
    __type(key, u32);
    __type(value, struct event);
    __uint(max_entries, 1);
} events_cache SEC(".maps");

static __always_inline struct event *
__get_or_init_event(void)
{
    struct event *ev, init = {};
    u32 key = 0;

    ev = bpf_map_lookup_elem(&events_cache, &key);
    if (likely(ev))
        return ev;

    bpf_map_update_elem(&events_cache, &key, &init, BPF_NOEXIST);
    return bpf_map_lookup_elem(&events_cache, &key);
}

static __always_inline struct event *
__get_event(void)
{
    u32 key = 0;
    return bpf_map_lookup_elem(&events_cache, &key);
}

static __always_inline struct event *
__get_and_del_event(void)
{
    struct event *ev;
    u32 key = 0;

    ev = bpf_map_lookup_elem(&events_cache, &key);
    if (unlikely(!ev))
        return NULL;

    bpf_map_delete_elem(&events_cache, &key);
    return ev;
}

static __always_inline u32
get_ethcmd(void *useraddr)
{
    u32 cmd;

    bpf_probe_read_user(&cmd, sizeof(cmd), useraddr);
    if (cmd == ETHTOOL_PERQUEUE)
        cmd = bpf_probe_read_user(&cmd, sizeof(cmd), useraddr + sizeof(cmd));

    return cmd;
}

static __always_inline int
__kp_dev_ethtool(void *ctx, struct net *net, struct ifreq *ifr, void *useraddr)
{
    struct event ev = {};

    ev.type = EVENT_TYPE_IOCTL;
    ev.ethcmd = get_ethcmd(useraddr);

    ev.pid = bpf_get_current_pid_tgid() >> 32;

    bpf_probe_read_kernel_str(ev.ifname, sizeof(ev.ifname), ifr->ifr_ifrn.ifrn_name);
    bpf_get_current_comm(ev.comm, sizeof(ev.comm));

    bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU, &ev, sizeof(ev));

    return BPF_OK;
}

SEC("kprobe/dev_ethtool")
int kp_dev_ethtool(struct pt_regs *ctx)
{
    struct net *net = (typeof(net))(void *)(u64) PT_REGS_PARM1(ctx);
    struct ifreq *ifr = (typeof(ifr))(void *)(u64) PT_REGS_PARM2(ctx);
    void *useraddr = (typeof(useraddr))(void *)(u64) PT_REGS_PARM3(ctx);
    return __kp_dev_ethtool(ctx, net, ifr, useraddr);
}

static __always_inline void
__get_dev_name(struct event *ev, struct ethnl_req_info *req)
{
    struct net_device *dev = BPF_CORE_READ(req, dev);

    if (likely(dev))
        bpf_probe_read_kernel_str(ev->ifname, sizeof(ev->ifname), dev->name);
}

SEC("kprobe/ethnl_default_doit")
int kp_ethnl_doit(struct pt_regs *ctx)
{
    struct genl_info *info = (typeof(info))(void *)(u64) PT_REGS_PARM2(ctx);
    u8 cmd = BPF_CORE_READ(info, genlhdr, cmd);
    struct event *ev = __get_or_init_event();

    if (unlikely(!ev))
        return BPF_OK;

    ev->type = EVENT_TYPE_GENL;
    ev->genlhdr_cmd = cmd;
    ev->req = NULL;

    return BPF_OK;
}

SEC("kprobe/ethnl_parse_header_dev_get")
int kp_ethnl_dev(struct pt_regs *ctx)
{
    struct ethnl_req_info *req = (typeof(req))(void *)(u64) PT_REGS_PARM1(ctx);
    struct event *ev = __get_event();

    if (unlikely(!ev))
        return BPF_OK;

    ev->req = req;

    return BPF_OK;
}

SEC("kretprobe/ethnl_parse_header_dev_get")
int krp_ethnl_dev(struct pt_regs *ctx)
{
    struct event *ev = __get_and_del_event();

    if (unlikely(!ev))
        return BPF_OK;

    if (likely(ev->req))
        __get_dev_name(ev, ev->req);

    ev->pid = bpf_get_current_pid_tgid() >> 32;
    bpf_get_current_comm(ev->comm, sizeof(ev->comm));

    bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU, ev, SIZEOF_EVENT);

    return BPF_OK;
}

char __license[] SEC("license") = "GPL";
