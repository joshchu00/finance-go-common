package indicator

import (
	"github.com/joshchu00/finance-go-common/cassandra"
	talib "github.com/markcheno/go-talib"
)

var (
	SMA0005 *Indicator = &Indicator{
		Type:   SMA,
		Column: cassandra.IndicatorColumnSMA0005,
		Period: 5,
	}
	SMA0010 *Indicator = &Indicator{
		Type:   SMA,
		Column: cassandra.IndicatorColumnSMA0010,
		Period: 10,
	}
	SMA0020 *Indicator = &Indicator{
		Type:   SMA,
		Column: cassandra.IndicatorColumnSMA0020,
		Period: 20,
	}
	SMA0060 *Indicator = &Indicator{
		Type:   SMA,
		Column: cassandra.IndicatorColumnSMA0060,
		Period: 60,
	}
	SMA0120 *Indicator = &Indicator{
		Type:   SMA,
		Column: cassandra.IndicatorColumnSMA0120,
		Period: 120,
	}
	SMA0240 *Indicator = &Indicator{
		Type:   SMA,
		Column: cassandra.IndicatorColumnSMA0240,
		Period: 240,
	}
)

func GetSMA(period int64) (idct *Indicator) {

	switch period {
	case 5:
		idct = SMA0005
	case 10:
		idct = SMA0010
	case 20:
		idct = SMA0020
	case 60:
		idct = SMA0060
	case 120:
		idct = SMA0120
	case 240:
		idct = SMA0240
	default:
		idct = &Indicator{
			Type:   SMA,
			Column: cassandra.IndicatorColumnUnknown,
			Period: period,
		}
	}

	return
}

func CalculateSMA(in []float64, period int64) (out []float64) {
	if len(in) < int(period) {
		out = make([]float64, len(in))
	} else {
		out = talib.Ma(in, int(period), talib.SMA)
	}
	return
}
