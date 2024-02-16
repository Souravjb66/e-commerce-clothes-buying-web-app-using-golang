package middleware
import(
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateHash(password *string)string{
	bytedata:=[]byte(*password)
	hashPassword,err:=bcrypt.GenerateFromPassword(bytedata,bcrypt.DefaultCost)
	if err!=nil{
		log.Print("problem in hashing")
	}
	return string(hashPassword)



}