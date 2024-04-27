-- name: DeleteBook :exec
DELETE
FROM Books
WHERE book_id = $1
RETURNING *;


-- name: DeleteUser :exec
DELETE
FROM Users
WHERE user_id = $1;

-- name: DeleteBorrowedBook :exec
DELETE
FROM BorrowedBooks
WHERE borrowed_id = $1;

-- name: TruncateDB :exec
TRUNCATE TABLE Books, Users, BorrowedBooks;