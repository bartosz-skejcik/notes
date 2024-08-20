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
func NewFiberServer(lc fx.Lifecycle, userHandlers *handlers.UserHandler, authMiddleware *middleware.AuthMiddleware, notebookHandlers *handlers.NotebookHandler, entryHandlers *handlers.EntryHandler, tagHandlers *handlers.TagHandler) *fiber.App {

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

    notebooks := api.Group("/notebooks")
    notebooks.Get("", authMiddleware.CheckAuth, notebookHandlers.GetNotebooks)
    notebooks.Post("", authMiddleware.CheckAuth, notebookHandlers.CreateNotebook)
    notebooks.Get("/:notebookId", authMiddleware.CheckAuth, notebookHandlers.GetNotebook)
    notebooks.Delete("/:notebookId", authMiddleware.CheckAuth, notebookHandlers.DeleteNotebook)

    // entries
    notebooks.Get("/:notebookId/entries", authMiddleware.CheckAuth, entryHandlers.GetEntries)
    notebooks.Post("/:notebookId/entries", authMiddleware.CheckAuth, entryHandlers.CreateEntry)
    notebooks.Patch("/:notebookId/entries/:entryId", authMiddleware.CheckAuth, entryHandlers.UpdateEntry)
    notebooks.Delete("/:notebookId/entries/:entryId", authMiddleware.CheckAuth, entryHandlers.DeleteEntry)
    notebooks.Post("/:notebookId/entries/:entryId/children", authMiddleware.CheckAuth, entryHandlers.CreateChildEntry)
    notebooks.Patch("/:notebookId/entries/:entryId/children/:childEntryId", authMiddleware.CheckAuth, entryHandlers.UpdateChildEntry)
    notebooks.Get("/:notebookId/entries/:entryId", authMiddleware.CheckAuth, entryHandlers.GetEntryById)
    notebooks.Get("/:notebookId/entries/:entryId/children", authMiddleware.CheckAuth, entryHandlers.GetEntryChildren)
    notebooks.Get("/:notebookId/entries/:entryId/photos", authMiddleware.CheckAuth, entryHandlers.GetEntryPhotos)
    notebooks.Post("/:notebookId/entries/:entryId/photos", authMiddleware.CheckAuth, entryHandlers.CreatePhoto)


    // tags
    tags := api.Group("/tags")
    tags.Get("", authMiddleware.CheckAuth, tagHandlers.GetTags)


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
            // config
            config.LoadEnvVars,
            // database
            database.CreatePostgresConnection,
            database.CreateRedisConnection,
            // auth
            auth.NewSessionManager,
            middleware.NewAuthMiddleware,
            // users
            handlers.NewUserHandler,
            storage.NewUserStorage,
            // notebooks
            storage.NewNotebookStorage,
            handlers.NewNotebookHandler,
            // entries
            storage.NewEntryStorage,
            handlers.NewEntryHandler,
            // tags
            storage.NewTagStorage,
            handlers.NewTagHandler,
        ),
        fx.Invoke(NewFiberServer),
    ).Run()
}