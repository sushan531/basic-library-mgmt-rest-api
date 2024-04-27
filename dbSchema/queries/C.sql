-- name: InsertBook :one
INSERT INTO Books (title, authors, publisher, publication_date, genre, description, language, edition, num_copies,
                   location, image_url, keywords)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: InsertUser :one
INSERT INTO Users (firstname,lastname, email, contact_number, address)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: InsertBorrowedBook :one
INSERT INTO BorrowedBooks (user_id, book_id, borrowed_date, returned_date, due_date)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
