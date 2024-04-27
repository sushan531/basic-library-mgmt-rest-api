CREATE TABLE IF NOT EXISTS Books
(
    book_id          SERIAL PRIMARY KEY,
    title            VARCHAR(255) NOT NULL UNIQUE,
    authors          VARCHAR(255) NOT NULL,
    publisher        VARCHAR(255),
    publication_date DATE,
    genre            VARCHAR(100),
    description      TEXT,
    language         VARCHAR(50),
    edition          INT,
    num_copies       INT,
    location         VARCHAR(100),
    image_url        VARCHAR(255),
    keywords         VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS Users
(
    user_id        SERIAL PRIMARY KEY,
    email          VARCHAR(255) NOT NULL UNIQUE,
    firstname      VARCHAR(255) NOT NULL,
    lastname       VARCHAR(255) NOT NULL,
    contact_number VARCHAR(50),
    address        TEXT         NOT NULL
);

CREATE TABLE IF NOT EXISTS BorrowedBooks
(
    borrowed_id   SERIAL PRIMARY KEY,
    user_id       INT  NOT NULL REFERENCES Users (user_id) ON DELETE CASCADE,
    book_id       INT  NOT NULL REFERENCES Books (book_id) ON DELETE CASCADE,
    borrowed_date DATE NOT NULL,
    returned_date DATE,
    due_date      DATE
);
