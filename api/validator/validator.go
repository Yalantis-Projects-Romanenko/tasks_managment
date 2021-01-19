package validator

import "github.com/go-playground/validator"

var (
	validate *validator.Validate
)

func Init() {
	validate = validator.New()
}

func Get() *validator.Validate {
	return validate
}

type ValidationError struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}

func ParseValidationErrors(err error) (result []ValidationError) {

	// this check is only needed when your code could produce
	// an invalid value for validation such as interface with nil
	// value most including myself do not usually have code like this.
	if _, ok := err.(*validator.InvalidValidationError); ok {
		result = append(result, ValidationError{
			"",
			err.Error(),
		})
		return
	}

	for _, err := range err.(validator.ValidationErrors) {
		result = append(result, ValidationError{
			err.Field(),
			err.Tag(),
		})
	}
	return result
}
