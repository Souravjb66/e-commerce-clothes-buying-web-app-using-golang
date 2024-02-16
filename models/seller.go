package models

type Seller struct{
	Id int `json:"id" gorm:"primary_Key"`
	Email_id string `json:"email_id" gorm:"not null"`
	Password string `json:"password" gorm:"not null" `
	Products []Product `json:"productid" gorm:"foreignKey:SellerId"`
	
}