package customers_db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/parin/msisdn-country-validator/domain"
	"github.com/parin/msisdn-country-validator/utils"
	"log"
	"net/http"
)

var (
	DB *sql.DB
	SqliteClient Client
	)

type Client struct {}

func init()  {
	var err error
	Client ,err:= sql.Open("sqlite3","./customer.db")
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	DB = Client
}

func (c  *Client) FetchCustomerInfo() ([]*domain.Customer,*utils.ApplicationError)  {
	customerInfo := []*domain.Customer{}
	rows , err := DB.Query("SELECT  name,phone FROM customer")
	if err != nil {
		log.Printf("Unable to fetch to row %v",err)
	}
	for rows.Next(){
		customer := &domain.Customer{}
		err :=  rows.Scan(&customer.Name,&customer.Msisdn)
		if err != nil {
			return nil,&utils.ApplicationError{
				Message: err.Error(),
				StatusCode: http.StatusInternalServerError,
				Code: "Unable to open db",
			}
		}
		customerInfo = append(customerInfo,customer)
	}
	return customerInfo ,nil
}
