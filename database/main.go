package database

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	_ "github.com/boltdb/bolt"
)

// ###########################################################################
// ###########################################################################
// Database
// ###########################################################################
// ###########################################################################

// Database encapsules everything needed for db access
type Database struct {
	db       *bolt.DB
	Instance string
	Path     string
}

// ###########################################################################

// NewDatabase is the constructor of a Database struct
func NewDatabase(path string, instance string) *Database {
	return &Database{
		nil,
		instance,
		path,
	}
}

// ---------------------------------------------------------------------------

// Open connects to the database
func (connection *Database) Open() {

	db, err := bolt.Open(connection.Path, 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		panic(err)
	}

	connection.db = db

	err = connection.db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("uptime"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte(connection.Instance))
		if err != nil {
			return fmt.Errorf("could not create instance bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		panic(fmt.Errorf("Failed to open database: %v", err))
	}
}

// ---------------------------------------------------------------------------

// Close closes to the database
func (connection *Database) Close() {
	if connection.db == nil {
		return
	}
	connection.db.Close()
}

// ---------------------------------------------------------------------------

// MarkUptime writes a new entry to the database
func (connection *Database) MarkUptime(startedAt time.Time, now time.Time) {
	uptimeStr := now.Sub(startedAt).String()
	nowStr := now.Format("2006-01-02 15:04:05")
	err := connection.db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("uptime")).Bucket([]byte(connection.Instance)).Put([]byte(nowStr), []byte(uptimeStr))
		if err != nil {
			return fmt.Errorf("could not insert value: %v", err)
		}
		return nil
	})
	if err != nil {
		panic(fmt.Errorf("Failed to update database: %v", err))
	}

	return
}
