package services

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3" // SQLite driver
    _ "github.com/lib/pq"           // PostgreSQL driver
    _ "github.com/denisenkom/go-mssqldb" // MSSQL driver
    _ "github.com/go-sql-driver/mysql" // MySQL driver
    _ "github.com/mattn/go-oci8" // Oracle driver
)

type DatabaseService struct {
    DB *sql.DB
}

func NewDatabaseService(driver, dataSource string) (*DatabaseService, error) {
    db, err := sql.Open(driver, dataSource)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    return &DatabaseService{DB: db}, nil
}

func (ds *DatabaseService) Close() error {
    return ds.DB.Close()
}

// Add additional methods for database operations here, e.g., CRUD operations for projects.