package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)
type userinfo struct {
	username string `form:"username"`
	password string `form:"password"'`
}
func main() {
	c:=gin.Default()
	c.LoadHTMLFiles("test.html")
	c.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"test.html",gin.H{})
	})
	c.POST("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"test.html",gin.H{})
		fmt.Println(c.Param("username"))
		var user userinfo
		if err:=c.ShouldBind(&user);err!=nil{
			fmt.Println("fail")
			return
		}
		fmt.Println(user.username)
	})
	c.Run(":9000")
}