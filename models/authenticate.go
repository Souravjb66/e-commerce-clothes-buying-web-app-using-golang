package models


type Auth struct{
	Id int `json:"id" gorm:"primary_Key"`
	Email_id string `json:"email_id" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Balance uint `json:"balance" `
	UserId int `json:"userid" gorm:"not null"`

}