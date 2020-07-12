package sink

import (
	"fmt"
	"log"
	"sensors/storage"

	uuid "github.com/satori/go.uuid"
)

//Repository is the interface to handle persistence of DHT22Reading
type Repository interface {
	Add(DHT22Reading) (DHT22Reading, error)
	Delete(DHT22Reading) error
	GetByID(uuid.UUID) (DHT22Reading, error)
}

type repository struct {
	database *storage.Database
}

//NewRepository creates a new DHT22Reading repository
func NewRepository(db *storage.Database) Repository {
	return &repository{database: db}
}

//Add a new DHT22Reading entry
func (r repository) Add(reading DHT22Reading) (DHT22Reading, error) {
	// var rdr DHT22Reading

	// var notExists bool
	// notExists = r.database.Where(DHT22Reading{Time: reading.Time}).First(&rdr).RecordNotFound()
	// if !notExists {
	// 	log.Printf("reading %v", r)
	// 	return rdr, &ErrReadingExists{inReading: reading, outReading: rdr}
	// }
	log.Printf("new reading: %+v\n", reading)
	result := r.database.Create(&reading)
	if err := result.Error; err != nil {
		return DHT22Reading{}, NewErrUnexpected(err)
	}

	return reading, nil
}

func (r repository) Delete(reading DHT22Reading) error {
	result := r.database.Delete(&reading)
	if err := result.Error; err != nil {
		return NewErrUnexpected(result.Error)
	}
	return nil
}

func (r repository) GetByID(id uuid.UUID) (DHT22Reading, error) {
	var reading DHT22Reading

	result := r.database.Where("device_id = ?", id.String()).First(&reading)
	if result.RecordNotFound() {
		msg := fmt.Sprintf("reading not found %v", id.String())
		return DHT22Reading{}, ErrReadingNotFound{message: msg}
	}

	if err := result.Error; err != nil {
		return DHT22Reading{}, err
	}
	return reading, nil
}
