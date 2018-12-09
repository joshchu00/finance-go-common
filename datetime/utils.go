package datetime

import (
	"time"
)

func GetTimestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func getTime(ts int64, loc *time.Location) time.Time {
	return time.Unix(ts/1e3, ts%1e3*1e6).In(loc)
}

func GetTimeString(ts int64, loc *time.Location) string {
	return getTime(ts, loc).String()
}

func GetDateString(ts int64, loc *time.Location) string {
	return getTime(ts, loc).Format("20060102")
}
