# benchmore

This package allows you to report On-CPU Time in addition to the wall time
measured by Go's builtin benchmarking framework. Example:

```go
import "github.com/felixge/benchmore"

func BenchmarkAdd(b *testing.B) {
	defer benchmore.ReportCPUTime(b)()

	for i := 0; i < b.N; i++ {
		Add(i, i)
	}
}

func Add(a, b int) int {
	return a + b
}
```
```
goos: darwin
goarch: amd64
pkg: github.com/felixge/benchmore
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkAdd-12    	1000000000	        0.2546 ns/op	        0.2545 cpu-ns/op
PASS
ok  	github.com/felixge/benchmore	0.429s
```

# More benchmore?

This package was created as a quick proof of concept response to a [tweet by @bwplotka](https://twitter.com/bwplotka/status/1490013375228698634).

In the future it could be extended to capture other metrics such as heap, rss, context switches, etc.

# License

MIT
