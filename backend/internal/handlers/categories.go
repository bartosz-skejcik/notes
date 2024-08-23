package handlers

import (
	"notes/internal/auth"
	"notes/internal/storage"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	Storage *storage.CategoryStorage
	SessionManager *auth.SessionManager
}

type categoryRequestBody struct {
	Name string `json:"name" validate:"required"`
	Color string `json:"color" validate:"required"`
}

type categorySuccessResponse struct {
	Id int `json:"id"`
}

func NewCategoryHandler(storage *storage.CategoryStorage, sessionManager *auth.SessionManager) *CategoryHandler {
	return &CategoryHandler{
		Storage: storage,
		SessionManager: sessionManager,
	}
}

// @Summary Get the user's categories
// @Description Get the user's categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {array} []storage.Category
// @Failure 400 {object} handlers.Response{Success: false, Message: string}
// @Failure 500 {object} handlers.Response{Success: false, Message: string}
// @Router /notebooks/{notebookId}/categories [get]
func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	// get the notebook id from the path
	notebookId, err := strconv.Atoi(c.Params("notebookId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the user data from the session
	user, err := h.SessionManager.GetSession(c.Get("Authorization")[7:])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	categories, err := h.Storage.GetCategories(notebookId, user.Id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(categories)
}

// @Summary Create a category
// @Description Create a category
// @Tags categories
// @Accept json
// @Produce json
// @Param notebookId path int true "The notebook's id"
// @Param category body categoryRequestBody true "The category's name"
// @Success 200 {object} categorySuccessResponse
// @Failure 400 {object} handlers.Response{Success: false, Message: string}
// @Failure 500 {object} handlers.Response{Success: false, Message: string}
// @Router /notebooks/{notebookId}/categories [post]
func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := h.SessionManager.GetSession(sessionId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the notebook id from the path
	notebookId, err := strconv.Atoi(c.Params("notebookId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the category info from the request body
	var category categoryRequestBody

	err = c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// validate the category struct
	validate := validator.New()
	err = validate.Struct(category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// create the category
	id, err := h.Storage.CreateCategory(&storage.NewCategory{
		NotebookId: notebookId,
		UserId: user.Id,
		Name: category.Name,
		Color: category.Color,
	})

	if err != nil {
		// Log the error or return it to the client
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Success: false, Message: err.Error()})
	}

	resp := categorySuccessResponse{
		Id: id,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}