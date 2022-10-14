package belajar_golang_validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
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

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "eko",
		Password: "eko",
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestStructCrossField(t *testing.T) {
	type RegisterUser struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	validate := validator.New()
	request := RegisterUser{
		Username:        "eko@example.com",
		Password:        "eko1234",
		ConfirmPassword: "eko1234",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Address: Address{
			City:    "",
			Country: "",
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"required,dive,required,min=3"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"",
			"X",
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"",
			"X",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SD Indonesia",
			},
			"SMP": {
				Name: "",
			},
			"": {
				Name: "",
			},
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
		Wallet    map[string]int    `validate:"dive,keys,required,endkeys,required,gt=1000"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"",
			"X",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SD Indonesia",
			},
			"SMP": {
				Name: "",
			},
			"": {
				Name: "",
			},
		},
		Wallet: map[string]int{
			"BCA":     1000000,
			"MANDIRI": 0,
			"":        1001,
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestAlias(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar,min=5"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id:     "123",
		Name:   "",
		Owner:  "",
		Slogan: "",
	}

	err := validate.Struct(seller)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestCustomValidationFunction(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	request := LoginRequest{
		Username: "EKOKK",
		Password: "",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestCustomValidationParameter(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	request := Login{
		Phone: "0904190424",
		Pin:   "123123",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err)
	}
}

func TestOrRule(t *testing.T) {
	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	request := Login{
		Username: "12345",
		Password: "ekoo",
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("field not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCrossFieldValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	type User struct {
		Username string `validate:"required,field_equals_ignore_case=Email|field_equals_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	user := User{
		Username: "eko@example.com",
		Email:    "eko@example.com",
		Phone:    "089999999999",
		Name:     "Eko",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegisterSuccess(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(RegisterRequest)

	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {
		// sukses
	} else {
		// gagal
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}
}

func TestStructLevelValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	request := RegisterRequest{
		Username: "089923942934",
		Email:    "eko@example.com",
		Phone:    "089923942934",
		Password: "rahasia",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}
