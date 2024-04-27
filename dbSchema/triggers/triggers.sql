-- Trigger function to prevent_duplicate_borrow of same book by same user.
CREATE OR REPLACE FUNCTION prevent_duplicate_borrow()
    RETURNS TRIGGER AS
$$
BEGIN
    -- Check if a user is trying to borrow a book they have already borrowed and haven't returned yet
    IF EXISTS (SELECT 1
               FROM BorrowedBooks AS existing
               WHERE existing.user_id = NEW.user_id
                 AND existing.book_id = NEW.book_id
                 AND existing.returned_date IS NULL) THEN
        -- Raise an exception if a duplicate borrow is detected
        RAISE EXCEPTION 'User with ID: % cannot borrow book with ID: % again. Book not yet returned.',
            NEW.user_id, NEW.book_id;
    END IF;
    -- Return the new row (assuming no exception is raised)
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER prevent_duplicate_insert
    BEFORE INSERT
    ON BorrowedBooks
    FOR EACH ROW
EXECUTE PROCEDURE prevent_duplicate_borrow();


-- Trigger function to update_book_copies when a book is borrowed by a user
CREATE OR REPLACE FUNCTION update_book_copies()
    RETURNS TRIGGER AS
$$
BEGIN
    IF NEW.book_id IN ( -- Check for available copies
        SELECT book_id
        FROM Books
        WHERE book_id = NEW.book_id -- Only check for the specific book being borrowed
          AND num_copies <= 0) THEN
        RAISE EXCEPTION 'Cannot borrow book with ID: % - No copies available.', NEW.book_id;
    END IF;

    UPDATE Books -- Update only the specific book
    SET num_copies = num_copies - 1
    WHERE book_id = NEW.book_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER update_book_num_copies
    AFTER INSERT
    ON BorrowedBooks
    FOR EACH ROW
EXECUTE PROCEDURE update_book_copies();


-- Trigger function to update_book_copies_on_return when a book is returned by a user.
CREATE OR REPLACE FUNCTION update_book_copies_on_return()
    RETURNS TRIGGER AS
$$
BEGIN
    IF OLD.returned_date IS NULL AND NEW.returned_date IS NOT NULL THEN -- Check for returned book
        UPDATE Books
        SET num_copies = num_copies + 1
        WHERE book_id = OLD.book_id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER update_book_num_copies_on_return
    AFTER UPDATE
    ON BorrowedBooks
    FOR EACH ROW
    WHEN (NEW.returned_date IS NOT NULL) -- Additional check in trigger definition
EXECUTE PROCEDURE update_book_copies_on_return();



