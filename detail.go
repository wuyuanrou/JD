package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func detail(c *gin.Context){
	GoodsId:= c.Request.URL.Query().Get("GoodsId")

	Str:="select id, name, price, sell, comment, GoodCommentRate, shopname, picture from goods.goods where id = ?"
	var g goods
	err:=db.QueryRow(Str, GoodsId).Scan(&g.Id, &g.Name, &g.Price, &g.Sell, &g.Comment, &g.GoodCommentRate, &g.Shop, &g.Picture)
	if err!= nil{
		fmt.Printf("Query failed, err:%v\n", err)
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
	})
}