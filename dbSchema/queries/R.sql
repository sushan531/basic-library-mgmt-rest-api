-- name: SearchBooks :many
SELECT *
FROM Books
WHERE LOWER(title) LIKE LOWER($1)
   OR LOWER(authors) LIKE LOWER($1);

-- name: ListAllBooks :many
SELECT *
FROM Books;

-- name: ListAllBooksNext :many
SELECT *
FROM Books
WHERE (
          book_id > $1
          )
order by book_id ASC
LIMIT $2;

-- name: ListAllBooksPrev :many
SELECT *
FROM Books
WHERE (
          book_id >= $1 AND book_id <= $2
          )
order by book_id ASC
LIMIT $3;

-- name: GetBookDetails :one
SELECT book_id, title, authors, genre
FROM Books
WHERE title = $1;


-- name: ListAllUsers :many
SELECT *
FROM Users;

-- name: ListUsersIds :many
SELECT user_id, email
FROM Users;

-- name: GetUser :one
SELECT *
FROM Users
WHERE LOWER(email) = LOWER($1);

-- name: ListAllBorrowedBooks :many
SELECT *
FROM BorrowedBooks;

-- name: CheckIfUserBorrowedBook :one
SELECT *
FROM BorrowedBooks
WHERE user_id = $1
  AND book_id = $2
  AND returned_date IS NULL;

-- name: CountRemainingBooks :one
SELECT SUM(num_copies)
FROM Books;