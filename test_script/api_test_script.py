import json
import pathlib
import requests
from tabulate import tabulate

URL = "http://0.0.0.0:8000"

headers = {
    'Content-Type': 'application/json'
}
DELETE_ID = ""


def connect_to_postgresql_and_truncate():
    """
    This function connects to the PostgreSQL database and truncates the data.
    """
    # Print a message indicating truncating the database
    print("\nTruncating database".upper())

    # Make a POST request to truncate the database
    requests.post("{0}/truncate".format(URL), headers=headers)

def read_json_data(data_file):
    """
    Read and load JSON data from a file.

    Args:
        data_file (str): The name of the JSON file.

    Returns:
        dict: The loaded JSON data.
    """
    root_path = pathlib.Path(__file__).parent.absolute()  # Get the absolute path of the directory containing this script
    data_path = root_path.joinpath(data_file)  # Construct the full path to the data file
    with open(data_path) as json_file:  # Open the data file
        data = json.load(json_file)["data"]  # Load the JSON data and extract the 'data' key
        return data  # Return the loaded data

import requests
import json

def insert_books_to_library(new_books="books.json"):
    """
    Insert new books into the library using an API call.

    Args:
    new_books (str): Path to the JSON file containing the new books data. Default is "books.json".
    """
    # Construct the URL for adding books
    url = f"{URL}/books/add"

    # Print a message indicating the API call
    print(f"\nCalling the API to insert books {url}")

    # Read the new books data from the JSON file
    books_data_list = read_json_data(new_books)

    # Iterate through each book and make a POST request to add it to the library
    for book in books_data_list:
        requests.post(url, headers=headers, data=json.dumps(book))

def insert_users_to_library():
    """
    Insert new users into the library using an API call.
    """
    url = f"{URL}/users/add"
    print(f"\nCalling the API to insert users {url}")  # Print a message indicating the API call

    # Read the new users data from the JSON file
    users_data_list = read_json_data("users.json")

    # Iterate through each user and make a POST request to add it to the library
    for user in users_data_list:
        requests.post(url, headers=headers, data=json.dumps(user))

def normalize_books_data(books_list):
    """
    Normalize the data in the books_list by extracting specific values from nested dictionaries.

    Args:
    books_list (list of dict): A list of dictionaries containing book information.

    Returns:
    list of dict: A list of dictionaries with normalized book information.
    """
    new_books_normalised = []

    for book in books_list:
        book["publisher"] = book["publisher"]["String"]
        book["publication_date"] = book["publication_date"]["Time"]
        book["genre"] = book["genre"]["String"]
        book["description"] = book["description"]["String"]
        book["language"] = book["language"]["String"]
        book["edition"] = book["edition"]["Int32"]
        book["num_copies"] = book["num_copies"]["Int32"]
        book["location"] = book["location"]["String"]
        book["image_url"] = book["image_url"]["String"]
        book["keywords"] = book["keywords"]["String"]

        new_books_normalised.append(book)

    return new_books_normalised


def _normalize_users_data(users_list):
    """
    Normalize the users' data by converting the 'contact_number' key value to a string.

    Args:
    users_list (list): A list of dictionaries containing user data.

    Returns:
    list: A list of dictionaries with the 'contact_number' value converted to a string.
    """
    users_list_normalised = []
    for user in users_list:
        user["contact_number"] = str(user["contact_number"])
        users_list_normalised.append(user)
    return users_list_normalised


def read_books_list():
    url = f"{URL}/books/list"
    print(f"\nCalling the API to read books {url}")
    books_list = requests.get(url, headers=headers)
    new_books_normalised = normalize_books_data(books_list.json()["data"])
    global DELETE_ID
    DELETE_ID = new_books_normalised[-1]["book_id"]
    print(tabulate(new_books_normalised, headers="keys", tablefmt='psql'))


