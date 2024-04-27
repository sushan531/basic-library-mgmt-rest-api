package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	path2 "path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

// DB : Global declaration of DB instance
var DB *sql.DB

// CreateDB creates a new database with the provided dbName, host, user, password, and port
func CreateDB(dbName, host, user, password, port string) error {
	// Set up the connection parameters
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s sslmode=disable",
		host, port, user, password)

	// Open a connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	// Create the database
	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		// If the database already exists, return nil
		if strings.Contains(err.Error(), "already exists") {
			return nil
		}
		return err
	}

	// Close the database connection
	err = db.Close()
	if err != nil {
		return err
	}

	// Return nil to indicate success
	return nil
}

// InitializeDB initializes the database schema by executing migration scripts in a specified folder
func InitializeDB(db *sql.DB) (bool, error) {
	// Get the path of the current file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	// Determine the folder containing the database schema migrations
	restServiceFolder := path2.Dir(path2.Dir(filename))
	dbSchemaFolder := filepath.Join(restServiceFolder, "dbSchema", "migrations")

	// Print the folder path for debugging purposes
	fmt.Println(dbSchemaFolder)

	// Open the schema folder
	files, err := os.ReadDir(dbSchemaFolder)
	if err != nil {
		return false, err
	}

	// Sort files lexicographically
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	// Execute migration scripts in order
	for _, file := range files {
		if file.IsDir() {
			continue // Skip subdirectories
		}
		filePath := filepath.Join(dbSchemaFolder, file.Name())
		script, err := os.ReadFile(filePath)
		if err != nil {
			return false, err
		}
		_, err = db.Exec(string(script))
		if err != nil {
			return false, err
		}
	}

	// Execute triggers script
	triggersFilePath := filepath.Join(restServiceFolder, "dbSchema", "triggers", "triggers.sql")
	triggers, _ := os.ReadFile(triggersFilePath)
	_, err = db.Exec(string(triggers))
	if err != nil {
		return false, err
	}

	return true, nil
}

// NewDB creates a new database connection
// dbName: the name of the database
// host: the host of the database
// user: the username for the database connection
// password: the password for the database connection
// port: the port for the database connection
func NewDB(dbName, host, user, password, port string) (*sql.DB, error) {
	// Construct the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	// Open a connection to the PostgreSQL database
	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return DB, nil
}

// GetDBInstance returns the singleton instance of the database connection
func GetDBInstance() *sql.DB {
	return DB
}
