# Copyright 2024 Leon Hwang.
# SPDX-License-Identifier: Apache-2.0

GOGEN := go generate
GOBUILD := GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath

ETHTOOLSNOOP_SRC := .
ETHTOOLSNOOP_BIN := ethtoolsnoop

ETHTOOLSNOOP_BPF_SRC := $(ETHTOOLSNOOP_SRC)/bpf/ethtool.c
ETHTOOLSNOOP_BPF_OBJ := $(ETHTOOLSNOOP_SRC)/ethtool_bpfel.o $(ETHTOOLSNOOP_SRC)/ethtool_bpfeb.o

.PHONY: build
.DEFAULT_GOAL := build

$(ETHTOOLSNOOP_BPF_OBJ): $(ETHTOOLSNOOP_BPF_SRC)
	$(GOGEN) .

$(ETHTOOLSNOOP_BIN): $(ETHTOOLSNOOP_SRC)
	$(GOBUILD) -o $@ $(ETHTOOLSNOOP_SRC)

build: $(ETHTOOLSNOOP_BPF_OBJ) $(ETHTOOLSNOOP_BIN)
