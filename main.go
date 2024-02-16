package main

import (
	"log"
	db "mainweb/database"
	"net/http"
	"github.com/gorilla/mux"
	service "mainweb/routes"

)
func main() {
	idb:=db.DBinstace{}
	idb.Connect()
	mydb:=idb.Db
	defer func(){
		lb,err:=mydb.DB()
		if err!=nil{
			log.Panic("db not close in last")
		}
		lb.Close()
	}()
	router := mux.NewRouter()
	callSave(router)
	callGet(router)
	frontend(router)

	http.ListenAndServe(":8080",router)

}
func callSave(router *mux.Router){
	save:=router.PathPrefix("/savedata").Subrouter()
	save.HandleFunc("/signup-user",service.Saveuser).Methods("POST")
	save.HandleFunc("/save-seller",service.SaveSeller).Methods("POST")
	save.HandleFunc("/save-buyer",service.SaveBuyer).Methods("POST")
	save.HandleFunc("/save-products",service.SaveProducts).Methods("POST")
	
	router.HandleFunc("/access",service.LoginUser).Methods("POST")


}
func callGet(router *mux.Router){
	
	get:=router.PathPrefix("/getdata").Subrouter()
	get.HandleFunc("/get-all-products",service.GetAllProduct).Methods("GET")
	get.HandleFunc("/get-all-buyer-product",service.GetAllBuyerProduct).Methods("GET")

}
func frontend(router *mux.Router){
	router.HandleFunc("/",service.Indexfile).Methods("GET")
	router.HandleFunc("/login",service.Loginfile).Methods("GET")
	router.HandleFunc("/products",service.Storefile).Methods("GET")
	router.HandleFunc("/card",service.Cardfile).Methods("GET")
}