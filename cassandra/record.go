package cassandra

import (
	"time"

	inf "gopkg.in/inf.v0"
)

type Record struct {
	Exchange string
	Symbol   string
	Period   string
	Datetime time.Time
	Name     string
	Open     *inf.Dec
	High     *inf.Dec
	Low      *inf.Dec
	Close    *inf.Dec
	Volume   int64
}

func (c *Client) InsertRecord(r *Record) (err error) {

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

func (c *Client) SelectAllRecord(exchange string, symbol string, period string) (rs []*Record, err error) {

	cql := "SELECT exchange, symbol, period, datetime, name, open, high, low, close, volume FROM record WHERE exchange = ? AND symbol = ? AND period = ? ORDER BY datetime ASC"

	iter := c.session.Query(
		cql,
		exchange,
		symbol,
		period,
	).Iter()

	var datetime time.Time
	var name string
	var open, high, low, close *inf.Dec
	var volume int64

	rs = make([]*Record, 0)

	for iter.Scan(&exchange, &symbol, &period, &datetime, &name, &open, &high, &low, &close, &volume) {
		rs = append(
			rs,
			&Record{
				Exchange: exchange,
				Symbol:   symbol,
				Period:   period,
				Datetime: datetime,
				Name:     name,
				Open:     open,
				High:     high,
				Low:      low,
				Close:    close,
				Volume:   volume,
			},
		)
	}

	err = iter.Close()
	if err != nil {
		return
	}

	return
}
