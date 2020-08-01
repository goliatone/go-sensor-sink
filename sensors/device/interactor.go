package device

import uuid "github.com/satori/go.uuid"

//Interactor implements interactions
type Interactor interface {
	Get() ([]Device, error)
	GetByID(uuid.UUID) (Device, error)
	Create(Device) (Device, error)
}

type interactor struct {
	repository Repository
}

//NewInteractor creates a new Interactor instance
func NewInteractor(deviceRepo Repository) Interactor {
	return &interactor{
		repository: deviceRepo,
	}
}

//GetByID will find a device by id
func (i interactor) GetByID(id uuid.UUID) (Device, error) {
	item, err := i.repository.GetByID(id)
	if err != nil {
		return Device{}, err
	}
	return item, nil
}

func (i interactor) Get() ([]Device, error) {
	items, err := i.repository.Get()
	if err != nil {
		return make([]Device, 0), err
	}
	return items, nil
}

func (i interactor) Create(item Device) (Device, error) {
	record, err := i.repository.Add(item)
	if err != nil {
		return Device{}, err
	}
	return record, nil
}
