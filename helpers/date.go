package helpers

import (
	"time"
	"wallester/models"
)

func FormatBirthDate(customer models.Customer) models.Customer {
		t,_ := time.Parse("2006-01-02T15:04:05Z",customer.BirthDate)
		t2 := t.Format("2006-01-02")
		customer.BirthDate = t2
	return customer
}