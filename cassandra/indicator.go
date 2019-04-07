package cassandra

import (
	"fmt"
	"time"

	inf "gopkg.in/inf.v0"
)

type IndicatorPartitionKey struct {
	Exchange string
	Symbol   string
	Period   string
}

type IndicatorPrimaryKey struct {
	IndicatorPartitionKey
	Datetime time.Time
}

type IndicatorRow struct {
	IndicatorPrimaryKey
	SMA0060 *inf.Dec
	SMA0120 *inf.Dec
	SMA0240 *inf.Dec
}

type IndicatorColumn string

const (
	IndicatorColumnUnknown IndicatorColumn = "unknown"
	IndicatorColumnSMA0060 IndicatorColumn = "sma0060"
	IndicatorColumnSMA0120 IndicatorColumn = "sma0120"
	IndicatorColumnSMA0240 IndicatorColumn = "sma0240"
)

func (c *Client) InsertIndicatorRow(i *IndicatorRow) (err error) {

	cql := "INSERT INTO indicator (exchange, symbol, period, datetime, sma0060, sma0120, sma0240) VALUES (?, ?, ?, ?, ?, ?, ?)"

	err = c.session.Query(
		cql,
		i.Exchange,
		i.Symbol,
		i.Period,
		i.Datetime,
		i.SMA0060,
		i.SMA0120,
		i.SMA0240,
	).Exec()

	return
}

func (c *Client) InsertIndicatorRowDecimalColumn(ipmk *IndicatorPrimaryKey, column IndicatorColumn, value *inf.Dec) (err error) {

	cql := fmt.Sprintf("INSERT INTO indicator (exchange, symbol, period, datetime, %s) VALUES (?, ?, ?, ?, ?)", column)

	err = c.session.Query(
		cql,
		ipmk.Exchange,
		ipmk.Symbol,
		ipmk.Period,
		ipmk.Datetime,
		value,
	).Exec()

	return
}

func (c *Client) SelectIndicatorRowsByPartitionKey(iptk *IndicatorPartitionKey) (irs []*IndicatorRow, err error) {

	cql := "SELECT exchange, symbol, period, datetime, sma0060, sma0120, sma0240 FROM indicator WHERE exchange = ? AND symbol = ? AND period = ? ORDER BY datetime ASC"

	iter := c.session.Query(
		cql,
		iptk.Exchange,
		iptk.Symbol,
		iptk.Period,
	).Iter()

	var exchange, symbol, period string
	var datetime time.Time
	var sma0060, sma0120, sma0240 *inf.Dec

	irs = make([]*IndicatorRow, 0)

	for iter.Scan(&exchange, &symbol, &period, &datetime, &sma0060, &sma0120, &sma0240) {
		irs = append(
			irs,
			&IndicatorRow{
				IndicatorPrimaryKey: IndicatorPrimaryKey{
					IndicatorPartitionKey: IndicatorPartitionKey{
						Exchange: exchange,
						Symbol:   symbol,
						Period:   period,
					},
					Datetime: datetime,
				},
				SMA0060: sma0060,
				SMA0120: sma0120,
				SMA0240: sma0240,
			},
		)
	}

	err = iter.Close()
	if err != nil {
		return
	}

	return
}
