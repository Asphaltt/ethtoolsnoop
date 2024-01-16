// Copyright 2024 Leon Hwang.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"os/signal"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/perf"
	"github.com/cilium/ebpf/rlimit"
	flag "github.com/spf13/pflag"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sys/unix"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang -no-strip -no-global-types ethtool ./bpf/ethtool.c -- -D__TARGET_ARCH_x86 -I./bpf/headers

var flags struct {
	debug bool
}

func init() {
	flag.BoolVar(&flags.debug, "debug", false, "debug mode")
	flag.Parse()
}

func main() {
	if err := unix.Setrlimit(unix.RLIMIT_NOFILE, &unix.Rlimit{
		Cur: 8192,
		Max: 8192,
	}); err != nil {
		log.Fatalf("failed to set temporary rlimit: %s", err)
	}
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatalf("Failed to set temporary rlimit: %s", err)
	}

	var obj ethtoolObjects
	if err := loadEthtoolObjects(&obj, nil); err != nil {
		log.Fatalf("Failed to load objects: %s", err)
	}
	defer obj.Close()

	if kp, err := link.Kprobe("dev_ethtool", obj.KpDevEthtool, nil); err != nil {
		log.Fatalf("Failed to create kprobe: %s", err)
	} else {
		defer kp.Close()
	}

	if krp, err := link.Kretprobe("ethnl_parse_header_dev_get", obj.KrpEthnlDev, nil); err != nil {
		log.Fatalf("Failed to create kretprobe: %s", err)
	} else {
		defer krp.Close()
	}

	if kp, err := link.Kprobe("ethnl_parse_header_dev_get", obj.KpEthnlDev, nil); err != nil {
		log.Fatalf("Failed to create kprobe: %s", err)
	} else {
		defer kp.Close()
	}

	if kp, err := link.Kprobe("ethnl_default_doit", obj.KpEthnlDoit, nil); err != nil {
		log.Fatalf("Failed to create kprobe: %s", err)
	} else {
		defer kp.Close()
	}

	ctx, stop := signal.NotifyContext(context.Background(), unix.SIGINT, unix.SIGTERM)
	defer stop()
	errg, ctx := errgroup.WithContext(ctx)

	reader, err := perf.NewReader(obj.Events, 4096)
	if err != nil {
		log.Fatalf("Failed to create perf event reader: %s", err)
	}

	errg.Go(func() error {
		<-ctx.Done()
		_ = reader.Close()
		return nil
	})

	errg.Go(func() error {
		return readEvent(ctx, reader)
	})

	if err := errg.Wait(); err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func readEvent(ctx context.Context, reader *perf.Reader) error {
	printHeader()

	var ev event
	for {
		record, err := reader.Read()
		if err != nil {
			if err == perf.ErrClosed {
				return nil
			}
			select {
			case <-ctx.Done():
				return nil
			default:
				return fmt.Errorf("failed to read record: %s", err)
			}
		}

		if record.LostSamples != 0 {
			log.Printf("Lost %d samples", record.LostSamples)
		}

		binary.Read(bytes.NewReader(record.RawSample), binary.LittleEndian, &ev)

		ev.print()

		select {
		case <-ctx.Done():
			return nil
		default:
		}
	}
}
