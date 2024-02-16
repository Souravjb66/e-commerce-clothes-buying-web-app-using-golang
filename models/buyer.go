package models

type Buyer struct{
	Id int `json:"id" gorm:"primary_Key"`
	Email_id string `json:"email_id" gorm:"not null"`
	Phone_no string `json:"phone_no" gorm:"not null"`
	ProductInstance []Product_instance `json:"productid" gorm:"many2many:buyer_product"`
	PaymentId int `json:"paymentid" gorm:"not null"`

}