# SyLog: Structured, Facility-Aware Logging for Go
![workflow](https://github.com/syniol/go-logger/actions/workflows/ci.yml/badge.svg)  

<p align='center'>
	<img style='max-width: 100%; width: 320px;' src='https://github.com/user-attachments/assets/865f2acc-30e0-4c5b-8ce6-069b3204b740' />
</p>

SyLog brings the battle-tested Facility and Severity methodology of Unix Syslog 
to modern Go applications. It provides a structured, JSON-first logging interface 
designed to give distributed systems the same level of granular categorisation 
that system administrators have relied on for decades.


## Why SyLog?
In the world of microservices, a simple string message is rarely enough. By implementing 
the facility concept, this package allows you to categorise logs not just by what happened, but by where it originated within your architecture (e.g., auth-service, payment-gateway, kernel).

 * **Contextual Tracing:** Automatically captures file locations and line numbers to pinpoint failures.
 * **RFC-Inspired Severities:** Moves beyond simple "Info/Error" levels to include Alert, Critical, and Emergency states.
 * **Structured Output:** Native JSON formatting out of the box, ready for ingestion by ELK, Splunk, or Grafana Loki.
 * **Developer-Centric API:** A clean, variadic interface that feels natural to Go developers.


## Quick Start
First you need to add and download SyLog dependency in your Go project.
```shell
go get github.com/syniol/go-logger
```

**Example code:**
```go
package main

import "github.com/syniol/go-logger"

func main() {
	sylog.LogAlert("pay microservice", "mock message for alert log")
}
```

**Resulting Output:**
```json
{
  "level": "alert",
  "facility": "pay microservice",
  "message": "mock message for alert log",
  "trace": [
    "location: '/Users/hadi/dev/golang/logger/logger.go' on line: 100"
  ],
  "timestamp": "2023-11-04T21:36:33Z"
}
```


## Log Schema Definition
Every log entry is emitted as a structured JSON object. This consistent schema ensures 
that your logs are machine-readable and ready for high-performance indexing in modern 
observability stacks.

| Field       | Type     | Description                                                                                                                                                                                                            |
|-------------|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `level`     | `string` | The severity level mapped from the invoked method (e.g., `alert`, `warn`, `info` ). Derived from **RFC 5424** standards.                                                                                               |
| `facility`  | `string` | The logical component or service name. Used as the primary partition for log filtering in distributed environments.                                                                                                    |
| `message`   | `string` | The primary log descriptor. This is captured from the first argument provided to the logging method.                                                                                                                   |
| `trace`     | `array`  | A collection of contextual metadata. The first element is automatically injected with the file path and line number of the caller. Subsequent elements contain any additional variadic arguments passed to the method. |
| `timestamp` | `string` | The UTC time of event generation, formatted according to the **ISO 8601 / RFC 3339** standard for universal compatibility.                                                                                             |

### A Design Note on the `trace` Array
The trace field is designed for **maximum forensic utility**. By capturing the exact 
line number where the log was triggered, it eliminates the guesswork during incident 
response. Because it is an array, it keeps your top-level JSON namespace clean while 
allowing you to pass as much supplemental context as needed.


## Severity Levels: The 8-Tier Standard
This package implements the full spectrum of Unix Syslog severity levels. Rather than passing a level
as a variable, SyLog provides dedicated methods for each. This ensures type safety and clear intent at
the call site.

### Which Method to Use?
Choosing the right level is critical for effective alerting and noise reduction in your monitoring dashboards.

| Method           | Level             | Syslog Code | Use Case                                                                        |
|------------------|-------------------|-------------|---------------------------------------------------------------------------------|
| `LogEmergency()` | **Emergency**     | 0           | **System is unusable.** Total failure. Everyone gets paged.                     |
| `LogAlert()`     | **Alert**         | 1           | **Immediate action required.** e.g., Primary database down.                     |                           
| `LogCritical()`  | **Critical**      | 2           | **Critical conditions.** e.g., Hard device errors or loss of primary functions. |
| `LogError()`     | **Error**         | 3           | **Error conditions.** Something failed, but the app is still running.           |                
| `LogWarning()`   | **Warning**       | 4           | **Warning conditions.** Potential issues that should be investigated.           |               
| `LogNotice()`    | **Notice**        | 5           | **Normal but significant.** Non-error events that deserve tracking.             |                   
| `LogInfo()`      | **Informational** | 6           | **Standard operational logs.** General ""heartbeat"" of the app.                |              
| `LogDebug()`     | **Debug**         | 7           | **Developer verbosity.** Detailed info for troubleshooting.                     |                             


### Implementation Mapping
In your JSON output, these methods automatically map to the `level` key. This allows your log aggregator
(Datadog, Grafana, etc.) to apply color-coding and trigger automated alerts based on the numerical priority.
```go
// Severity 1: Trigger SRE page
sylog.LogAlert("vault-service", "Credential rotation failed - Unauthorised access detected")

// Severity 6: Standard audit trail
sylog.LogInfo("user-service", "User ID 505 logged in successfully")
```
>**Architect's Tip:** In a production environment, you should typically set your log sink to ignore `LogDebug` and
`LogInfo` to save on storage costs, while ensuring `LogWarning` and above are always captured.

### The Unopinionated Philosophy
SyLog is designed to be a transparent sidecar to your application logic. Unlike other logging libraries that may
force a `panic()` or `os.Exit()` when a high-severity log is triggered, this package is entirely non-intrusive.

#### Execution Control
Whether you call `LogInfo()` or `LogEmergency()`, the package will:
1. Capture the trace and timestamp.
2. Format the JSON payload.
3. Emit the log to the configured output.
4. Return control to your application.

#### Why this matters
* **Testing:** Your unit tests won't unexpectedly terminate when testing error-handling paths.
* **Graceful Shutdowns:** You can log a Critical failure and still allow your application to close database connections or finish inflight requests before exiting.
* **Separation of Concerns:** We believe a logger’s job is to report, not to decide the lifecycle of your software.

#### Example: Handling an Emergency
```go
if err != nil {
    // The logger records the disaster...
    sylog.LogEmergency("kernel", "Memory limit reached, unable to continue")
    
    // ...but YOU decide how to crash.
    panic("System Unstable") 
}
```


## Core Concepts: The Power of "Facility"
In systems architecture, clarity is the antidote to complexity. When adapting the Syslog
Facility concept for modern microservices, we are essentially moving from "What happened?"
to "Who is responsible?"

In traditional Unix, facilities were predefined (like `auth`, `cron`, or `kern`). 
In SyLog package, the Facility becomes a dynamic tag representing the Service Identity.
Unlike traditional loggers that only categorise by severity (Level), SyLog introduces 
the Facility parameter as a first-class citizen.

### What is a Facility?
In this package, a Facility represents the Identity of the Component. While the level 
tells you the urgency of the message, the `facility` tells you exactly which microservice 
or sub-system generated it.

### Why this matters for Observability
By explicitly tagging logs with a facility, you provide high-cardinality metadata that 
modern observability tools (CloudWatch, Prometheus, Datadog) use for indexing and filtering.

| Advantage              | Benefit                                                                                                                  |
|------------------------|--------------------------------------------------------------------------------------------------------------------------|
| **Granular Filtering** | Instantly isolate all logs from the `billing-engine` without parsing complex message strings.                            |
| **Routing Logic**      | Route logs from the `auth` facility to a high-security storage bucket while sending `frontend` logs to standard storage. |
| **Alerting Accuracy**  | Set up monitors that trigger only when `alert` level logs appear specifically within the `database-proxy` facility.      |

### How to choose a Facility name
For the best results in distributed environments, we recommend following a consistent naming 
convention for your facilities:

 * **Service Name:** e.g., payment-microservice
 * **Sub-system:** e.g., worker-pool or cache-layer
 * **Environment (Optional):** e.g., prod-api-gateway

### Implementation Example
Instead of generic logging:
```go
// Vague: Where did this happen?
sylog.LogError("Database connection timed out")
```

Use the **Facility-first** approach:
```go
// Precise: The 'order-processor' is the culprit.
sylog.LogError("order-processor", "Database connection timed out")
```
When this reaches **CloudWatch** or **Datadog**, you can simply run a query like facility: 
"order-processor" to see the entire lifecycle of that specific service's behavior.

## Performance Benchmarks
The document that tracks the performance of the logger package is published in details 
at the root of repository [`BENCHMARKS.md`](https://github.com/syniol/go-logger/blob/main/BENCHMARKS.md).


| Benchmark Case            | Iterations | Memory Allocations | Bytes Allocations | Operation Allocations |
|---------------------------|------------|--------------------|-------------------|-----------------------|
| BenchmarkSyniolLogger-8   | 	 1891555	 | 611.1 ns/op	       | 664 B/op	         | 9 allocs/op           |
| BenchmarkSlogJSON-8       | 	 2352092	 | 510.8 ns/op	       | 0 B/op	           | 0 allocs/op           |
| BenchmarkSlogWithSource-8 | 	 1000000	 | 1053 ns/op	        | 584 B/op	         | 6 allocs/op           |


#### Credits
Author: [Hadi Tajallaei](mailto:hadi@syniol.com)

Copyright &copy; 2023-2026 Syniol Limited. All rights Reserved.
