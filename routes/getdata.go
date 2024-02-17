package routes

import (
	"encoding/json"
	"log"
	"mainweb/models"
	"net/http"
	base "mainweb/database"

	
)


func GetAllBuyerProduct(w http.ResponseWriter,r *http.Request){
    type buyer struct{
		Email_id string `json:"email_id" gorm:"not null"`
	    Phone_no string `json:"phone_no" gorm:"not null"`
		ProductInstance []models.Product_instance 
	}
	var buyerdata buyer
	products:=buyerdata.ProductInstance
	r.ParseForm()
	err:=base.Mydb.Db.Preload("ProductInstance").Find(&buyerdata,r.FormValue("EMAIL_ID"))
	if err!=nil{
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)


}
func GetAllProduct(w http.ResponseWriter,r *http.Request){
	type product struct{
		Brand string `json:"brand" gorm:"not null"`
	    Name string `json:"name" gorm:"not null"`
	    Price uint `json:"price" gorm:"not null"`
	    Total uint `json:"total" gorm:"not null"`
	}
	products:=[]product{}
	if err:=base.Mydb.Db.Find(&products);err!=nil{
		log.Println(err)
		
	}
    w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}