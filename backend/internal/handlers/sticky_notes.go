package handlers

import (
	"notes/internal/auth"
	"notes/internal/storage"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type StickyNoteHandler struct {
	Storage *storage.StickyNoteStorage
	SessionManager *auth.SessionManager
}

func NewStickyNoteHandler(storage *storage.StickyNoteStorage, sessionManager *auth.SessionManager) *StickyNoteHandler {
	return &StickyNoteHandler{
		Storage: storage,
		SessionManager: sessionManager,
	}
}

type stickyNoteRequestBody struct {
	Title string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type stickyNoteSuccessResponse struct {
	Id int `json:"id"`
}

// CreateStickyNote

func (s *StickyNoteHandler) CreateStickyNote(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := s.SessionManager.GetSession(sessionId)
	if err != nil {
		return c.JSON(Response{Success: false, Message: "invalid session"})
	}

	// get the sticky note info from the request body
	var stickyNote stickyNoteRequestBody

	err = c.BodyParser(&stickyNote)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the notebook id from the path
	notebookId, err := strconv.Atoi(c.Params("notebookId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the category id from the path
	categoryId, err := strconv.Atoi(c.Params("categoryId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// validate the sticky note struct
	validate := validator.New()
	err = validate.Struct(stickyNote)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// create the sticky note
	id, err := s.Storage.CreateStickyNote(&storage.NewStickyNote{
		NotebookID: notebookId,
		CategoryID: categoryId,
		AuthorID: user.Id,
		Title: stickyNote.Title,
		Content: stickyNote.Content,
	})

	if err != nil {
		// Log the error or return it to the client
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Success: false, Message: err.Error()})
	}

	resp := stickyNoteSuccessResponse{
		Id: id,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

// GetAllStickyNotesForUser

func (s *StickyNoteHandler) GetStickyNotesForUserCategory(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := s.SessionManager.GetSession(sessionId)
	if err != nil {
		return c.JSON(Response{Success: false, Message: "invalid session"})
	}

	// get the category id from the path
	categoryId, err := strconv.Atoi(c.Params("categoryId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the sticky notes for the user and category
	stickyNotes, err := s.Storage.GetStickyNotesForUserCategory(user.Id, categoryId)
	if err != nil {
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(stickyNotes)
}

// GetStickyNotesForUserCategory

func (s *StickyNoteHandler) GetAllStickyNotesForUser(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := s.SessionManager.GetSession(sessionId)
	if err != nil {
		return c.JSON(Response{Success: false, Message: "invalid session"})
	}

	// get the sticky notes for the user
	stickyNotes, err := s.Storage.GetAllStickyNotesForUser(user.Id)
	if err != nil {
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(stickyNotes)
}