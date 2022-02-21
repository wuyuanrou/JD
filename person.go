package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//加入购物车和加入关注所要参数太多

func AddMoney(c *gin.Context){
	add :=c.Request.URL.Query().Get("add")
	id:=c.Request.URL.Query().Get("id")

	newadd, _:= strconv.Atoi(add)
	Str1:="select money from users.users where id = ?"
	var u user
	err:=db.QueryRow(Str1, id).Scan(&u.Money)
	if err!=nil{
		fmt.Printf("Scan failed, err:%v\n", err)
		return
	}

	nmoney:=u.Money+newadd
	Str:="update users.users set money = ? where id = ?"
	ret, err:=db.Exec(Str, nmoney, id)
	if err!=nil{
		fmt.Printf("Update failed, err:%v\n", err)
		return
	}
	c.JSON(200, gin.H{
		"result": ret,
	})
}
//用的是用户列表

//用的是关注列表

func AddLike(c *gin.Context){
	id:=c.PostForm("id")
	GoodsId:=c.PostForm("GoodsId")
	price:=c.PostForm("price")
	shopname:=c.PostForm("shopname")
	name:=c.PostForm("name")

	Str:="insert into users.likes(id, GoodsId, price, shopname, name) values(?, ?, ?, ?, ?)"
	ret, err:=db.Exec(Str, id, GoodsId, price, shopname, name)
	if err!=nil{
		fmt.Printf("Add failed, err:%v\n", err)
		return
	}
	n, err2:=ret.RowsAffected()
	if err2!=nil{
		fmt.Printf("Rows affected failed, err:%v\n", err2)
		return
	}
	c.JSON(200, gin.H{
		"Affected Rows": n,
	})
}

func LikeList(c *gin.Context){
	id:=c.Request.URL.Query().Get("id")

	Str:="select GoodsId, name, price ,shopname from users.likes where id=?"//用户的喜爱列表
	rows, err:=db.Query(Str, id)
	if err!=nil{
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var g goods
		err:=rows.Scan(&g.Id, &g.Name, &g.Price, &g.Shop, &g.Picture)
		if err!=nil{
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		c.JSON(200, gin.H{
			"GoodsId": g.Id,
			"name": g.Name,
			"price": g.Price,
			"shop": g.Shop,
			"picture": g.Picture,
		})
	}
}
//用的是关注列表

//用的是购物车列表


func CarList(c *gin.Context){
	id:=c.Request.URL.Query().Get("id")

	Str:="select GoodsId, name, price, shopname, picture from car where id=?"//用户的购物车列表
	rows, err:=db.Query(Str, id)
	if err!=nil{
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var g goods
		err:=rows.Scan(&g.Id, &g.Name, &g.Price, &g.Shop, &g.Picture)
		if err!=nil{
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		c.JSON(200, gin.H{
			"GoodsId": g.Id,
			"name": g.Name,
			"price": g.Price,
			"shop": g.Shop,
			"picture": g.Picture,
		})
	}
}

func AddCar(c *gin.Context){
	id:=c.PostForm("id")
	GoodsId:=c.PostForm("GoodsId")
	price:=c.PostForm("price")
	shopName:=c.PostForm("shopName")
	picture := c.PostForm("picture")

	Str:="insert into car(id, GoodsId, price, shopname, picture) values(?, ?, ?, ?, ?)"
	ret, err:=db.Exec(Str, id, GoodsId, price, shopName, picture)
	if err!=nil{
		fmt.Printf("Add failed, err:%v\n", err)
		return
	}
	n, err2:=ret.RowsAffected()
	if err2!=nil{
		fmt.Printf("Rows affected wrong, err:%v\n", err2)
		return
	}
	c.JSON(200, gin.H{
		"Affected Rows": n,
	})
}

func sum(c *gin.Context){
	id:=c.Request.URL.Query().Get("id")
	GoodsId:=c.Request.URL.Query().Get("GoodsId")

	Str:="update users.car set flag = 1 where GoodsId=? and id=?"
	ret, err:=db.Exec(Str, GoodsId, id)
	if err!=nil{
		fmt.Printf("Update failed, err:%v\n", err)
		return
	}
	n, err2:=ret.RowsAffected()
	if err2!=nil{
		fmt.Printf("Rows affected failed, err:%v\n", err2)
		return
	}
	c.JSON(200, gin.H{
		"Affected Rows": n,
	})

	var total int
	Str2:="select sum(price) as total from users.car where flag = 1 and id=?"
	err3:=db.QueryRow(Str2, id).Scan(&total)
	if err3!=nil{
		fmt.Printf("Sum failed, err:%v\n", err3)
		return
	}
	c.JSON(200, gin.H{
		"Total": total,
	})
}

func Delete(c *gin.Context){
	id:=c.Request.URL.Query().Get("id")
	GoodsId:=c.Request.URL.Query().Get("GoodsId")

	Str:="delete from car where GoodsId=? and id=?"
	ret, err:=db.Exec(Str, GoodsId, id)
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
//用的是购物车列表
