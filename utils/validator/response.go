package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrResponse struct {
	Errors []string `json:"errors"`
}

func ToErrResponse(err error) *ErrResponse {
	if fieldError, ok := err.(validator.ValidationErrors); ok {
		resp := ErrResponse{
			Errors: make([]string, len(fieldError)),
		}

		for i, err := range fieldError {
			switch err.Tag() {

			case "required":
				resp.Errors[i] = fmt.Sprintf("The field %s is required", err.Field())
			case "max":
				resp.Errors[i] = fmt.Sprintf("The field %s must be less than %s", err.Field(), err.Param())
			case "url":
				resp.Errors[i] = fmt.Sprintf("The field %s must be a valid URL", err.Field())
			case "alphaspace":
				resp.Errors[i] = fmt.Sprintf("The field %s must contain only letters and spaces", err.Field())
			case "datetime":
				if err.Param() == "2006-01-02" {
					resp.Errors[i] = fmt.Sprintf("The field %s must be a valid date (YYYY-MM-DD)", err.Field())
				} else {
					resp.Errors[i] = fmt.Sprintf("The field %s must be a valid date and time (YYYY-MM-DD HH:MM:SS)", err.Field())
				}
			default:
				resp.Errors[i] = fmt.Sprintf("The field %s is invalid", err.Field())

			}
		}

		return &resp
	}

	return nil
}
