package routes

import (
	config "golang-api/configs"
	controller_user "golang-api/controllers/user"
	"golang-api/httpserver"
	repository_user "golang-api/repositories/user"
	service_user "golang-api/services/user"

	"gorm.io/gorm"
)

func Setup(route httpserver.Router,
	db *gorm.DB, config config.Config) {

	// initial repositories
	userRepo := repository_user.New(db)

	// initial services

	userSvc := service_user.NewUserService(userRepo)

	// initial controller
	userController := controller_user.NewCancelController(userSvc)

	v1 := route.Group("/v1")

	user := v1.Group("/user")
	{
		user.GET("", userController.List)
		user.GET("/:id", userController.GetUser)
		user.POST("", httpserver.DbTransactionMiddleware(userController.CreateUser, db))
		// user.DELETE("")
		// user.PATCH("")
		// user.POST("")
		// user.PUT("")
	}

}
