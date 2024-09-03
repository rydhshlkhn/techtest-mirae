package resthandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rydhshlkhn/techtest-mirae/domain"
	"github.com/rydhshlkhn/techtest-mirae/usecase"
)

type StockHandler struct {
	uc usecase.StockUsecase
}

func NewStockHandler(app *fiber.App, uc usecase.StockUsecase) {
	handler := &StockHandler{uc: uc}
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/stock", handler.createStock)
	v1.Get("/stock", handler.getAllStock)
	v1.Get("/stock/:id", handler.getStockByID)
	v1.Put("/stock/:id", handler.updateStock)
	v1.Delete("/stock/:id", handler.deleteStock)
}

func (h *StockHandler) createStock(c *fiber.Ctx) (err error) {
	var stockPayload = new(domain.Stock)
	if err = c.BodyParser(&stockPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "can not parse json",
		})
	}

	stock := domain.Stock{
		Name:      stockPayload.Name,
		Code:      stockPayload.Code,
		Price:     stockPayload.Price,
		Frequency: stockPayload.Frequency,
		Volume:    stockPayload.Volume,
	}

	if err = h.uc.CreateStock(&stock); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(domain.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "success",
		Data:    stock,
	})
}

func (h *StockHandler) getAllStock(c *fiber.Ctx) (err error) {
	stocks := h.uc.GetAllStock()

	return c.Status(fiber.StatusOK).JSON(domain.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    stocks,
	})
}

func (h *StockHandler) getStockByID(c *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "invalid stock ID",
		})
	}

	stock, err := h.uc.GetStockByID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    stock,
	})
}

func (h *StockHandler) updateStock(c *fiber.Ctx) (err error) {
	var stock domain.Stock
	if err = c.BodyParser(&stock); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "can not parse json",
		})
	}

	stock_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "invalid stock ID",
		})
	}

	stock.ID = stock_id

	err = h.uc.UpdateStock(&stock)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    stock,
	})
}

func (h *StockHandler) deleteStock(c *fiber.Ctx) (err error) {
	stock_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "invalid stock ID",
		})
	}

	err = h.uc.DeleteStock(stock_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "success",
	})
}
