package restapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Visszaadja az összes ételt az adatbázisból
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
		if err := rows.Scan(&food.EtelID, &food.Nev, &food.Leiras, &food.KepPath, &food.Ar); err != nil {
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

// Egy ételt visszaad URL-ben megadott ID alapján.
func GetFood(c *gin.Context) {
	var food food
	err := Db.QueryRow("SELECT * FROM Etelek WHERE EtelID = ?", c.Param("id")).Scan(&food.EtelID, &food.Nev, &food.Leiras, &food.KepPath, &food.Ar)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.JSON(http.StatusOK, food)
}

// Visszaadja az összes rendelt étel mennyiségét és annak adatát a Vásárló ID-je alapján.
func GetAllFoodByCustomer(c *gin.Context) {
	orderRows, err := Db.Query("SELECT * FROM Rendelesek WHERE VasarloId = ? ORDER BY VasarloID, Darab", c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	var foods []foodWithVolume

	defer orderRows.Close()
	for orderRows.Next() {
		var order order
		if err := orderRows.Scan(&order.VasarloID, &order.EtelID, &order.Ar, &order.Darab); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		var food food
		err := Db.QueryRow("SELECT * FROM Etelek WHERE EtelID = ?", order.EtelID).Scan(&food.EtelID, &food.Nev, &food.Leiras, &food.KepPath, &food.Ar)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
			return
		}
		food.Ar = order.Ar
		foodWithVolume := foodWithVolume{Food: food, Volume: order.Darab}
		foods = append(foods, foodWithVolume)
	}
	if foods == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empty database"})
		return
	}

	if err := orderRows.Err(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foods)
}

// Létrehoz egy új ételt a body-ban megadott adatok alapján.
func PostFood(c *gin.Context) {
	var newFood food
	if err := c.BindJSON(&newFood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foodResult, err := Db.Exec("INSERT INTO Etelek (Nev, Leiras, Kep, Ar) VALUES (?, ?, ?, ?)", newFood.Nev, newFood.Leiras, newFood.KepPath, newFood.Ar)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	insertedId, err := foodResult.LastInsertId()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	newFood.EtelID = int(insertedId)

	c.IndentedJSON(http.StatusCreated, newFood)
}

// Módosít egy ételt ID alapján.
func UpdateFood(c *gin.Context) {
	var isFoodExists bool = false
	err := Db.QueryRow("SELECT EXISTS(SELECT * FROM Etelek WHERE EtelID = ?)", c.Param("id")).Scan(&isFoodExists)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if !isFoodExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	var newFood food
	if err := c.BindJSON(&newFood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = Db.Exec("UPDATE Etelek SET Nev=?, Leiras=?, Kep=?, Ar=? WHERE EtelID=?", newFood.Nev, newFood.Leiras, newFood.KepPath, newFood.Ar, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	newFood.EtelID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newFood)
}

// Kitöröl egy ételt ID alapján.
func DeleteFood(c *gin.Context) {
	foodResult, err := Db.Exec("DELETE FROM Etelek WHERE EtelID=?", c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, err := foodResult.RowsAffected()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if rowsAffected <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, c.Param("id"))
}
