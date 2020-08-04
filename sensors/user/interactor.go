package user

import uuid "github.com/satori/go.uuid"

type Interactor interface {
	Create(User) (User, error)
	Read() ([]User, error)
	Update(User) error
	Delete(uuid.UUID) error
	GetByID(uuid.UUID) (User, error)
}

type interactor struct {
	repository Repository
}

func NewInteractor(userRepo Repository) Interactor {
	return &interactor{
		repository: userRepo,
	}
}

func (i interactor) GetByID(id uuid.UUID) (User, error) {
	item, err := i.repository.GetByID(id)
	if err != nil {
		return User{}, err
	}
	return item, nil
}

func (i interactor) Read() ([]User, error) {
	records, err := i.repository.Get()
	if err != nil {
		return make([]User, 0), err
	}
	return records, nil
}

func (i interactor) Create(item User) (User, error) {
	record, err := i.repository.Add(item)
	if err != nil {
		return User{}, err
	}
	return record, nil
}

func (i interactor) Update(item User) error {
	return i.repository.Update(item)
}

func (i interactor) Delete(id uuid.UUID) error {
	return i.repository.DeleteByID(id)
}
