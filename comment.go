package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
)

//评价表（区）有用户id(id)，商品id(GoodsId),  用户名， 评价内容

func AddComments(c *gin.Context){
	id:=c.PostForm("id")
	name:=c.PostForm("name")
	context:=c.PostForm("context")
	GoodsId:= c.PostForm("GoodsId")
	shopname:=c.PostForm("shopname")

	Str:="insert into comment.comment(id, name, context, GoodsId, shopname) values (?, ?, ?, ?, ?)"
	_, err:=db.Exec(Str, id, name, context, GoodsId, shopname)
	if err!=nil{
		fmt.Printf("Insert failed, err:%v\n", err)
		return
	}
}

func showcomments(c *gin.Context){
	GoodsId:=c.Request.URL.Query().Get("GoodsId")

	Str:="select id, name, context from comment.comment where GoodsId = ?"
	ret, err:=db.Query(Str, GoodsId)
	if err!=nil{
		 fmt.Printf("Scan failed, err:%v\n", err)
		return
	}
	defer ret.Close()
	for ret.Next(){
		var k comment
		err2:=ret.Scan(&k.Id, &k.Name, &k.Context)
		if err2!=nil{
			fmt.Printf("Scan failed, err:%v\n", err2)
			return
		}
		c.JSON(200, gin.H{
			"id": k.Id,
			"name": k.Name,
			"context": k.Context,
		})
	}
}