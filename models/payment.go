package models

type Payment struct{
	Id int `json:"id" gorm:"primary_Key"`
	Amount uint `json:"amount" gorm:"not null"`
	Buyerinstance Buyer `json:"buyerinstance" gorm:"foreignKey:PaymentId"`

}