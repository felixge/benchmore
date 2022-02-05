package cputime

import (
	"syscall"
	"testing"
	"time"
)

// ReportCPUTime takes a *testing.B, starts a CPU timer, and returns a function
// to stop the timer and report the amount of passed time to b.ReportMetric().
func ReportCPUTime(b *testing.B) func() {
	start, err := cpuTime()
	if err != nil {
		b.Fatal(err)
	}
	return func() {
		end, err := cpuTime()
		if err != nil {
			b.Fatal(err)
		}
		dt := end - start
		b.ReportMetric(float64(dt)/float64(b.N), "cpu-ns/op")
	}
}

// cpuTime returns the amount of time the Go process has been running on the
// cpuTime since it's been started.
func cpuTime() (time.Duration, error) {
	var rusage syscall.Rusage
	if err := syscall.Getrusage(syscall.RUSAGE_SELF, &rusage); err != nil {
		return 0, err
	}
	d := time.Duration(rusage.Stime.Nano()) + time.Duration(rusage.Utime.Nano())
	return d, nil
}
