package domain

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"github.com/parin/msisdn-country-validator/utils"
)
// Database is implemented to fetch data from storage
type Database interface {
	Close() error
	FetchCustomers() ([]*Customer, *utils.ApplicationError)
}
//SqliteDatabase implements Database  Interface
type SqliteDatabase struct {
	Database
	db *sql.DB
}


//NewCustomerDatabase creates a new customer database access object
func NewCustomerDatabase(database string) (*SqliteDatabase, error) {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		return nil, fmt.Errorf("unable to open db: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}
	return &SqliteDatabase{ db: db }, nil
}

// Close terminates the database connection
func (c *SqliteDatabase) Close() error {
	return c.db.Close()
}
// FetchCustomers gets customers from  Customer database
func (c *SqliteDatabase) FetchCustomers() ([]*Customer, *utils.ApplicationError) {
	log.Printf("trying to fetch customer")
	customerInfo := []*Customer{}
	rows, err := c.db.Query("SELECT name, phone FROM customer")
	if err != nil {
		return nil, &utils.ApplicationError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "Unable to open db",
		}
	}
	defer rows.Close()
	for rows.Next() {
		customer := &Customer{}
		err := rows.Scan(&customer.Name, &customer.Msisdn)
		if err != nil {
			return nil, &utils.ApplicationError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
				Code:       "Unable to open db",
			}
		}
		customerInfo = append(customerInfo, customer)
	}
	return customerInfo, nil
}
