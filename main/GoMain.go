package main

import (
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
	//router := httprouter.New()

	log.Println("go server start ...")
	SqlDB, err := initDatabase()
	controller.Dinit(SqlDB, err)
	http.HandleFunc("/login", controller.UserLogin)
	http.HandleFunc("/delete", controller.RemoveUser)
	http.HandleFunc("/update", controller.ModifyUser)
	http.HandleFunc("/userById", controller.FindUserById)

	// router.POST("/userById", controller.FindUserById)
	http.ListenAndServe(":8080", nil)

	//http.ListenAndServe(":8080", router)
	// x-www-form-urlencoded; charset=UTF-8

}


func initDatabase() (*sql.DB, error) {
	mysqlDriver := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=UTF8", username, password, host, port, database)
	fmt.Println("数据库driver: ",mysqlDriver)
	SqlDB, err := sql.Open("mysql", mysqlDriver)
	log.Print("数据库初始化完成！")
//	engine := gin.Default()
//	engine.Use(Cors())
	return SqlDB,err
}



// 允许跨域中间件
/*func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")  // 允许跨域访问(实际上只有这个才是控制是否允许跨域访问)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")  // 允许的访问的方法
		c.Header("Access-Control-Allow-Credentials", "true")  // 跨域是否允许携带cookies
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()
	}
}*/