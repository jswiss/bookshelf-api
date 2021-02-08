package database

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}
type BorrowBookTxParams struct {
	Book   int32 `json:"book"`
	Friend int32 `json:"friend"`
}

type BorrowBookTxResult struct {
	BorrowedBook BorrowedBook `json:borrowed_book`
	Book         int32        `json:"book"`
	Friend       int32        `json:"friend"`
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (store *Store) BorrowBookTx(ctx context.Context, arg BorrowBookTxParams) (BorrowBookTxResult, error) {
	var result BorrowBookTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.BorrowedBook, err = q.CreateBorrowedBook(ctx, CreateBorrowedBookParams{
			Book:   arg.Book,
			Friend: arg.Friend,
		})
		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}
