package route

import (
	"Oracle/controller"
	"Oracle/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(controller.CorsMiddleware())
	r.Use(sessions.Sessions("mysession", store))
	r.Use(middleware.JwtAuthMiddleware())
	r.OPTIONS("/data/gradeData", controller.CorsMiddleware())
	r.POST("/data/gradeData", controller.GradeDataHandle)
	r.OPTIONS("/data/findDataById", controller.CorsMiddleware())
	r.POST("/data/findDataById", controller.FindDataByIdHandle)
	r.OPTIONS("/data/feedBack", controller.CorsMiddleware())
	r.POST("/data/feedBack", controller.Feedback)
	r.OPTIONS("/data/useData", controller.CorsMiddleware())
	r.POST("/data/useData", controller.UseDataHandle)
	r.OPTIONS("/data/getStartTime", controller.CorsMiddleware())
	r.POST("/data/getStartTime", controller.GetStartTimeHandle)

	r.POST("/users/register", controller.RegisterHandler)
	r.POST("/users/login", controller.LoginHandler)

	return r
}
