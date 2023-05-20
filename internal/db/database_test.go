package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wigit-gh/webapp/internal/config"
	"gorm.io/gorm"
)

var (
	// conf is the configuration we'll use to connect to our db during testing.
	conf = config.Config{
		DBUser: "root",
		DBPass: "WWApp-dev-pwd-0", // modify this to be your local root password but change back before pushing
		DBHost: "localhost",
		DBPort: "3306",
		DBName: "wwapp_dev_db_test", // make sure the create this db. there's a script in /scripts
	}

	// dsn is the domain string name to connect to the database.
	dsn = NewDatabaseDSN(conf)
)

// TestNewDatabaseDSN tests that the function returns a valid dsn value.
func TestNewDatabaseDSN(t *testing.T) {
	assert := assert.New(t)

	assert.True(dsn != "")
	assert.Contains(dsn, conf.DBUser)
	assert.Contains(dsn, conf.DBPass)
	assert.Contains(dsn, conf.DBHost)
	assert.Contains(dsn, conf.DBPort)
	assert.Contains(dsn, conf.DBName)
}

// TestCreateDBConnection makes sure the object returned by the createDBConnection
// function is of the correct type.
func TestCreateDBConnection(t *testing.T) {
	assert := assert.New(t)
	var gormDB *gorm.DB

	// On success.
	session, err := createDBConnection(dsn)
	assert.NotNil(session)
	assert.Nil(err)
	assert.IsType(gormDB, session)

	// On failure
	assert.Panics(func() { createDBConnection("nonsense-string") })
}

// TestNewDB checks that the function returns the custom database type.
func TestNewDB(t *testing.T) {
	assert := assert.New(t)
	var customDBtype *DB

	// On success
	db, err := NewDB(dsn)
	assert.NotNil(db)
	assert.Nil(err)
	assert.IsType(customDBtype, db)

	// On wrong input
	assert.Panics(func() { NewDB("nonsense-string") })
}
