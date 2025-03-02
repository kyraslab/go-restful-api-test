package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kyraslab/go-restful-api-test/exception"
	"github.com/kyraslab/go-restful-api-test/model/web"
	"github.com/kyraslab/go-restful-api-test/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

// Create Category
func (controller *CategoryControllerImpl) Create(c *fiber.Ctx) error {
	categoryCreateRequest := new(web.CategoryCreateRequest)
	if err := c.BodyParser(categoryCreateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	categoryResponse, err := controller.CategoryService.Create(c.Context(), *categoryCreateRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   categoryResponse,
	})
}

// Update Category
func (controller *CategoryControllerImpl) Update(c *fiber.Ctx) error {
	categoryUpdateRequest := new(web.CategoryUpdateRequest)
	if err := c.BodyParser(categoryUpdateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	id, err := strconv.ParseUint(c.Params("categoryId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Invalid Category ID",
			Data:   err.Error(),
		})
	}
	categoryUpdateRequest.Id = id

	categoryResponse, err := controller.CategoryService.Update(c.Context(), *categoryUpdateRequest)
	if err != nil {
		if _, ok := err.(exception.NotFoundError); ok {
			return c.Status(fiber.StatusNotFound).JSON(web.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: "Not Found",
				Data:   err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	})
}

// Delete Category
func (controller *CategoryControllerImpl) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("categoryId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Invalid Category ID",
			Data:   err.Error(),
		})
	}

	err = controller.CategoryService.Delete(c.Context(), id)
	if err != nil {
		if _, ok := err.(exception.NotFoundError); ok {
			return c.Status(fiber.StatusNotFound).JSON(web.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: "Not Found",
				Data:   err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "Deleted Successfully",
	})
}

// Find Category By ID
func (controller *CategoryControllerImpl) FindById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("categoryId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Invalid Category ID",
			Data:   err.Error(),
		})
	}

	categoryResponse, err := controller.CategoryService.FindById(c.Context(), id)
	if err != nil {
		if _, ok := err.(exception.NotFoundError); ok {
			return c.Status(fiber.StatusNotFound).JSON(web.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: "Not Found",
				Data:   err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	})
}

// Find All Categories
func (controller *CategoryControllerImpl) FindAll(c *fiber.Ctx) error {
	categoryResponses, err := controller.CategoryService.FindAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   categoryResponses,
	})
}
