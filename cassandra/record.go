package cassandra

import (
	"time"

	inf "gopkg.in/inf.v0"
)

type RecordPartitionKey struct {
	Exchange string
	Symbol   string
	Period   string
}

type RecordPrimaryKey struct {
	RecordPartitionKey
	Datetime time.Time
}

type RecordRow struct {
	RecordPrimaryKey
	Name   string
	Open   *inf.Dec
	High   *inf.Dec
	Low    *inf.Dec
	Close  *inf.Dec
	Volume int64
}

func (c *Client) InsertRecordRow(r *RecordRow) (err error) {

	cql := "INSERT INTO record (exchange, symbol, period, datetime, name, open, high, low, close, volume) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	err = c.session.Query(
		cql,
		r.Exchange,
		r.Symbol,
		r.Period,
		r.Datetime,
		r.Name,
		r.Open,
		r.High,
		r.Low,
		r.Close,
		r.Volume,
	).Exec()

	return
}

func (c *Client) SelectRecordRowsByPartitionKey(rptk *RecordPartitionKey) (rrs []*RecordRow, err error) {

	cql := "SELECT exchange, symbol, period, datetime, name, open, high, low, close, volume FROM record WHERE exchange = ? AND symbol = ? AND period = ? ORDER BY datetime ASC"

	iter := c.session.Query(
		cql,
		rptk.Exchange,
		rptk.Symbol,
		rptk.Period,
	).Iter()

	var exchange, symbol, period string
	var datetime time.Time
	var name string
	var open, high, low, close *inf.Dec
	var volume int64

	rrs = make([]*RecordRow, 0)

	for iter.Scan(&exchange, &symbol, &period, &datetime, &name, &open, &high, &low, &close, &volume) {
		rrs = append(
			rrs,
			&RecordRow{
				RecordPrimaryKey: RecordPrimaryKey{
					RecordPartitionKey: RecordPartitionKey{
						Exchange: exchange,
						Symbol:   symbol,
						Period:   period,
					},
					Datetime: datetime,
				},
				Name:   name,
				Open:   open,
				High:   high,
				Low:    low,
				Close:  close,
				Volume: volume,
			},
		)
	}

	err = iter.Close()
	if err != nil {
		return
	}

	return
}
