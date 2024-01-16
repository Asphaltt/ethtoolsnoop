// Copyright 2024 Leon Hwang.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"strings"
)

type ethIoctlCmd uint16

const (
	ETHTOOL_GSET     = 0x00000001 /* DEPRECATED, Get settings. */
	ETHTOOL_SSET     = 0x00000002 /* DEPRECATED, Set settings. */
	ETHTOOL_GDRVINFO = 0x00000003 /* Get driver info. */
	ETHTOOL_GREGS    = 0x00000004 /* Get NIC registers. */
	ETHTOOL_GWOL     = 0x00000005 /* Get wake-on-lan options. */
	ETHTOOL_SWOL     = 0x00000006 /* Set wake-on-lan options. */
	ETHTOOL_GMSGLVL  = 0x00000007 /* Get driver message level */
	ETHTOOL_SMSGLVL  = 0x00000008 /* Set driver msg level. */
	ETHTOOL_NWAY_RST = 0x00000009 /* Restart autonegotiation. */
	ETHTOOL_GLINK    = 0x0000000a
	ETHTOOL_GEEPROM  = 0x0000000b /* Get EEPROM data */
	ETHTOOL_SEEPROM  = 0x0000000c /* Set EEPROM data. */
	// ETHTOOL_XXX 		 = 0x0000000d /* 0x0d is unused. */
	ETHTOOL_GCOALESCE     = 0x0000000e /* Get coalesce config */
	ETHTOOL_SCOALESCE     = 0x0000000f /* Set coalesce config. */
	ETHTOOL_GRINGPARAM    = 0x00000010 /* Get ring parameters */
	ETHTOOL_SRINGPARAM    = 0x00000011 /* Set ring parameters. */
	ETHTOOL_GPAUSEPARAM   = 0x00000012 /* Get pause parameters */
	ETHTOOL_SPAUSEPARAM   = 0x00000013 /* Set pause parameters. */
	ETHTOOL_GRXCSUM       = 0x00000014 /* Get RX hw csum enable (ethtool_value) */
	ETHTOOL_SRXCSUM       = 0x00000015 /* Set RX hw csum enable (ethtool_value) */
	ETHTOOL_GTXCSUM       = 0x00000016 /* Get TX hw csum enable (ethtool_value) */
	ETHTOOL_STXCSUM       = 0x00000017 /* Set TX hw csum enable (ethtool_value) */
	ETHTOOL_GSG           = 0x00000018 /* Get scatter-gather enable */
	ETHTOOL_SSG           = 0x00000019 /* Set scatter-gather enable */
	ETHTOOL_TEST          = 0x0000001a /* execute NIC self-test. */
	ETHTOOL_GSTRINGS      = 0x0000001b /* get specified string set */
	ETHTOOL_PHYS_ID       = 0x0000001c /* identify the NIC */
	ETHTOOL_GSTATS        = 0x0000001d /* get NIC-specific statistics */
	ETHTOOL_GTSO          = 0x0000001e /* Get TSO enable (ethtool_value) */
	ETHTOOL_STSO          = 0x0000001f /* Set TSO enable (ethtool_value) */
	ETHTOOL_GPERMADDR     = 0x00000020 /* Get permanent hardware address */
	ETHTOOL_GUFO          = 0x00000021 /* Get UFO enable (ethtool_value) */
	ETHTOOL_SUFO          = 0x00000022 /* Set UFO enable (ethtool_value) */
	ETHTOOL_GGSO          = 0x00000023 /* Get GSO enable (ethtool_value) */
	ETHTOOL_SGSO          = 0x00000024 /* Set GSO enable (ethtool_value) */
	ETHTOOL_GFLAGS        = 0x00000025 /* Get flags bitmap(ethtool_value) */
	ETHTOOL_SFLAGS        = 0x00000026 /* Set flags bitmap(ethtool_value) */
	ETHTOOL_GPFLAGS       = 0x00000027 /* Get driver-private flags bitmap */
	ETHTOOL_SPFLAGS       = 0x00000028 /* Set driver-private flags bitmap */
	ETHTOOL_GRXFH         = 0x00000029 /* Get RX flow hash configuration */
	ETHTOOL_SRXFH         = 0x0000002a /* Set RX flow hash configuration */
	ETHTOOL_GGRO          = 0x0000002b /* Get GRO enable (ethtool_value) */
	ETHTOOL_SGRO          = 0x0000002c /* Set GRO enable (ethtool_value) */
	ETHTOOL_GRXRINGS      = 0x0000002d /* Get RX rings available for LB */
	ETHTOOL_GRXCLSRLCNT   = 0x0000002e /* Get RX class rule count */
	ETHTOOL_GRXCLSRULE    = 0x0000002f /* Get RX classification rule */
	ETHTOOL_GRXCLSRLALL   = 0x00000030 /* Get all RX classification rule */
	ETHTOOL_SRXCLSRLDEL   = 0x00000031 /* Delete RX classification rule */
	ETHTOOL_SRXCLSRLINS   = 0x00000032 /* Insert RX classification rule */
	ETHTOOL_FLASHDEV      = 0x00000033 /* Flash firmware to device */
	ETHTOOL_RESET         = 0x00000034 /* Reset hardware */
	ETHTOOL_SRXNTUPLE     = 0x00000035 /* Add an n-tuple filter to device */
	ETHTOOL_GRXNTUPLE     = 0x00000036 /* deprecated */
	ETHTOOL_GSSET_INFO    = 0x00000037 /* Get string set info */
	ETHTOOL_GRXFHINDIR    = 0x00000038 /* Get RX flow hash indir'n table */
	ETHTOOL_SRXFHINDIR    = 0x00000039 /* Set RX flow hash indir'n table */
	ETHTOOL_GFEATURES     = 0x0000003a /* Get device offload settings */
	ETHTOOL_SFEATURES     = 0x0000003b /* Change device offload settings */
	ETHTOOL_GCHANNELS     = 0x0000003c /* Get no of channels */
	ETHTOOL_SCHANNELS     = 0x0000003d /* Set no of channels */
	ETHTOOL_SET_DUMP      = 0x0000003e /* Set dump settings */
	ETHTOOL_GET_DUMP_FLAG = 0x0000003f /* Get dump settings */
	ETHTOOL_GET_DUMP_DATA = 0x00000040 /* Get dump data */
	ETHTOOL_GET_TS_INFO   = 0x00000041 /* Get time stamping and PHC info */
	ETHTOOL_GMODULEINFO   = 0x00000042 /* Get plug-in module information */
	ETHTOOL_GMODULEEEPROM = 0x00000043 /* Get plug-in module eeprom */
	ETHTOOL_GEEE          = 0x00000044 /* Get EEE settings */
	ETHTOOL_SEEE          = 0x00000045 /* Set EEE settings */
	ETHTOOL_GRSSH         = 0x00000046 /* Get RX flow hash configuration */
	ETHTOOL_SRSSH         = 0x00000047 /* Set RX flow hash configuration */
	ETHTOOL_GTUNABLE      = 0x00000048 /* Get tunable configuration */
	ETHTOOL_STUNABLE      = 0x00000049 /* Set tunable configuration */
	ETHTOOL_GPHYSTATS     = 0x0000004a /* get PHY-specific statistics */
	ETHTOOL_PERQUEUE      = 0x0000004b /* Set per queue options */
	ETHTOOL_GLINKSETTINGS = 0x0000004c /* Get ethtool_link_settings */
	ETHTOOL_SLINKSETTINGS = 0x0000004d /* Set ethtool_link_settings */
	ETHTOOL_PHY_GTUNABLE  = 0x0000004e /* Get PHY tunable configuration */
	ETHTOOL_PHY_STUNABLE  = 0x0000004f /* Set PHY tunable configuration */
	ETHTOOL_GFECPARAM     = 0x00000050 /* Get FEC settings */
	ETHTOOL_SFECPARAM     = 0x00000051 /* Set FEC settings */
)

