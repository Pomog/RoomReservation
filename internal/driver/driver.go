package driver

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// DB holds the database connection pool.
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const (
	maxOpenDBConn = 10
	maxIdleDBConn = 5
	maxDBlifetime = 5 * time.Minute
)

// ConnectSQL creates a connection pool for PostgteSQL database.
func ConnectSQL(dsn string) (*DB, error) {
	db, err := NewDatabese(dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBlifetime)

	// check connection
	if err = testDB(db); err != nil {
		return nil, err
	}

	dbConn.SQL = db

	return dbConn, nil
}

// NewDatabese creates a new database connection.
func NewDatabese(dsn string) (*sql.DB, error) {
	// connect to database
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")
	return conn, nil
}

// testDB tries to ping database to check if connection is alive.
func testDB(db *sql.DB) error {
	// check connection
	if err := db.Ping(); err != nil {
		return err
	}
	log.Println("Connected to database")
	return nil
}
