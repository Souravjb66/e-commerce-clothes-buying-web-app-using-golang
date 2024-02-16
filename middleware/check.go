package middleware

import (
	// "encoding/json"
	"log"
	"mainweb/models"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)
func SellerEmail(email *string,db *gorm.DB)bool{
	err:=db.First(&models.Seller{},*email)
	if err!=nil{
		log.Println(err)
		return false

	}
	return true


}
func PaymentDone(amount *uint,email *string,db *gorm.DB)bool{
	type userauth struct{  //can require to add all firlds
		Balance uint `json:"balance"`

	}
	var auth models.Auth
	user:=userauth{}
	err:=db.Find(&user,"email_id=?",*email)
	if err!=nil{
		log.Println("user not exist and you check out payment done method ")
		return false
	}
	
	total:=user.Balance
	if total<*amount{
		log.Println("low balance in your account")
		return false
		
	}
	total=total-*amount
	
	auth.Balance=total
	db.Save(&auth)

	return true

}
func FindByEmail(email *string,db *gorm.DB)int{
	type seller struct{
		Id int `json:"id" gorm:"primary_Key"`
	    Email_id string `json:"email_id" gorm:"not null"`
	    Password string `json:"password" gorm:"not null" `
	    Products []models.Product `json:"productid" gorm:"foreignKey:SellerId"`

	}
	var slr seller
	err:=db.First(&slr,*email)
	if err!=nil{
		log.Println(err)
	}
	return slr.Id

}
func BuyerExist(email *string,db *gorm.DB)bool{
	err:=db.First(&models.Auth{},*email)
	if err!=nil{
		log.Println(err)
		return false
	}
	return true
}
func Authenticate(email *string,password *string,db *gorm.DB)bool{
	type auth struct{
		Email_id string `json:"email_id"`
		Password string `json:"password"`
		Balance uint `json:"balance"`
	}
	
	var myauth auth
	bytepass:=[]byte(*password)
	ar:=db.Find(&myauth,*email)
	if ar!=nil{
		log.Println("in authenticate!!!",ar)
		return false
	}

	err:=bcrypt.CompareHashAndPassword([]byte(myauth.Password),bytepass)
	if err!=nil{
		log.Println("wrong password",err)
		return false
	}
	return true

}