var ethIoctlCmds = []string{
	"",
	"ETHTOOL_GSET",
	"ETHTOOL_SSET",
	"ETHTOOL_GDRVINFO",
	"ETHTOOL_GREGS",
	"ETHTOOL_GWOL",
	"ETHTOOL_SWOL",
	"ETHTOOL_GMSGLVL",
	"ETHTOOL_SMSGLVL",
	"ETHTOOL_NWAY_RST",
	"ETHTOOL_GLINK",
	"ETHTOOL_GEEPROM",
	"ETHTOOL_SEEPROM",
	"", // Missing
	"ETHTOOL_GCOALESCE",
	"ETHTOOL_SCOALESCE",
	"ETHTOOL_GRINGPARAM",
	"ETHTOOL_SRINGPARAM",
	"ETHTOOL_GPAUSEPARAM",
	"ETHTOOL_SPAUSEPARAM",
	"ETHTOOL_GRXCSUM",
	"ETHTOOL_SRXCSUM",
	"ETHTOOL_GTXCSUM",
	"ETHTOOL_STXCSUM",
	"ETHTOOL_GSG",
	"ETHTOOL_SSG",
	"ETHTOOL_TEST",
	"ETHTOOL_GSTRINGS",
	"ETHTOOL_PHYS_ID",
	"ETHTOOL_GSTATS",
	"ETHTOOL_GTSO",
	"ETHTOOL_STSO",
	"ETHTOOL_GPERMADDR",
	"ETHTOOL_GUFO",
	"ETHTOOL_SUFO",
	"ETHTOOL_GGSO",
	"ETHTOOL_SGSO",
	"ETHTOOL_GFLAGS",
	"ETHTOOL_SFLAGS",
	"ETHTOOL_GPFLAGS",
	"ETHTOOL_SPFLAGS",
	"ETHTOOL_GRXFH",
	"ETHTOOL_SRXFH",
	"ETHTOOL_GGRO",
	"ETHTOOL_SGRO",
	"ETHTOOL_GRXRINGS",
	"ETHTOOL_GRXCLSRLCNT",
	"ETHTOOL_GRXCLSRULE",
	"ETHTOOL_GRXCLSRLALL",
	"ETHTOOL_SRXCLSRLDEL",
	"ETHTOOL_SRXCLSRLINS",
	"ETHTOOL_FLASHDEV",
	"ETHTOOL_RESET",
	"ETHTOOL_SRXNTUPLE",
	"ETHTOOL_GRXNTUPLE",
	"ETHTOOL_GSSET_INFO",
	"ETHTOOL_GRXFHINDIR",
	"ETHTOOL_SRXFHINDIR",
	"ETHTOOL_GFEATURES",
	"ETHTOOL_SFEATURES",
	"ETHTOOL_GCHANNELS",
	"ETHTOOL_SCHANNELS",
	"ETHTOOL_SET_DUMP",
	"ETHTOOL_GET_DUMP_FLAG",
	"ETHTOOL_GET_DUMP_DATA",
	"ETHTOOL_GET_TS_INFO",
	"ETHTOOL_GMODULEINFO",
	"ETHTOOL_GMODULEEEPROM",
	"ETHTOOL_GEEE",
	"ETHTOOL_SEEE",
	"ETHTOOL_GRSSH",
	"ETHTOOL_SRSSH",
	"ETHTOOL_GTUNABLE",
	"ETHTOOL_STUNABLE",
	"ETHTOOL_GPHYSTATS",
	"ETHTOOL_PERQUEUE",
	"ETHTOOL_GLINKSETTINGS",
	"ETHTOOL_SLINKSETTINGS",
	"ETHTOOL_PHY_GTUNABLE",
	"ETHTOOL_PHY_STUNABLE",
	"ETHTOOL_GFECPARAM",
	"ETHTOOL_SFECPARAM",
}

