package graph

import (
	"errors"

	"github.com/jmcvetta/neoism"
)

type Tx struct {
	client *Client
	tx     *neoism.Tx
	Errors *[]neoism.TxError
}

// Execute a single query.
func (t *Tx) Execute(q *neoism.CypherQuery) error {
	qs := []*neoism.CypherQuery{q}
	return t.ExecuteMany(qs)
}

func (t *Tx) ExecuteMany(qs []*neoism.CypherQuery) error {
	if t.tx == nil {
		tx, err := t.client.db.Begin(qs)
		t.tx = tx
		t.Errors = &t.tx.Errors

		// Since statements are pre-defined, none of them should cause
		// unexpected errors
		if len(t.tx.Errors) > 0 {
			return errors.New(t.tx.Errors[0].Message)
		}

		return err
	}

	return t.tx.Query(qs)
}

// Rollback the transaction if open.
func (t *Tx) Rollback() error {
	if t.tx != nil {
		return t.tx.Rollback()
	}

	return nil
}

// Commit the transaction if open.
func (t *Tx) Commit() error {
	if t.tx != nil {
		return t.tx.Commit()
	}

	return nil
}
