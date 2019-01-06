package cassandra

import (
	"time"

	inf "gopkg.in/inf.v0"
)

type Record struct {
	exchange string
	symbol   string
	period   string
	datetime time.Time
	name     string
	open     *inf.Dec
	high     *inf.Dec
	low      *inf.Dec
	close    *inf.Dec
	volume   int64
}

func (c *Client) InsertRecord(r *Record) (err error) {

	cql := "INSERT INTO record (exchange, symbol, period, datetime, name, open, high, low, close, volume) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	err = c.session.Query(
		cql,
		r.exchange,
		r.symbol,
		r.period,
		r.datetime,
		r.name,
		r.open,
		r.high,
		r.low,
		r.close,
		r.volume,
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
				exchange: exchange,
				symbol:   symbol,
				period:   period,
				datetime: datetime,
				name:     name,
				open:     open,
				high:     high,
				low:      low,
				close:    close,
				volume:   volume,
			},
		)
	}

	err = iter.Close()
	if err != nil {
		return
	}

	return
}
