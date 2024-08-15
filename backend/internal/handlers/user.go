package handlers

import (
	"fmt"
	"notes/internal/auth"
	"notes/internal/storage"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	Storage *storage.UserStorage
	SessionManager *auth.SessionManager
}

func NewUserHandler(storage *storage.UserStorage, sessionManager *auth.SessionManager) *UserHandler {
	return &UserHandler{
		Storage: storage,
		SessionManager: sessionManager,
	}
}

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

type signInRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type userRequestBody struct {
	DisplayName string `json:"display_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

type signUpSuccessResponse struct {
	Id int `json:"id"`
}

// @Summary Sign up a user
// @Description Sign up a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body userRequestBody true "The user's first name, last name, email, and password"
// @Success 200 {object} signUpSuccessResponse
// @Failure 400 {object} handlers.SuccessResponse{Success: false}
// @Failure 500 {object} handlers.SuccessResponse{Success: false}
// @Router /auth/sign-up [post]
func (u *UserHandler) SignUpUser(c *fiber.Ctx) error {
	// get the info from the request body
	var user userRequestBody

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// validate the user struct
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Success: false, Message: err.Error()})
	}

	// create the user
	id, err := u.Storage.CreateUser(&storage.NewUser{
		DisplayName: user.DisplayName,
		Email:     user.Email,
		Password:  string(hashedPassword),
	})

	if err != nil {
		// Log the error or return it to the client
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Success: false, Message: err.Error()})
	}

	// generate the user's session
	sessionId, err := u.SessionManager.CreateSession(&auth.UserSession{
		Id:        id,
		DisplayName: user.DisplayName,
		Email:     user.Email,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Success: false, Message: err.Error()})
	}

	// set the session id as a header
	c.Response().Header.Set("Authorization", fmt.Sprintf("Bearer %s", sessionId))

	resp := signUpSuccessResponse{
		Id: id,
	}
	return c.JSON(resp)
}

// @Summary Sign in a user
// @Description Sign in a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body signInRequestBody true "The user's email and password"
// @Success 200 {object} handlers.SuccessResponse
// @Failure 400 {object} handlers.SuccessResponse{Success: false}
// @Failure 500 {object} handlers.SuccessResponse{Success: false}
// @Router /auth/sign-in [post]
func (handler *UserHandler) SignInUser(c *fiber.Ctx) error {
	var user signInRequestBody

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	fmt.Println(user)

	// validate the user struct
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: err.Error()})
	}

	// sign the user in
	sessionId, err := handler.SessionManager.SignIn(user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Success: false, Message: err.Error()})
	}

	// set the session id as a header
	c.Response().Header.Set("Authorization", fmt.Sprintf("Bearer %d", sessionId.Id))

	return c.JSON(Response{Success: true, Message: "Successfully signed in"})
}

// @Summary Sign out a user
// @Description Sign out a user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} handlers.SuccessResponse
// @Failure 400 {object} handlers.SuccessResponse{Success: false}
// @Failure 500 {object} handlers.SuccessResponse{Success: false}
// @Router /auth/sign-out [post]
func (u *UserHandler) SignOutUser(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// ensure the session header is not empty and in the correct format
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Message: "invalid session header"})
	}

	// get the session id
	sessionId := sessionHeader[7:]

	// delete the session
	err := u.SessionManager.SignOut(sessionId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Success: false, Message: err.Error()})
	}

	return c.JSON(Response{Success: true, Message: "Successfully signed out"})
}

// @Summary Get the user's info
// @Description Get the user's info
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} auth.UserSession
// @Failure 400 {object} handlers.SuccessResponse{Success: false}
// @Failure 500 {object} handlers.SuccessResponse{Success: false}
// @Router /auth/me [get]
func (u *UserHandler) GetUserInfo(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// get the session id
	sessionId := sessionHeader[7:]

	// get the user data from the session
	user, err := u.SessionManager.GetSession(sessionId)
	if err != nil {
		return c.JSON(Response{Success: false, Message: err.Error()})
	}

	return c.JSON(user)
}