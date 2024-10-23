//go:build !debugtracer

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package support // import "github.com/open-telemetry/opentelemetry-ebpf-profiler/support"

// debugtracer_dummy.go satisfies build requirements where the eBPF debug tracers
// file does not exist.
var debugTracerData []byte
