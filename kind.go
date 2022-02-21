package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//给大类编号，分为type_id(大类)和parent_id(小类)
//再创建一个关于分类id的表便于查询

func KindSearch(c *gin.Context){
	TypeId:=c.Request.URL.Query().Get("TypeId")

	Str:="select id, price, sell, comment, GoodCommentRate, name, picture, shopname from goods.goods where TypeId=?"
	rows, err:=db.Query(Str, TypeId)
	if err!=nil{
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var g goods
		err:=rows.Scan(&g.Id, &g.Price, &g.Sell, &g.Comment, &g.GoodCommentRate, &g.Name, &g.Picture, &g.Shop)
		if err!=nil{
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		c.JSON(200, gin.H{
			"GoodsId": g.Id,
			"price": g.Price,
			"sell": g.Sell,
			"comment": g.Comment,
			"GoodCommentRate": g.GoodCommentRate,
			"name": g.Name,
			"picture":g.Picture,
			"shopname": g.Shop,
		})	}
}