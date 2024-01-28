package main

//main
import (
	"RakaminProject/handler"
	"RakaminProject/initializer"
	"RakaminProject/middleware"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDb()
}

func main() {

	r := gin.Default()

	userHandler := handler.NewUserHandler()
	photoHandler := handler.NewPhotoHandler()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to RakaminProject")
	})

	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")

	{
		userRoutes.POST("/register", userHandler.AddUser)
		userRoutes.POST("/signin", userHandler.SignInUser)
	}

	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	{
		userProtectedRoutes.GET("/", userHandler.GetUser)
		userProtectedRoutes.GET("/photo", userHandler.GetAllPhoto)
		userProtectedRoutes.PUT("/:user", userHandler.UpdateUser)
		userProtectedRoutes.DELETE("/:user", userHandler.DeleteUser)
	}

	photoRoutes := apiRoutes.Group("/photos", middleware.AuthorizeJWT())
	{
		photoRoutes.GET("/:photo", photoHandler.GetPhoto)
		photoRoutes.POST("/", photoHandler.AddPhoto)
		photoRoutes.PUT("/:photo", photoHandler.UpdatePhoto)
		photoRoutes.DELETE("/:photo", photoHandler.DeletePhoto)
	}

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