func (cmd ethIoctlCmd) String() string {
	if int(cmd) < len(ethIoctlCmds) {
		return ethIoctlCmds[cmd]
	}

	return fmt.Sprintf("Unknown[%x]", int(cmd))
}

var ethtoolOptions = map[string]string{
	"--get-phy-tunable": "--get-phy-tunable(Get PHY tunable)",
	"--get-tunable":     "--get-tunable(Get tunable)",
	"--phy-statistics":  "--phy-statistics(Show phy statistics)",
	"--reset":           "--reset(Reset components)",
	"--set-eee":         "--set-eee(Set EEE settings)",
	"--set-fec":         "--set-fec(Set FEC settings)",
	"--set-phy-tunable": "--set-phy-tunable(Set PHY tunable)",
	"--set-priv-flags":  "--set-priv-flags(Set private flags)",
	"--set-pse":         "--set-pse(Set Power Sourcing Equipment settings)",
	"--set-tunable":     "--set-tunable(Set tunable)",
	"--show-eee":        "--show-eee(Show EEE settings)",
	"--show-fec":        "--show-fec(Show FEC settings)",
	"--show-priv-flags": "--show-priv-flags(Query private flags)",
	"-A":                "-A|--pause(Set pause options)",
	"-C":                "-C|--coalesce(Set coalesce options)",
	"-E":                "-E|--change-eeprom(Change bytes in device EEPROM)",
	"-G":                "-G|--set-ring(Set RX/TX ring parameters)",
	"-K":                "-K|--features|--offload(Set protocol offload and other features)",
	"-L":                "-L|--set-channels(Set Channels)",
	"-N":                "-N|-U|--config-nfc|--config-ntuple(Configure Rx network flow classification options or rules)",
	"-P":                "-P|--show-permaddr(Show permanent hardware address)",
	"-Q":                "-Q|--per-queue(Apply per-queue command.)",
	"-S":                "-S|--statistics(Show adapter statistics)",
	"-T":                "-T|--show-time-stamping(Show time stamping capabilities)",
	"-W":                "-W|--set-dump(Set dump flag of the device)",
	"-X":                "-X|--set-rxfh-indir|--rxfh(Set Rx flow hash indirection table and/or RSS hash key)",
	"-a":                "-a|--show-pause(Show pause options)",
	"-c":                "-c|--show-coalesce(Show coalesce options)",
	"-d":                "-d|--register-dump(Do a register dump)",
	"-e":                "-e|--eeprom-dump(Do a EEPROM dump)",
	"-f":                "-f|--flash(Flash firmware image from the specified file to a region on the device)",
	"-g":                "-g|--show-ring(Query RX/TX ring parameters)",
	"-i":                "-i|--driver(Show driver information)",
	"-k":                "-k|--show-features|--show-offload(Get state of protocol offload and other features)",
	"-l":                "-l|--show-channels(Query Channels)",
	"-m":                "-m|--dump-module-eeprom|--module-info(Query/Decode Module EEPROM information and optical diagnostics if available)",
	"-n":                "-n|-u|--show-nfc|--show-ntuple(Show Rx network flow classification options or rules)",
	"-p":                "-p|--identify(Show visible port identification (e.g. blinking))",
	"-r":                "-r|--negotiate(Restart N-WAY negotiation)",
	"-s":                "-s|--change(Change generic options)",
	"-t":                "-t|--test(Execute adapter self test)",
	"-w":                "-w|--get-dump(Get dump flag, data)",
	"-x":                "-x|--show-rxfh-indir|--show-rxfh(Show Rx flow hash indirection table and/or RSS hash key)",
	"<default>":         "<default>(Display standard information about device)",
}

