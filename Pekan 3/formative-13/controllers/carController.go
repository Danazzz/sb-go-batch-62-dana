package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

    "github.com/gin-gonic/gin"
)

type Car struct {
	ID 		int 	`json:"id"`
	Brand 	string 	`json:"brand"`
	Model 	string 	`json:"model"`
	Price 	int 	`json:"price"`
}

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func CreateCar(ctx *gin.Context) {
	var car Car

	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.AbortWithError(http.StatusBadRequest,err)
		return
	}
	err := db.QueryRow("INSERT INTO cars (brand, model, price) VALUES ($1, $2, $3) RETURNING id",
		car.Brand, car.Model, car.Price).Scan(&car.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"car": car,
	})
}

func UpdateCar(ctx *gin.Context) {
	id := ctx.Param("id")
	var car Car

	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE cars SET brand=$1, model=$2, price=$3 WHERE id=$4",
		car.Brand, car.Model, car.Price, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully updated", id),
	})
}

func GetCar(ctx *gin.Context) {
	id := ctx.Param("id")
	var car Car

	err := db.QueryRow("SELECT id, brand, model, price FROM cars WHERE id = $1", id).Scan(&car.ID, &car.Brand, &car.Model, &car.Price)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car":car,
	})
}

func DeleteCar(ctx *gin.Context)  {
	id := ctx.Param("id")

	_, err := db.Exec("DELETE FROM cars WHERE id=$1", id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully deleted", id),
	})
}