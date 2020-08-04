package data

import uuid "github.com/satori/go.uuid"

//SensorReading holds sensor data
type SensorReading struct {
}

//UserContract for User operations
type UserContract struct {
	UserID uuid.UUID
}

//ChanNewUsers handles new users
type ChanNewUsers struct {
	Channel chan UserContract
	Reader  <-chan UserContract
	Writer  chan<- UserContract
}
