package dto

type CustomerResponse struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	City        string `json:"city,omitempty"`
	Zipcode     string `json:"zipcode,omitempty"`
	DateOfBirth string `json:"date_of_birth,omitempty"`
	Status      string `json:"status,omitempty"`
}
