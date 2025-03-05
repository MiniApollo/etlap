package restapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCustomer(c *gin.Context) {
	rows, err := Db.Query("SELECT * FROM Vasarlok")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	var customers []customer

	defer rows.Close()
	for rows.Next() {
		var customer customer
		if err := rows.Scan(&customer.VasarloID, &customer.Nev, &customer.Email, &customer.Telefonszam, &customer.LeadasiIdo, &customer.ElkeszultIdo, &customer.Osszeg); err != nil {
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

func GetAllCustomerByOrder(c *gin.Context) {
	var isDoneParam string = c.DefaultQuery("isDone", "false")

	isDone, err := strconv.ParseBool(isDoneParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkNullQuery := "SELECT Vasarlok.* FROM Rendelesek INNER JOIN Vasarlok ON Rendelesek.VasarloID=Vasarlok.VasarloID WHERE Vasarlok.ElkeszultIdo is null GROUP BY VasarloID"
	if isDone {
		checkNullQuery = "SELECT Vasarlok.* FROM Rendelesek INNER JOIN Vasarlok ON Rendelesek.VasarloID=Vasarlok.VasarloID WHERE Vasarlok.ElkeszultIdo is not null GROUP BY VasarloID"
	}

	rows, err := Db.Query(checkNullQuery)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	var customers []customer

	defer rows.Close()
	for rows.Next() {
		var customer customer
		if err := rows.Scan(&customer.VasarloID, &customer.Nev, &customer.Email, &customer.Telefonszam, &customer.LeadasiIdo, &customer.ElkeszultIdo, &customer.Osszeg); err != nil {
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
	err := Db.QueryRow("SELECT * FROM Vasarlok WHERE VasarloID = ?", c.Param("id")).Scan(&Customer.VasarloID, &Customer.Nev, &Customer.Email, &Customer.Telefonszam, &Customer.LeadasiIdo, &Customer.ElkeszultIdo, &Customer.Osszeg)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, Customer)
}
