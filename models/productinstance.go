package models


type Product_instance struct{
	
	Id int `json:"id" gorm:"primary_Key"`
	Brand string `json:"brand" gorm:"not null"`
	Name string `json:"name" gorm:"not null"`
	Price uint `json:"price" gorm:"not null"`
	ProductBuyer []Buyer `json:"productbuyer" gorm:"many2many:buyer_product"`
	

}