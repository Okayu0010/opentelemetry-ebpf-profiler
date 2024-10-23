package libpf // import "github.com/open-telemetry/opentelemetry-ebpf-profiler/libpf"

// PID represent Unix Process ID (pid_t)
type PID uint32

func (p PID) Hash32() uint32 {
	return uint32(p)
}
