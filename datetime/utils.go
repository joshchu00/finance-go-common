package datetime

import (
	"time"
)

func GetMilliSecond(t time.Time) int64 {
	return t.UnixNano() / 1e6
}
