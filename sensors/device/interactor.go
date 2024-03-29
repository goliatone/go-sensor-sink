package device

import uuid "github.com/satori/go.uuid"

//Interactor implements interactions
type Interactor interface {
	Create(Device) (Device, error)
	Read() ([]Device, error)
	Update(Device) error
	Delete(uuid.UUID) error
	GetByID(uuid.UUID) (Device, error)
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

func (i interactor) Read() ([]Device, error) {
	records, err := i.repository.Get()
	if err != nil {
		return make([]Device, 0), err
	}
	return records, nil
}

func (i interactor) Create(item Device) (Device, error) {
	record, err := i.repository.Add(item)
	if err != nil {
		return Device{}, err
	}
	return record, nil
}

func (i interactor) Update(item Device) error {
	err := i.repository.Update(item)
	return err
}

func (i interactor) Delete(id uuid.UUID) error {
	err := i.repository.DeleteByID(id)
	return err
}
