package domain

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/parin/msisdn-country-validator/utils"
	"net/http"
)

var (
	DB *sql.DB
	CustomerDao customerDao
)

type customerDao struct {

}

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
func (c  customerDao) FetchCustomers() ([]*Customer,*utils.ApplicationError)  {
	customerInfo := []*Customer{}
	rows , err := DB.Query("SELECT  name,phone FROM customer")
	if err != nil {
		return nil,&utils.ApplicationError{
			Message: err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "Unable to open db",
		}
	}
	defer rows.Close()
	for rows.Next(){
		customer := &Customer{}
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
