//go:build dummy

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package support // import "github.com/open-telemetry/opentelemetry-ebpf-profiler/support"

// support_dummy.go satisfies build requirements where the eBPF tracers file does
// not exist.

var tracerData []byte
