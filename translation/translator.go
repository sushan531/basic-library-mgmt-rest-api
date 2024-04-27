package translation

import (
	"LMS/orm"
	"LMS/requestHandlers/models"
	"database/sql"
	"time"
)

// TranslateBookToDBObject maps a MiniBook to an InsertBookParams.
//
// Takes a MiniBook as input and returns an InsertBookParams.
func TranslateBookToDBObject(book models.MiniBook) orm.InsertBookParams {
	return orm.InsertBookParams{
		Title:           book.Title,
		Authors:         book.Author,
		Publisher:       sql.NullString{String: book.Publisher, Valid: book.Publisher != ""},
		PublicationDate: sql.NullTime{Time: book.PublicationDate, Valid: !book.PublicationDate.IsZero()},
		Genre:           sql.NullString{String: book.Genre, Valid: book.Genre != ""},
		Description:     sql.NullString{String: book.Description, Valid: book.Description != ""},
		Language:        sql.NullString{String: book.Language, Valid: book.Language != ""},
		Edition:         sql.NullInt32{Int32: book.Edition, Valid: book.Edition != 0},
		NumCopies:       sql.NullInt32{Int32: book.NumCopies, Valid: book.NumCopies != 0},
		Location:        sql.NullString{String: book.Location, Valid: book.Location != ""},
		Keywords:        sql.NullString{String: book.Keywords, Valid: book.Keywords != ""},
		ImageUrl:        sql.NullString{String: book.ImageURL, Valid: book.ImageURL != ""},
	}
}

// TranslateBookToConditionalUpdateObject translates a MiniBook struct into a ConditionalUpdateBookParams struct
// based on certain rules for updating each field. It returns the ConditionalUpdateBookParams object.
func TranslateBookToConditionalUpdateObject(bookId int32, book models.MiniBook) orm.ConditionalUpdateBookParams {
	// Initialize the updateObject with empty values
	updateObject := orm.ConditionalUpdateBookParams{}

	// Set the BookID directly
	updateObject.BookID = bookId

	// Update the Title field if it's not empty
	if book.Title != "" {
		updateObject.Title = book.Title
		updateObject.Column1 = 1
	}

	// Update the Authors field if it's not empty
	if book.Author != "" {
		updateObject.Authors = book.Author
		updateObject.Column3 = 1
	}

	// Update the Publisher field if it's not empty
	if book.Publisher != "" {
		updateObject.Publisher = sql.NullString{String: book.Publisher, Valid: true}
		updateObject.Column5 = 1
	}

	// Update the PublicationDate field if it's not the zero value
	if book.PublicationDate != (time.Time{}) {
		updateObject.PublicationDate = sql.NullTime{Time: book.PublicationDate, Valid: true}
		updateObject.Column7 = 1
	}

	// Update the Genre field if it's not empty
	if book.Genre != "" {
		updateObject.Genre = sql.NullString{String: book.Genre, Valid: true}
		updateObject.Column9 = 1
	}

	// Update the Description field if it's not empty
	if book.Description != "" {
		updateObject.Description = sql.NullString{String: book.Description, Valid: true}
		updateObject.Column11 = 1
	}

	// Update the Language field if it's not empty
	if book.Language != "" {
		updateObject.Language = sql.NullString{String: book.Language, Valid: true}
		updateObject.Column13 = 1
	}

	// Update the Edition field if it's not the zero value
	if book.Edition != 0 {
		updateObject.Edition = sql.NullInt32{Int32: book.Edition, Valid: true}
		updateObject.Column15 = 1
	}

	// Update the NumCopies field if it's not the zero value
	if book.NumCopies != 0 {
		updateObject.NumCopies = sql.NullInt32{Int32: book.NumCopies, Valid: true}
		updateObject.Column17 = 1
	}

	// Update the Location field if it's not empty
	if book.Location != "" {
		updateObject.Location = sql.NullString{String: book.Location, Valid: true}
		updateObject.Column19 = 1
	}

	// Update the Keywords field if it's not empty
	if book.Keywords != "" {
		updateObject.Keywords = sql.NullString{String: book.Keywords, Valid: true}
		updateObject.Column23 = 1
	}

	// Update the ImageUrl field if it's not empty
	if book.ImageURL != "" {
		updateObject.ImageUrl = sql.NullString{String: book.ImageURL, Valid: true}
		updateObject.Column21 = 1
	}

	return updateObject
}

// TranslateUserToDBObject maps a MiniUser to an InsertUserParams.
//
// Takes a MiniUser as input and returns an InsertUserParams.
func TranslateUserToDBObject(user models.MiniUser) orm.InsertUserParams {
	insertUserParams := orm.InsertUserParams{
		Email:     user.Email,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Address:   user.Address,
	}

	// Update the ContactNumber field if it's not empty
	if user.ContactNumber != "" {
		insertUserParams.ContactNumber = sql.NullString{String: user.ContactNumber, Valid: true}
	}

	return insertUserParams
}

// TranslateBorrowToDBObject takes a MiniBorrowedBook and user details and translates them into an InsertBorrowedBookParams object.
func TranslateBorrowToDBObject(borrow models.MiniBorrowedBook, user orm.User, borrowedBook orm.GetBookDetailsRow) orm.InsertBorrowedBookParams {
	// If the BorrowedDate is not set, default it to the current time in UTC
	if borrow.BorrowedDate.IsZero() {
		borrow.BorrowedDate = time.Now().UTC()
	}
	// Create and return an InsertBorrowedBookParams object with the translated values
	return orm.InsertBorrowedBookParams{
		UserID:       user.UserID,
		BookID:       borrowedBook.BookID,
		BorrowedDate: borrow.BorrowedDate,
		DueDate:      sql.NullTime{Time: borrow.BorrowedDate.AddDate(0, 2, 0), Valid: true},
	}
}

// TranslateBorrowToConditionalUpdateObject translates a BorrowedBook struct into a ConditionalUpdateBorrowedBookParams struct
// based on certain rules for updating each field. It returns the ConditionalUpdateBorrowedBookParams object.
func TranslateBorrowToConditionalUpdateObject(borrowed orm.Borrowedbook) orm.ConditionalUpdateBorrowedBookParams {
	// Initialize the updateObject with empty values
	updateObject := orm.ConditionalUpdateBorrowedBookParams{}

	// Set the BorrowedID directly
	updateObject.BorrowedID = borrowed.BorrowedID
	updateObject.Column1 = 1

	// Update the ReturnedDate field based on its validity
	if borrowed.ReturnedDate.Valid {
		updateObject.ReturnedDate = sql.NullTime{Time: borrowed.ReturnedDate.Time, Valid: true}
	} else {
		updateObject.ReturnedDate = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	}

	// Update the DueDate field based on its validity
	if borrowed.DueDate.Valid {
		updateObject.DueDate = sql.NullTime{Time: borrowed.DueDate.Time, Valid: true}
		updateObject.Column3 = 1
	}

	return updateObject
}
