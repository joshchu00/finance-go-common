package indicator

import (
	"github.com/joshchu00/finance-go-common/cassandra"
)

type IndicatorType string

const (
	SMA IndicatorType = "sma"
)

type Indicator struct {
	Type   IndicatorType
	Column cassandra.IndicatorColumn
	Value  int64
}
