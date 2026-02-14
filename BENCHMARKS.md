# Performance Benchmarks
This document tracks the performance of the logger package to ensure it 
remains suitable for high-throughput production environments. We measure 
execution time, memory usage, and heap allocations.


## Methodology
Benchmarks are executed using the standard Go testing tool: 
```shell
go test -bench=. -benchmem -count=10 > bench_results.txt
```

The results below are generated on:

 * **OS:** Darwin _(macOS)_
 * **Architecture:** ARM64
 * **CPU:** Apple M2
 * **Go Version:** 1.24

| Benchmark Case            | Iterations | Memory Allocations   | Bytes Allocations | Operation Allocations |
|---------------------------|------------|----------------------|-------------------|-----------------------|
| BenchmarkSyniolLogger-8   | 	 2792167  | 	       407.3 ns/op	 | 336 B/op	         | 6 allocs/op           |
| BenchmarkSyniolLogger-8   | 	 2944899	 | 408.6 ns/op	         | 336 B/op	         | 6 allocs/op           |
| BenchmarkSyniolLogger-8   | 	 2927808	 | 409.0 ns/op	         | 336 B/op	         | 6 allocs/op           |
| BenchmarkSyniolLogger-8   | 	 2939458	 | 409.0 ns/op	         | 336 B/op	         | 6 allocs/op           |
| BenchmarkSyniolLogger-8   | 	 2938980	 | 408.9 ns/op	         | 336 B/op	         | 6 allocs/op           |
| BenchmarkSyniolLogger-8   | 	 2927814	 | 411.8 ns/op	         | 336 B/op	         | 6 allocs/op           |
| BenchmarkSyniolLogger-8   | 	 2787464	 | 410.9 ns/op	         | 336 B/op	         | 6 allocs/op           |
| BenchmarkSyniolLogger-8   | 	 2920656	 | 409.7 ns/op	         | 336 B/op	         | 6 allocs/op           |
| BenchmarkSyniolLogger-8   | 	 2928950	 | 411.4 ns/op	         | 336 B/op	         | 6 allocs/op           |
| BenchmarkSyniolLogger-8   | 	 2881113	 | 410.9 ns/op	         | 336 B/op	         | 6 allocs/op           |
| BenchmarkSlogJSON-8       | 	 2355129	 | 508.6 ns/op	         | 0 B/op	           | 0 allocs/op           |
| BenchmarkSlogJSON-8       | 	 2355932	 | 513.1 ns/op	         | 0 B/op	           | 0 allocs/op           |
| BenchmarkSlogJSON-8       | 	 2373483	 | 511.4 ns/op	         | 0 B/op	           | 0 allocs/op           |
| BenchmarkSlogJSON-8       | 	 2355658	 | 508.8 ns/op	         | 0 B/op	           | 0 allocs/op           |
| BenchmarkSlogJSON-8       | 	 2359935	 | 510.0 ns/op	         | 0 B/op	           | 0 allocs/op           |
| BenchmarkSlogJSON-8       | 	 2353203	 | 507.1 ns/op	         | 0 B/op	           | 0 allocs/op           |
| BenchmarkSlogJSON-8       | 	 2361937	 | 517.4 ns/op	         | 0 B/op	           | 0 allocs/op           |
| BenchmarkSlogJSON-8       | 	 2343535	 | 518.6 ns/op	         | 0 B/op	           | 0 allocs/op           |
| BenchmarkSlogJSON-8       | 	 2300044	 | 516.1 ns/op	         | 0 B/op	           | 0 allocs/op           |
| BenchmarkSlogJSON-8       | 	 2342082	 | 515.3 ns/op	         | 0 B/op	           | 0 allocs/op           |
| BenchmarkSlogWithSource-8 | 	 1000000	 | 1056 ns/op	          | 584 B/op	         | 6 allocs/op           |
| BenchmarkSlogWithSource-8 | 	 1000000	 | 1044 ns/op	          | 584 B/op	         | 6 allocs/op           |
| BenchmarkSlogWithSource-8 | 	 1000000	 | 1044 ns/op	          | 584 B/op	         | 6 allocs/op           |
| BenchmarkSlogWithSource-8 | 	 1000000	 | 1042 ns/op	          | 584 B/op	         | 6 allocs/op           |
| BenchmarkSlogWithSource-8 | 	 1000000	 | 1042 ns/op	          | 584 B/op	         | 6 allocs/op           |
| BenchmarkSlogWithSource-8 | 	 1000000	 | 1042 ns/op	          | 584 B/op	         | 6 allocs/op           |
| BenchmarkSlogWithSource-8 | 	 1000000	 | 1042 ns/op	          | 584 B/op	         | 6 allocs/op           |
| BenchmarkSlogWithSource-8 | 	 1000000	 | 1043 ns/op	          | 584 B/op	         | 6 allocs/op           |
| BenchmarkSlogWithSource-8 | 	 1000000	 | 1051 ns/op	          | 584 B/op	         | 6 allocs/op           |
| BenchmarkSlogWithSource-8 | 	 1000000	 | 1054 ns/op	          | 584 B/op	         | 6 allocs/op           |

```shell
PASS
ok  github.com/syniol/go-logger	44.622s
```


### Comparison: logger vs. slog (Standard Library)
> **Note on the Trace Field:** Our logger captures the caller's file and line 
number automatically. While this adds a small overhead compared to standard 
logging, it significantly reduces MTTR (Mean Time to Recovery) by providing 
instant forensic context.


#### Analysis of Allocations
The primary source of allocations in this package comes from:

 1. **Stack Tracing:** The `runtime.Caller` lookup required for the trace field.


##### Running Benchmarks Locally
To reproduce these results, run the following command from the root of the repository.
