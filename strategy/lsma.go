package strategy

import (
	"github.com/joshchu00/finance-go-common/cassandra"
)

var (
	LSMA *Strategy = &Strategy{
		Column: cassandra.StrategyColumnLSMA,
	}
)

type LSMAValue string

const (
	LSMANIL  LSMAValue = ""
	LSMABUY  LSMAValue = "buy"
	LSMASELL LSMAValue = "sell"
)

type LSMAInput struct {
	SMA0060 float64
	SMA0120 float64
	SMA0240 float64
}

func CalculateLSMA(in []*LSMAInput) (out []LSMAValue) {

	out = make([]LSMAValue, 0)

	currentStatus := LSMANIL

	for i, n := range in {

		status := LSMASELL

		if (n.SMA0060 >= n.SMA0120) && (n.SMA0120 >= n.SMA0240) {
			status = LSMABUY
		}

		value := LSMANIL

		if i != 0 && currentStatus != status {
			value = status
		}

		out = append(out, value)

		currentStatus = status
	}

	return
}
