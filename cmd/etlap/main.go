// Main fájl elindítja az etlap webszervert.
//
// A szerver az alábbi lépésekkel indul el:
//  1. A környezeti változók beolvasása.
//  2. Csatlakozás a MySQL szerverhez.
//  3. CORS konfiguráció beállítása.
//  4. API végpontok regisztrálása.
//  5. A webszerver indítása a konfigurált porton.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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
		Addr:   os.Getenv("DB_PORT"),
		DBName: os.Getenv("DBName"),
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

	// Cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")

	router.Use(cors.New(config))

	// User API
	routes := []string{"/", "/etelek", "/kosar", "/admin"}
	for _, url := range routes {
		router.StaticFile(url, "./assets/index.html")
	}

	router.StaticFile("/favicon.jpg", "./assets/favicon.jpg")
	router.Static("/assets", "./assets/")

	router.GET("/food", restapi.GetAllFood)
	router.GET("/food/:id", restapi.GetFood)
	router.POST("/order", restapi.PostOrder)

	// Admin API
	router.GET("/auth", restapi.CheckAdminPassword)

	router.POST("/food", restapi.CheckAdminToken, restapi.PostFood)
	router.PATCH("/food/:id", restapi.CheckAdminToken, restapi.UpdateFood)
	router.DELETE("/food/:id", restapi.CheckAdminToken, restapi.DeleteFood)

	router.GET("/order", restapi.CheckAdminToken, restapi.GetAllOrders)
	router.GET("/order/:id", restapi.CheckAdminToken, restapi.GetOrder)
	router.GET("/order/customer", restapi.CheckAdminToken, restapi.GetAllCustomerByOrder)
	router.GET("/order/food/:id", restapi.CheckAdminToken, restapi.GetAllFoodByCustomer)

	router.GET("/customer", restapi.CheckAdminToken, restapi.GetAllCustomer)
	router.GET("/customer/:id", restapi.CheckAdminToken, restapi.GetCustomer)
	router.PATCH("/customer/:id", restapi.CheckAdminToken, restapi.CompleteOrder)

	// 404 page
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
		c.File("./assets/index.html")
	})

	port := os.Getenv("PORT")
	if len(port) <= 0 {
		port = "localhost:8080"
	}
	router.Run(port)
}
