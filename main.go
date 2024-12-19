package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

func main() {
	fmt.Println("Running Server...")

	file, err := os.Open("2018_Central_Park_Squirrel_Census_-_Squirrel_Data_20241218.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var records []SquirrelData
	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, records)
	})

	router.GET("/id/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		data, err := findByID(records, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "Error",
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, data)
		}
	})

	router.GET("/age/:age", func(ctx *gin.Context) {
		age := ctx.Param("age")
		data, err := findByAge(records, age)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "Error",
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, data)
		}
	})

	router.GET("/primaryFurColor/:color", func(ctx *gin.Context) {
		color := ctx.Param("color")
		data, err := findByPrimaryFurColor(records, color)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "Error",
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, data)
		}
	})

	router.GET("/highlightFurColor/:color", func(ctx *gin.Context) {
		color := ctx.Param("color")
		data, err := findByHighlightFurColor(records, color)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "Error",
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, data)
		}
	})

	router.GET("/condition/:cond", func(ctx *gin.Context) {
		condition := ctx.Param("cond")
		data, err := findSquirrelByCondition(records, condition)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "Error",
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, data)
		}
	})

	router.Run()
}
