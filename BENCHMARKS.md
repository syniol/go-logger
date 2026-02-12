# Performance Benchmarks
This document tracks the performance of the logger package to ensure it 
remains suitable for high-throughput production environments. We measure 
execution time, memory usage, and heap allocations.


## Methodology
Benchmarks are executed using the standard Go testing tool: 
`go test -bench=. -benchmem -count=10 > bench_results.txt`

The results below are generated on:

 * **OS:** macOS / Linux (specify your dev machine)
 * **Architecture:** ARM64 / AMD64
 * **Go Version:** 1.2x


### Comparison: logger vs. slog (Standard Library)
> **Note on the Trace Field:** Our logger captures the caller's file and line 
number automatically. While this adds a small overhead compared to standard 
logging, it significantly reduces MTTR (Mean Time to Recovery) by providing 
instant forensic context.


#### Analysis of Allocations
The primary source of allocations in this package comes from:

 1. **Interface Conversion:** Using `...interface{}` for variadic arguments.
 2. **Stack Tracing:** The `runtime.Caller` lookup required for the trace field.


##### Running Benchmarks Locally
To reproduce these results, run the following command from the root of the repository.