var ethIoctlCmdMsgs = map[ethIoctlCmd]string{
	ETHTOOL_GSET:          "",
	ETHTOOL_SSET:          "",
	ETHTOOL_GDRVINFO:      "-d,-e,-i",
	ETHTOOL_GREGS:         "-d",
	ETHTOOL_GWOL:          "<default>,-s",
	ETHTOOL_SWOL:          "-s",
	ETHTOOL_GMSGLVL:       "<default>,-s",
	ETHTOOL_SMSGLVL:       "-s",
	ETHTOOL_NWAY_RST:      "-r",
	ETHTOOL_GLINK:         "<default>",
	ETHTOOL_GEEPROM:       "-e",
	ETHTOOL_SEEPROM:       "-E",
	ETHTOOL_GCOALESCE:     "-c,-C",
	ETHTOOL_SCOALESCE:     "-C",
	ETHTOOL_GRINGPARAM:    "-g,-G",
	ETHTOOL_SRINGPARAM:    "-G",
	ETHTOOL_GPAUSEPARAM:   "-a,-A",
	ETHTOOL_SPAUSEPARAM:   "-A",
	ETHTOOL_GRXCSUM:       "",
	ETHTOOL_SRXCSUM:       "",
	ETHTOOL_GTXCSUM:       "",
	ETHTOOL_STXCSUM:       "",
	ETHTOOL_GSG:           "",
	ETHTOOL_SSG:           "",
	ETHTOOL_TEST:          "-t",
	ETHTOOL_GSTRINGS:      "-x,-S",
	ETHTOOL_PHYS_ID:       "-p",
	ETHTOOL_GSTATS:        "-S",
	ETHTOOL_GTSO:          "",
	ETHTOOL_STSO:          "",
	ETHTOOL_GPERMADDR:     "-P",
	ETHTOOL_GUFO:          "",
	ETHTOOL_SUFO:          "",
	ETHTOOL_GGSO:          "",
	ETHTOOL_SGSO:          "",
	ETHTOOL_GFLAGS:        "-C",
	ETHTOOL_SFLAGS:        "-K",
	ETHTOOL_GPFLAGS:       "--show-priv-flags",
	ETHTOOL_SPFLAGS:       "--set-priv-flags",
	ETHTOOL_GRXFH:         "-n",
	ETHTOOL_SRXFH:         "-N",
	ETHTOOL_GGRO:          "",
	ETHTOOL_SGRO:          "",
	ETHTOOL_GRXRINGS:      "-n,-x,-X",
	ETHTOOL_GRXCLSRLCNT:   "-n",
	ETHTOOL_GRXCLSRULE:    "",
	ETHTOOL_GRXCLSRLALL:   "",
	ETHTOOL_SRXCLSRLDEL:   "",
	ETHTOOL_SRXCLSRLINS:   "",
	ETHTOOL_FLASHDEV:      "-f",
	ETHTOOL_RESET:         "--reset",
	ETHTOOL_SRXNTUPLE:     "",
	ETHTOOL_GRXNTUPLE:     "",
	ETHTOOL_GSSET_INFO:    "--phy-statistics,-t,-x,-S",
	ETHTOOL_GRXFHINDIR:    "",
	ETHTOOL_SRXFHINDIR:    "",
	ETHTOOL_GFEATURES:     "-C",
	ETHTOOL_SFEATURES:     "-K",
	ETHTOOL_GCHANNELS:     "-l,-L",
	ETHTOOL_SCHANNELS:     "-L",
	ETHTOOL_SET_DUMP:      "-W",
	ETHTOOL_GET_DUMP_FLAG: "-w",
	ETHTOOL_GET_DUMP_DATA: "-w",
	ETHTOOL_GET_TS_INFO:   "-T",
	ETHTOOL_GMODULEINFO:   "-m",
	ETHTOOL_GMODULEEEPROM: "-m",
	ETHTOOL_GEEE:          "--show-eee",
	ETHTOOL_SEEE:          "--set-eee",
	ETHTOOL_GRSSH:         "-x,-X",
	ETHTOOL_SRSSH:         "-X",
	ETHTOOL_GTUNABLE:      "--get-tunable",
	ETHTOOL_STUNABLE:      "--set-tunable",
	ETHTOOL_GPHYSTATS:     "--phy-statistics",
	ETHTOOL_PERQUEUE:      "",
	ETHTOOL_GLINKSETTINGS: "-K",
	ETHTOOL_SLINKSETTINGS: "",
	ETHTOOL_PHY_GTUNABLE:  "--get-phy-tunable",
	ETHTOOL_PHY_STUNABLE:  "--set-phy-tunable",
	ETHTOOL_GFECPARAM:     "--show-fec",
	ETHTOOL_SFECPARAM:     "--set-fec",
}

