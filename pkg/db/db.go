package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

// InitDB initializes the database
func InitDB(dbConfig ...*mysql.Config) (*DB, error) {
	// dbConfig is optional
	if len(dbConfig) == 0 {
		dbConfig = append(dbConfig, &mysql.Config{
			User:                 "root",
			Passwd:               "password",
			Net:                  "tcp",
			Addr:                 "0.0.0.0:3306",
			DBName:               "awesome_expense_tracker",
			AllowNativePasswords: true,
		})
	}

	// connect to db
	db, err := connect(dbConfig[0])
	if err != nil {
		return nil, err
	}

	// migrate db
	err = migrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// NewDB creates a new database
func NewDB(driverName, dataSourceName string) (*DB, error) {
	// connect to db
	db, err := Open(
		driverName,
		dataSourceName,
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// DB is a wrapper around sql.DB
type DB struct {
	*sql.DB
}

// Open opens a new database connection
func Open(driverName, dataSourceName string) (*DB, error) {

	// open db
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	// wrap db
	wrappedDB := &DB{db}

	return wrappedDB, nil
}

// Ping pings the database
func (db *DB) Ping() error {
	return db.DB.Ping()
}

// Close closes the database
func (db *DB) Close() error {
	return db.DB.Close()
}

// connect connects to the database
func connect(dbConfig *mysql.Config) (*DB, error) {
	// connect to db
	db, err := NewDB(
		"mysql",
		dbConfig.FormatDSN(),
	)
	if err != nil {
		return nil, err
	}
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			log.Println("connected to db")
			return db, nil
		}
		log.Printf("Error: %v", err)
		log.Printf("failed to connect to db, retrying in 1 second (attempt %d)", i+1)
		time.Sleep(time.Second)
	}

	return nil, fmt.Errorf("failed to connect to db")
}

// migrate migrates the database
func migrate(db *DB) error {
	// create users table
	err := createUserTable(db)
	if err != nil {
		return err
	}

	// create expenses table
	err = createExpenseTable(db)
	if err != nil {
		return err
	}

	log.Println("migrated db")

	return nil
}

// createUserTable creates the users table
func createUserTable(db *DB) error {
	// create users table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	)`)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	return nil
}

// createExpenseTable creates the expenses table
func createExpenseTable(db *DB) error {
	// create expenses table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		amount NUMERIC(10, 2) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	)`)
	if err != nil {
		return fmt.Errorf("failed to create expenses table: %v", err)
	}

	return nil
}
