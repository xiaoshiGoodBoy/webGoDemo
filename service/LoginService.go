package service


import(
	"GoDemo/goDao"
	"database/sql"
	"log"
)

func loginJudge(name string, password string){


}

func RegistryUser(db *sql.DB,err error,username string , password string){
	goDao.InsertData(db, err, username, password)
	log.Println("service")
}

func RemoveUser(db *sql.DB,err error,id string) int {
	return goDao.DeleteData(db,err,id)
}

func ModifyUser(db *sql.DB, err error, id string, name string, password string) int {
	return goDao.UpdateUser(db,err,id,name,password)
}