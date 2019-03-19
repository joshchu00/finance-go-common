package cassandra

type StrategyPrimaryKey struct {
	Name string
}

type StrategyRow struct {
	StrategyPrimaryKey
	Description string
}

func (c *Client) InsertStrategyRow(s *StrategyRow) (err error) {

	cql := "INSERT INTO strategy (name, description) VALUES (?, ?)"

	err = c.session.Query(
		cql,
		s.Name,
		s.Description,
	).Exec()

	return
}

func (c *Client) SelectStrategyRows() (srs []*StrategyRow, err error) {

	cql := "SELECT name, description FROM strategy"

	iter := c.session.Query(
		cql,
	).Iter()

	var name, description string

	srs = make([]*StrategyRow, 0)

	for iter.Scan(&name, &description) {
		srs = append(
			srs,
			&StrategyRow{
				StrategyPrimaryKey: StrategyPrimaryKey{
					Name: name,
				},
				Description: description,
			},
		)
	}

	err = iter.Close()
	if err != nil {
		return
	}

	return
}
