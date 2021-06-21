package service

import (
	"github.com/dbielecki97/banking/domain"
	"github.com/dbielecki97/banking/dto"
	"github.com/dbielecki97/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewDefaultCustomerService(repository domain.CustomerRepository) *DefaultCustomerService {
	return &DefaultCustomerService{repository}
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	cs, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	customers := make([]dto.CustomerResponse, 0)
	for _, c := range cs {
		customers = append(customers, c.ToDto())
	}
	return customers, err
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	customer := c.ToDto()

	return &customer, err
}
