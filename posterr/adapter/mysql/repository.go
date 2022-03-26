package mysql

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"strider-backend-test.com/log"
)

type Repository interface {
	BeginTransaction(ctx context.Context) (tx Transaction, err error)
	Commit(ctx context.Context, tx Transaction)
	Rollback(ctx context.Context, tx Transaction)
	QueryContext(ctx context.Context, query string, args ...interface{}) (row *sql.Rows, err error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type mySQL struct {
	db     *sql.DB
	logger *logrus.Entry
}

func New(db *sql.DB) Repository {
	return &mySQL{db: db, logger: log.GetLogger()}
}

func (m *mySQL) BeginTransaction(ctx context.Context) (tx Transaction, err error) {
	return m.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (m *mySQL) Commit(ctx context.Context, tx Transaction) {
	if err := tx.Commit(); err != nil {
		m.logger.Errorf("error while commiting transaction: %s", err.Error())
	}
}

func (m *mySQL) Rollback(ctx context.Context, tx Transaction) {
	if err := tx.Rollback(); err != nil {
		m.logger.Errorf("error while rolling back transaction: %s", err.Error())
	}
}

func (m *mySQL) QueryContext(ctx context.Context, query string, args ...interface{}) (row *sql.Rows, err error) {
	return m.db.QueryContext(ctx, query, args...)
}

func (m *mySQL) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return m.db.QueryRowContext(ctx, query, args...)
}
