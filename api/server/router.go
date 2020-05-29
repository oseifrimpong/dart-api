package server

import (
	"dart-api/api/delivery"
	"dart-api/api/service"

	"github.com/gin-gonic/gin"
)

func (ds *dserver) MapRoutes() {
	// Group routes for v1

	apiV1 := ds.router.Group("api/v1")
	ds.userRoutes(apiV1)
}

func (ds *dserver) userRoutes(api *gin.RouterGroup) {
	userRoutes := api.Group("/users")
	{
		var userSvc service.Service
		ds.cont.Invoke(func(u service.Service) {
			userSvc = u
		})
		usr := delivery.NewUserController(userSvc)

		userRoutes.GET("/", usr.GetAll)
		userRoutes.POST("/", usr.Create)
		userRoutes.GET("/:id", usr.GetByID)
		userRoutes.PUT("/:id", usr.Update)
		userRoutes.DELETE("/:id", usr.Delete)
	}
}
