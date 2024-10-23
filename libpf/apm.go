// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package libpf // import "github.com/open-telemetry/opentelemetry-ebpf-profiler/libpf"

type APMSpanID [8]byte
type APMTraceID [16]byte
type APMTransactionID = APMSpanID

var InvalidAPMSpanID = APMSpanID{0, 0, 0, 0, 0, 0, 0, 0}
