package readings

import (
	"sensors/sink"

	"github.com/gofiber/fiber"
)

//Create will create and return a reading
func Create(domain sink.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		var item sink.DHT22Reading
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

		ctx.Status(fiber.StatusOK).JSON(sink.ReadResponse(record))
	}
}

//List returns readings
func List(domain sink.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		qs := sink.NewSearchParameters(ctx)

		records, err := domain.List(qs)
		if err != nil {
			errHTTP := ErrResponse(err)
			ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}
		ctx.Status(fiber.StatusOK).JSON(sink.ListResponse(records))
	}
}

//ListByBucket will return a list of readings
func ListByBucket(domain sink.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		bucket := ctx.Params("bucket")
		records, err := domain.ListByBucket(bucket)
		if err != nil {
			errHTTP := ErrResponse(err)
			ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}
		ctx.Status(fiber.StatusOK).JSON(sink.AggregateResponse(records))
	}
}
