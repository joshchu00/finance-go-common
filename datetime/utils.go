package datetime

import (
	"time"
)

func GetMilliSecond(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func GetDateString(t time.Time) string {
	return t.Format("20060102")
}
