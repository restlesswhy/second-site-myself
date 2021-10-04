package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var router *gin.Engine
var connection *sql.DB
var connectionString = "host=127.0.0.1 port=5432 user=exezz password=1234 dbname=exezz sslmode=disable"

func main() {
	var e error

	connection, e = sql.Open("postgres", connectionString)
	if e != nil {
		panic(e)
	}
	defer connection.Close()

	router = gin.Default()
	router.Static("/assets/", "static/")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/registration", reg)
	router.GET("/authorization", auth)
	router.GET("/", index)
	router.POST("/user/reg", userRegistration)
	router.POST("/user/auth", userAuthorization)
	_ = router.Run(":8080")
}

func reg(c *gin.Context) {
	c.HTML(200, "registration.html", gin.H{})
}

func auth(c *gin.Context) {
	c.HTML(200, "authorization.html", gin.H{})
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"Role": "manager",
	})
}

func userRegistration(c *gin.Context) {
	var user User

	e := c.BindJSON(&user)
	if e != nil {
		c.JSON(200, gin.H{
			"error": e.Error(),
		})
		return
	}

	e = user.Create()
	if e != nil {
		c.JSON(200, gin.H{
			"error": "cant",
		})
		return
	}

	c.JSON(200, gin.H{
		"Error": nil,
	})
}

func userAuthorization(c *gin.Context) {
	var user User

	e := c.BindJSON(&user)
	if e != nil {
		c.JSON(200, gin.H{
			"error": e.Error(),
		})
		return
	}

	e = user.Select()
	if e != nil {
		c.JSON(200, gin.H{
			"error": "cant auth",
		})
		return
	}

	c.JSON(200, gin.H{
		"Error": nil,
	})
}