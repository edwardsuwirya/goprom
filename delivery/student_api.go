package delivery

import (
	"enigmacamp.com/goprom/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

type StudentApi struct {
	publicRoute *gin.RouterGroup
}

func NewStudentApi(publicRoute *gin.RouterGroup) *StudentApi {
	studentApi := new(StudentApi)
	studentApi.publicRoute = publicRoute
	studentApi.initRouter()
	return studentApi
}
func (api *StudentApi) initRouter() {
	userRoute := api.publicRoute.Group("/student")
	userRoute.GET("/:idcard", api.getStudentById)
	userRoute.POST("", api.createStudent)
}

func (api *StudentApi) getStudentById(c *gin.Context) {
	name := c.Param("idcard")
	c.JSON(200, gin.H{
		"message": name,
	})
}
func (api *StudentApi) createStudent(c *gin.Context) {
	var student model.Student
	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Created student : %v", student),
	})
}
