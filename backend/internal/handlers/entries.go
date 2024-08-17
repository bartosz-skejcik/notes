package handlers

import (
	"fmt"
	"notes/internal/auth"
	"notes/internal/storage"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type EntryHandler struct {
	Storage *storage.EntryStorage
	SessionManager *auth.SessionManager
}

// type NewEntry struct {
// 	NotebookID int `json:"notebook_id"`
// 	Title      string `json:"title"`
// 	Content    string `json:"content"`
// 	Timestamp  time.Time
// 	HasPhoto   bool `json:"has_photo"`
// 	TagId int `json:"tag_id"`
// }

// type NewChildEntry struct {
// 	NotebookID int `json:"notebook_id"`
// 	Title      string `json:"title"`
// 	Content    string `json:"content"`
// 	ParentEntryId int `json:"parent_entry_id"`
// 	Timestamp  time.Time
// 	TagId int `json:"tag_id"`
// }

// type Entry struct {
// 	ID        int `json:"id"`
// 	NotebookID int `json:"notebook_id"`
// 	Title     string `json:"title"`
// 	Content   string `json:"content"`
// 	Timestamp time.Time `json:"timestamp"`
// 	HasPhoto bool `json:"has_photo"`
// 	TagId int `json:"tag_id"`
// }

type entryRequestBody struct {
	NotebookId int
	Title string `json:"title" validate:"required"`
	Content string `json:"content"`
	Role string `json:"role"`
	HasPhoto bool `json:"has_photo"`
	TagId int `json:"tag_id"`
}

type childEntryRequestBody struct {
	NotebookId int
	Title string `json:"title" validate:"required"`
	Content string `json:"content"`
	Role string `json:"role"`
	ParentEntryId int `json:"parent_entry_id" validate:"required"`
	TagId int `json:"tag_id"`
}

type photoRequestBody struct {
	ImageData []byte `json:"image_data" validate:"required"`
	MimeType string `json:"mime_type" validate:"required"`
}

type entrySuccessResponse struct {
	Id int `json:"id"`
}

type photoSuccessResponse struct {
	Id int `json:"id"`
}

func NewEntryHandler(storage *storage.EntryStorage, sessionManager *auth.SessionManager) *EntryHandler {
	return &EntryHandler{
		Storage: storage,
		SessionManager: sessionManager,
	}
}

// @Summary Get the user's entries for a notebook
// @Description Get the user's entries for a notebook
// @Tags entries
// @Accept json
// @Produce json
// @Param notebookId path int true "The notebook's id"
// @Success 200 {array} []storage.Entry
// @Failure 400 {object} handlers.Response{Success: false, Message: string}
// @Failure 500 {object} handlers.Response{Success: false, Message: string}
// @Router /notebooks/{notebookId}/entries [get]

func (e *EntryHandler) GetEntries(c *fiber.Ctx) error {
	// get the notebook id from the path
	notebookId, err := strconv.Atoi(c.Params("notebookId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := e.SessionManager.GetSession(sessionId)
	if err != nil {
		fmt.Println("siema")
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	entries, err := e.Storage.GetEntriesForNotebook(user.Id, notebookId)
	if err != nil {
		fmt.Println("elo")
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(entries)
}

// @Summary Create an entry
// @Description Create an entry
// @Tags entries
// @Accept json
// @Produce json
// @Param notebookId path int true "The notebook's id"
// @Param entry body entryRequestBody true "The entry's info"
// @Success 200 {object} entrySuccessResponse
// @Failure 400 {object} handlers.Response{Success: false, Message: string}
// @Failure 500 {object} handlers.Response{Success: false, Message: string}
// @Router /notebooks/{notebookId}/entries [post]
func (e *EntryHandler) CreateEntry(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := e.SessionManager.GetSession(sessionId)
	if err != nil {
		return c.JSON(Response{Success: false, Message: "invalid session"})
	}

	// get the notebook id from the path
	notebookId, err := strconv.Atoi(c.Params("notebookId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the entry info from the request body
	var entry entryRequestBody

	err = c.BodyParser(&entry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// validate the entry struct
	validate := validator.New()
	err = validate.Struct(entry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// create the entry
	newEntry, err := e.Storage.CreateEntry(user.Id, &storage.NewEntry{
		NotebookID: notebookId,
		Title: entry.Title,
		Content: entry.Content,
		HasPhoto: entry.HasPhoto,
		Role: &entry.Role,
		TagId: &entry.TagId,
	})

	if err != nil {
		// Log the error or return it to the client
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Success: false, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(newEntry)
}

// @Summary Create a child entry
// @Description Create a child entry
// @Tags entries
// @Accept json
// @Produce json
// @Param entry body childEntryRequestBody true "The entry's info"
// @Success 200 {object} entrySuccessResponse
// @Failure 400 {object} handlers.Response{Success: false, Message: string}
// @Failure 500 {object} handlers.Response{Success: false, Message: string}
// @Router /notebooks/{notebookId}/entries/{entryId}/children [post]
func (e *EntryHandler) CreateChildEntry(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := e.SessionManager.GetSession(sessionId)
	if err != nil {
		fmt.Println("kutas")
		return c.JSON(Response{Success: false, Message: "invalid session kutasie"})
	}

	// get the notebook id from the path
	notebookId, err := strconv.Atoi(c.Params("notebookId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the entry id from the path
	entryId, err := strconv.Atoi(c.Params("entryId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the entry info from the request body
	var entry childEntryRequestBody

	err = c.BodyParser(&entry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// validate the entry struct
	validate := validator.New()
	err = validate.Struct(entry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// create the entry
	newEntry, err := e.Storage.CreateChildEntry(user.Id, &storage.NewChildEntry{
		NotebookID: notebookId,
		Title: entry.Title,
		Content: entry.Content,
		Role: &entry.Role,
		ParentEntryId: &entryId,
		TagId: &entry.TagId,
	})

	if err != nil {
		// Log the error or return it to the client
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Success: false, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(newEntry)
}

// @Summary Get the children entreis for an entry
// @Description Get the children entreis for an entry
// @Tags entries
// @Accept json
// @Produce json
// @Param entryId path int true "The entry's id"
// @Success 200 {array} []storage.Entry
// @Failure 400 {object} handlers.Response{Success: false, Message: string}
// @Failure 500 {object} handlers.Response{Success: false, Message: string}
// @Router /notebooks/{notebookId}/entries/{entryId}/children [get]
func (e *EntryHandler) GetEntryChildren(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := e.SessionManager.GetSession(sessionId)
	if err != nil {
		fmt.Println("siema")
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	// get the notebook id from the path
	notebookId, err := strconv.Atoi(c.Params("notebookId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the entry id from the path
	entryId, err := strconv.Atoi(c.Params("entryId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the children entries
	children, err := e.Storage.GetEntryChildren(user.Id, notebookId, entryId)
	if err != nil {
		fmt.Println("elo")
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(children)
}

// @Summary Get entry by id
// @Description Get entry by id
// @Tags entries
// @Accept json
// @Produce json
// @Param entryId path int true "The entry's id"
// @Success 200 {object} storage.Entry
// @Failure 400 {object} handlers.Response{Success: false, Message: string}
// @Failure 500 {object} handlers.Response{Success: false, Message: string}
// @Router /notebooks/{notebookId}/entries/{entryId} [get]
func (e *EntryHandler) GetEntryById(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := e.SessionManager.GetSession(sessionId)
	if err != nil {
		fmt.Println("siema")
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	// get the entry id from the path
	entryId, err := strconv.Atoi(c.Params("entryId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the notebook id from the path
	notebookId, err := strconv.Atoi(c.Params("notebookId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the children entries
	entry, err := e.Storage.GetEntryById(user.Id, notebookId, entryId)
	if err != nil {
		fmt.Println("elo")
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(entry)
}

// @Summary Get the children entreis for an entry
// @Description Get the children entreis for an entry
// @Tags entries
// @Accept json
// @Produce json
// @Param entryId path int true "The entry's id"
// @Success 200 {array} []storage.Entry
// @Failure 400 {object} handlers.Response{Success: false, Message: string}
// @Failure 500 {object} handlers.Response{Success: false, Message: string}
// @Router /notebooks/{notebookId}/entries/{entryId}/photos [get]
func (e *EntryHandler) GetEntryPhotos(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := e.SessionManager.GetSession(sessionId)
	if err != nil {
		fmt.Println("siema")
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	// get the entry id from the path
	entryId, err := strconv.Atoi(c.Params("entryId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the children entries
	children, err := e.Storage.GetEntryPhotos(user.Id, entryId)
	if err != nil {
		fmt.Println("elo")
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(children)
}

// @Summary Create a photo
// @Description Create a photo
// @Tags entries
// @Accept json
// @Produce json
// @Param entryId path int true "The entry's id"
// @Param photo body photoRequestBody true "The photo's info"
// @Success 200 {object} photoSuccessResponse
// @Failure 400 {object} handlers.Response{Success: false, Message: string}
// @Failure 500 {object} handlers.Response{Success: false, Message: string}
// @Router /notebooks/{notebookId}/entries/{entryId}/photos [post]
func (e *EntryHandler) CreatePhoto(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := e.SessionManager.GetSession(sessionId)
	if err != nil {
		fmt.Println("siema")
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	// get the entry id from the path
	entryId, err := strconv.Atoi(c.Params("entryId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// get the photo info from the request body
	var photo photoRequestBody

	err = c.BodyParser(&photo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// validate the photo struct
	validate := validator.New()
	err = validate.Struct(photo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// create the photo
	id, err := e.Storage.CreatePhoto(entryId, &storage.NewPhoto{
		AuthorId: user.Id,
		ImageData: photo.ImageData,
		MimeType: photo.MimeType,
	})

	if err != nil {
		// Log the error or return it to the client
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Success: false, Message: err.Error()})
	}

	resp := photoSuccessResponse{
		Id: id,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}