package services

import (
	"fmt"
	"github.com/parin/msisdn-country-validator/domain"
	"github.com/parin/msisdn-country-validator/utils"
)

var (
	CustService  CustomerService
)
type  CustomerService struct {
	db domain.Database
}

//NewCustomerService creates a  new  customer  service
func NewCustomerService(db domain.Database) (*CustomerService, error) {
	if db == nil {
		return nil, fmt.Errorf("invalid database : %v", db)
	}
	return &CustomerService{
		db: db,
	}, nil
}
//GetCustomer gets customer from database
func (c CustomerService)  GetCustomers()  ([]*domain.Customer,*utils.ApplicationError) {
	//log.Printf("trying to fetch customer")
	var  validCustomers []*domain.Customer
	customers, err := c.db.FetchCustomers()
	if err != nil {
		return nil, err
	}
	for _ ,customer := range customers {
		customer.Validate()
		validCustomers = append(validCustomers,customer)
	}
	return validCustomers , nil
}

