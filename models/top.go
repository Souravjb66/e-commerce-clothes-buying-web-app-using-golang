package models

type Product struct{
	Id int `json:"id" gorm:"primary_Key"`
	Brand string `json:"brand" gorm:"not null"`
	Name string `json:"name" gorm:"not null"`
	Price uint `json:"price" gorm:"not null"`
	Total uint `json:"total" gorm:"not null"`
	SellerId int `json:"sellerid" gorm:"not null"` 


}