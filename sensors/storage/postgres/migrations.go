package postgres

import (
	"sensors/auth"
	"sensors/device"
	"sensors/storage"
)

//Migrate updates he db with new columns, tables and the like
func Migrate(database *storage.Database) {

	//Enable uuid extension
	database.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	database.DB.AutoMigrate(
		auth.User{},
		device.Device{},
	)
}
