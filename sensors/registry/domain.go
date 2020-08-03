package registry

import (
	"sensors"
	"sensors/auth"
	"sensors/device"
	"sensors/sink"
	"sensors/storage"
)

//Domain exposes our actions
type Domain struct {
	Auth     auth.Interactor
	Users    auth.Repository
	Devices  device.Interactor
	Readings sink.Interactor
}

//NewDomain will create a new domain
func NewDomain(config sensors.Config, database *storage.Database) *Domain {
	userRepo := auth.NewRepository(database)
	sinkRepo := sink.NewRepository(database)
	deviceRepo := device.NewRepository(database)

	return &Domain{
		Users:    userRepo,
		Readings: sink.NewInteractor(sinkRepo),
		Devices:  device.NewInteractor(deviceRepo),
		Auth:     auth.NewInteractor(config.Auth, userRepo),
	}
}
