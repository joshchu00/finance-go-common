package cassandra

import (
	"fmt"
	"time"
)

type AttentionPartitionKey struct {
	Exchange string
	Symbol   string
	Period   string
}

type AttentionPrimaryKey struct {
	AttentionPartitionKey
	Datetime time.Time
}

type AttentionRow struct {
	AttentionPrimaryKey
	LSMA string
}

type AttentionColumn string

const (
	AttentionColumnLSMA AttentionColumn = "lsma"
)

func (c *Client) InsertAttentionRow(a *AttentionRow) (err error) {

	cql := "INSERT INTO attention (exchange, symbol, period, datetime, lsma) VALUES (?, ?, ?, ?, ?)"

	err = c.session.Query(
		cql,
		a.Exchange,
		a.Symbol,
		a.Period,
		a.Datetime,
		a.LSMA,
	).Exec()

	return
}

func (c *Client) InsertAttentionRowTextColumn(apmk *AttentionPrimaryKey, column AttentionColumn, value string) (err error) {

	cql := fmt.Sprintf("INSERT INTO attention (exchange, symbol, period, datetime, %s) VALUES (?, ?, ?, ?, ?)", column)

	err = c.session.Query(
		cql,
		apmk.Exchange,
		apmk.Symbol,
		apmk.Period,
		apmk.Datetime,
		value,
	).Exec()

	return
}

func (c *Client) SelectAttentionRowsByPartitionKey(aptk *AttentionPartitionKey) (ars []*AttentionRow, err error) {

	cql := "SELECT exchange, symbol, period, datetime, lsma FROM attention WHERE exchange = ? AND symbol = ? AND period = ? ORDER BY datetime ASC"

	iter := c.session.Query(
		cql,
		aptk.Exchange,
		aptk.Symbol,
		aptk.Period,
	).Iter()

	var exchange, symbol, period string
	var datetime time.Time
	var lsma string

	ars = make([]*AttentionRow, 0)

	for iter.Scan(&exchange, &symbol, &period, &datetime, &lsma) {
		ars = append(
			ars,
			&AttentionRow{
				AttentionPrimaryKey: AttentionPrimaryKey{
					AttentionPartitionKey: AttentionPartitionKey{
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
