package cassandra

import (
	"fmt"
	"time"
)

type StrategyPartitionKey struct {
	Exchange string
	Symbol   string
	Period   string
}

type StrategyPrimaryKey struct {
	StrategyPartitionKey
	Datetime time.Time
}

type StrategyRow struct {
	StrategyPrimaryKey
	LSMA string
}

type StrategyColumn string

const (
	StrategyColumnUnknown StrategyColumn = "unknown"
	StrategyColumnLSMA    StrategyColumn = "lsma"
)

func (c *Client) InsertStrategyRow(s *StrategyRow) (err error) {

	cql := "INSERT INTO strategy (exchange, symbol, period, datetime, lsma) VALUES (?, ?, ?, ?, ?)"

	err = c.session.Query(
		cql,
		s.Exchange,
		s.Symbol,
		s.Period,
		s.Datetime,
		s.LSMA,
	).Exec()

	return
}

func (c *Client) InsertStrategyRowStringColumn(spmk *StrategyPrimaryKey, column StrategyColumn, value string) (err error) {

	cql := fmt.Sprintf("INSERT INTO strategy (exchange, symbol, period, datetime, %s) VALUES (?, ?, ?, ?, ?)", column)

	err = c.session.Query(
		cql,
		spmk.Exchange,
		spmk.Symbol,
		spmk.Period,
		spmk.Datetime,
		value,
	).Exec()

	return
}

func (c *Client) SelectStrategyRowByPrimaryKey(spmk *StrategyPrimaryKey) (sr *StrategyRow, err error) {

	cql := "SELECT exchange, symbol, period, datetime, lsma FROM strategy WHERE exchange = ? AND symbol = ? AND period = ? AND datetime = ? LIMIT 1"

	query := c.session.Query(
		cql,
		spmk.Exchange,
		spmk.Symbol,
		spmk.Period,
		spmk.Datetime,
	)

	var exchange, symbol, period string
	var datetime time.Time
	var lsma string

	err = query.Scan(&exchange, &symbol, &period, &datetime, &lsma)
	if err != nil {
		return
	}

	sr = &StrategyRow{
		StrategyPrimaryKey: StrategyPrimaryKey{
			StrategyPartitionKey: StrategyPartitionKey{
				Exchange: exchange,
				Symbol:   symbol,
				Period:   period,
			},
			Datetime: datetime,
		},
		LSMA: lsma,
	}

	return
}

func (c *Client) SelectStrategyRowsByPartitionKey(sptk *StrategyPartitionKey) (srs []*StrategyRow, err error) {

	cql := "SELECT exchange, symbol, period, datetime, lsma FROM strategy WHERE exchange = ? AND symbol = ? AND period = ? ORDER BY datetime ASC"

	iter := c.session.Query(
		cql,
		sptk.Exchange,
		sptk.Symbol,
		sptk.Period,
	).Iter()

	var exchange, symbol, period string
	var datetime time.Time
	var lsma string

	srs = make([]*StrategyRow, 0)

	for iter.Scan(&exchange, &symbol, &period, &datetime, &lsma) {
		srs = append(
			srs,
			&StrategyRow{
				StrategyPrimaryKey: StrategyPrimaryKey{
					StrategyPartitionKey: StrategyPartitionKey{
						Exchange: exchange,
						Symbol:   symbol,
						Period:   period,
					},
					Datetime: datetime,
				},
				LSMA: lsma,
			},
		)
	}

	err = iter.Close()
	if err != nil {
		return
	}

	return
}
