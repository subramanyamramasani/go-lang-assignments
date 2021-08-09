package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rizalgowandy/go-rest-api/database"
	_ "github.com/subramanyam/go-rest-api/docs/ginsimple" // you need to update github.com/rizalgowandy/go-swag-sample with your own project path
	"github.com/subramanyam/go-rest-api/model"
	"github.com/subramanyam/go-rest-api/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
/*
func main() {
	// Gin instance
	r := gin.New()

	// Routes
	r.GET("employee", EmpGet)
	//r.POST("employee", HealthCheck)
	//r.PUT("employee/:id", HealthCheck)
	//r.DELETE("employee/:id", HealthCheck)

	url := ginSwagger.URL("http://localhost:8060/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Start server
	if err := r.Run(":8060"); err != nil {
		log.Fatal(err)
	}
} */
var err error

func main() {
	database.DB, err = gorm.Open("mysql", database.DbURL(database.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer database.DB.Close()
	database.DB.AutoMigrate(&model.Employee{})
	//r := gin.New()
	r := routes.Setuprouter()

	r.GET("/", EmpGet)
	r.POST("/", EmpCreate)

	url := ginSwagger.URL("http://localhost:8021/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	if err := r.Run(":8021"); err != nil {
		log.Fatal(err)
		r.Run()
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]

func EmpGet(c *gin.Context) {
	var emp []model.Employee
	err := model.GetAllEmp(&emp)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, emp)
	}
}

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
