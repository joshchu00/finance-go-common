package strategy

import (
	"github.com/joshchu00/finance-go-common/cassandra"
)

var (
	SSMA *Strategy = &Strategy{
		Column: cassandra.StrategyColumnSSMA,
	}
)

type SSMAValue string

const (
	SSMANIL  SSMAValue = ""
	SSMABUY  SSMAValue = "buy"
	SSMASELL SSMAValue = "sell"
)

type SSMAInput struct {
	SMA0005 float64
	SMA0010 float64
	SMA0020 float64
}

func CalculateSSMA(in []*SSMAInput) (out []SSMAValue) {

	out = make([]SSMAValue, 0)

	currentStatus := SSMANIL

	for i, n := range in {

		status := SSMASELL

		if (n.SMA0005 >= n.SMA0010) && (n.SMA0010 >= n.SMA0020) {
			status = SSMABUY
		}

		value := SSMANIL

		if i != 0 && currentStatus != status {
			value = status
		}

		out = append(out, value)

		currentStatus = status
	}

	return
}
