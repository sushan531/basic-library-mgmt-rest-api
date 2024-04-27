package main

import (
	"LMS/config"
	"LMS/requestHandlers/handlers"
	"LMS/requestHandlers/validators"
	"LMS/storage"
	"github.com/labstack/echo/v4"
	"gopkg.in/bluesuncorp/validator.v9"
	"net/http"
)

// initiateApp initializes the Echo application, configures routes, sets up database connections, and returns the fully configured Echo app.
//
// No parameters.
// Returns *echo.Echo.
func initiateApp() *echo.Echo {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to LMS, a cool Library Management System!")
	})
	cfg, err := config.ReadConfig("test")
	if err != nil {
		panic(err)
	}
	err = storage.CreateDB(cfg.PGDATABASE, cfg.Host, cfg.User, cfg.Password, cfg.Port)
	if err != nil {
		panic(err)
	}
	storage.DB, err = storage.NewDB(cfg.PGDATABASE, cfg.Host, cfg.User, cfg.Password, cfg.Port)
	if err != nil {
		panic(err)
	}
	_, err = storage.InitializeDB(storage.DB)
	if err != nil {
		panic(err)
	}
	app.Validator = &validators.CustomValidator{Validator: validator.New()}
	return app
}

func runServer() {
	app := initiateApp()

	// Book related endpoints in "books" group
	books := app.Group("/books")
	books.POST("/add", handlers.AddBook)
	books.GET("/list", handlers.ListBooks)
	books.GET("/list-paged", handlers.ListAllBooksPaginated)
	books.GET("/search", handlers.SearchBookByTitleOrAuthor)
	books.POST("/update", handlers.UpdateBook)
	books.DELETE("/delete", handlers.DeleteBook)
	books.GET("/count", handlers.CountRemainingBooks)

	// Users related endpoints in "users" group
	users := app.Group("/users")
	users.POST("/add", handlers.AddUser)
	users.GET("/list", handlers.ListUsers)

	// Borrow related endpoints in "borrow" group
	borrow := app.Group("/borrow")
	borrow.POST("/add", handlers.AddBorrow)
	borrow.GET("/list", handlers.ListBorrow)
	borrow.POST("/return", handlers.ReturnBorrow)

	app.POST("/truncate", handlers.TruncateDB)

	// Listen Server in 0.0.0.0:8000
	app.Logger.Fatal(app.Start(":8000"))
}

// main initiates the application, sets up different endpoint groups for books, users, and borrow,
// and starts the server on port 1323.
//
// No parameters.
// No return values.
func main() {
	runServer()
}
