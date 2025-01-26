package controllers

import (
	"net/http"
	"sb-go-quiz/databases"
	"sb-go-quiz/repositories"
	"sb-go-quiz/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

var err error

func CFindCategories(c *gin.Context) {
	var statusCode int = http.StatusOK
	var result gin.H

	categories, err := repositories.RFindCategories(databases.DbConnection)

	if err.Message != "" {
		statusCode = err.Status
		result = gin.H{
			"result": err.Message,
		}
		c.JSON(statusCode, result)
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(statusCode, result)

}

func CCreateCategories(c *gin.Context) {
	var statusCode int = http.StatusOK
	var result gin.H
	var category structs.SCategories

	errBind := c.ShouldBindJSON(&category)

	user, _ := c.Get("user")

	userStruct, _ := user.(structs.SUsers)

	category.CreatedBy = userStruct.ID
	category.ModifiedBy = userStruct.ID

	if errBind != nil {
		result = gin.H{
			"result": err.Error(),
		}
	}

	err := repositories.RCreateCategories(databases.DbConnection, category)

	if err.Message != "" {
		statusCode = err.Status
		result = gin.H{
			"result": err.Message,
		}
	} else {
		result = gin.H{
			"result": "Success",
		}
	}

	c.JSON(statusCode, result)

}

func CUpdateCategories(c *gin.Context) {
	var statusCode int = http.StatusOK
	var result gin.H
	var category structs.SCategories
	id, _ := strconv.Atoi(c.Param("id"))

	errBind := c.ShouldBindJSON(&category)

	if errBind != nil {
		result = gin.H{
			"result": err.Error(),
		}
	}

	user, _ := c.Get("user")

	userStruct, _ := user.(structs.SUsers)

	category.CreatedBy = userStruct.ID
	category.ModifiedBy = userStruct.ID

	category.ID = id

	err := repositories.RUpdateCategories(databases.DbConnection, category)

	if err.Message != "" {
		statusCode = err.Status
		result = gin.H{
			"result": err.Message,
		}
	} else {
		result = gin.H{
			"result": "Success",
		}
	}

	c.JSON(statusCode, result)
}

func CDeleteCategories(c *gin.Context) {
	var statusCode int = http.StatusOK
	var result gin.H
	var category structs.SCategories
	id, _ := strconv.Atoi(c.Param("id"))

	category.ID = id

	err := repositories.RDeleteCategories(databases.DbConnection, category.ID)

	if err.Message != "" {
		statusCode = err.Status
		c.JSON(statusCode, gin.H{
			"result": err.Message,
		})
		return
	} else {
		result = gin.H{
			"result": "Success",
		}
	}

	c.JSON(statusCode, result)

}

func CFindCategory(c *gin.Context) {
	var statusCode int = http.StatusOK
	var result gin.H
	var category structs.SCategories
	id, _ := strconv.Atoi(c.Param("id"))

	category.ID = id

	data, err := repositories.RFindCategory(databases.DbConnection, category.ID)

	if err.Message != "" {
		statusCode = err.Status
		result = gin.H{
			"result": err.Message,
		}
	} else {
		result = gin.H{
			"result": data,
		}
	}

	c.JSON(statusCode, result)
}

func CFindBooksByCategory(c *gin.Context) {
	var statusCode int = http.StatusOK
	var result gin.H
	var books []structs.SBooks

	id, _ := strconv.Atoi(c.Param("id"))

	books, err := repositories.RFindBooksByCategory(databases.DbConnection, id)

	if err.Message != "" {
		statusCode = err.Status
		result = gin.H{
			"result": err.Message,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(statusCode, result)
}
