package devices

import (
	"sensors/device"

	"github.com/gofiber/fiber"
	uuid "github.com/satori/go.uuid"
)

//GetByID returns a device by ID
func GetByID(domain device.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		id, err := uuid.FromString(ctx.Params("id"))
		if err != nil {
			errHTTP := ErrResponse(err)
			_ = ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		record, err := domain.GetByID(id)
		if err != nil {
			errHTTP := ErrResponse(err)
			_ = ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		ctx.Status(fiber.StatusOK).JSON(device.DeviceResponse(record))
	}
}

//Get returns a list of devices
func Get(domain device.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		records, err := domain.Get()
		if err != nil {
			errHTTP := ErrResponse(err)
			ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}
		ctx.Status(fiber.StatusOK).JSON(device.ListDevicesResponse(records))
	}
}

//Create will create and return a device
func Create(domain device.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		var item device.Device
		if err := ctx.BodyParser(&item); err != nil {
			errHTTP := ErrResponse(err)
			ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		record, err := domain.Create(item)
		if err != nil {
			errHTTP := ErrResponse(err)
			ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		ctx.Status(fiber.StatusOK).JSON(device.DeviceResponse(record))
	}
}
