package repository

import (
	"gorm.io/gorm"
	"wallester/models"
)

type CustomerRepository interface {
	ShowCustomer() ([]models.Customer,int64)
	UpdateCustomer(customer *models.Customer) error
	InsertCustomer(customer *models.Customer) error
	SearchCustomer(firstname string, lastname string) ([]models.Customer,int64)
	FindById(id string) (models.Customer, error)
	LockedLock(customer *models.Customer) error
	LockedUnlock(customer *models.Customer) error
}
type customerRepository struct {
db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository{
	return &customerRepository {
		db: db,
	}


}

func (r *customerRepository) ShowCustomer() ([]models.Customer, int64) {
	var customers []models.Customer
	var count int64
	r.db.Find(&customers).Count(&count)
	return customers, count
}

func (r *customerRepository) UpdateCustomer(customer *models.Customer) error {
	r.db.Save(&customer)
	return nil

}

func (r *customerRepository) FindById(id string) (models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("id = ?", id).First(&customer).Error
	return customer, err
}

func (r *customerRepository) InsertCustomer(customer *models.Customer) error{
	r.db.Create(&customer)
	return nil
}
func (r *customerRepository) LockedLock(customer *models.Customer) error {
	r.db.Model(&customer).Update("locked", true)
	return nil
}
func (r *customerRepository) LockedUnlock(customer *models.Customer) error {
	r.db.Model(&customer).Update("locked", false)
	return nil
}


func (r *customerRepository) SearchCustomer(firstname string, lastname string) ([]models.Customer,int64){
	var customer []models.Customer
	var count int64
	r.db.Order("id DESC").Where("first_name = ? OR last_name = ?", firstname, lastname).Find(&customer).Count(&count)
	return customer, count
}