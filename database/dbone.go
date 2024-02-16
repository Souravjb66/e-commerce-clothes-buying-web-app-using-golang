package database
import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	table "mainweb/models"
)

type DBinstace struct{
	Db *gorm.DB
}
var Mydb DBinstace
func(instance *DBinstace) Connect(){
	dsn:="root:sourav@90###@tcp(localhost:3306)/clouthstore?parseTime=true"
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Panic("problem in db connection--!!")

	}
	Mydb.Db=db
	err=db.AutoMigrate(&table.User{},&table.Seller{},&table.Product{},&table.Product_instance{},&table.Payment{},&table.Buyer{},&table.Auth{})
	if err!=nil{
		log.Print("table not created ----!!")
	}
	
}

