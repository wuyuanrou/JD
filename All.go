package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func main(){
	var r = gin.Default()
	initDB()

	r.GET("/register", Register)
	r.GET("/login", Login)


	Final:=r.Group("/users", cookie())
	{

		Final.GET("/home", home)
		Final.GET("/AddMoney", AddMoney)
		Final.POST("/AddLike", AddLike)
		Final.GET("/LikeList", LikeList)
		Final.POST("/AddCar", AddCar)
		Final.GET("/CarList", CarList)
		Final.GET("/Delete", Delete)

		Final.GET("/Search", search)
		Final.GET("/priceDesc", priceDesc)
		Final.GET("/sellDesc", sellDesc)
		Final.GET("/commentDesc", commentDesc)

		Final.GET("/Detail", detail)
		Final.GET("/ShowComment", showcomments)
		Final.POST("/AddComment", AddComments)

		Final.GET("/Show", show)
		Final.GET("/Pricedesc", Pricedesc)
		Final.GET("/Selldesc", Selldesc)
		Final.GET("/Commentdesc", Commentdesc)

	}
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
