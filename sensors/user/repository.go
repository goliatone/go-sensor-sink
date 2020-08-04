package user

import (
	"fmt"
	"log"
	"sensors/storage"

	uuid "github.com/satori/go.uuid"
)

//Repository user repo
type Repository interface {
	Add(User) (User, error)
	Get() ([]User, error)
	Delete(User) error
	DeleteByID(uuid.UUID) error
	GetByID(uuid.UUID) (User, error)
	Update(User) error
}

type repository struct {
	database *storage.Database
}

//NewRepository create new repo
func NewRepository(db *storage.Database) Repository {
	return &repository{
		database: db,
	}
}

func (r repository) Add(item User) (User, error) {
	var model User
	record := r.database.Where(User{ID: item.ID}).First(&model)
	if record.RecordNotFound() == false {
		log.Printf("cant create record already exists %v", model)
		return model, &ErrRecordExists{inUser: item, outUser: model}
	}

	result := r.database.Where(User{ID: item.ID}).Assign(item).FirstOrCreate(&model)
	if err := result.Error; err != nil {
		return User{}, NewErrUnexpected(err)
	}

	return model, nil
}

func (r repository) Delete(user User) error {
	result := r.database.Delete(&user)
	if err := result.Error; err != nil {
		return NewErrUnexpected(result.Error)
	}
	return nil
}

func (r repository) DeleteByID(id uuid.UUID) error {
	var record User
	result := r.database.Where("id = ?", id.String()).First(&record)
	if result.RecordNotFound() {
		msg := fmt.Sprintf("record with id %v not found", id.String())
		return ErrRecordNotFound{message: msg}
	}
	if err := result.Error; err != nil {
		return NewErrUnexpected(result.Error)
	}

	result = r.database.Delete(&record)
	if err := result.Error; err != nil {
		return NewErrUnexpected(result.Error)
	}
	return nil
}

func (r repository) Get() ([]User, error) {
	var records []User
	r.database.Find(&records)
	return records, nil
}

func (r repository) GetByID(id uuid.UUID) (User, error) {
	var record User
	result := r.database.Where("id = ?", id.String()).First(&record)
	if result.RecordNotFound() {
		msg := fmt.Sprintf("record with id %v not found", id.String())
		return User{}, ErrRecordNotFound{message: msg}
	}

	if err := result.Error; err != nil {
		return User{}, err
	}
	return record, nil
}

func (r repository) Update(record User) error {
	var model User
	result := r.database.Model(&model).Omit("id").Update(record)
	return result.Error
}
