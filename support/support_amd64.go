//go:build amd64 && !dummy

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package support // import "github.com/open-telemetry/opentelemetry-ebpf-profiler/support"

import (
	_ "embed"
)

//go:embed ebpf/tracer.ebpf.release.amd64
var tracerData []byte
