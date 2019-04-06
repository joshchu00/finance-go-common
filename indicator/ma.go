package indicator

import (
	"github.com/joshchu00/finance-go-common/cassandra"
)

var (
	SMA0060 *Indicator = &Indicator{
		Type:   SMA,
		Column: cassandra.IndicatorColumnSMA0060,
		Value:  60,
	}
	SMA0120 *Indicator = &Indicator{
		Type:   SMA,
		Column: cassandra.IndicatorColumnSMA0120,
		Value:  120,
	}
	SMA0240 *Indicator = &Indicator{
		Type:   SMA,
		Column: cassandra.IndicatorColumnSMA0240,
		Value:  240,
	}
)
