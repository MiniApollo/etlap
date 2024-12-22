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
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	restapi.AdminPassword = os.Getenv("ADMINPASS")
	restapi.JwtSecret = []byte(os.Getenv("JWT_SECRET"))

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

	// User API
	router.GET("/food", restapi.GetAllFood)
	router.GET("/food/:id", restapi.GetFood)
	router.POST("/order", restapi.PostOrder)

	// Admin API
	router.GET("/admin", restapi.CheckAdminPassword)

	router.POST("/food", restapi.CheckAdminToken, restapi.PostFood)
	router.PATCH("/food/:id", restapi.CheckAdminToken, restapi.UpdateFood)
	router.DELETE("/food/:id", restapi.CheckAdminToken, restapi.DeleteFood)

	router.GET("/order", restapi.CheckAdminToken, restapi.GetAllOrders)
	router.GET("/order/:id", restapi.CheckAdminToken, restapi.GetOrder)
	router.GET("/order/customer", restapi.CheckAdminToken, restapi.GetAllCustomerByOrder)
	router.GET("/order/food/:id", restapi.CheckAdminToken, restapi.GetAllFoodByCustomer)
	router.DELETE("/order/:id", restapi.CheckAdminToken, restapi.DeleteOrder)

	router.GET("/customer", restapi.CheckAdminToken, restapi.GetAllCustomer)
	router.GET("/customer/:id", restapi.CheckAdminToken, restapi.GetCustomer)
	router.DELETE("/customer/:id", restapi.CheckAdminToken, restapi.DeleteCustomer)

	router.Run("localhost:8080")
}
