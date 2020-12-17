package validators

import (
	"github.com/bearbin/go-age"
	"github.com/go-playground/validator/v10"
	"time"
)

func CustomerGender(fl validator.FieldLevel) bool {
	return fl.Field().String() == "Male" || fl.Field().String() == "Female"
}

func CustomerBirthDate(fl validator.FieldLevel) bool {
	t,_ := time.Parse("2006-01-02",fl.Field().String())
	return age.Age(t) >= 18 && age.Age(t) <= 60
}

