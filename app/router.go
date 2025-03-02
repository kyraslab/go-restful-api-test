package app

import (
	"github.com/aronipurwanto/go-restful-api/controller"
	"github.com/aronipurwanto/go-restful-api/middleware"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	categoryController controller.CategoryController
	productController  controller.ProductController
}

func NewController(categoryController controller.CategoryController, productController controller.ProductController) Controller {
	return Controller{categoryController, productController}
}

func NewRouter(app *fiber.App, c Controller) {
	authMiddleware := middleware.NewAuthMiddleware()
	api := app.Group("/api", authMiddleware)

	categories := api.Group("/categories")
	categories.Get("/", c.categoryController.FindAll)
	categories.Get("/:categoryId", c.categoryController.FindById)
	categories.Post("/", c.categoryController.Create)
	categories.Put("/:categoryId", c.categoryController.Update)
	categories.Delete("/:categoryId", c.categoryController.Delete)

	products := api.Group("/products")
	products.Get("/", c.productController.FindAll)
	products.Get("/:categoryId", c.productController.FindById)
	products.Post("/", c.productController.Create)
	products.Put("/:categoryId", c.productController.Update)
	products.Delete("/:categoryId", c.productController.Delete)
}
