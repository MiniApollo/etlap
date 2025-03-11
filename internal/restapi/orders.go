package restapi

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(c *gin.Context) {
	rows, err := Db.Query("SELECT * FROM Rendelesek ORDER BY VasarloID, EtelID")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	var orders []order

	defer rows.Close()
	for rows.Next() {
		var order order
		if err := rows.Scan(&order.VasarloID, &order.EtelID, &order.Ar, &order.Darab); err != nil {
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	var orders []order

	defer rows.Close()
	for rows.Next() {
		var order order
		if err := rows.Scan(&order.VasarloID, &order.EtelID, &order.Ar, &order.Darab); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil || orders == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func PostOrder(c *gin.Context) {
	var newFullOrder fullOrder
	if err := c.BindJSON(&newFullOrder); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCustomer, err := Db.Exec("INSERT INTO Vasarlok (Nev, Email, Telefonszam) VALUES (?, ?, ?)", newFullOrder.Customer.Nev, newFullOrder.Customer.Email, newFullOrder.Customer.Telefonszam)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	// _ ignore set variable
	// https://stackoverflow.com/questions/27764421/what-is-underscore-comma-in-a-go-declaration
	// Get ID, because auto generated
	newFullOrder.Customer.VasarloID, _ = newCustomer.LastInsertId()

	for index, order := range newFullOrder.Foods {
		var food food
		err := Db.QueryRow("SELECT * FROM Etelek WHERE EtelID = ?", order.EtelID).Scan(&food.EtelID, &food.Nev, &food.Leiras, &food.KepPath, &food.Ar)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		_, err = Db.Exec("INSERT INTO Rendelesek (VasarloID, EtelID, Ar, Darab) VALUES (?, ?, ?, ?)", newFullOrder.Customer.VasarloID, order.EtelID, food.Ar, order.Volume)
		if err != nil {
			fmt.Println("Failed to insert into Rendelesek table at: ", index)
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
	}

	var sumPrice float64
	err = Db.QueryRow("SELECT SUM(Ar*Darab) FROM Rendelesek WHERE VasarloID = ?", newFullOrder.Customer.VasarloID).Scan(&sumPrice)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	_, err = Db.Exec("UPDATE Vasarlok SET Osszeg = ? WHERE VasarloID = ?", sumPrice, newFullOrder.Customer.VasarloID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	newFullOrder.Customer.Osszeg = sql.NullFloat64{Float64: sumPrice, Valid: true}

	c.IndentedJSON(http.StatusCreated, newFullOrder)
}

func CompleteOrder(c *gin.Context) {
	var isDone map[string]bool
	if err := c.BindJSON(&isDone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isDoneDateTime sql.NullString
	if isDone["Elkeszult"] {
		// About why this date format
		// https://pkg.go.dev/time#Time.Format
		isDoneDateTime = sql.NullString{String: time.Now().Format("2006-01-02 15:04:05"), Valid: true}
	}

	_, err := Db.Exec("UPDATE Vasarlok SET ElkeszultIdo=? WHERE VasarloID=?", isDoneDateTime, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, c.Param("id"))
}
