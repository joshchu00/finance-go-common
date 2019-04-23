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
	Name string
	SSMA string
	LSMA string
}

type StrategyColumn string

const (
	StrategyColumnUnknown StrategyColumn = "unknown"
	StrategyColumnName    StrategyColumn = "name"
	StrategyColumnSSMA    StrategyColumn = "ssma"
	StrategyColumnLSMA    StrategyColumn = "lsma"
)

func (c *Client) InsertStrategyRow(s *StrategyRow) (err error) {

	cql := "INSERT INTO strategy (exchange, symbol, period, datetime, name, ssma, lsma) VALUES (?, ?, ?, ?, ?, ?, ?)"

	err = c.session.Query(
		cql,
		s.Exchange,
		s.Symbol,
		s.Period,
		s.Datetime,
		s.Name,
		s.SSMA,
		s.LSMA,
	).Exec()

	return
}

func (c *Client) InsertStrategyRowName(spmk *StrategyPrimaryKey, name string) (err error) {

	cql := "INSERT INTO strategy (exchange, symbol, period, datetime, name) VALUES (?, ?, ?, ?, ?)"

	err = c.session.Query(
		cql,
		spmk.Exchange,
		spmk.Symbol,
		spmk.Period,
		spmk.Datetime,
		name,
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

	cql := "SELECT exchange, symbol, period, datetime, name, ssma, lsma FROM strategy WHERE exchange = ? AND symbol = ? AND period = ? AND datetime = ? LIMIT 1"

	query := c.session.Query(
		cql,
		spmk.Exchange,
		spmk.Symbol,
		spmk.Period,
		spmk.Datetime,
	)

	var exchange, symbol, period string
	var datetime time.Time
	var name string
	var ssma, lsma string

	err = query.Scan(&exchange, &symbol, &period, &datetime, &name, &ssma, &lsma)
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
		Name: name,
		SSMA: ssma,
		LSMA: lsma,
	}

	return
}

func (c *Client) SelectStrategyRowsByPartitionKey(sptk *StrategyPartitionKey) (srs []*StrategyRow, err error) {

	cql := "SELECT exchange, symbol, period, datetime, name, ssma, lsma FROM strategy WHERE exchange = ? AND symbol = ? AND period = ? ORDER BY datetime ASC"

	iter := c.session.Query(
		cql,
		sptk.Exchange,
		sptk.Symbol,
		sptk.Period,
	).Iter()

	var exchange, symbol, period string
	var datetime time.Time
	var name string
	var ssma, lsma string

	srs = make([]*StrategyRow, 0)

	for iter.Scan(&exchange, &symbol, &period, &datetime, &name, &ssma, &lsma) {
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
				Name: name,
				SSMA: ssma,
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
