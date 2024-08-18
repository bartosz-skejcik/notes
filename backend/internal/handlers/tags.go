package handlers

import (
	"notes/internal/storage"

	"github.com/gofiber/fiber/v2"
)

type TagHandler struct {
	Storage *storage.TagStorage
}

func NewTagHandler(storage *storage.TagStorage) *TagHandler {
	return &TagHandler{
		Storage: storage,
	}
}

func (h *TagHandler) GetTags(c *fiber.Ctx) error {
	tags, err := h.Storage.GetAllTags()
	if err != nil {
		return err
	}

	return c.JSON(tags)
}