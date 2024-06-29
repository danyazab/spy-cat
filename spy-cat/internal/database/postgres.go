package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // <------------ here
	"spy-cat/internal/configs"
)

// PostgSQLDB represents the implementation of the PostgreSQL database
type PostgreSQLDB struct {
	db *sql.DB
}

// NewPostgreSQLDB creates a new instance of the PostgreSQL database
func NewPostgreSQLDB(cfg *configs.Config) (*PostgreSQLDB, error) {

	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName))
	if err != nil {
		return nil, err
	}

	return &PostgreSQLDB{db: db}, nil
}

// Query executes a query against the database and returns the resulting strings
func (p *PostgreSQLDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return p.db.Query(query, args...)
}

// QueryRow method implementation for PostgreSQLDB
func (p *PostgreSQLDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return p.db.QueryRow(query, args...)
}

// Exec executes a database query that does not return a result
func (p *PostgreSQLDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return p.db.Exec(query, args...)
}

// Prepare prepares a query to the database for execution and returns a *sql.Stmt object
func (p *PostgreSQLDB) Prepare(query string) (*sql.Stmt, error) {
	return p.db.Prepare(query)
}

// Close closes the connection to the database
func (p *PostgreSQLDB) Close() error {
	return p.db.Close()
}
