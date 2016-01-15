package db

import (
	"fmt"
	"time"
)

type Entity interface {
	GetId() *int64
	SetId(int64)
}

type Cursor struct {
	//
}

type SentTransaction struct {
	Id            *int64    `db:"id"`
	Status        string    `db:"status"` // pending/success/failure
	Source        string    `db:"source"`
	SubmittedAt   time.Time `db:"submitted_at"`
	SucceededAt   time.Time `db:"succeeded_at"`
	OperationType string    `db:"operation_type"`
	Ledger        *uint64   `db:"ledger"`
	EnvelopeXdr   string    `db:"enveloper_xdr"`
	ResultXdr     *string   `db:"result_xdr"`
}

func (st *SentTransaction) GetId() *int64 {
	return st.Id
}

func (st *SentTransaction) SetId(id int64) {
	st.Id = &id
}

func (st *SentTransaction) MarkSucceeded() {
	st.SucceededAt = time.Now()
}

func GetInsertQuery(objectType string) (query string, err error) {
	switch objectType {
	case "*db.SentTransaction":
		query = `
		INSERT INTO SentTransaction
			(status,source,submitted_at,succeeded_at,operation_type,enveloper_xdr)
		VALUES
			(:status,:source,:submitted_at,:succeeded_at,:operation_type,:enveloper_xdr)`
	default:
		err = fmt.Errorf("No INSERT query for: %s (must be a pointer)", objectType)
	}
	return
}

func GetUpdateQuery(objectType string) (query string, err error) {
	switch objectType {
	case "*db.SentTransaction":
		query = `
		UPDATE SentTransaction SET
			status = :status,
			source = :source,
			submitted_at = :submitted_at,
			succeeded_at = :succeeded_at,
			operation_type = :operation_type,
			enveloper_xdr = :enveloper_xdr
		WHERE
			id = :id
		`
	default:
		err = fmt.Errorf("No UPDATE query for: %s (must be a pointer)", objectType)
	}
	return
}
