package cputime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_cpuTime(t *testing.T) {
	start, err := cpuTime()
	require.NoError(t, err)

	goal := 100 * time.Millisecond
	cpuHog(goal)

	stop, err := cpuTime()
	require.NoError(t, err)

	dt := stop - start
	require.Greater(t, dt, goal/2)
	require.Less(t, dt, goal*2)
}

func cpuHog(d time.Duration) {
	done := time.After(d)
	for {
		select {
		case <-done:
			return
		default:
			_ = struct{}{} // do nothing
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	defer ReportCPUTime(b)()

	for i := 0; i < b.N; i++ {
		add(i, i)
	}
}

func add(a, b int) int {
	return a + b
}
