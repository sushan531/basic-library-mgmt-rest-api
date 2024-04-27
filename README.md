## LMS - Library Management System REST API

This project provides a REST API for managing library loans and book information.

**Prerequisites**

* **Golang:** Download and install Golang from the official website: [https://go.dev/](https://go.dev/).
* **Docker:** Install Docker if you haven't
  already: [https://docs.docker.com/engine/install/](https://docs.docker.com/engine/install/).
* **Python3:** Install Python if you haven't
  already: [https://www.python.org/downloads/](https://www.python.org/downloads/). Also install the requests and tabulate library
  which is required to run tests `pip install requests; pip install tabulate`
* **PostgreSQL:** This project utilizes a PostgreSQL database. You can either install PostgreSQL directly or leverage
  Docker to run it as a container.

**Running the Server**

**1. Install Dependencies**

- Golang: Follow the installation instructions on the official website linked above.
- Docker: Refer to the Docker documentation for installation steps specific to your operating system.

**2. Start the PostgreSQL Database**

**Install: Using Docker**

- Make sure you are in the projects root directory i.e `LMS` in this case.
- Open a terminal and navigate to the project's root directory (e.g., `cd LMS`).
- Start a PostgreSQL container in the background with:

  ```bash
  docker-compose up -d
  ``` 

**3. Build the REST API Server**

- Open a terminal and navigate to the project's root directory (e.g., `cd LMS`).
- Build the server using the command:

  ```bash
  go build .
  ```
  This should also install all the required packages such as **sqlc**, **echo**, etc.

This will create an executable file named `LMS` or `LMS.exe` depending on your system.

**4. Run the Server**

- Execute the compiled file:

    - Linux/macOS: `./LMS`
    - Windows: `start LMS.exe`

**5. Run automated test**
- All the test related scripts and data are located in the test_script folder
- Install the required libraries using 
    - ` pip install -r test_script/requirements.txt`
    - 
- Execute the python file:
    - `python3 test_script/api_test_script.py` or
    - `python test_script/api_test_script.py`
  

- You should be able to see all the responses sequentially

**5. Testing out endpoints in POSTMAN**
- If you want to test out the endpoints in postman you can import the file `LMS.postman_collection.json` and test out the endpoints individually.
- You will have to set the value of `URL` to `http://localhost:8000` or `http://0.0.0.0:8000` in the environment.

**Additional Notes**

* This document assumes basic familiarity with Golang and Docker and Python.
* You might need to create new python virtual env if not already exists to install requests library.