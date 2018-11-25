package datetime

import (
	"time"
)

func GetMilliSecond(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func GetTime(ms int64) time.Time {
	return time.Unix(ms/1e3, ms%1e3*1e6)
}

func GetDateString(t time.Time) string {
	return t.Format("20060102")
}
