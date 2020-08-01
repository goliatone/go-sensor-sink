package registry

import (
	"sensors"
	"sensors/auth"
	"sensors/device"
	"sensors/sink"
	"sensors/storage"
)

type Domain struct {
	Auth  auth.Interactor
	Users auth.Repository
	// Devices  device.Repository
	Devices  device.Interactor
	Readings sink.Repository
}

//NewDomain will create a new domain
func NewDomain(config sensors.Config, database *storage.Database) *Domain {
	userRepo := auth.NewRepository(database)
	deviceRepo := device.NewRepository(database)

	return &Domain{
		Users:    userRepo,
		Devices:  device.NewInteractor(deviceRepo),
		Readings: sink.NewRepository(database),
		Auth:     auth.NewInteractor(config.Auth, userRepo),
	}
}
