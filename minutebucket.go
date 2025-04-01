package minutebucket

import (
	"time"
)

// TODO
// minutesPerBucket defines how many minutes each bucket is created by
var minutesPerBucket int64 = 1

// secondsPerMinute defines the number of seconds in a minute.
var secondsPerMinute int64 = 60

// millisPerMinute defines the number of milliseconds in a minute.
var millisPerMinute int64 = 60000

// microsPerMinute defines the number of microseconds in a minute.
var microsPerMinute int64 = 60000000

// nanosPerMinute defines the number of nanoseconds in a minute.
var nanosPerMinute int64 = 60000000000

// Now calculates the current minute bucket based on the current UTC time.
// It returns the Unix timestamp divided by the number of seconds in a minute.
func Now() int64 {
	return time.Now().UTC().Unix() / (secondsPerMinute * minutesPerBucket)
}

// ToMinuteBucket converts a given time.Time object to its corresponding minute bucket.
// It calculates the bucket by dividing the time's UTC Unix timestamp by the number of seconds in a minute.
//
// Args:
//
//	t (time.Time): The time to convert.
//
// Returns:
//
//	(int64): The minute bucket representation of the time.
func ToMinuteBucket(t time.Time) int64 {
	return t.UTC().Unix() / (secondsPerMinute * minutesPerBucket)
}

// ToUnix converts a minute bucket value back to a Unix timestamp (in seconds).
// It represents the start time of that minute bucket.
//
// Args:
//
//	minuteBucket (int64): The minute bucket value.
//
// Returns:
//
//	(int64): The corresponding Unix timestamp in seconds.
func ToUnix(minuteBucket int64) int64 {
	return minuteBucket * secondsPerMinute
}

// ToUnixMilli converts a minute bucket value back to a Unix timestamp in milliseconds.
// It represents the start time of that minute bucket.
//
// Args:
//
//	minuteBucket (int64): The minute bucket value.
//
// Returns:
//
//	(int64): The corresponding Unix timestamp in milliseconds.
func ToUnixMilli(minuteBucket int64) int64 {
	return minuteBucket * millisPerMinute
}

// ToUnixMicro converts a minute bucket value back to a Unix timestamp in microseconds.
// It represents the start time of that minute bucket.
//
// Args:
//
//	minuteBucket (int64): The minute bucket value.
//
// Returns:
//
//	(int64): The corresponding Unix timestamp in microseconds.
func ToUnixMicro(minuteBucket int64) int64 {
	return minuteBucket * microsPerMinute
}

// ToUnixNano converts a minute bucket value back to a Unix timestamp in nanoseconds.
// It represents the start time of that minute bucket.
//
// Args:
//
//	minuteBucket (int64): The minute bucket value.
//
// Returns:
//
//	(int64): The corresponding Unix timestamp in nanoseconds.
func ToUnixNano(minuteBucket int64) int64 {
	return minuteBucket * nanosPerMinute
}
