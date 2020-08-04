package registry

import (
	"sensors"
	"sensors/auth"
	"sensors/device"
	"sensors/sink"
	"sensors/storage"
	"sensors/user"
)

//Domain exposes our actions
type Domain struct {
	Auth     auth.Interactor
	Users    user.Interactor
	Devices  device.Interactor
	Readings sink.Interactor
}

//NewDomain will create a new domain
func NewDomain(config sensors.Config, database *storage.Database, channels *Channels) *Domain {
	authRepo := auth.NewRepository(database)
	sinkRepo := sink.NewRepository(database)
	userRepo := user.NewRepository(database)
	deviceRepo := device.NewRepository(database)

	return &Domain{
		Users:    user.NewInteractor(userRepo),
		Readings: sink.NewInteractor(sinkRepo),
		Devices:  device.NewInteractor(deviceRepo),
		Auth:     auth.NewInteractor(config.Auth, authRepo, channels.ChannelNewUsers),
	}
}
