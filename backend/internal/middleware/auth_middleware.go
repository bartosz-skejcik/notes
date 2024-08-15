package middleware

import (
	"notes/internal/auth"
	"notes/internal/storage"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	Storage *storage.UserStorage
	SessionManager *auth.SessionManager
}

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

func NewAuthMiddleware(storage *storage.UserStorage, sessionManager *auth.SessionManager) *AuthMiddleware {
	return &AuthMiddleware{
		Storage: storage,
		SessionManager: sessionManager,
	}
}

func (a *AuthMiddleware) CheckAuth(c *fiber.Ctx) error {
	sessionHeader := c.Get("Authorization")

	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: "invalid session header"})
	}

	sessionId := sessionHeader[7:]

	_, err := a.SessionManager.GetSession(sessionId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Success: false, Message: err.Error()})
	}

	return c.Next()
}