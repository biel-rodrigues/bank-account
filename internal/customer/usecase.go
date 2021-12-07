package customer

import (
	"github.com/bank-account/internal/dtos/request"
	"gorm.io/gorm"
)

func create(DB *gorm.DB, dto request.Customer) (response *Customer, err error) {
	c := dtoToCustomer(dto)

	err = DB.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func read(DB *gorm.DB, id interface{}) Customer {
	customer := Customer{}
	DB.Find(&customer, id)
	return customer
}

func update(DB *gorm.DB, id interface{}, dto request.Customer) {
	var customer Customer
	DB.Where("id=?", id).Find(&customer)

	customer.Name = dto.Name
	customer.Cpf = dto.Cpf
	DB.Save(&customer)
}

func delete(DB *gorm.DB, id interface{}) {
	customer := Customer{}
	DB.Where("id=?", id).Delete(&customer)
}
