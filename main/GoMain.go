package main

import(
	"GoDemo/controller"
	"database/sql"
	"fmt"
	// "github.com/kardianos/govendor/context"
	"log"
	"net/http"
)

const (
	host string = "127.0.0.1"
	port string = "3306"
	username string = "root"
	password string = ""
	database string = "test"
)

func main() {
	log.Println("go server start ...")
	SqlDB, err := initDatabase()
	controller.Dinit(SqlDB, err)
	http.HandleFunc("/login", controller.UserLogin)
	http.HandleFunc("/delete", controller.RemoveUser)
	http.HandleFunc("/update", controller.ModifyUser)
	http.ListenAndServe(":8080", nil)
	// x-www-form-urlencoded; charset=UTF-8

}


func initDatabase() (*sql.DB, error) {
	mysqlDriver := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=UTF8", username, password, host, port, database)
	fmt.Println("数据库driver: ",mysqlDriver)
	SqlDB, err := sql.Open("mysql", mysqlDriver)
	log.Print("数据库初始化完成！")
	return SqlDB,err
}
