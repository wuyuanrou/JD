package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


var db = &sql.DB{}



type user struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Password int `json:"password"`
	Money int `json:"money"`
}

type goods struct{
	Kind int `json:"kind"`
	Price int `json:"price"`
	Id int `json:"id"`
	Sell int   `json:"sell"`
	Comment int `json:"comment"`
	GoodCommentRate int `json:"good_comment_rate"`
	Shop string `json:"shop"`
	Name string `json:"name"`
	Picture string `json:"picture"`
	TypeId int `json:"type_id"`
}

type comment struct{
	Id int `json:"id"`
	Name string `json:"name"`
	GoodsId int `json:"goods_id"`
	Context string `json:"context"`
	Shopname string `json:"shopname"`
}

type order struct {
	GoodsId int `json:"goods_id"`
	Name string `json:"name"`
	Shopname string `json:"shopname"`
	Situ int `json:"situ"`
	Price int `json:"price"`
}

func initDB(){
	var err error
	connStr := "root:123456@tcp(localhost)/users"
	db, err = sql.Open("mysql", connStr)
	if err!=nil{
		log.Fatal(err)
		return
	}//开数据库
}

func cookie()gin.HandlerFunc {
	return func(c *gin.Context) {
		v, err := c.Cookie("gin_cookie")
		if err == nil {
			if v == "value_cookie" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "未登录"})
		c.Abort()
		return
	}
}

func home(c *gin.Context){
	c.JSON(200, gin.H{
		"data": "home",
	})
}






