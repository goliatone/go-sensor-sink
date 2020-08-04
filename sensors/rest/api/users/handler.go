package users

import (
	"sensors/user"

	"github.com/gofiber/fiber"
	uuid "github.com/satori/go.uuid"
)

func Read(domain user.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		recordID := uuid.FromStringOrNil(ctx.Params("id"))
		record, err := domain.GetByID(recordID)
		if err != nil {
			errHTTP := ErrResponse(err)
			ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		ctx.Status(fiber.StatusOK).JSON(user.ReadResponse(record))
	}
}

func List(domain user.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		records, err := domain.Read()
		if err != nil {
			errHTTP := ErrResponse(err)
			ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}
		ctx.Status(fiber.StatusOK).JSON(user.ListResponse(records))
	}
}

func Create(domain user.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		var item user.User
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
		ctx.Status(fiber.StatusOK).JSON(user.ReadResponse(record))
	}
}

func Update(domain user.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		recordID := uuid.FromStringOrNil(ctx.Params("id"))

		var item user.User
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

		ctx.Status(fiber.StatusOK).JSON(user.ReadResponse(item))
	}
}

func Delete(domain user.Interactor) func(*fiber.Ctx) {
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
