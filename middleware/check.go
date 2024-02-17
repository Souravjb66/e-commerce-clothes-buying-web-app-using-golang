package middleware

import (
	// "encoding/json"
	
	"log"
	"mainweb/models"
	"sync"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
func SellerEmail(email *string,db *gorm.DB)bool{
	type seller struct{
		Email_id string `json:"email_id" gorm:"not null"`
	}
	var sl1 seller
	err:=db.Where("email_id = ?", *email).First(&sl1)
	if err!=nil{
		log.Println(err)
		
	}
    var semail string =sl1.Email_id
	if semail==""{
		log.Println("empty string")
		return false
	}

	return true
}
func PaymentDone(amount *uint,email *string,db *gorm.DB)bool{
	// type userauth struct{  //can require to add all firlds
	// 	Balance uint `json:"balance"`
	// }
	
	var auth models.Auth
	// user:=userauth{}
	err:=db.Find(&auth,"email_id=?",*email)
	if err!=nil{
		// log.Println("user not exist and you check out payment done method ")
		log.Println(err)
		// return false
	}
	
	total:=auth.Balance
	if total<*amount{
		log.Println("low balance in your account")
		return false
		
	}
	var wg sync.Mutex
	func(){
		wg.Lock()
		defer wg.Unlock()
		total=total-*amount

	}()
	
	auth.Balance=total
	db.Save(&auth)

	return true

}
func FindByEmail(sendpass *string,email *string,db *gorm.DB)models.Seller{
	
	seller:=models.Seller{}
	err:=db.Where("email_id=?",*email).First(&seller)
	if err!=nil{
		log.Println(err)
	}
	if seller.Email_id==""{
		log.Println("wrong email")
		return seller
	}
	newvalue:=bcrypt.CompareHashAndPassword([]byte(seller.Password),[]byte(*sendpass))
	if newvalue!=nil{
		log.Println("problem in seller password in hash compare!!!")
		log.Println(newvalue)
		
	}
	
	log.Println(seller.Id)
	return seller

}
func BuyerExist(email *string,db *gorm.DB)bool{
	// type buyer struct{
	// 	Email_id string `json:"email_id" gorm:"not null"`
	//     Password string `json:"password" gorm:"not null"`
	// }
	// var minauth buyer
	log.Println("fir :",*email)
	var mx models.Auth
	err:=db.Where("email_id =?",*email).Find(&mx)
	if err!=nil{
		log.Println("email id !!!!!!")
		log.Println(err)
		
	}
	log.Println("email ::",mx.Email_id)
	if mx.Email_id==""{
		log.Println("not a valid buyer")
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
	ar:=db.Where("email_id=?",*email).Find(&myauth)
	if ar!=nil{
		log.Println("in authenticate!!!",ar)
		
	}
	if myauth.Email_id==""{
		return false
	}

	err:=bcrypt.CompareHashAndPassword([]byte(myauth.Password),bytepass)
	if err!=nil{
		log.Println("wrong password",err)
		return false
	}
	return true

}