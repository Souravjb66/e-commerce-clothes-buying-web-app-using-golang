package main

import (
	"log"
	db "mainweb/database"
	service "mainweb/routes"
	"net/http"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)
func main() {
	idb:=db.DBinstace{}
	idb.Connect()
	mydb:=db.Mydb.Db
	var wg sync.WaitGroup
	defer func(){
		lb,err:=mydb.DB()
		if err!=nil{
			log.Panic("db not close in last")
		}
		lb.Close()
	}()

	router := mux.NewRouter()
	callSave(router,&wg)
	callGet(router,&wg)
	frontend(router)
    wg.Wait()

    headers:=handlers.AllowedHeaders([]string{"Content-Type","Authorization"})
    methods:=handlers.AllowedMethods([]string{"GET","HEAD","POST","PUT","OPTIONS"})
    origins:=handlers.AllowedOrigins([]string{"*"})


	log.Fatal(http.ListenAndServe(":8080",handlers.CORS(headers,methods,origins)(router)))


	

}
func callSave(router *mux.Router,wg *sync.WaitGroup){
	wg.Add(1)
	func(){
		defer wg.Done()
		save:=router.PathPrefix("/savedata").Subrouter()
	    save.HandleFunc("/signup-user",service.Saveuser).Methods("POST")
	    save.HandleFunc("/save-seller",service.SaveSeller).Methods("POST")
	    save.HandleFunc("/save-buyer",service.SaveBuyer).Methods("POST")
	    save.HandleFunc("/save-products",service.SaveProducts).Methods("POST")
	    //this is to login then get all product 
	    router.HandleFunc("/access",service.LoginUser).Methods("POST")

	}()
	


}
func callGet(router *mux.Router,wg *sync.WaitGroup){
	wg.Add(1)
	func(){
		defer wg.Done()
		get:=router.PathPrefix("/getdata").Subrouter()
	    get.HandleFunc("/get-all-products",service.GetAllProduct).Methods("GET")
	    get.HandleFunc("/get-all-buyer-product",service.GetAllBuyerProduct).Methods("GET")
	}()
	
}
func frontend(router *mux.Router){
	
	router.HandleFunc("/",service.Indexfile).Methods("GET")
	router.HandleFunc("/login",service.Loginfile).Methods("GET")
	router.HandleFunc("/products",service.Storefile).Methods("GET")
	router.HandleFunc("/card",service.Cardfile).Methods("GET")
	router.HandleFunc("/create-seller",service.Createseller).Methods("GET")
	router.HandleFunc("/add-products",service.AddproductFile).Methods("GET")

}