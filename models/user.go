package models

type User struct{
	Id int `json:"id" gorm:"prinmary_Key"`
	First_name string `json:"first_name" gorm:"not null"`
	Last_name string `json:"last_name" gorm:"not null"`
	Phone_no string `json:"phone_no" gorm:"not null"`
	Email_id string `json:"email_id" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Authentication Auth `json:"auth" gorm:"foreignKey:UserId"`
}