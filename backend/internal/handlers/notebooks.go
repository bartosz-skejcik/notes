package handlers

import (
	"fmt"
	"notes/internal/auth"
	"notes/internal/storage"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type NotebookHandler struct {
	Storage *storage.NotebookStorage
	SessionManager *auth.SessionManager
}

type notebookRequestBody struct {
	Name string `json:"name" validate:"required"`
}

type notebookSuccessResponse struct {
	Id int `json:"id"`
}

func NewNotebookHandler(storage *storage.NotebookStorage, sessionManager *auth.SessionManager) *NotebookHandler {
	return &NotebookHandler{
		Storage: storage,
		SessionManager: sessionManager,
	}
}

// @Summary Get the user's notebooks
// @Description Get the user's notebooks
// @Tags notebooks
// @Accept json
// @Produce json
// @Success 200 {array} []storage.Notebook
// @Failure 400 {object} handlers.Response{Success: false}
// @Failure 500 {object} handlers.Response{Success: false}
// @Router /notebooks [get]
func (n *NotebookHandler) GetNotebooks(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := n.SessionManager.GetSession(sessionId)
	if err != nil {
		fmt.Println("siema")
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	notebooks, err := n.Storage.GetNotebooks(user.Id)
	if err != nil {
		fmt.Println("elo")
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(notebooks)
}

// @Summary Create a notebook
// @Description Create a notebook
// @Tags notebooks
// @Accept json
// @Produce json
// @Param notebook body notebookRequestBody true "The notebook's name"
// @Success 200 {object} notebookSuccessResponse
// @Failure 400 {object} handlers.Response{Success: false}
// @Failure 500 {object} handlers.Response{Success: false}
// @Router /notebooks [post]
func (n *NotebookHandler) CreateNotebook(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := n.SessionManager.GetSession(sessionId)
	if err != nil {
		return c.JSON(Response{Success: false, Message: "invalid session"})
	}

	// get the notebook info from the request body
	var notebook notebookRequestBody

	err = c.BodyParser(&notebook)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// validate the notebook struct
	validate := validator.New()
	err = validate.Struct(notebook)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// create the notebook
	id, err := n.Storage.CreateNotebook(&storage.NewNotebook{
		UserId: user.Id,
		Name: notebook.Name,
	})

	if err != nil {
		// Log the error or return it to the client
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Success: false, Message: err.Error()})
	}

	resp := notebookSuccessResponse{
		Id: id,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @Summary Get a notebook
// @Description Get a notebook
// @Tags notebooks
// @Accept json
// @Produce json
// @Param notebookId path int true "The notebook's id"
// @Success 200 {object} storage.Notebook
// @Failure 400 {object} handlers.Response{Success: false}
// @Failure 500 {object} handlers.Response{Success: false}
// @Router /notebooks/{notebookId} [get]
func (n *NotebookHandler) GetNotebook(c *fiber.Ctx) error {
	// get the notebook id from the path
	notebookId, err := strconv.Atoi(c.Params("notebookId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the notebook data from the storage
	notebook, err := n.Storage.GetNotebookById(notebookId)
	if err != nil {
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(notebook)
}