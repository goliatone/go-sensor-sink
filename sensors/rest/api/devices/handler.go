package devices

import (
	"sensors/device"

	"github.com/gofiber/fiber"
	uuid "github.com/satori/go.uuid"
)

//GetByID returns a device by ID
func Read(domain device.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		recordID := uuid.FromStringOrNil(ctx.Params("id"))
		record, err := domain.GetByID(recordID)
		if err != nil {
			errHTTP := ErrResponse(err)
			ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		ctx.Status(fiber.StatusOK).JSON(device.ReadResponse(record))
	}
}

//List returns a list of devices
func List(domain device.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		records, err := domain.Read()
		if err != nil {
			errHTTP := ErrResponse(err)
			ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}
		ctx.Status(fiber.StatusOK).JSON(device.ListResponse(records))
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

		ctx.Status(fiber.StatusOK).JSON(device.ReadResponse(record))
	}
}

//Update will update a given record
func Update(domain device.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		recordID := uuid.FromStringOrNil(ctx.Params("id"))

		var item device.Device
		item, err := domain.GetByID(recordID)
		if err != nil {
			errHTTP := ErrResponse(err)
			ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		if err := ctx.BodyParser(&item); err != nil {
			errHTTP := ErrResponse(err)
			ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		if err := domain.Update(item); err != nil {
			errHTTP := ErrResponse(err)
			ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		ctx.Status(fiber.StatusOK).JSON(device.ReadResponse(item))
	}
}

//Delete will delete by id
func Delete(domain device.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		recordID := uuid.FromStringOrNil(ctx.Params("id"))

		if err := domain.Delete(recordID); err != nil {
			errHTTP := ErrResponse(err)
			ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		ctx.Status(fiber.StatusOK).Send("OK")
	}
}
