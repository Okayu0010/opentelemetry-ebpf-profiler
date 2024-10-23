//go:build arm64

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tpbase // import "github.com/open-telemetry/opentelemetry-ebpf-profiler/tpbase"

func GetAnalyzers() []Analyzer {
	return arm64GetAnalyzers()
}
