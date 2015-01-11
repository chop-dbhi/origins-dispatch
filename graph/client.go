package graph

import "github.com/jmcvetta/neoism"

type Client struct {
	Uri string
	db  *neoism.Database
}

func (c *Client) connect() error {
	if c.db == nil {
		db, err := neoism.Connect(c.Uri)

		if err != nil {
			return err
		}

		c.db = db
	}

	return nil
}

func (c *Client) Tx() (*Tx, error) {
	err := c.connect()

	if err != nil {
		return nil, err
	}

	return &Tx{
		client: c,
	}, nil
}

func (c *Client) ExecuteMany(qs []*neoism.CypherQuery) error {
	tx, err := c.Tx()

	if err != nil {
		return err
	}

	if err = tx.ExecuteMany(qs); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (c *Client) Execute(q *neoism.CypherQuery) error {
	qs := []*neoism.CypherQuery{q}

	return c.ExecuteMany(qs)
}

func (c *Client) ExecuteStatement(s string) error {
	q := neoism.CypherQuery{
		Statement: s,
	}

	return c.Execute(&q)
}
