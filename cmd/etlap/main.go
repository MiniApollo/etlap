package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/MiniApollo/etlap/internal/restapi"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Make db global variable so we can access it from anywhere
// https://stackoverflow.com/questions/34046194/how-to-pass-arguments-to-router-handlers-in-golang-using-gin-web-framework
var db *sql.DB

func main() {
	errEnv := godotenv.Load("database.env")
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "etlap",
		// Needed for mysql.config
		// https://stackoverflow.com/questions/70757210/how-do-i-connect-to-a-mysql-instance-without-using-the-password
		AllowNativePasswords: true,
	}
	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to database!")
	restapi.Db = db

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/food", restapi.GetAllFood)
	router.GET("/food/:id", restapi.GetFood)
	router.POST("/food", restapi.PostFood)
	router.PATCH("/food/:id", restapi.UpdateFood)
	router.DELETE("/food/:id", restapi.DeleteFood)

	router.GET("/order", restapi.GetAllOrders)
	router.GET("/order/:id", restapi.GetOrder)
	router.POST("/order", restapi.PostOrder)
	router.DELETE("/order/:id", restapi.DeleteOrder)

	router.GET("/customer", restapi.GetAllCustomers)
	router.GET("/customer/:id", restapi.GetCustomer)
	router.DELETE("/customer/:id", restapi.DeleteCustomer)

	router.Run("localhost:8080")
}
