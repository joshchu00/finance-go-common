package datetime

import (
	"time"
)

func NowMilliSecond() int64 {
	return time.Now().UnixNano() / 1e6
}
