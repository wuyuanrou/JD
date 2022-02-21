package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//订单有用户id, 商品id，商品名，价格， 商铺名， 订单状态
//待付款001， 待收货002， 待评价003， 退换售后004， 确认收货005

func OrderList(c *gin.Context){
	id:=c.Request.URL.Query().Get("id")

	Str:="select GoodsId, name, price, shopname, situ from users.order where id = ?"
	rows, err:=db.Query(Str, id)
	if err!=nil{
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var o order
		err:=rows.Scan(&o.GoodsId, &o.Name, &o.Price, &o.Shopname, &o.Situ)
		if err!=nil{
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		c.JSON(200, gin.H{
			"GoodsId": o.GoodsId,
			"Name": o.Name,
			"Price": o.Price,
			"Shopname": o.Shopname,
			"Situation": o.Situ,
		})
	}
}

func DeleteOrder(c *gin.Context){
	id:=c.Request.URL.Query().Get("id")
	GoodsId:=c.Request.URL.Query().Get("GoodsId")

	Str:="delete from users.order where id=? and GoodsId =?"
	ret, err:=db.Exec(Str, id, GoodsId)
	if err!=nil{
		fmt.Printf("Delete failed, err:%v\n", err)
		return
	}

	n, err2:=ret.RowsAffected()
	if err2!=nil{
		fmt.Printf("Get RowsAffected failed, err:%v\n", err)
		return
	}
	c.JSON(200, gin.H{
		"result": n,
	})
}

func UpdateSitu(c *gin.Context){
	id:=c.Request.URL.Query().Get("id")
	GoodsId:=c.Request.URL.Query().Get("GoodsId")

	Str:="update users.order set situ=? where id=? and GoodsId=?"
	ret, err:=db.Exec(Str, 005, id, GoodsId)
	if err!=nil{
		fmt.Printf("Update failed, err:%v\n", err)
		return
	}

	n, err2:=ret.RowsAffected()
	if err2!=nil{
		fmt.Printf("Get RowsAffected failed, err:%v\n", err)
		return
	}
	c.JSON(200, gin.H{
		"Rows Affected": n,
	})
}

func ReadyToPay(c *gin.Context){
	id:=c.Request.URL.Query().Get("id")
	GoodsId:=c.Request.URL.Query().Get("GoodsId")

	Str:="select GoodsId, name, price, shopname, situ from users.order where (id, GoodsId, situ)=(?, ?, ?)"
	rows, err:=db.Query(Str, id, GoodsId, 001)
	if err!=nil{
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var o order
		err:=rows.Scan(&o.GoodsId, &o.Name, &o.Price, &o.Shopname, &o.Situ)
		if err!=nil{
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		c.JSON(200, gin.H{
			"GoodsId": o.GoodsId,
			"Name": o.Name,
			"Price": o.Price,
			"Shopname": o.Shopname,
			"Situation": &o.Situ,
		})
	}
}

func ReadyToReceive(c *gin.Context){
	id:=c.Request.URL.Query().Get("id")
	GoodsId:=c.Request.URL.Query().Get("GoodsId")

	Str:="select GoodsId, name, price, shopname, situ from users.order where (id, GoodsId, situ)=(?, ?, ?)"
	rows, err:=db.Query(Str, id, GoodsId, 002)
	if err!=nil{
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var o order
		err:=rows.Scan(&o.GoodsId, &o.Name, &o.Price, &o.Shopname, &o.Situ)
		if err!=nil{
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		c.JSON(200, gin.H{
			"GoodsId": o.GoodsId,
			"Name": o.Name,
			"Price": o.Price,
			"Shopname": o.Shopname,
			"Situation": o.Situ,
		})
	}
}

func ReadyToComment(c *gin.Context){
	id:=c.Request.URL.Query().Get("id")
	GoodsId:=c.Request.URL.Query().Get("GoodsId")

	Str:="select GoodsId, name, price, shopname, situ from users.order where (id, GoodsId, situ)=(?, ?, ?)"
	rows, err:=db.Query(Str, id, GoodsId, 003)
	if err!=nil{
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var o order
		err:=rows.Scan(&o.GoodsId, &o.Name, &o.Price, &o.Shopname, &o.Situ)
		if err!=nil{
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		c.JSON(200, gin.H{
			"GoodsId": o.GoodsId,
			"Name": o.Name,
			"Price": o.Price,
			"Shopname": o.Shopname,
			"Situation": o.Situ,
		})
	}
}

func AfterBuy(c *gin.Context){
	id:=c.Request.URL.Query().Get("id")
	GoodsId:=c.Request.URL.Query().Get("GoodsId")

	Str:="select GoodsId, name, price, shopname, situ from users.order where (id, GoodsId, situ)=(?, ?, ?)"
	rows, err:=db.Query(Str, id, GoodsId, 004)
	if err!=nil{
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var o order
		err:=rows.Scan(&o.GoodsId, &o.Name, &o.Price, &o.Shopname, &o.Situ)
		if err!=nil{
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		c.JSON(200, gin.H{
			"GoodsId": o.GoodsId,
			"Name": o.Name,
			"Price": o.Price,
			"Shopname": o.Shopname,
			"Situation": o.Situ,
		})
	}
}


