package main

import (
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/tomo-micco/HouseholdHelpApp/internal/handlers"

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

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	dbConfig := mysql.Config{
		Net:                  os.Getenv("DB_NET"),
		Addr:                 os.Getenv("DB_ADDR") + ":" + os.Getenv("DB_PORT"),
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  nil,
	}

	db := sqlx.MustOpen("mysql", dbConfig.FormatDSN())
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Error cannot connect to database")
		return
	}

	// TODO: DB接続してハンドラーへのDIを行う
	usersHandler := handlers.NewUsersHandler(db)
	r.GET("/users", usersHandler.GetUsers)
	r.GET("/users/:id", usersHandler.GetById)
	r.POST("/users", usersHandler.CreateUser)
	r.PUT("/users", usersHandler.UpdateUser)
	r.DELETE("/users/:id", usersHandler.DeleteUser)

	r.Run()
}
