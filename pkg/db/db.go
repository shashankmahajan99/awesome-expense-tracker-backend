package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	migrate "github.com/golang-migrate/migrate/v4"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
			MultiStatements:      true,
		})
	}

	// connect to db
	db, err := connect(dbConfig[0])
	if err != nil {
		return nil, err
	}

	// migrate db
	err = migrateDb(db)
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
func migrateDb(db *DB) error {
	// Replace the path with the path to your migrations directory
	driver, err := migrate_mysql.WithInstance(db.DB, &migrate_mysql.Config{
		DatabaseName: "awesome_expense_tracker",
	})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/db/migration",
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	if err == migrate.ErrNoChange {
		log.Println("no changes in the database schema")
	} else {
		log.Println("migrated db")
	}

	return nil
}
