package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

var birthdatedata = []struct {
	have string
	want bool
}{
	{"1999-03-23", true},
	{"gsdolsdfoldfsolfd", false},
	{"alasask@fmdlr.ru", false},
	{"'-0||'", false},
	{"1950-02-02", false},
}

func TestBirthDate(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("birth", CustomerBirthDate)

	for _, item := range birthdatedata {
		err := validate.Var(item.have, "birth")
		if item.want {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

var genders = []struct {
	have string
	want bool
} {
	{"Male", true},
	{"Female", true},
	{"gsdolsdfoldfsolfd", false},
	{"alasask@fmdlr.ru", false},
	{"'-0||'", false},
	{"1950-02-02", false},
}
func TestGender(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("genderval", CustomerGender)

	for _, item := range genders {
		err := validate.Var(item.have, "genderval")
		if item.want {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}