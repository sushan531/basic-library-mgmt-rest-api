package models

import "time"

type MiniUser struct {
	FirstName     string `json:"firstname" validate:"required"`
	LastName      string `json:"lastname" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	ContactNumber string `json:"contactNumber"`
	Address       string `json:"address" validate:"required"`
}

type MiniBook struct {
	Title           string    `json:"title" validate:"required"`
	Author          string    `json:"author" validate:"required"`
	Publisher       string    `json:"publisher" validate:"required"`
	PublicationDate time.Time `json:"publicationDate"`
	Genre           string    `json:"genre"`
	Description     string    `json:"description"`
	Language        string    `json:"language"`
	Edition         int32     `json:"edition"`
	NumCopies       int32     `json:"numCopies" validate:"required,gte=0"`
	Location        string    `json:"location"`
	ImageURL        string    `json:"imageURL" validate:"required,url"`
	Keywords        string    `json:"keywords"`
}

type MiniBorrowedBook struct {
	UserEmail    string    `json:"userEmail"  validate:"required,email"`
	BookTitle    string    `json:"bookTitle" validate:"required"`
	BorrowedDate time.Time `json:"borrowedDate"`
	DueDate      time.Time `json:"dueDate"`
}

type MiniReturnBook struct {
	UserEmail    uint      `json:"userEmail"  validate:"required,email"`
	BookTitle    uint      `json:"title" validate:"required"`
	ReturnedDate time.Time `json:"returnedDate" validate:"required"`
}

type DefaultResponseModel struct {
	Msg string `json:"message"`
}

type DataResponseModel struct {
	Data interface{} `json:"data"`
}

type PaginatedResponse struct {
	PreviousToken string `json:"previous_token"`
	NextToken     string `json:"next_token"`
	Data          any    `json:"data"`
}
