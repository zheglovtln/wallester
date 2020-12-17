package models

type Customer struct {
ID uint `gorm:"primaryKey" json:"id"`
FirstName string `validate:"max=100,required" gorm:"size:100;not null;" json:"firstname" binding:"required"`
LastName string `validate:"max=100,required" gorm:"size:100;not null" json:"lastname" binding:"required"`
BirthDate string `validate:"birth" gorm:"not null,required; type:date" json:"birthdate" binding:"required"`
Gender string `validate:"genderval,required" gorm:"not null" json:"gender" binding:"required"`
Email string `validate:"required,email" gorm:"not null"; json:"email" binding:"required"`
Address string `validate:"max=200" gorm:"size:200" json:"address" binding:"required"`
Locked bool `gorm:"not null;type:bool;default:false" json:"locked"`
}

func (c *Customer) TableName() string {
	return "customer"
}

//ID primary, required, AutoInc
//FirstName string required length 100
//LastName string required length 100
//BirthDate datetime required 18 to 60
//Gender string required allowed(Male,Female)
//Email string required (email@mail.com format)
//Address string optional length 200