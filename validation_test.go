package belajar_golang_validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	user := "eko"

	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateTwoVariables(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "rahasia"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()
	user := "12345"

	err := validate.Var(user, "required,numeric")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()
	user := "99999999999999"

	err := validate.Var(user, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "eko@example.com",
		Password: "eko1234",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}
