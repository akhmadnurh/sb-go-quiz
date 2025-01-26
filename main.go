package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sb-go-quiz/controllers"
	"sb-go-quiz/databases"
	"sb-go-quiz/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appPort := os.Getenv("APP_PORT")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`,
		host, port, user, password, dbname, sslmode,
	)
	fmt.Println(psqlInfo)

	DB, err = sql.Open("postgres", psqlInfo)
	defer DB.Close()

	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	databases.DBMigrate(DB)

	router := gin.Default()

	api := router.Group("/api")
	api.Use(middlewares.MAuth(DB))
	{
		api.GET("/categories", controllers.CFindCategories)
		api.GET("/categories/:id", controllers.CFindCategory)
		api.GET("/categories/:id/books", controllers.CFindBooksByCategory)
		api.POST("/categories", controllers.CCreateCategories)
		api.PUT("/categories/:id", controllers.CUpdateCategories)
		api.DELETE("/categories/:id", controllers.CDeleteCategories)
		api.GET("/books", controllers.CFindBooks)
		api.GET("/books/:id", controllers.CFindBook)
		api.POST("/books", controllers.CCreateBook)
		api.PUT("/books/:id", controllers.CUpdateBook)
		api.DELETE("/books/:id", controllers.CDeleteBook)
	}

	router.Run(fmt.Sprintf(":%s", appPort))

}
