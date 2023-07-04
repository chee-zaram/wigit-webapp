package db

import (
	"fmt"
	// _ "time/tzdata" if we make use of select time zone we may have to uncomment this and move to main

	"github.com/chee-zaram/wigit-webapp/backend/internal/config"
	"github.com/chee-zaram/wigit-webapp/backend/internal/logger"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is a type used to extend the functionality of the DB object.
type DB struct {
	*gorm.DB
}

// Connector servers as a global link to the database.
var Connector *DB

// Query method is called to send request to the database within a scoped session.
//
// It takes in a function which carries out a database operation within the scope of
// the `tx` transaction parameter.
func (db *DB) Query(fn func(tx *gorm.DB) error) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := fn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// NewDatabaseDSN returns a string to be used an domain string name when
// creating a database connection.
func NewDatabaseDSN(conf config.Config) string {
	return fmt.Sprintf(
		// we use Local time here but don't forget to make it use gh time before prod
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DBUser,
		conf.DBPass,
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
	)
}

// NewDB takes in a domain string name and returns a pointer to a
// newly configured DB object with connection to the database and an error if any occurred.
func NewDB(dsn string) (*DB, error) {
	session, err := createDBConnection(dsn)
	if err != nil {
		return nil, err
	}

	// Set up connection pool, etc. as needed

	return &DB{session}, nil
}

// createConnection sets up connection to the database and creates a session.
// It returns the session and an error if any occurred.
func createDBConnection(dsn string) (*gorm.DB, error) {
	// Open connection to database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.SetGORMLogToFile()})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create or migrate tables based on the structs.
	if err = db.AutoMigrate(GetSchemas()); err != nil {
		return nil, fmt.Errorf("failed to automigrate tables: %w", err)
	}

	// Create session
	session := db.Session(&gorm.Session{FullSaveAssociations: true})
	if session.Error != nil {
		return nil, fmt.Errorf("failed to create a new session: %w", session.Error)
	}

	return session, nil
}

// GetConnector returns a ready connector to the database.
func GetConnector(conf config.Config) *DB {
	dbConnector, err := NewDB(NewDatabaseDSN(conf))
	if err != nil {
		log.Panic().Err(err).Msg("failed to get db connector")
	}

	return dbConnector
}
