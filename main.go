package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kyraslab/go-restful-api-test/app"
	"github.com/kyraslab/go-restful-api-test/controller"
	"github.com/kyraslab/go-restful-api-test/helper"
	"github.com/kyraslab/go-restful-api-test/model/domain"
	"github.com/kyraslab/go-restful-api-test/repository"
	"github.com/kyraslab/go-restful-api-test/service"
	"log"
)

func main() {

	server := fiber.New()

	// Initialize Database
	db := app.NewDB()

	// Run Auto Migration (Opsional, bisa dihapus jika tidak diperlukan)
	err := db.AutoMigrate(&domain.Category{})
	helper.PanicIfError(err)

	// Initialize Validator
	validate := validator.New()

	// Initialize Repository, Service, and Controller
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository, validate)
	categoryController := controller.NewCategoryController(categoryService)
	productRepository := repository.NewCategoryRepository(db)
	productService := service.NewCategoryService(productRepository, validate)
	productController := controller.NewCategoryController(productService)

	controllers := app.NewController(categoryController, productController)

	// Setup Routes
	app.NewRouter(server, controllers)

	// Start Server
	log.Println("Server running on port 8080")
	err = server.Listen(":8080")
	helper.PanicIfError(err)
}
