package device

import (
	"fmt"
	"log"
	"sensors/storage"

	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	Add(Device) (Device, error)
	Get() ([]Device, error)
	Delete(Device) error
	DeleteByID(uuid.UUID) error
	GetByID(uuid.UUID) (Device, error)
	GetByHardwareID(string) (Device, error)
	Update(Device) error
}

type repository struct {
	database *storage.Database
}

func NewRepository(db *storage.Database) Repository {
	return &repository{database: db}
}

func (r repository) Add(device Device) (Device, error) {
	var d Device

	record := r.database.Where(Device{HardwareID: device.HardwareID}).First(&d)
	if record.RecordNotFound() == false {
		log.Printf("cant create device already exists %v", d)
		return d, &ErrDeviceExists{inDevice: device, outDevice: d}
	}

	result := r.database.Where(Device{HardwareID: device.HardwareID}).Assign(device).FirstOrCreate(&d)
	if err := result.Error; err != nil {
		return Device{}, NewErrUnexpected(err)
	}

	return d, nil
}

func (r repository) Delete(device Device) error {
	result := r.database.Delete(&device)
	if err := result.Error; err != nil {
		return NewErrUnexpected(result.Error)
	}
	return nil
}

func (r repository) DeleteByID(id uuid.UUID) error {
	var device Device
	result := r.database.Where("id = ?", id.String()).First(&device)
	if result.RecordNotFound() {
		msg := fmt.Sprintf("device with id %v not found", id.String())
		return ErrDeviceNotFound{message: msg}
	}
	if err := result.Error; err != nil {
		return NewErrUnexpected(result.Error)
	}
	result = r.database.Delete(&device)
	if err := result.Error; err != nil {
		return NewErrUnexpected(result.Error)
	}
	return nil
}

func (r repository) Get() ([]Device, error) {
	var devices []Device
	r.database.Find(&devices)
	return devices, nil
}

func (r repository) GetByID(id uuid.UUID) (Device, error) {
	var device Device

	result := r.database.Where("id = ?", id.String()).First(&device)

	if result.RecordNotFound() {
		msg := fmt.Sprintf("device with id %v not found", id.String())
		return Device{}, ErrDeviceNotFound{message: msg}
	}

	if err := result.Error; err != nil {
		return Device{}, err
	}

	return device, nil
}

func (r repository) GetByHardwareID(hid string) (Device, error) {
	var device Device
	result := r.database.Where("hardware_id = ?", hid).First(&device)

	if result.RecordNotFound() {
		msg := fmt.Sprintf("device with hardware id %v not found", hid)
		return Device{}, ErrDeviceNotFound{message: msg}
	}

	if err := result.Error; err != nil {
		return Device{}, err
	}

	return device, nil
}

func (r repository) Update(device Device) error {
	var d Device
	result := r.database.Model(&d).Omit("id").Update(device)

	return result.Error
}
