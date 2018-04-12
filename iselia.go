package iselia

import (
	"github.com/tchalo/iselia/database"
	service "github.com/tchalo/iselia/service"
)

type IseliaService struct {
	Customer service.CustomerService
}

func NewIselia(db database.Database) *IseliaService {
	return &IseliaService{
		Customer: *service.NewCustomerService(db),
	}
}
