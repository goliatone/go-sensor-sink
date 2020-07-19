package sink

import (
	"fmt"
	"log"
	"sensors/storage"
	"strconv"
	"time"

	"github.com/gofiber/fiber"
	uuid "github.com/satori/go.uuid"
)

//Repository is the interface to handle persistence of DHT22Reading
type Repository interface {
	Add(DHT22Reading) (DHT22Reading, error)
	Delete(DHT22Reading) error
	Get(SearchParameters) ([]DHT22Reading, error)
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
	if reading.Time.IsZero() == true {
		reading.Time = time.Now()
	}

	// log.Printf("new reading: %+v\n", reading)
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

//SearchParameters query builder for search
type SearchParameters struct {
	Page  int
	Limit int
	Order string
}

//NewSearchParameters build new SearchParameters from context
func NewSearchParameters(c *fiber.Ctx) SearchParameters {
	sp := SearchParameters{
		Page:  1,
		Limit: 10,
		Order: "time desc",
	}

	if c.Query("order") != "" {
		sp.Order = c.Query("order")
	}

	if c.Query("page") != "" {
		p, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			log.Println("error converting page")
		} else {
			sp.Page = p
		}
	}

	if c.Query("limit") != "" {
		l, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			log.Println("error converting limit")
		} else {
			sp.Limit = l
		}
	}

	log.Printf("order: %v\n", c.Query("order"))

	return sp
}

//Offset returns page * limit
func (s SearchParameters) Offset() int {
	return s.Page * s.Limit
}

func (r repository) Get(qs SearchParameters) ([]DHT22Reading, error) {
	var readings []DHT22Reading
	// r.database.LogMode(true)
	// r.database.Offset(40).Limit(10).Find(&readings)

	offset := qs.Offset()
	tx := r.database.Offset(offset)
	tx = tx.Limit(qs.Limit)
	tx = tx.Order(qs.Order)

	tx.LogMode(true)
	tx.Find(&readings)

	return readings, nil
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
