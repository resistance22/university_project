package db

import (
	"github.com/jackc/pgx/v5"
)

type Store struct {
	Queries *Queries
	db      *pgx.Conn
}

func NewPGXStore(db *pgx.Conn) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// func (s *Store) execTX(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbError := tx.Rollback(ctx); rbError != nil {
// 			return fmt.Errorf("tx error: %v, rb error %v", err, rbError)
// 		}
// 		return err
// 	}
// 	return tx.Commit(ctx)
// }
