package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // database driver
)

// Database is my database struct
type Database struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	Name         string `json:"name,omitempty"`
	Host         string `json:"host,omitempty"`
	DatabaseType string `json:"databaseType,omitempty"`
}

func (database *Database) formatSQLConnection() string {
	return database.Username + ":" + database.Password + "@tcp(" + database.Host + ":3308)/" + database.Name
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
