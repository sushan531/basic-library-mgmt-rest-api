-- name: ConditionalUpdateBook :one
UPDATE Books
SET title            = CASE
                           WHEN $1::INT = 1 THEN $2
                           ELSE title
    END,
    authors          = CASE
                           WHEN $3::INT = 1 THEN $4
                           ELSE authors
        END,
    publisher        = CASE
                           WHEN $5::INT = 1 THEN $6
                           ELSE publisher
        END,
    publication_date = CASE
                           WHEN $7::INT = 1 THEN $8
                           ELSE publication_date
        END,
    genre            = CASE
                           WHEN $9::INT = 1 THEN $10
                           ELSE genre
        END,
    description      = CASE
                           WHEN $11::INT = 1 THEN $12
                           ELSE description
        END,
    language         = CASE
                           WHEN $13::INT = 1 THEN $14
                           ELSE language
        END,
    edition          = CASE
                           WHEN $15::INT = 1 THEN $16
                           ELSE edition
        END,
    num_copies       = CASE
                           WHEN $17::INT = 1 THEN $18
                           ELSE num_copies
        END,
    location         = CASE
                           WHEN $19::INT = 1 THEN $20
                           ELSE location
        END,
    image_url        = CASE
                           WHEN $21::INT = 1 THEN $22
                           ELSE image_url
        END,
    keywords         = CASE
                           WHEN $23::INT = 1 THEN $24
                           ELSE keywords
        END
WHERE book_id = $25
RETURNING *;


-- name: ConditionalUpdateUser :one
UPDATE Users
SET firstname       = CASE
                         WHEN $1::INT = 1 THEN $2
                         ELSE firstname
    END,
    lastname       = CASE
                         WHEN $3::INT = 1 THEN $4
                         ELSE lastname
        END,
    email          = CASE
                         WHEN $5::INT = 1 THEN $6
                         ELSE email
        END,
    contact_number = CASE
                         WHEN $7::INT = 1 THEN $8
                         ELSE contact_number
        END,
    address        = CASE
                         WHEN $9::INT = 1 THEN $10
                         ELSE address
        END
WHERE user_id = $11
RETURNING *;


-- name: ConditionalUpdateBorrowedBook :one
UPDATE BorrowedBooks
SET returned_date = CASE
                        WHEN $1::INT = 1 THEN $2
                        ELSE returned_date
    END,
    due_date      = CASE
                        WHEN $3::INT = 1 THEN $4
                        ELSE due_date
        END
WHERE borrowed_id = $5
RETURNING *;