def read_books_paginated(limit=3, token=""):
    url = f"{URL}/books/list-paged?limit={limit}&token={token}"
    print(f"\nCalling the API to read books using pagination {url}")
    books = requests.get(url, headers=headers)
    books = books.json()
    prev_token = books["previous_token"]
    next_token = books["next_token"]
    normalised_books = normalize_books_data(books["data"])
    print(tabulate(normalised_books, headers="keys", tablefmt='psql'))
    print(f"\nPrevious token: {prev_token}\nNext token: {next_token}")
    return prev_token, next_token


def read_users_list():
    url = f"{URL}/users/list"
    print(f"\nCalling the API to read users {url}")
    users_list = requests.get(url, headers=headers)
    users_list_normalised = _normalize_users_data(users_list.json()["data"])
    print(tabulate(users_list_normalised, headers="keys", tablefmt='psql'))


def delete_books_from_library(delete_id):
    url = f"{URL}/books/delete?book_id={delete_id}"
    print(f"\nCalling the API to delete book {url}")
    requests.delete(url, headers=headers)


def search_book_by_title(title):
    url = f"{URL}/books/search?search-text={title}"
    print(f"\nCalling the API to insert books {url}")
    response = requests.get(url, headers=headers)
    normalised_data = normalize_books_data(response.json()["data"])
    print(tabulate(normalised_data, headers="keys", tablefmt='psql'))


def count_books():
    url = f"{URL}/books/count"
    print(f"\nCalling the API to count books {url}")
    response = requests.get(url, headers=headers)
    print(response.json()["data"])


def borrowing_book():
    url = f"{URL}/borrow/add"
    print(f"\nCalling the API to borrow books {url}")
    borrowed = requests.post(url, headers=headers, data=json.dumps(
        {
            "userEmail": "tonystark@mail.com",
            "bookTitle": "Lord Of The Rings"
        }
    ))
    if borrowed.status_code == 200:
        print(tabulate([borrowed.json()["data"]], headers="keys", tablefmt='psql'))
    else:
        print(borrowed.json())


def returning_book():
    url = f"{URL}/borrow/return"
    print(f"\nCalling the API to return borrowed book {url}")
    borrowed = requests.post(url, headers=headers, data=json.dumps(
        {
            "userEmail": "tonystark@mail.com",
            "bookTitle": "Lord Of The Rings"
        }
    ))
    if borrowed.status_code == 200:
        print(tabulate([borrowed.json()], headers="keys", tablefmt='psql'))
    else:
        print(borrowed.json())


if __name__ == '__main__':
    print("Running Test Script\n".upper())

    connect_to_postgresql_and_truncate()

    print("\nInserting and displaying books to library".upper())
    insert_books_to_library()
    read_books_list()

    print("\nInserting and displaying users to library".upper())
    insert_users_to_library()
    read_users_list()

    print(f"\nRemoving a book from Library with id: {DELETE_ID}".upper())
    delete_books_from_library(DELETE_ID)
    read_books_list()

    print("\nSearching book by title : `Lord Of The Rings`".upper())
    search_book_by_title("Lord Of The Rings")

    print("\nCalculating the total number of books.".upper())
    count_books()

    print("\nBorrowing a book.`Lord Of The Rings`".upper())
    borrowing_book()

    print("\nUser trying to borrow the same book again will return error ".upper())
    borrowing_book()

    print("\nThe new count of the book should be decreased by 1.".upper())
    search_book_by_title("Lord Of The Rings")

    print("\nReturning the book `Lord Of The Rings`".upper())
    returning_book()

    print("\nNow the user can try to borrow `Lord Of The Rings` book and it should be borrowed.".upper())
    borrowing_book()

    print("\nAdding more data to the books table.".upper())
    insert_books_to_library("more_books.json")

    print("\nReading books paginated using the paginated API".upper())
    prev_token, next_token = read_books_paginated(limit=3, token="")

    print("\nReading books paginated using the paginated API with NEXT TOKEN".upper())
    prev_token, next_token = read_books_paginated(limit=3, token=next_token)

    print("\nReading books paginated using the paginated API with PREVIOUS TOKEN\nThis should print the previous page of books.".upper())
    prev_token, next_token = read_books_paginated(limit=3, token=prev_token)
