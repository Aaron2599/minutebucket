# MinuteBucket: Precise Go Time Conversion to Minute Intervals

The `minutebucket` Go package provides robust utilities for seamlessly converting Go's standard `time.Time` objects into discrete minute intervals. It simplifies time management by representing each minute as a unique `int64` value, signifying the total number of minutes elapsed since the Unix epoch (January 1, 1970 UTC). This technique is ideal for time bucketing and efficient time-series data handling in Go applications.

## **Optimize Your Time-Based Operations:**

* **Accurate Time Conversion:** Convert any `time.Time` into its corresponding "minute bucket" integer (`int64`).
* **Current Time Bucket:** Easily fetch the current minute bucket using `minutebucket.Now()`.
* **Reverse Conversion:** Transform a minute bucket `int64` back into standard Unix time formats (seconds, milliseconds, microseconds, nanoseconds).

## **Why Use MinuteBucket? Common Applications:**

* **Efficient Database Partitioning:** Create optimized, time-based partition keys for databases (e.g., Cassandra, InfluxDB, TimescaleDB) to group data by minute, enhancing query performance.
* **Consistent Time Bucketing:** Generate uniform, minute-based identifiers or keys for caching, logging, or distributed systems.
* **Time-Series Analysis:** Simplify the aggregation and analysis of data within distinct one-minute intervals. Ideal for monitoring, analytics, and event processing.
* **Go Time Management:** A fundamental tool for developers needing precise control over time granularity in Go projects.

## Installation
```bash
go get github.com/Aaron2599/minutebucket
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/Aaron2599/minutebucket"
	"time"
)

func main() {
	var currentMinute = minutebucket.Now()
	fmt.Printf("Current minuteBucket Epoch: %d\n", currentMinute)

	fmt.Println("---")

	var someDate = time.Date(2025, time.April, 15, 9, 45, 30, 0, time.UTC)
	var bucket = minutebucket.ToMinuteBucket(someDate)
	fmt.Printf("minuteBucket Epoch for %s: %d\n", someDate.Format(time.RFC3339), bucket)

	fmt.Println("--- Converting Current minuteBucket Epoch Back ---")

	var Unix = minutebucket.ToUnix(currentMinute)
	fmt.Printf("%d -> Unix Seconds: %d\n", currentMinute, Unix)

	var UnixMilli = minutebucket.ToUnixMilli(currentMinute)
	fmt.Printf("%d -> Unix Millis: %d\n", currentMinute, UnixMilli)

	var UnixMicro = minutebucket.ToUnixMicro(currentMinute)
	fmt.Printf("%d -> Unix Micros: %d\n", currentMinute, UnixMicro)

	var UnixNano = minutebucket.ToUnixNano(currentMinute)
	fmt.Printf("%d -> Unix Nanos:  %d\n", currentMinute, UnixNano)
}
```


