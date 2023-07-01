package model

import "github.com/go-playground/validator/v10"

type FizzbuzzRequest struct {
	FirstInteger  int    `json:"firstInteger" validate:"required,gte=0"`
	SecondInteger int    `json:"secondInteger" validate:"required,gte=0"`
	Limit         int    `json:"limit" validate:"required,gte=0"`
	FirstString   string `json:"firstString" validate:"required,min=1"`
	SecondString  string `json:"secondString" validate:"required,min=1"`
}

// Constructor
func NewFizzbuzzRequest(firstInteger int, secondInteger int, limit int, firstString string, secondString string) FizzbuzzRequest {
	return FizzbuzzRequest{
		FirstInteger:  firstInteger,
		SecondInteger: secondInteger,
		Limit:         limit,
		FirstString:   firstString,
		SecondString:  secondString,
	}
}

func ValidateFizzbuzzRequest(request FizzbuzzRequest) error {
	validate := validator.New()
	return validate.Struct(request)
}
