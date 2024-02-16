package routes

import (
	"encoding/json"
	"log"
	
	mdware "mainweb/middleware"
	"mainweb/models"
	"net/http"
	"strconv"
    "html/template"
	// "gorm.io/gorm"
	base "mainweb/database"
)



func Saveuser(w http.ResponseWriter,r *http.Request){//here both user & auth save
	r.ParseForm()
	password:=r.FormValue("password")
	newpassword:=mdware.CreateHash(&password)
	Uname:=r.FormValue("first_name")+" "+r.FormValue("last_name")
	balance,er:=strconv.Atoi(r.FormValue("balance"))
	if er!=nil{
		log.Println("priblem apear")
		return
	}
	user:=models.User{
		First_name: r.FormValue("first_name"),
		Last_name: r.FormValue("last_name"),
		Phone_no: r.FormValue("phone_no"),
		Email_id: r.FormValue("email_id"),
		Password: newpassword,
		Authentication: models.Auth{
			Email_id: r.FormValue("email_id"),
			Password: newpassword,
			Balance: uint(balance),
		},

	}
	userlogs:=base.Mydb.Db.Create(&user)
	if userlogs!=nil{
		
		log.Println(userlogs)
	}
	
	json.NewEncoder(w).Encode(Uname)  //sending user name to frontend


}

func SaveSeller(w http.ResponseWriter,r *http.Request){ //saving only seller we update product
	r.ParseForm()
	fpass:=r.FormValue("password")
	femail:=r.FormValue("email_id")
	getbool:=mdware.SellerEmail(&femail,base.Mydb.Db)
	if getbool{
		json.NewEncoder(w).Encode("email alredy registered")
		return
	}
	password:=mdware.CreateHash(&fpass)
	seller:=models.Seller{
		Email_id: r.FormValue("email_id"),
		Password: password,
	}
	sellerlogs:=base.Mydb.Db.Create(&seller)
	if sellerlogs!=nil{
		log.Println(sellerlogs)
	}
	json.NewEncoder(w).Encode("seller is created")

}
func SaveProducts(w http.ResponseWriter,r *http.Request){

	r.ParseForm()
	price,er1:=strconv.Atoi(r.FormValue("price"))
	total,er2:=strconv.Atoi(r.FormValue("total"))
	if er1!=nil{
		log.Println(er1)
	}
	if er2!=nil{
		log.Println(er2)
	}
	
	var Product []models.Product 
	var email string = r.FormValue("email_id")
	myprod:=models.Product{
		Name: r.FormValue("name"),
		Brand: r.FormValue("brand"),
		Price:uint(price),
		Total: uint(total),
	}
	getprid:=mdware.FindByEmail(&email,base.Mydb.Db)
	
	Product=append(Product, myprod)
	sellertable:=models.Seller{
		Id:getprid,  //giving primary key to identify that email row
		Products: Product,
	}
    prodlogs:=base.Mydb.Db.Save(&sellertable)
	// prodlogs:=db.Create(&product)
	if prodlogs!=nil{
		log.Println(prodlogs)
	}
	json.NewEncoder(w).Encode(r.FormValue("name"))

}

func SaveBuyer(w http.ResponseWriter,r *http.Request){
	type buy struct{
	    Email_id string `json:"email_id" gorm:"not null"`
	    Phone_no string `json:"phone_no" gorm:"not null"`
		Product []models.Product_instance `json:"product" gorm:"not null"`
		Amount uint `json:"amount" gorm:"not null"`
		// Brand string `json:"brand" gorm:"not null"`
		// Name string `json:"name" gorm:"not null"`
	}
	var collect buy
	torf:=mdware.BuyerExist(&collect.Email_id,base.Mydb.Db)
	if !torf{
		log.Println("wrong email--!")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&collect)
	value:=collect.Amount
    prob:=mdware.PaymentDone(&value,&collect.Email_id,base.Mydb.Db)
	if !prob{
		log.Println("please select correct")
		return
	}
	
	buyer:=models.Buyer{
		Email_id: collect.Email_id,
		Phone_no: collect.Phone_no,
		ProductInstance:collect.Product ,
	}
	// var paybil models.Payment
	paybil:=models.Payment{
		Amount: collect.Amount,
		Buyerinstance: buyer,
	}
	
	base.Mydb.Db.Create(&buyer)
	base.Mydb.Db.Create(&paybil)
	json.NewEncoder(w).Encode("payment done")
	
}
func LoginUser(w http.ResponseWriter,r *http.Request){
	email:=r.FormValue("email_id")
	password:=r.FormValue("password")
	result:=mdware.Authenticate(&email,&password,base.Mydb.Db)
	if !result{
		log.Println("check your input")
		return
	}
	http.Redirect(w,r,"/products",http.StatusSeeOther)
}



func Indexfile(w http.ResponseWriter,r *http.Request){
	temp,err:=template.ParseFiles("templates/index.html")
	if err!=nil{
		http.Error(w,"html file not found",http.StatusNotFound)
	}
	temp.Execute(w,nil)
}
func Storefile(w http.ResponseWriter,r *http.Request){
	temp,err:=template.ParseFiles("templates/store.html")
	if err!=nil{
		http.Error(w,"html file not found",http.StatusNotFound)

	}
	temp.Execute(w,nil)
}
func Loginfile(w http.ResponseWriter,r *http.Request){
	temp,err:=template.ParseFiles("templates/login.html")
	if err!=nil{
		http.Error(w,"html file not found",http.StatusNotFound)
	}
	temp.Execute(w,nil)
}
func Cardfile(w http.ResponseWriter,r *http.Request){
	temp,_:=template.ParseFiles("templates/card.html")
	temp.Execute(w,nil)
}