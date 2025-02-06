package httpserver

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func (r UrlRequest) Validation() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	
	err := validate.Struct(r)
	if err != nil {
		validationError := err.(validator.ValidationErrors)
		// Разобраться как привести к типу
		msg := msgValidator(validationError)

		return fmt.Errorf(msg)
	}
	
	return nil
}

func msgValidator(fe validator.ValidationErrors) string {
	var msg string

	for _, field := range fe {
		switch(field.Tag()){
		case "url":
			msg = "field doesn`t type url"
			break
		case "required":
			msg = fmt.Sprintf("field is required: %s", field.Field())
			break
		default:
			msg = "validation error"
			break 
		}
	}

	return msg
}