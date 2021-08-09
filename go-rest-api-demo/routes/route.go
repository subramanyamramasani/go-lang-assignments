package routes

import (
	//	"newback/controller"

	"github.com/gin-gonic/gin"
	"github.com/rizalgowandy/go-swag-sample/controller"
)

//SetupRouter ... Configure routes
func Setuprouter() *gin.Engine {

	r := gin.Default()

	emp1 := r.Group("/employee-api")
	{
		emp1.GET("employee", controller.EmpGet)
		emp1.POST("employee", controller.EmpCreate)
		emp1.GET("employee/:id", controller.GetEmpByID)
		emp1.PUT("employee/:id", controller.EmpUpdate)
		emp1.DELETE("employee/:id", controller.EmpDelete)
	}
	return r
}
