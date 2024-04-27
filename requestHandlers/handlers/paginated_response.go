package handlers

import (
	"LMS/orm"
	"LMS/requestHandlers/models"
	"LMS/requestHandlers/utils"
	"LMS/storage"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func ListAllBooksPaginated(c echo.Context) error {
	// Get a database instance
	db := storage.GetDBInstance()
	queries := orm.New(db)
	limitStr := c.QueryParam("limit")
	limit, _ := strconv.Atoi(limitStr)
	limit32 := int32(limit)
	token := c.QueryParam("token")
	if token != "" {
		command, startId, endId, err := utils.DecodePaginationCursor(token)
		if err != nil {
			return c.JSON(http.StatusOK, models.PaginatedResponse{
				NextToken:     "",
				PreviousToken: "",
				Data:          "",
			})
		} else {
			switch command {
			case utils.PageTypeNextPage:
				params := orm.ListAllBooksNextParams{
					BookID: startId,
					Limit:  limit32,
				}
				asc, _ := queries.ListAllBooksNext(context.Background(), params)
				if len(asc) >= 1 {
					first := asc[0]
					last := asc[len(asc)-1]
					return c.JSON(http.StatusOK, models.PaginatedResponse{
						NextToken:     utils.EncodePaginationCursor(utils.PageTypeNextPage, last.BookID, 000000000),
						PreviousToken: utils.EncodePaginationCursor(utils.PageTypePreviousPage, first.BookID-limit32, last.BookID),
						Data:          asc,
					})
				} else {
					return c.JSON(http.StatusOK, models.PaginatedResponse{
						// todo Even if the data is empty there should be a cursor pointing to the previous page.
						// This will be done by fetching the (largest primary key) = `end` and (largest primary key - limit) = `start`
					})
				}
			case utils.PageTypePreviousPage:
				params := orm.ListAllBooksPrevParams{
					BookID:   startId,
					BookID_2: endId,
					Limit:    limit32,
				}
				desc, _ := queries.ListAllBooksPrev(context.Background(), params)
				first := desc[0]
				last := desc[len(desc)-1]
				previousToken := ""
				if first.BookID-limit32 <= 1000 {
					previousToken = ""
				} else {
					previousToken = utils.EncodePaginationCursor(utils.PageTypePreviousPage, first.BookID-limit32, first.BookID)
				}
				return c.JSON(http.StatusOK, models.PaginatedResponse{
					NextToken:     utils.EncodePaginationCursor(utils.PageTypeNextPage, last.BookID, 000000000),
					PreviousToken: previousToken,
					Data:          desc,
				})
			}
		}
	}
	params := orm.ListAllBooksNextParams{
		BookID: 0,
		Limit:  limit32,
	}
	asc, _ := queries.ListAllBooksNext(context.Background(), params)
	if len(asc) == 0 {
		return c.JSON(http.StatusOK, models.PaginatedResponse{})
	}
	last := asc[len(asc)-1]
	return c.JSON(http.StatusOK, models.PaginatedResponse{
		NextToken:     utils.EncodePaginationCursor(utils.PageTypeNextPage, last.BookID, 000000000),
		PreviousToken: "",
		Data:          asc,
	})
}
