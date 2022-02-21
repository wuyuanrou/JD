package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func search(c *gin.Context){
	word:=c.Request.URL.Query().Get("word")

	Str:="select id, name, price, sell, comment, GoodCommentRate, shopname, picture from goods.goods where name LIKE '%电视%'"
	rows,err:=db.Query(Str, word)
	if err!= nil{
		fmt.Printf("Load falied, err:%v\n", err)
		return
	}

	defer rows.Close()
	for rows.Next(){
		var g goods
		err2:=rows.Scan(&g.Id, &g.Name, &g.Price, &g.Sell, &g.Comment, &g.GoodCommentRate, &g.Shop, &g.Picture)
		if err2!=nil{
			fmt.Printf("scan failed, err:%v\n", err2)
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
}

func priceDesc(c *gin.Context){
	//word:=c.Request.URL.Query().Get("word")

	Str:="select id, name, price, sell, comment, GoodCommentRate, shopname, picture from goods.goods where name LIKE '%电视%' order by price desc"
	rows,err:=db.Query(Str)
	if err!= nil{
		fmt.Printf("Load falied, err:%v\n", err)
		return
	}

	defer rows.Close()
	for rows.Next(){
		var g goods
		err2:=rows.Scan(&g.Id, &g.Name, &g.Price, &g.Sell, &g.Comment, &g.GoodCommentRate, &g.Shop, &g.Picture)
		if err2!=nil{
			fmt.Printf("scan failed, err:%v\n", err2)
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

func commentDesc(c *gin.Context){
	//word:=c.Request.URL.Query().Get("word")

	Str:="select id, name, price, sell, comment, GoodCommentRate, shopname, picture from goods.goods where name LIKE '%电视%' order by comment desc"
	rows,err:=db.Query(Str)
	if err!= nil{
		fmt.Printf("Load falied, err:%v\n", err)
		return
	}

	defer rows.Close()
	for rows.Next(){
		var g goods
		err2:=rows.Scan(&g.Id, &g.Name, &g.Price, &g.Sell, &g.Comment, &g.GoodCommentRate, &g.Shop, &g.Picture)
		if err2!=nil{
			fmt.Printf("scan failed, err:%v\n", err2)
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

func sellDesc(c *gin.Context){
	//word:=c.Request.URL.Query().Get("word")

	Str:="select id, name, price, sell, comment, GoodCommentRate, shopname, picture from goods.goods where name LIKE '%电视%' order by sell desc"
	rows,err:=db.Query(Str)
	if err!= nil{
		fmt.Printf("Load falied, err:%v\n", err)
		return
	}

	defer rows.Close()
	for rows.Next(){
		var g goods
		err2:=rows.Scan(&g.Id, &g.Name, &g.Price, &g.Sell, &g.Comment, &g.GoodCommentRate, &g.Shop, &g.Picture)
		if err2!=nil{
			fmt.Printf("scan failed, err:%v\n", err2)
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