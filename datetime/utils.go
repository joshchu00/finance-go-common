package datetime

import (
	"time"
)

func GetTimestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func getTime(ts int64) time.Time {
	return time.Unix(ts/1e3, ts%1e3*1e6)
}

func GetTimestampString(ts int64) string {
	return getTime(ts).String()
}

func GetDateString(ts int64) string {
	return getTime(ts).Format("20060102")
}
