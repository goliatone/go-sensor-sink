package sink

//Interactor CRUD interactions
type Interactor interface {
	Create(DHT22Reading) (DHT22Reading, error)
	List(SearchParameters) ([]DHT22Reading, error)
	ListByBucket(string) ([]SensorReadingAggregate, error)
	Update(DHT22Reading) error
	Delete(DHT22Reading) error
}

type interactor struct {
	repository Repository
}

//NewInteractor creates a new Interactor
func NewInteractor(sinkRepo Repository) Interactor {
	return &interactor{
		repository: sinkRepo,
	}
}

func (i interactor) Create(item DHT22Reading) (DHT22Reading, error) {
	record, err := i.repository.Add(item)
	if err != nil {
		return DHT22Reading{}, err
	}
	return record, nil
}

func (i interactor) List(query SearchParameters) ([]DHT22Reading, error) {
	records, err := i.repository.Get(query)
	if err != nil {
		return make([]DHT22Reading, 0), err
	}
	return records, nil
}

func (i interactor) ListByBucket(bucket string) ([]SensorReadingAggregate, error) {
	records, err := i.repository.GetAggregateByBucket(bucket)
	if err != nil {
		return make([]SensorReadingAggregate, 0), err
	}
	return records, nil
}

func (i interactor) Update(item DHT22Reading) error {
	return i.repository.Update(item)
}

func (i interactor) Delete(item DHT22Reading) error {
	return i.repository.Delete(item)
}
