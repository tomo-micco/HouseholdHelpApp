package main

import (
	_ "github.com/go-sql-driver/mysql"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// TODO: DB接続してハンドラーへのDIを行う
	// usersHandler := handlers.NewUsersHandler(db)
	// r.GET("/users", usersHandler.GetUsers)
	// r.GET("/users/:id", usersHandler.GetById)
	// r.POST("/users", usersHandler.CreateUser)
	// r.PUT("/users", usersHandler.UpdateUser)
	// r.DELETE("/users/:id", usersHandler.DeleteUser)

	r.Run()
}
