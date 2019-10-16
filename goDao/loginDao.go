package goDao

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"

)
/*
const (
	host string = "127.0.0.1"
	username string = "root"
	password string = ""
	port string = "3306"
	database string = "test"
)

var (
	SqlDB *sql.DB
	err error
)

func databaseInit() (*sql.DB,error) {
	DBStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",username,password,host,port,database)
	fmt.Println("数据库连接： ",DBStr)
	return sql.Open("mysql", DBStr)
}

func selectUserById(id int){
	sqlStr := "select * from tab_user where id = ? "

	SqlDB, err := databaseInit()
	if err != nil{
		fmt.Println("连接失败，错误原因：", err)
	}
	stmt,_ := SqlDB.Prepare(sqlStr)
	rows, err := stmt.Query(id)
	if err != nil{
		fmt.Println("查询异常，错误原因：", err)
	}

	for rows.Next(){
		//封装结构体
	}

}
*/

func InsertData(db *sql.DB,err error,name string, password string){
	if err != nil{
		fmt.Println("连接失败：", err)
	}
	strSql := "insert into tab_user (loginName, userPassword) values (?,?)"
	stmt,_ :=db.Prepare(strSql)
	stmt.Exec(name, password)
	log.Println("插入成功！")
}



func DeleteData(db *sql.DB, err error , id string) int {
	SqlStr := "delete from tab_user where id = ?"
	stmt, err := db.Prepare(SqlStr)
	if err != nil {
		fmt.Println("删除操作运行错误：", err)
	}
	res,_ := stmt.Exec(id)
	fmt.Println("删除的返回值： ",res)

	return 0
}


func UpdateUser(db *sql.DB, err error, id string, name string, password string) int {
	SqlStr := "update tab_user set loginName= ?, userPassword= ? where id= ?"
	stmt,err := db.Prepare(SqlStr);
	if err != nil {
		fmt.Println("更新时异常: ", err)
	}
	stmt.Exec(name, password, id)
	return 0
}