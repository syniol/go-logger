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
of severity levels, which provided the first standardized framework for categorizing log entries based on their 
impact or urgency.

The following are the levels defined by Syslog in descending order of severity:

 * Emergency `(emerg/fatal)`: _indicates that the system is unusable and requires immediate attention._
 * Alert `(alert)`: _indicates that immediate action is necessary to resolve a critical issue._
 * Critical `(crit)`: _signifies critical conditions in the program that demand intervention to prevent system failure._
 * Error `(error)`: _indicates error conditions that impair some operation but are less severe than critical situations._
 * Warning `(warn)`: _signifies potential issues that may lead to errors or unexpected behavior in the future if not addressed._
 * Notice `(notice)`: _applies to normal but significant conditions that may require monitoring._
 * Informational `(info)`: _includes messages that provide a record of the normal operation of the system._
 * Debug `(debug)`: _intended for logging detailed information about the system for debugging purposes._


#### Credits
Copyright &copy; 2023 Syniol Limited. All rights Reserved.