func init() {
	for k, v := range ethIoctlCmdMsgs {
		if v == "" {
			continue
		}

		var msgs []string
		for _, opt := range strings.Split(v, ",") {
			if o := ethtoolOptions[opt]; o != "" {
				msgs = append(msgs, o)
			}
		}

		if len(msgs) > 0 {
			ethIoctlCmdMsgs[k] = strings.Join(msgs, ", ")
		}
	}
}

func (cmd ethIoctlCmd) Message() string {
	return ethIoctlCmdMsgs[cmd]
}

type ethGenlCmd uint8

const (
	ETHTOOL_MSG_USER_NONE ethGenlCmd = iota
	ETHTOOL_MSG_STRSET_GET
	ETHTOOL_MSG_LINKINFO_GET
	ETHTOOL_MSG_LINKINFO_SET
	ETHTOOL_MSG_LINKMODES_GET
	ETHTOOL_MSG_LINKMODES_SET
	ETHTOOL_MSG_LINKSTATE_GET
	ETHTOOL_MSG_DEBUG_GET
	ETHTOOL_MSG_DEBUG_SET
	ETHTOOL_MSG_WOL_GET
	ETHTOOL_MSG_WOL_SET
	ETHTOOL_MSG_FEATURES_GET
	ETHTOOL_MSG_FEATURES_SET
	ETHTOOL_MSG_PRIVFLAGS_GET
	ETHTOOL_MSG_PRIVFLAGS_SET
	ETHTOOL_MSG_RINGS_GET
	ETHTOOL_MSG_RINGS_SET
	ETHTOOL_MSG_CHANNELS_GET
	ETHTOOL_MSG_CHANNELS_SET
	ETHTOOL_MSG_COALESCE_GET
	ETHTOOL_MSG_COALESCE_SET
	ETHTOOL_MSG_PAUSE_GET
	ETHTOOL_MSG_PAUSE_SET
	ETHTOOL_MSG_EEE_GET
	ETHTOOL_MSG_EEE_SET
	ETHTOOL_MSG_TSINFO_GET
	ETHTOOL_MSG_CABLE_TEST_ACT
	ETHTOOL_MSG_CABLE_TEST_TDR_ACT
	ETHTOOL_MSG_TUNNEL_INFO_GET
	ETHTOOL_MSG_FEC_GET
	ETHTOOL_MSG_FEC_SET
	ETHTOOL_MSG_MODULE_EEPROM_GET
	ETHTOOL_MSG_STATS_GET
	ETHTOOL_MSG_PHC_VCLOCKS_GET
	ETHTOOL_MSG_MODULE_GET
	ETHTOOL_MSG_MODULE_SET
	ETHTOOL_MSG_PSE_GET
	ETHTOOL_MSG_PSE_SET
	ETHTOOL_MSG_RSS_GET
	ETHTOOL_MSG_PLCA_GET_CFG
	ETHTOOL_MSG_PLCA_SET_CFG
	ETHTOOL_MSG_PLCA_GET_STATUS
	ETHTOOL_MSG_MM_GET
	ETHTOOL_MSG_MM_SET
)

