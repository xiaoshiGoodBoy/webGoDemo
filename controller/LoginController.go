package controller

import (
	"GoDemo/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	SqlDB *sql.DB
	err error
)


func Dinit(db *sql.DB, err error){
	SqlDB  = db
	err =  err
}


func UserLogin(w http.ResponseWriter, r *http.Request ){
	w.Header().Set("Access-Control-Allow-Origin","*")
	r.ParseForm()
	fmt.Println(r.Form["username"], "------用户名")
	fmt.Println(r.Form["password"], "------密码")
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	service.RegistryUser(SqlDB, err ,username, password)
	w.Write([]byte("loading success ...... "))
}

func RemoveUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	service.RemoveUser(SqlDB, err, r.Form.Get("id"))
	w.Write([]byte(string("success")))
}

func ModifyUser(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	service.ModifyUser(SqlDB, err, r.Form.Get("id"),r.Form.Get("username"), r.Form.Get("password"))
	log.Println("修改执行完成！")
}

func FindUserById(w http.ResponseWriter, r *http.Request){
	//跨域请求需要添加此项
	w.Header().Set("Access-Control-Allow-Origin","*")

	r.ParseForm()
	studentPojo := service.FindUser(SqlDB, err, r.Form.Get("id"))
	fmt.Println(studentPojo, "-.--.-controller" )
	by,err :=json.Marshal(studentPojo)
	if err != nil{
		fmt.Println("错误信息：", err)
	}
	w.Write(by)
}

/*func FindUserById(c *gin.Context){

	studentPojo := service.FindUser(SqlDB, err, c.Param("id"))
	fmt.Println(studentPojo, "-.--.-controller" )
	c.Header("Content-Type", "application/json;charset=utf8")
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":   studentPojo,
	})
}*/