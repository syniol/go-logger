# Performance Benchmarks
This document tracks the performance of the logger package to ensure it 
remains suitable for high-throughput production environments. We measure 
execution time, memory usage, and heap allocations.


## Methodology
Benchmarks are executed using the standard Go testing tool: 
`go test -bench=. -benchmem -count=10 > bench_results.txt`

The results below are generated on:

 * **OS:** macOS darwin
 * **Architecture:** ARM64
 * **CPU:** Apple M2
 * **Go Version:** 1.24

   BenchmarkSyniolLogger-8     	 1891555	       611.1 ns/op	     664 B/op	       9 allocs/op
   BenchmarkSyniolLogger-8     	 1957856	       613.0 ns/op	     664 B/op	       9 allocs/op
   BenchmarkSyniolLogger-8     	 1957063	       615.2 ns/op	     664 B/op	       9 allocs/op
   BenchmarkSyniolLogger-8     	 1886391	       621.1 ns/op	     664 B/op	       9 allocs/op
   BenchmarkSyniolLogger-8     	 1942413	       614.8 ns/op	     664 B/op	       9 allocs/op
   BenchmarkSyniolLogger-8     	 1951993	       618.5 ns/op	     664 B/op	       9 allocs/op
   BenchmarkSyniolLogger-8     	 1935925	       617.1 ns/op	     664 B/op	       9 allocs/op
   BenchmarkSyniolLogger-8     	 1897650	       624.7 ns/op	     664 B/op	       9 allocs/op
   BenchmarkSyniolLogger-8     	 1945797	       630.0 ns/op	     664 B/op	       9 allocs/op
   BenchmarkSyniolLogger-8     	 1942968	       628.3 ns/op	     664 B/op	       9 allocs/op
   BenchmarkSlogJSON-8         	 2352092	       510.8 ns/op	       0 B/op	       0 allocs/op
   BenchmarkSlogJSON-8         	 2353771	       509.9 ns/op	       0 B/op	       0 allocs/op
   BenchmarkSlogJSON-8         	 2380887	       510.9 ns/op	       0 B/op	       0 allocs/op
   BenchmarkSlogJSON-8         	 2371741	       507.7 ns/op	       0 B/op	       0 allocs/op
   BenchmarkSlogJSON-8         	 2364392	       513.5 ns/op	       0 B/op	       0 allocs/op
   BenchmarkSlogJSON-8         	 2346122	       508.3 ns/op	       0 B/op	       0 allocs/op
   BenchmarkSlogJSON-8         	 2375124	       504.0 ns/op	       0 B/op	       0 allocs/op
   BenchmarkSlogJSON-8         	 2372458	       505.7 ns/op	       0 B/op	       0 allocs/op
   BenchmarkSlogJSON-8         	 2375211	       502.5 ns/op	       0 B/op	       0 allocs/op
   BenchmarkSlogJSON-8         	 2354427	       504.1 ns/op	       0 B/op	       0 allocs/op
   BenchmarkSlogWithSource-8   	 1000000	      1053 ns/op	     584 B/op	       6 allocs/op
   BenchmarkSlogWithSource-8   	 1000000	      1067 ns/op	     584 B/op	       6 allocs/op
   BenchmarkSlogWithSource-8   	 1000000	      1058 ns/op	     584 B/op	       6 allocs/op
   BenchmarkSlogWithSource-8   	 1000000	      1052 ns/op	     584 B/op	       6 allocs/op
   BenchmarkSlogWithSource-8   	 1000000	      1053 ns/op	     584 B/op	       6 allocs/op
   BenchmarkSlogWithSource-8   	 1000000	      1054 ns/op	     584 B/op	       6 allocs/op
   BenchmarkSlogWithSource-8   	 1000000	      1053 ns/op	     584 B/op	       6 allocs/op
   BenchmarkSlogWithSource-8   	 1000000	      1050 ns/op	     584 B/op	       6 allocs/op
   BenchmarkSlogWithSource-8   	 1000000	      1050 ns/op	     584 B/op	       6 allocs/op
   BenchmarkSlogWithSource-8   	 1000000	      1053 ns/op	     584 B/op	       6 allocs/op
   PASS
   ok  	github.com/syniol/go-logger	46.628s

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
