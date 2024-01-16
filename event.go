// Copyright 2024 Leon Hwang.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"unsafe"

	"github.com/tklauser/ps"
)

const (
	eventTypeIoctl = 1
	eventTypeGenl  = 2
)

type event struct {
	Type     uint8
	GenlCmd  ethGenlCmd
	IoctlCmd ethIoctlCmd
	Pid      uint32
	Ifname   [16]byte
	Comm     [16]byte
}

func nullStr(b []byte) string {
	off := 0
	for ; off < len(b) && b[off] != 0; off++ {
	}

	return unsafe.String(&b[0], off)
}

func (e *event) ifname() string {
	return nullStr(e.Ifname[:])
}

func (e *event) getProcessName(pid int) string {
	p, err := ps.FindProcess(pid)
	if err != nil {
		return nullStr(e.Comm[:])
	}

	if p.Command() == "ethtool" {
		process := e.getProcessName(p.PPID())
		return fmt.Sprintf("ethtool(parent %d:%s)", p.PPID(), process)
	}

	return p.Command()
}

func printHeader() {
	fmt.Printf("%-16s %8s:%-32s %-30s %s\n", "Interface", "PID", "Process", "IOCTL_CMD/GENL_CMD", "ethtool args")
}

func (e *event) print() {
	var (
		cmd string
		msg string
	)

	if e.Type == eventTypeIoctl {
		cmd = e.IoctlCmd.String()
		if flags.debug {
			msg = "from ioctl"
		} else {
			msg = e.IoctlCmd.Message()
		}
	} else {
		cmd = e.GenlCmd.String()
		if flags.debug {
			msg = "from genl"
		} else {
			msg = e.GenlCmd.Message()
		}
	}

	process := e.getProcessName(int(e.Pid))
	fmt.Printf("%-16s %8d:%-32s %-30s %s\n", e.ifname(), e.Pid, process, cmd, msg)
}
