package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
)


func show(c *gin.Context){
	shop:=c.Request.URL.Query().Get("shop")
	Str:="select id, price, sell, comment, GoodCommentRate, name, picture from goods.goods where shopname=?"
	rows, err:=db.Query(Str, shop)
	if err!=nil{
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var g goods
		err:=rows.Scan(&g.Id, &g.Name, &g.Price, &g.Sell, &g.Comment, &g.GoodCommentRate, &g.Picture)
		if err!=nil{
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		c.JSON(200, gin.H{
			"GoodsId": g.Id,
			"name": g.Name,
			"price": g.Price,
			"sell": g.Sell,
			"comment": g.Comment,
			"GoodCommentRate": g.GoodCommentRate,
			"picture":g.Picture,
		})	}
}

func Pricedesc(c *gin.Context){
	shop:=c.Request.URL.Query().Get("shop")

	Str:="select id, name, price, sell, comment, GoodCommentRate, picture from goods.goods where shopname=? order by price desc"
	rows, err:=db.Query(Str, shop)
	if err!=nil{
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var g goods
		err:=rows.Scan(&g.Id, &g.Name, &g.Price, &g.Sell, &g.Comment, &g.GoodCommentRate, &g.Picture)
		if err!=nil{
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		c.JSON(200, gin.H{
			"GoodsId": g.Id,
			"name": g.Name,
			"price": g.Price,
			"sell": g.Sell,
			"comment": g.Comment,
			"GoodCommentRate": g.GoodCommentRate,
			"picture":g.Picture,
		})	}
}

func Commentdesc(c *gin.Context){
	shop:=c.Request.URL.Query().Get("shop")

	Str:="select id, name, price, sell, comment, GoodCommentRate, picture from goods.goods where shopname=? order by comment desc"
	rows, err:=db.Query(Str, shop)
	if err!=nil{
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var g goods
		err:=rows.Scan(&g.Id, &g.Name, &g.Price, &g.Sell, &g.Comment, &g.GoodCommentRate, &g.Picture)
		if err!=nil{
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		c.JSON(200, gin.H{
			"GoodsId": g.Id,
			"name": g.Name,
			"price": g.Price,
			"sell": g.Sell,
			"comment": g.Comment,
			"GoodCommentRate": g.GoodCommentRate,
			"picture":g.Picture,
		})	}
}

func Selldesc(c *gin.Context){
	shop:=c.Request.URL.Query().Get("shop")

	Str:="select id, name, price, sell, comment, GoodCommentRate, picture from goods.goods where shopname=? order by sell desc "
	rows, err:=db.Query(Str, shop)
	if err!=nil{
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var g goods
		err:=rows.Scan(&g.Id, &g.Name, &g.Price, &g.Sell, &g.Comment, &g.GoodCommentRate, &g.Picture)
		if err!=nil{
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		c.JSON(200, gin.H{
			"GoodsId": g.Id,
			"name": g.Name,
			"price": g.Price,
			"sell": g.Sell,
			"comment": g.Comment,
			"GoodCommentRate": g.GoodCommentRate,
			"picture":g.Picture,
		})	}
}


