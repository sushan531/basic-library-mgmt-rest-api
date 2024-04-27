package handlers

import (
	"LMS/orm"
	"LMS/requestHandlers/models"
	"LMS/storage"
	tls "LMS/translation"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// AddBook is a function that adds a book to the database.
//
// It takes in a context object and returns an error.
func AddBook(c echo.Context) (err error) {
	// Create a new MiniBook model
	book := new(models.MiniBook)

	// Get the database instance
	db := storage.GetDBInstance()
	queries := orm.New(db)

	// Bind the request payload to the book model
	if err := c.Bind(book); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Validate the book model
	if err = c.Validate(book); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Translate the book model to a database object
	tBook := tls.TranslateBookToDBObject(*book)

	// Insert the book into the database
	insertedBook, err := queries.InsertBook(context.Background(), tBook)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the inserted book as JSON response
	return c.JSON(http.StatusOK, models.DataResponseModel{
		Data: insertedBook,
	})
}

// AddUser is a function to add a new user to the database
func AddUser(c echo.Context) (err error) {
	// Create a new MiniUser model
	user := new(models.MiniUser)
	// Get the database instance
	db := storage.GetDBInstance()
	// Create ORM queries instance
	queries := orm.New(db)

	// Bind the request data to the user model
	if err := c.Bind(user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Validate the user model
	if err = c.Validate(user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Translate MiniUser to database object
	tUser := tls.TranslateUserToDBObject(*user)

	// Insert the user into the database
	insertedUser, err := queries.InsertUser(context.Background(), tUser)
	if err != nil {
		// Check for duplicate key error
		if strings.Contains(err.Error(), "duplicate key") {
			return echo.NewHTTPError(http.StatusInternalServerError, "User already exists")
		}
		// Return internal server error for other errors
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the inserted user as JSON response
	return c.JSON(http.StatusOK, models.DataResponseModel{
		Data: insertedUser,
	})
}

// AddBorrow adds a borrowed book to the database
func AddBorrow(c echo.Context) error {
	// Create a new MiniBorrowedBook instance
	borrow := new(models.MiniBorrowedBook)
	// Get the database instance
	db := storage.GetDBInstance()
	// Create ORM queries
	queries := orm.New(db)

	// Bind the request body to the borrow instance
	if err := c.Bind(borrow); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Validate the borrow instance
	if err := c.Validate(borrow); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Get the user from the database
	user, err := queries.GetUser(context.Background(), borrow.UserEmail)
	if err != nil {
		return err
	}

	// Get the book details from the database
	details, err := queries.GetBookDetails(context.Background(), borrow.BookTitle)
	if err != nil {
		return err
	}

	// Translate the borrow instance to a database object
	tBorrow := tls.TranslateBorrowToDBObject(*borrow, user, details)

	// Insert the borrowed book into the database
	insertedBorrow, err := queries.InsertBorrowedBook(context.Background(), tBorrow)
	if err != nil {
		// Handle duplicate key error
		if strings.Contains(err.Error(), "duplicate key") {
			return echo.NewHTTPError(http.StatusExpectationFailed, "Book already borrowed")
		}
		// Handle internal server error
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the inserted borrow as JSON response
	return c.JSON(http.StatusOK, models.DataResponseModel{
		Data: insertedBorrow,
	})
}

func TruncateDB(c echo.Context) error {
	db := storage.GetDBInstance()
	// Create ORM queries
	queries := orm.New(db)
	err := queries.TruncateDB(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "Database truncated")
}
