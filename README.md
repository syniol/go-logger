# SyLog
Go Logger package adopted Cloud Native principles in core.


## Usage Example
Please see below for example code.

```go
package main

import "github.com/syniol/go-logger"

func main() {
	sylog.LogAlert("mock message for alert log")
}
```

you could see JSON output below:

```json
{
  "level": "alert",
  "message": "mock message for alert log",
  "trace": [
    "location: file: '/Users/hadi/dev/golang/logger/logger.go' on line: 100"
  ],
  "timestamp": "2023-11-04T21:36:33Z"
}
```

### Log levels
Log levels for software applications have a rich history dating back to the 1980s. One of the earliest and most 
influential logging solutions for Unix systems, [Syslog](https://en.wikipedia.org/wiki/Syslog) , introduced a range 
of severity levels, which provided the first standardised framework for categorising log entries based on their 
impact or urgency.

The following are the levels defined by Syslog in descending order of severity:

 * __Emergency__ `(emerg/fatal)`: _System is unusable and requires immediate attention._
 * __Alert__ `(alert)`: _Immediate action is necessary to resolve a critical issue._
 * __Critical__ `(crit)`: _Critical conditions in the program that demand intervention to prevent system failure._
 * __Error__ `(error)`: _Error conditions that impair some operation but are less severe than critical situations._
 * __Warning__ `(warn)`: _Potential issues that may lead to errors or unexpected behavior in the future if not addressed._
 * __Notice__ `(notice)`: _Applies to normal but significant conditions that may require monitoring._
 * __Informational__ `(info)`: _Includes messages that provide a record of the normal operation of the system._
 * __Debug__ `(debug)`: _Logging detailed information about the system for debugging purposes._


#### Credits
Copyright &copy; 2023 Syniol Limited. All rights Reserved.
