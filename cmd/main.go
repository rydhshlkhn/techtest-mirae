package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rydhshlkhn/techtest-mirae/config"
	"github.com/rydhshlkhn/techtest-mirae/delivery/resthandler"
	"github.com/rydhshlkhn/techtest-mirae/infra"
	"github.com/rydhshlkhn/techtest-mirae/repository"
	"github.com/rydhshlkhn/techtest-mirae/usecase"
)

func main() {
	config.LoadEnv()
	db := infra.IniitDB()
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	stockRepo := repository.NewStockRepository(db)
	stockUsecase := usecase.NewStockUsecase(stockRepo)
	resthandler.NewStockHandler(app, stockUsecase)

	app.Listen(":8000")
}
