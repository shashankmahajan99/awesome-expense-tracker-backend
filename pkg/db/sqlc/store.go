package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	migrate "github.com/golang-migrate/migrate/v4"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Store interface {
	Querier
	RegisterUser(ctx context.Context, arg CreateUserParams) (RegisterUserResult, error)
	DeleteUser(ctx context.Context, username string) error
	ListUserByUsername(ctx context.Context, username string) (User, error)
}

// MySQLStore is a wrapper around sql.MySQLStore
type MySQLStore struct {
	*sql.DB
	*Queries
}

// InitDB initializes the database
func InitDB(dbConfig ...*mysql.Config) (Store, error) {
	// dbConfig is optional
	if len(dbConfig) == 0 {
		dbConfig = append(dbConfig, &mysql.Config{
			User:                 "root",
			Passwd:               "password",
			Net:                  "tcp",
			Addr:                 "localhost:3306",
			DBName:               "awesome_expense_tracker",
			AllowNativePasswords: true,
			MultiStatements:      true,
			ParseTime:            true,
		})
	}
	if dbUser := os.Getenv("MYSQLUSER"); dbUser != "" {
		dbConfig[0].User = dbUser
	}
	if dbPassword := os.Getenv("MYSQLPASSWORD"); dbPassword != "" {
		dbConfig[0].Passwd = dbPassword
	}
	if dbAddr := os.Getenv("MYSQLADDR"); dbAddr != "" {
		dbConfig[0].Addr = dbAddr
	}
	if dbName := os.Getenv("MYSQLDBNAME"); dbName != "" {
		dbConfig[0].DBName = dbName
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
func NewDB(driverName, dataSourceName string) (*MySQLStore, error) {
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

// Open opens a new database connection
func Open(driverName, dataSourceName string) (*MySQLStore, error) {

	// open db
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	// wrap db
	wrappedDB := &MySQLStore{DB: db, Queries: New(db)}

	return wrappedDB, nil
}

// Ping pings the database
func (db *MySQLStore) Ping() error {
	return db.DB.Ping()
}

// connect connects to the database
func connect(dbConfig *mysql.Config) (*MySQLStore, error) {
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
func migrateDb(db *MySQLStore) error {
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

// ExecTx executes a function within a database transaction
func (store *MySQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.BeginTx(ctx, nil)
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
