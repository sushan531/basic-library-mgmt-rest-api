package handlers

import (
	"LMS/orm"
	"LMS/requestHandlers/models"
	"LMS/storage"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// DeleteBook deletes a book from the database based on the provided book_id parameter.
//
// Parameters:
//
//	c: echo.Context - the Echo context for the HTTP request
//
// Returns:
//
//	error - returns an error if the deletion process fails
func DeleteBook(c echo.Context) (err error) {
	// Get database instance
	db := storage.GetDBInstance()

	// Initialize ORM queries
	queries := orm.New(db)

	// Retrieve book_id from query parameter
	bookId := c.QueryParam("book_id")

	// Check if book_id is empty
	if bookId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing book_id parameter")
	}

	// Convert book_id to integer
	bookIdInt, _ := strconv.Atoi(bookId)

	// Delete the book from the database
	err = queries.DeleteBook(context.Background(), int32(bookIdInt))

	// Handle deletion error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Prepare response model
	respModel := models.DefaultResponseModel{Msg: "Book deleted successfully"}

	// Return JSON response
	return c.JSON(http.StatusOK, models.DataResponseModel{
		Data: respModel,
	})
}
