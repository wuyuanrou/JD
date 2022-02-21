package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
)


func Register(c *gin.Context){
	name := c.Request.URL.Query().Get("name")
	password := c.Request.URL.Query().Get("password")

	Str1 := "insert into users.users(name, password) values(?, ?)"
	stmt,err:=db.Prepare(Str1)
	if err!=nil{
		fmt.Printf("Prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_,err = stmt.Exec(name, password)
	if err!=nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}//注册

	Str2:="select id from users.users where (name, password) = (?, ?)"
	stmt2, err:=db.Prepare(Str2)
	if err!=nil{
		fmt.Printf("Prepare id failed, err:%v\n", err)
		return
	}
	defer stmt2.Close()
	rows,err:=stmt2.Query(name, password)
	if err!=nil{
		fmt.Printf("Query failed, err=%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var u user
		err:=rows.Scan(&u.Id)
		if err!=nil{
			fmt.Printf("Scan failed, err=%v\n", err)
			return
		}

		c.JSON(200, gin.H{
			"success": true,
			"code": 200,
			"id": u.Id,
		})
	}//返回id账号
}



func Login(c *gin.Context){
	id:= c.Request.URL.Query().Get("id")
	password:=c.Request.URL.Query().Get("password")

	Str:="select id, name, money  from users.users where (id, password) = (?, ?)"
	var u user
	err:=db.QueryRow(Str, id, password).Scan(&u.Id, &u.Name, &u.Money)
	if err!=nil{
		fmt.Printf("Scan failed, err:%v\n", err)
		return
	}

	c.SetCookie("gin_cookie", "value_cookie", 120, "/", "localhost", false, true)
	//c.String(200, "Login success")

	c.JSON(200, gin.H{
		"success": true,
		"code":200,
		"id":u.Id,
		"name":u.Name,
		"money":u.Money,
	})

}//登录

//c.Request.URL.Query().Get
