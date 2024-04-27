package handlers

import (
	"LMS/orm"
	"LMS/requestHandlers/models"
	"LMS/storage"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ListBooks retrieves a list of books from the database and returns them as JSON.
//
// It takes a context as a parameter and returns an error.
func ListBooks(c echo.Context) error {
	// Get a database instance
	db := storage.GetDBInstance()

	// Create ORM queries
	queries := orm.New(db)

	// Retrieve all books from the database
	books, err := queries.ListAllBooks(context.Background())
	if err != nil {
		// Return internal server error if there's an error
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the list of books as JSON
	return c.JSON(http.StatusOK, models.DataResponseModel{
		Data: books,
	})
}

// SearchBookByTitleOrAuthor searches books by title or author.
// It takes an echo Context as input and returns an error.
func SearchBookByTitleOrAuthor(c echo.Context) (err error) {
	// Get database instance
	db := storage.GetDBInstance()

	// Initialize ORM queries
	queries := orm.New(db)

	// Get search text from query parameter
	searchText := c.QueryParam("search-text")

	// Search books based on search text
	books, err := queries.SearchBooks(context.Background(), searchText)
	if err != nil {
		// Return internal server error if search fails
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return JSON response with search results
	return c.JSON(http.StatusOK, models.DataResponseModel{
		Data: books,
	})
}

// ListUsers retrieves a list of users from the database and returns them as JSON.
//
// It takes an echo Context as input and returns an error.
func ListUsers(c echo.Context) (err error) {
	// Get database instance
	db := storage.GetDBInstance()

	// Initialize ORM queries
	queries := orm.New(db)

	// Retrieve all users from the database
	users, err := queries.ListAllUsers(context.Background())
	if err != nil {
		// Return internal server error if there's an error
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the list of users as JSON
	return c.JSON(http.StatusOK, models.DataResponseModel{
		Data: users,
	})
}

// CountRemainingBooks counts the remaining books in the database and returns the count as JSON.
//
// Parameters:
//
//	c: echo.Context - the Echo context for the HTTP request
//
// Returns:
//
//	error - returns an error if the counting process fails
func CountRemainingBooks(c echo.Context) (err error) {
	// Get database instance
	db := storage.GetDBInstance()

	// Create ORM queries
	queries := orm.New(db)

	// Count the remaining books
	count, err := queries.CountRemainingBooks(context.Background())
	if err != nil {
		// Return internal server error if there's an error
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the count of remaining books as JSON
	return c.JSON(http.StatusOK, models.DataResponseModel{Data: count})
}

// ListBorrow retrieves a list of all borrowed books from the database and returns it as a JSON response.
func ListBorrow(c echo.Context) error {
	// Get a database instance
	db := storage.GetDBInstance()

	// Create queries using the ORM
	queries := orm.New(db)

	// Retrieve the list of borrowed books
	borrow, err := queries.ListAllBorrowedBooks(context.Background())
	if err != nil {
		// Return an internal server error if there's an error
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the list of borrowed
	return c.JSON(http.StatusOK, models.DataResponseModel{
		Data: borrow,
	})
}