var ethGenlCmds = []string{
	"ETHTOOL_MSG_USER_NONE",
	"ETHTOOL_MSG_STRSET_GET",
	"ETHTOOL_MSG_LINKINFO_GET",
	"ETHTOOL_MSG_LINKINFO_SET",
	"ETHTOOL_MSG_LINKMODES_GET",
	"ETHTOOL_MSG_LINKMODES_SET",
	"ETHTOOL_MSG_LINKSTATE_GET",
	"ETHTOOL_MSG_DEBUG_GET",
	"ETHTOOL_MSG_DEBUG_SET",
	"ETHTOOL_MSG_WOL_GET",
	"ETHTOOL_MSG_WOL_SET",
	"ETHTOOL_MSG_FEATURES_GET",
	"ETHTOOL_MSG_FEATURES_SET",
	"ETHTOOL_MSG_PRIVFLAGS_GET",
	"ETHTOOL_MSG_PRIVFLAGS_SET",
	"ETHTOOL_MSG_RINGS_GET",
	"ETHTOOL_MSG_RINGS_SET",
	"ETHTOOL_MSG_CHANNELS_GET",
	"ETHTOOL_MSG_CHANNELS_SET",
	"ETHTOOL_MSG_COALESCE_GET",
	"ETHTOOL_MSG_COALESCE_SET",
	"ETHTOOL_MSG_PAUSE_GET",
	"ETHTOOL_MSG_PAUSE_SET",
	"ETHTOOL_MSG_EEE_GET",
	"ETHTOOL_MSG_EEE_SET",
	"ETHTOOL_MSG_TSINFO_GET",
	"ETHTOOL_MSG_CABLE_TEST_ACT",
	"ETHTOOL_MSG_CABLE_TEST_TDR_ACT",
	"ETHTOOL_MSG_TUNNEL_INFO_GET",
	"ETHTOOL_MSG_FEC_GET",
	"ETHTOOL_MSG_FEC_SET",
	"ETHTOOL_MSG_MODULE_EEPROM_GET",
	"ETHTOOL_MSG_STATS_GET",
	"ETHTOOL_MSG_PHC_VCLOCKS_GET",
	"ETHTOOL_MSG_MODULE_GET",
	"ETHTOOL_MSG_MODULE_SET",
	"ETHTOOL_MSG_PSE_GET",
	"ETHTOOL_MSG_PSE_SET",
	"ETHTOOL_MSG_RSS_GET",
	"ETHTOOL_MSG_PLCA_GET_CFG",
	"ETHTOOL_MSG_PLCA_SET_CFG",
	"ETHTOOL_MSG_PLCA_GET_STATUS",
	"ETHTOOL_MSG_MM_GET",
	"ETHTOOL_MSG_MM_SET",
}

