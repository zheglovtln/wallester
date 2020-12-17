package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"wallester/helpers"
	"wallester/models"
	"wallester/repository"
	"wallester/validators"
)
type CustomerInterface interface {
	TableView(c *gin.Context)
	SearchView(c *gin.Context)
	EditView(c *gin.Context)
	CreateView(c *gin.Context)

	ShowCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	InsertCustomer(c *gin.Context)
	SearchCustomer(c *gin.Context)
}

type controller struct {
	repo repository.CustomerRepository
}

var validate *validator.Validate
var pageSize int


func New(repo repository.CustomerRepository) CustomerInterface {
	pageSize = 4
	validate = validator.New()
	validate.RegisterValidation("genderval", validators.CustomerGender)
	validate.RegisterValidation("birth", validators.CustomerBirthDate)
	return &controller{
		repo: repo,
	}
}

//Views
func (x controller)TableView(c *gin.Context) {
	var search string
	var queryType string
	if c.PostForm("FirstName") != "" || c.PostForm("LastName") != "" {
		queryType = "GET"
		search = "/api/customers/search?FirstName=" + c.PostForm("FirstName") + "&LastName=" + c.PostForm("LastName")
	} else {
		queryType = "POST"
		search = "/api/customers"
	}
	c.HTML(http.StatusOK, "index", gin.H{
		"search": search,
		"type": queryType,
	})
}

func (x controller)SearchView(c *gin.Context) {
	c.HTML(http.StatusOK, "search", gin.H{

	})
}
func (x controller)EditView(c *gin.Context) {
	customer, err := x.repo.FindById(c.Param("id"))
	if customer.Locked == true || err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}
	customer = helpers.FormatBirthDate(customer)
	c.HTML(http.StatusOK, "edit", gin.H{
		"customer": customer,
		"request": "/api/customers/"+c.Param("id"),
	})
}

func (x controller)CreateView(c *gin.Context) {

	c.HTML(http.StatusOK, "add", gin.H{
	"request": "/api/customers/add",
	})
}

func (x controller)ShowCustomer(c *gin.Context) {
	customers,_ := x.repo.ShowCustomer()
c.JSON(http.StatusOK, gin.H {
	"data": customers,
})
}

//Methods for API
func (x controller) InsertCustomer(c *gin.Context) {
	var customer models.Customer
	c.ShouldBindWith(&customer, binding.FormMultipart)
	err := validate.Struct(customer)
	if err != nil {
		c.JSON(http.StatusOK,err.Error())
		return
	}


	err = x.repo.InsertCustomer(&customer)
	if err != nil {

	}
	c.JSON(http.StatusOK, "Insert successful!")
}

func (x controller) UpdateCustomer(c *gin.Context) {
	customer, _ := x.repo.FindById(c.Param("id"))
	c.ShouldBindWith(&customer, binding.FormMultipart)
	err := validate.Struct(customer)
	if err != nil {
	c.JSON(http.StatusBadRequest,gin.H{
		"error": err.Error(),
	})
	return
	}


	err = x.repo.UpdateCustomer(&customer)
	if err != nil {
		fmt.Println("Updated Failed")
	}
	c.JSON(http.StatusOK, "Data successfully updated!")
}

func (x controller) SearchCustomer(c *gin.Context) {
	customers, _ := x.repo.SearchCustomer(c.Query("FirstName"), c.Query("LastName"))
	c.JSON(http.StatusOK, gin.H {
		"data": customers,
	})
}