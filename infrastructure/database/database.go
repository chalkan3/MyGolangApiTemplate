package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // database driver
)

// Database is my database struct
type Database struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	DatabaseType string `json:"databaseType"`
}

func (database *Database) formatSQLConnection() string {
	return database.Username + ":" + database.Password + "@/" + database.Name
}

// GetConnection get a database connection
func (database *Database) GetConnection() *sql.DB {
	db, err := sql.Open(database.DatabaseType, database.formatSQLConnection())
	if err != nil {
		panic(err.Error())
	}

	return db
}

// NewDatabase IoC
func NewDatabase() *Database {
	return GetManagerAppConfig().Database
}
