package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"GoDemo/service"
)

var (
	SqlDB *sql.DB
	err error
)


func Dinit(db *sql.DB, err error){
	SqlDB  = db
	err =  err
}


func UserLogin(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form["username"], "------用户名")
	fmt.Println(r.Form["password"], "------密码")
	w.Write([]byte("login success..... "))
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	service.RegistryUser(SqlDB, err ,username, password)

}

func RemoveUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	res := service.RemoveUser(SqlDB, err, r.Form.Get("id"))
	w.Write([]byte(string(res)))
}

func ModifyUser(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	service.ModifyUser(SqlDB, err, r.Form.Get("id"),r.Form.Get("username"), r.Form.Get("password"))
	log.Println("修改执行完成！")
}