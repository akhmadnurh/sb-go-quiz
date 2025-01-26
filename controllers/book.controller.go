package controllers

import (
	"net/http"
	"sb-go-quiz/databases"
	"sb-go-quiz/repositories"
	"sb-go-quiz/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CFindBooks(c *gin.Context) {
	statusCode := http.StatusOK

	books, err := repositories.RFindBooks(databases.DbConnection)

	if err.Message != "" {
		statusCode = err.Status
		c.JSON(statusCode, gin.H{
			"result": err.Message,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"result": books,
	})
}

func CFindBook(c *gin.Context) {
	statusCode := http.StatusOK

	id, _ := strconv.Atoi(c.Param("id"))

	book, err := repositories.RFindBook(databases.DbConnection, id)

	if err.Message != "" {
		statusCode = err.Status
		c.JSON(statusCode, gin.H{
			"result": err.Message,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"result": book,
	})
}

func CCreateBook(c *gin.Context) {
	statusCode := http.StatusOK

	user, _ := c.Get("user")
	userStruct, _ := user.(structs.SUsers)

	var book structs.SBooks
	errBind := c.ShouldBindJSON(&book)

	if errBind != nil {
		statusCode = http.StatusBadRequest
		c.JSON(statusCode, gin.H{
			"result": errBind.Error(),
		})
		return
	}

	if book.ReleaseYear < 1980 && book.ReleaseYear > 2024 {
		statusCode = http.StatusBadRequest
		c.JSON(statusCode, gin.H{
			"result": "Release year must be between 1980 and 2024",
		})
		return
	}

	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	book.CreatedBy = userStruct.ID
	book.ModifiedBy = userStruct.ID

	err := repositories.RCreateBook(databases.DbConnection, book)

	if err.Message != "" {
		statusCode = err.Status
		c.JSON(statusCode, gin.H{
			"result": err.Message,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"result": "Success",
	})

}

func CUpdateBook(c *gin.Context) {
	statusCode := http.StatusOK

	id, _ := strconv.Atoi(c.Param("id"))

	user, _ := c.Get("user")
	userStruct, _ := user.(structs.SUsers)

	var book structs.SBooks
	errBind := c.ShouldBindJSON(&book)

	if errBind != nil {
		statusCode = http.StatusBadRequest
		c.JSON(statusCode, gin.H{
			"result": errBind.Error(),
		})
		return
	}

	if book.ReleaseYear < 1980 && book.ReleaseYear > 2024 {
		statusCode = http.StatusBadRequest
		c.JSON(statusCode, gin.H{
			"result": "Release year must be between 1980 and 2024",
		})
		return
	}

	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	book.ID = id
	book.ModifiedBy = userStruct.ID

	err := repositories.RUpdateBook(databases.DbConnection, book)

	if err.Message != "" {
		statusCode = err.Status
		c.JSON(statusCode, gin.H{
			"result": err.Message,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"result": "Success",
	})
}

func CDeleteBook(c *gin.Context) {
	statusCode := http.StatusOK

	id, _ := strconv.Atoi(c.Param("id"))

	err := repositories.RDeleteBook(databases.DbConnection, id)

	if err.Message != "" {
		statusCode = err.Status
		c.JSON(statusCode, gin.H{
			"result": err.Message,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"result": "Success",
	})
}
