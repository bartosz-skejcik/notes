package main

import (
	"context"
	"fmt"

	"notes/config"
	"notes/internal/auth"
	"notes/internal/database"
	"notes/internal/handlers"
	"notes/internal/middleware"
	"notes/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"

	_ "notes/docs"
)

//		@title			Go Notes API
//		@version		0.1
//		@description	This is an API for a note taking app
//		@securityDefinitions.apikey	ApiKeyAuth
//		@in							header
//		@name						Authorization
//		@description				Token in Bearer format to authenticate the user
//		@host		localhost:3000
//	    @BasePath	/api
func NewFiberServer(lc fx.Lifecycle, userHandlers *handlers.UserHandler, authMiddleware *middleware.AuthMiddleware) *fiber.App {

    // Initialize a new Fiber app
    app := fiber.New()

    // Enable CORS
    app.Use(cors.New())

    // Enable logger
    app.Use(logger.New())

    // Create a group for the API routes
    api := app.Group("/api")

    // Create a group for auth routes which is nested under the API group
    auth := api.Group("/auth")

    // Create routes for the auth endpoints
    auth.Post(("/sign-up"), userHandlers.SignUpUser)
	auth.Post("/sign-in", userHandlers.SignInUser)
	auth.Get("/me", authMiddleware.CheckAuth, userHandlers.GetUserInfo)
	auth.Post("/sign-out", userHandlers.SignOutUser)


    // Health check route
	app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "ok"})
	})

    app.Get("/swagger/*", swagger.HandlerDefault)


    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c *fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, World 👋!")
    })

    lc.Append(fx.Hook{
        OnStart: func(ctx context.Context) error {
            fmt.Println("Starting fiber server on port 3000")
            go app.Listen(":3000")
            return nil
        },
        OnStop: func(ctx context.Context) error {
            return app.Shutdown()
        },
    })

    return app
}

func main() {
    fx.New(
        fx.Provide(
            config.LoadEnvVars,
            database.CreatePostgresConnection,
            database.CreateRedisConnection,
            auth.NewSessionManager,
            handlers.NewUserHandler,
            storage.NewUserStorage,
            middleware.NewAuthMiddleware,
        ),
        fx.Invoke(NewFiberServer),
    ).Run()
}