package postgres

import (
	"sensors"
	"sensors/storage"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var once sync.Once

//NewDatabase creates a new Database object
func NewDatabase(config sensors.Config) (*storage.Database, error) {
	var err error

	db := new(storage.Database)

	var conn *gorm.DB
	once.Do(func() {
		conn, err = gorm.Open("postgres", config.DB.String("disable"))
		//check config if enabled
		// conn.LogMode(true)
	})

	if err != nil {
		return nil, err
	}

	db.DB = conn

	return db, nil
}
