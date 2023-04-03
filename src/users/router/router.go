package router

import (
	"e_commerce-microservices/src/users/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	UserSignUp(c *gin.Context)
	UserSignIn(c *gin.Context)
	GetUserProfile(c *gin.Context)
	Logout(c *gin.Context)
}

type Middleware interface {
	IsAuthorized(c *gin.Context)
	CORSMiddleware() gin.HandlerFunc
}


type User_Mircoservice struct {
	api Controller
	mid Middleware
}

func NewUserService(api Controller) *User_Mircoservice {
	return &User_Mircoservice{
		api: api,
		mid: middleware.NewUserMiddleware(),
	}
}

func enableUserRoutes(router *gin.Engine, us *User_Mircoservice) {

	// User routes
	router.POST("/signup", us.api.UserSignUp)

	// all routes here are for admin user group only
	router.POST("/login", us.api.UserSignIn)
	router.GET("/profile", us.mid.IsAuthorized, us.api.GetUserProfile)
	router.GET("/logout", us.api.Logout)
}

func (us *User_Mircoservice) newUserRouter() *gin.Engine {
	router := gin.Default()

	enableUserRoutes(router, us)
	
	return router
}

func (us *User_Mircoservice) Start() {
	router := us.newUserRouter()
	log.Println("Starting User Microservice on port 4002")
	router.Run(":4002")
}