package cassandra

import (
	"time"

	"github.com/gocql/gocql"
	inf "gopkg.in/inf.v0"
)

type Client struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

func NewClient(hosts string, keyspace string) (c *Client, err error) {

	c = &Client{}

	c.cluster = gocql.NewCluster(hosts)
	c.cluster.Keyspace = keyspace
	c.cluster.Consistency = gocql.Quorum

	c.session, err = c.cluster.CreateSession()

	return
}

func (c *Client) InsertRecord(
	exchange string,
	symbol string,
	period string,
	datetime time.Time,
	name string,
	open *inf.Dec,
	high *inf.Dec,
	low *inf.Dec,
	close *inf.Dec,
	volume int64,
) (err error) {

	r := &record{
		exchange,
		symbol,
		period,
		datetime,
		name,
		open,
		high,
		low,
		close,
		volume,
	}

	err = r.insert(c.session)

	return
}

func (c *Client) Close() {
	c.session.Close()
}
