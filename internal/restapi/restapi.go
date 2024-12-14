package restapi

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// https://stackoverflow.com/questions/35038864/how-to-access-global-variables
var Db *sql.DB
var AdminPassword string
var JwtSecret []byte

type customer struct {
	VasarloID   int64  `json:"VasarloID"` // Not needed, because of auto increment
	Nev         string `json:"Nev" binding:"required"`
	Email       string `json:"Email" binding:"required"`
	Telefonszam string `json:"Telefonszam" binding:"required"`
}

type orderRow struct {
	VasarloID int `json:"VasarloID"`
	EtelID    int `json:"EtelID"`
}

type order struct {
	Customer customer `json:"Customer" binding:"required"`
	Foods    []food   `json:"Foods" binding:"required"`
}

type food struct {
	EtelID int     `json:"EtelID" binding:"required"`
	Nev    string  `json:"Nev" binding:"required"`
	Leiras string  `json:"Leiras" `
	Kep    string  `json:"Kep" `
	Ar     float64 `json:"Ar" binding:"required"`
}

// Capital Letters mean public
func GetAllFood(c *gin.Context) {
	rows, err := Db.Query("SELECT * FROM Etelek")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	var foods []food

	defer rows.Close()
	for rows.Next() {
		var food food
		if err := rows.Scan(&food.EtelID, &food.Nev, &food.Leiras, &food.Kep, &food.Ar); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		foods = append(foods, food)
	}
	if foods == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empty database"})
		return
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foods)
}

func GetFood(c *gin.Context) {
	var food food
	err := Db.QueryRow("SELECT * FROM Etelek WHERE EtelID = ?", c.Param("id")).Scan(&food.EtelID, &food.Nev, &food.Leiras, &food.Kep, &food.Ar)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	c.JSON(http.StatusOK, food)
}

func PostFood(c *gin.Context) {

}

func UpdateFood(c *gin.Context) {

}

func DeleteFood(c *gin.Context) {

}

func GetAllOrders(c *gin.Context) {
	rows, err := Db.Query("SELECT * FROM Rendelesek ORDER BY VasarloID, EtelID")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	var orders []orderRow

	defer rows.Close()
	for rows.Next() {
		var order orderRow
		if err := rows.Scan(&order.VasarloID, &order.EtelID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		orders = append(orders, order)
	}
	// Rows is not nil but rows.Err() is nil if database is empty
	// when err := rows.Err() nil pointer created and err.Error()
	// tries to dereference it
	// Needed othervise nil pointer dereference
	if orders == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empty database"})
		return
	}

	if err := rows.Err(); err != nil || orders == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	rows, err := Db.Query("SELECT * FROM Rendelesek WHERE VasarloId = ? ORDER BY VasarloID, EtelID", c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}
	var orders []orderRow

	defer rows.Close()
	for rows.Next() {
		var order orderRow
		if err := rows.Scan(&order.VasarloID, &order.EtelID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil || orders == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func PostOrder(c *gin.Context) {
	// 4 adatot lementeni adatbázisba: név, email, telefonszám, ételek listája
	// Létrehozni egy új vevőt
	// Kapcsoló táblába berakni az új vevő ID-ét és eggyesével az ételeket

	// TODO: Better Error handling
	// Check for empty foods list
	var newOrder order
	if err := c.BindJSON(&newOrder); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCustomer, err := Db.Exec("INSERT INTO Vasarlok (Nev, Email, Telefonszam) VALUES (?, ?, ?)", newOrder.Customer.Nev, newOrder.Customer.Email, newOrder.Customer.Telefonszam)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	// _ ignore set variable
	// https://stackoverflow.com/questions/27764421/what-is-underscore-comma-in-a-go-declaration
	newOrder.Customer.VasarloID, _ = newCustomer.LastInsertId()

	for i, e := range newOrder.Foods {
		_, err := Db.Exec("INSERT INTO Rendelesek (VasarloID, EtelID) VALUES (?, ?)", newOrder.Customer.VasarloID, e.EtelID)
		if err != nil {
			fmt.Println("Failed to insert into Rendelesek table at: ", i)
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusCreated, newOrder)
}

func DeleteOrder(c *gin.Context) {

}

func GetAllCustomers(c *gin.Context) {
	rows, err := Db.Query("SELECT * FROM Vasarlok")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}
	var customers []customer

	defer rows.Close()
	for rows.Next() {
		var customer customer
		if err := rows.Scan(&customer.VasarloID, &customer.Nev, &customer.Email, &customer.Telefonszam); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		customers = append(customers, customer)
	}
	if customers == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empty database"})
		return
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customers)
}

// There are some example users, skip some IDs when using the API
func GetCustomer(c *gin.Context) {
	var Customer customer
	err := Db.QueryRow("SELECT * FROM Vasarlok WHERE VasarloID = ?", c.Param("id")).Scan(&Customer.VasarloID, &Customer.Nev, &Customer.Email, &Customer.Telefonszam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}
	c.JSON(http.StatusOK, Customer)
}

func DeleteCustomer(c *gin.Context) {

}

func CheckAdminPassword(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Authorization header"})
		return
	}

	password, err := base64.StdEncoding.DecodeString(authParts[1])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if string(password) != AdminPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}
	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := generatedToken.SignedString(JwtSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate token"})
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}

func CheckAdminToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	// c.AbortWithStatus needed for checker functions
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Authorization header"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authToken, err := base64.StdEncoding.DecodeString(authParts[1])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-ErrorChecking
	accessToken, err := jwt.Parse(string(authToken), func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JwtSecret), nil
	})
	if err != nil || !accessToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error1": err.Error()})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	claims, ok := accessToken.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error2": err.Error()})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Succesful check, go to api function
	c.Next()
}
