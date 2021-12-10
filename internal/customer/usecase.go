package customer

import (
	"github.com/bank-account/internal/dtos/request"
	"gorm.io/gorm"
)

func create(DB *gorm.DB, DTO request.Customer) (response *Customer, err error) {
	c := convert(DTO)

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

func update(DB *gorm.DB, id interface{}, DTO request.Customer) {
	var customer Customer
	DB.Where("id=?", id).Find(&customer)

	customer.Name = DTO.Name
	customer.Cpf = DTO.Cpf
	DB.Save(&customer)
}

func delete(DB *gorm.DB, id interface{}) {
	customer := Customer{}
	DB.Where("id=?", id).Delete(&customer)
}
