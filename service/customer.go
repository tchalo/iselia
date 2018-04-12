package service

import (
	"fmt"
	"github.com/tchalo/iselia/database"
)

type Customer struct {
	Username string
	Password string
}

type CustomerService struct {
	db string
}

func NewCustomerService(database database.Database) *CustomerService {
	return &CustomerService{
		db: database.Mysql,
	}
}

func (cs *CustomerService) Register(data map[string]interface{}) error {
	fmt.Println(data["param1"])
	return nil
}
