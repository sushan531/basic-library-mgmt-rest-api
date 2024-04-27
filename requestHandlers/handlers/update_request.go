package handlers

import (
	"LMS/orm"
	"LMS/requestHandlers/models"
	"LMS/storage"
	tls "LMS/translation"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

// UpdateBook updates a book based on the book_id query parameter
func UpdateBook(c echo.Context) error {
	// Get the database instance
	db := storage.GetDBInstance()

	// Extract the book_id query parameter and convert it to an integer
	bookId := c.QueryParam("book_id")
	bookIdInt, _ := strconv.Atoi(bookId)

	// Create a new ORM instance using the database
	queries := orm.New(db)

	// Create a new MiniBook instance
	book := new(models.MiniBook)

	// Bind the request body to the MiniBook instance
	if err := c.Bind(book); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Translate the MiniBook to a conditional update object
	tBook := tls.TranslateBookToConditionalUpdateObject(int32(bookIdInt), *book)

	// Perform a conditional update on the book and get the updated book
	updatedBook, err := queries.ConditionalUpdateBook(context.Background(), tBook)
	if err != nil {
		// Return an internal server error if there's an error during the update
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the updated book as JSON
	return c.JSON(http.StatusOK, models.DataResponseModel{
		Data: updatedBook,
	})
}

// ReturnBorrow handles the return of a borrowed book by a user
func ReturnBorrow(c echo.Context) (err error) {
	// Create a new MiniBorrowedBook instance
	borrow := new(models.MiniBorrowedBook)
	// Get a connection to the database
	db := storage.GetDBInstance()
	queries := orm.New(db)

	// Bind the request body to the borrow struct
	if err := c.Bind(borrow); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// Validate the borrow struct
	if err = c.Validate(borrow); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Get the user based on the provided email
	user, err := queries.GetUser(context.Background(), borrow.UserEmail)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Get the details of the book based on the provided title
	details, err := queries.GetBookDetails(context.Background(), borrow.BookTitle)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return echo.NewHTTPError(http.StatusNotFound, "Book not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Check if the user has already borrowed the book
	borrowed, err := queries.CheckIfUserBorrowedBook(context.Background(), orm.CheckIfUserBorrowedBookParams{
		UserID: user.UserID, BookID: details.BookID},
	)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return echo.NewHTTPError(http.StatusNotFound, "Book not borrowed")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Check if the book has already been returned
	if borrowed.ReturnedDate.Valid == true {
		return echo.NewHTTPError(http.StatusExpectationFailed, "Book already returned")
	}

	// Translate the borrow struct to a conditional update object
	tBorrow := tls.TranslateBorrowToConditionalUpdateObject(borrowed)

	// Update the borrowed book record in the database
	updatedBorrow, err := queries.ConditionalUpdateBorrowedBook(context.Background(), tBorrow)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the updated borrowed book as JSON response
	return c.JSON(http.StatusOK, models.DataResponseModel{
		Data: updatedBorrow,
	})
}
