# ethtoolsnoop

`ethtoolsnoop` is a tool for tracing the execution of `ethtool`.

## Use example

```bash
# echo Execute `ethtool -i enp0s1; ethtool -l enp0s1; ethtool -g enp0s1` in another terminal.
# ./ethtoolsnoop
Interface             PID:Process                          IOCTL_CMD/GENL_CMD             ethtool args
enp0s1              11198:ethtool(parent 6373:zsh)         ETHTOOL_GDRVINFO               -d|--register-dump(Do a register dump), -e|--eeprom-dump(Do a EEPROM dump), -i|--driver(Show driver information)
enp0s1              11199:ethtool(parent 6373:zsh)         ETHTOOL_MSG_CHANNELS_GET       -l|--show-channels(Query Channels)
enp0s1              11200:ethtool(parent 6373:zsh)         ETHTOOL_MSG_RINGS_GET          -g|--show-ring(Query RX/TX ring parameters)
```

In the output:

- First column is the interface name.
- Second column is the PID and process name of the process that called
  `ethtool`'s `ioctl()` syscall or sent `ethtool`'s genetlink message, and the
  PID and process name of the parent process if the tracee process is `ethtool`.
- Third column is the underneath command for kernel to execute, including ways
  of `ioctl()` syscall and genetlink message.
- Fourth column is the arguments of `ethtool` command, which may be
  corresponding to the third column. *But this column maybe incomplete.*

## Download

Please download the latest release from this repo's release page.

`ethtoolsnoop` is expected to run on Linux kernel 5.2 and later with BTF support.

## Intenals

`ethtoolsnoop` uses `kprobe` on `dev_ethtool()` to trace the execution of
`ethtool`'s `ioctl()` syscall.

And it uses `kprobe` on `ethnl_default_doit()`, `ethnl_parse_header_dev_get()`
and `kretprobe` on `ethnl_parse_header_dev_get()` to trace the execution of
`ethtool`'s genetlink message.

## License

`ethtoolsnoop` is licensed under the Apache 2.0 license, and its bpf code is
licensed under the GPL 2.0 license.
