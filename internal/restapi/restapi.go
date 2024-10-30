package restapi

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// https://stackoverflow.com/questions/35038864/how-to-access-global-variables
var Db *sql.DB

type customer struct {
	VasarloID   int    `json:"VasarloID"`
	Nev         string `json:"Nev"`
	Email       string `json:"Email"`
	Telefonszam int    `json:"Telefonszam"`
}

type order struct {
	VasarloID int `json:"VasarloID"`
	EtelID    int `json:"EtelID"`
}

type food struct {
	EtelID int     `json:"EtelID"`
	Nev    string  `json:"Nev"`
	Leiras string  `json:"Leiras"`
	Kep    string  `json:"Kep"`
	Ar     float64 `json:"Ar"`
}

// Capital Letters mean public
func GetAllFood(c *gin.Context) {
	rows, err := Db.Query("SELECT * FROM Etelek")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}
	var foods []food

	defer rows.Close()
	for rows.Next() {
		var food food
		if err := rows.Scan(&food.EtelID, &food.Nev, &food.Leiras, &food.Kep, &food.Ar); err != nil {
			log.Fatal(err)
		}
		foods = append(foods, food)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
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
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}
	var orders []order

	defer rows.Close()
	for rows.Next() {
		var order order
		if err := rows.Scan(&order.VasarloID, &order.EtelID); err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil || orders == nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	rows, err := Db.Query("SELECT * FROM Rendelesek WHERE VasarloId = ? ORDER BY VasarloID, EtelID", c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}
	var orders []order

	defer rows.Close()
	for rows.Next() {
		var order order
		if err := rows.Scan(&order.VasarloID, &order.EtelID); err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, customers)
}

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
