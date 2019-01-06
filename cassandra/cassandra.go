package cassandra

import (
	"github.com/gocql/gocql"
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

func (c *Client) Close() {
	c.session.Close()
}
