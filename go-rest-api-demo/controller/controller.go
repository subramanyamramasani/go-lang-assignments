package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/subramanyam/go-rest-api/model"
)

//Retrieving Entire data of all Employees
func EmpGet(c *gin.Context) {
	var emp []model.Employee
	err := model.GetAllEmp(&emp)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, emp)
	}
}

//Creating Employee
func EmpCreate(c *gin.Context) {
	var emp model.Employee
	c.BindJSON(&emp)
	err := model.CreateEmp(&emp)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, emp)
	}
}

//Retrieving Employee data by id
func GetEmpByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var emp model.Employee
	err := model.GetEmpByID(&emp, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, emp)
	}
}

//Updating Employee
func EmpUpdate(c *gin.Context) {
	var emp model.Employee
	id := c.Params.ByName("id")
	err := model.GetEmpByID(&emp, id)
	if err != nil {
		c.JSON(http.StatusNotFound, emp)
	}
	c.BindJSON(&emp)
	err = model.EmpUpdate(&emp, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, emp)
	}
}

//deleting employee
func EmpDelete(c *gin.Context) {
	var emp model.Employee
	id := c.Params.ByName("id")
	err := model.EmpDelete(&emp, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
