package sink

import "time"

//DHT22Reading is the struct that captures the raw reading of our sensors
type DHT22Reading struct {
	Time        time.Time `json:"time,omit_empty" gorm:"time;primary_key"`
	Hardware    string    `json:"hardware" gorm:"hardware;primary_key"`
	Humidity    float32   `json:"humidity" gorm:"humidity"`
	Temperature float32   `json:"temperature" gorm:"temperature"`
}

//TableName update our table name
func (DHT22Reading) TableName() string {
	return "dht_readings"
}

//SensorReadingAggregate represents aggregated sensor readings of different time buckets
type SensorReadingAggregate struct {
	Hardware    string  `json:"hardware" gorm:"device;index:hardware"`
	Humidity    float32 `json:"humidity" gorm:"humidity"`
	HumidityAvg float32 `json:"humidity_avg" gorm:"humidity_avg"`
	HumidityMax float32 `json:"humidity_max" gorm:"humidity_max"`
	HumidityMin float32 `json:"humidity_min" gorm:"humidity_min"`

	Temperature    float32 `json:"temperature" gorm:"temperature"`
	TemperatureAvg float32 `json:"temperature_avg" gorm:"temperature_avg"`
	TemperatureMax float32 `json:"temperature_max" gorm:"temperature_max"`
	TemperatureMin float32 `json:"temperature_min" gorm:"temperature_min"`

	Bucket string `json:"bucket"`
}