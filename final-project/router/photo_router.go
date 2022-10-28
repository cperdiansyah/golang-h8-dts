package router

import (
	"github.com/gin-gonic/gin"
	"github.com/cperdiansyah/golang-h8-dts/final-project/app"
	"github.com/cperdiansyah/golang-h8-dts/final-project/controller"
	"github.com/cperdiansyah/golang-h8-dts/final-project/middleware"
	"github.com/cperdiansyah/golang-h8-dts/final-project/repository"
	"github.com/cperdiansyah/golang-h8-dts/final-project/service"
)

func PhotoRouter(router *gin.Engine) {
	db := app.NewDB()

	repo := repository.NewPhotoRepository(db)
	srv := service.NewPhotoService(repo)
	ctrl := controller.NewPhotoController(srv)

	photoRouter := router.Group("/photos", middleware.AuthMiddleware())

	{
		photoRouter.POST("/", ctrl.Create)
		photoRouter.GET("/", ctrl.GetAll)
		{
			photoRouter.PUT("/:id", middleware.PhotoMiddleware(srv), ctrl.Update)
			photoRouter.DELETE("/:id", middleware.PhotoMiddleware(srv), ctrl.Delete)
		}
	}
}
