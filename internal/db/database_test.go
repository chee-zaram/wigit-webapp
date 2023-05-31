package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wigit-gh/webapp/internal/config"
	"github.com/wigit-gh/webapp/internal/logging"
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

	// Database object
	db = GetConnector(conf)

	// All table names
	tableNames = struct {
		users    string
		products string
		orders   string
		bookings string
		items    string
		services string
		slots    string
	}{
		users: "users", products: "products", orders: "orders", bookings: "bookings",
		items: "items", services: "services", slots: "slots",
	}
)

// TestMain runs before any test in this package runs.
// It sets all environment variables to ensure all tests can be carried out,
// and resets the variables when tests are done.
func TestMain(m *testing.M) {
	logging.ConfigureLogger("dev")

	// Drop all existing tables to start from clean slate
	db.Exec("DROP TABLES IF EXISTS users, orders, bookings, services, slots, products, items;")

	// Call NewDB to perform automigration
	db, _ = NewDB(dsn)

	// Run all tests
	exitCode := m.Run()

	// Return the exitCode
	os.Exit(exitCode)
}

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

	session, err := createDBConnection(dsn)
	assert.NotNil(session)
	assert.Nil(err)
	assert.IsType(gormDB, session)
}

// TestCreateDBConnectionPanic tests that the function panics with bad argument.
func TestCreateDBConnectionPanic(t *testing.T) {
	assert.Panics(t, func() { createDBConnection("nonsense-string") })
}

// TestNewDB checks that the function returns the custom database type.
func TestNewDB(t *testing.T) {
	assert := assert.New(t)
	var customDBtype *DB

	assert.NotNil(db)
	assert.IsType(customDBtype, db)
}

// TestNewDBPanic tests that the function panics with bad argument.
func TestNewDBPanic(t *testing.T) {
	assert.Panics(t, func() { NewDB("nonsense-string") })
}

// TestNewDBAutoMigration checks for the automigration of tables in the database
func TestNewDBAutoMigration(t *testing.T) {
	assert := assert.New(t)

	tables := []string{"users", "orders", "items", "bookings", "products", "services", "slots"}
	query := "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = ? AND table_name = ?"

	for _, table := range tables {
		var count int
		if err := db.Query(func(tx *gorm.DB) error {
			return tx.Raw(query, conf.DBName, table).Scan(&count).Error
		}); err != nil {
			t.Fatal("failed send query")
		}

		assert.True(count == 1)
	}
}

// TestBeforeCreateHook performs a create operation on the database.
func TestBeforeCreateHook(t *testing.T) {
	assert := assert.New(t)
	isFree := new(bool)
	slot := new(Slot)
	dbSlot := new(Slot)

	dateString := "Mon, 26 Jan 2024"
	timeString := "10:50 AM"
	*slot = Slot{
		DateString: &dateString,
		TimeString: &timeString,
		IsFree:     isFree,
	}

	if err := slot.SaveToDB(); err != nil {
		t.Fatal("failed to create slot in database")
	}

	if err := db.Query(func(tx *gorm.DB) error {
		return tx.Model(slot).First(dbSlot, slot.ID).Error
	}); err != nil {
		t.Fatal("failed to fetch user from slot database")
	}

	assert.NotNil(dbSlot)
	assert.NotNil(dbSlot.ID)
	assert.Equal(dbSlot.ID, slot.ID)
}
