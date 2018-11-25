package cassandra

import (
	"time"

	"github.com/gocql/gocql"
	inf "gopkg.in/inf.v0"
)

type record struct {
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

var (
	insertRecordCQL = "INSERT INTO record (exchange, symbol, period, datetime, name, open, high, low, close, volume) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
)

func (r *record) insert(session *gocql.Session) (err error) {

	err = session.Query(
		insertRecordCQL,
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