func (cmd ethGenlCmd) String() string {
	if int(cmd) < len(ethGenlCmds) {
		return ethGenlCmds[cmd]
	}

	return fmt.Sprintf("Unknown[%x]", int(cmd))
}

var ethGenlCmdMsgs = map[ethGenlCmd]string{
	ETHTOOL_MSG_USER_NONE:          "",
	ETHTOOL_MSG_STRSET_GET:         "-k",
	ETHTOOL_MSG_LINKINFO_GET:       "<default>",
	ETHTOOL_MSG_LINKINFO_SET:       "",
	ETHTOOL_MSG_LINKMODES_GET:      "<default>",
	ETHTOOL_MSG_LINKMODES_SET:      "",
	ETHTOOL_MSG_LINKSTATE_GET:      "<default>",
	ETHTOOL_MSG_DEBUG_GET:          "<default>",
	ETHTOOL_MSG_DEBUG_SET:          "",
	ETHTOOL_MSG_WOL_GET:            "<default>",
	ETHTOOL_MSG_WOL_SET:            "",
	ETHTOOL_MSG_FEATURES_GET:       "-k",
	ETHTOOL_MSG_FEATURES_SET:       "-K",
	ETHTOOL_MSG_PRIVFLAGS_GET:      "--show-priv-flags",
	ETHTOOL_MSG_PRIVFLAGS_SET:      "",
	ETHTOOL_MSG_RINGS_GET:          "-g",
	ETHTOOL_MSG_RINGS_SET:          "-G",
	ETHTOOL_MSG_CHANNELS_GET:       "-l",
	ETHTOOL_MSG_CHANNELS_SET:       "-L",
	ETHTOOL_MSG_COALESCE_GET:       "-c",
	ETHTOOL_MSG_COALESCE_SET:       "-C",
	ETHTOOL_MSG_PAUSE_GET:          "-a",
	ETHTOOL_MSG_PAUSE_SET:          "-A",
	ETHTOOL_MSG_EEE_GET:            "--show-eee",
	ETHTOOL_MSG_EEE_SET:            "--set-eee",
	ETHTOOL_MSG_TSINFO_GET:         "-T",
	ETHTOOL_MSG_CABLE_TEST_ACT:     "",
	ETHTOOL_MSG_CABLE_TEST_TDR_ACT: "",
	ETHTOOL_MSG_TUNNEL_INFO_GET:    "",
	ETHTOOL_MSG_FEC_GET:            "--show-fec",
	ETHTOOL_MSG_FEC_SET:            "--set-fec",
	ETHTOOL_MSG_MODULE_EEPROM_GET:  "-m",
	ETHTOOL_MSG_STATS_GET:          "",
	ETHTOOL_MSG_PHC_VCLOCKS_GET:    "",
	ETHTOOL_MSG_MODULE_GET:         "",
	ETHTOOL_MSG_MODULE_SET:         "",
	ETHTOOL_MSG_PSE_GET:            "",
	ETHTOOL_MSG_PSE_SET:            "",
	ETHTOOL_MSG_RSS_GET:            "",
	ETHTOOL_MSG_PLCA_GET_CFG:       "",
	ETHTOOL_MSG_PLCA_SET_CFG:       "",
	ETHTOOL_MSG_PLCA_GET_STATUS:    "",
	ETHTOOL_MSG_MM_GET:             "",
	ETHTOOL_MSG_MM_SET:             "",
}

func init() {
	for k, v := range ethGenlCmdMsgs {
		if v == "" {
			continue
		}

		var msgs []string
		for _, msg := range strings.Split(v, ",") {
			if m := ethtoolOptions[msg]; m != "" {
				msgs = append(msgs, m)
			}
		}

		if len(msgs) > 0 {
			ethGenlCmdMsgs[k] = strings.Join(msgs, ", ")
		}
	}
}

func (cmd ethGenlCmd) Message() string {
	return ethGenlCmdMsgs[cmd]
}
