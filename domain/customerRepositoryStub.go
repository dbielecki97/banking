package domain

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func NewCustomerRepositoryStub() *CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:          "1001",
			Name:        "Dawid",
			City:        "Warsaw",
			Zipcode:     "15-222",
			DateOfBirth: "2000-01-01",
			Status:      "1",
		}, {
			Id:          "1002",
			Name:        "Michał",
			City:        "Warsaw",
			Zipcode:     "111111",
			DateOfBirth: "2000-01-01",
			Status:      "1",
		},
	}
	return &CustomerRepositoryStub{customers: customers}

}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}
