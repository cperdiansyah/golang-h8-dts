package router

import (
	"github.com/gin-gonic/gin"
	"github.com/cperdiansyah/golang-h8-dts/final-project/app"
	"github.com/cperdiansyah/golang-h8-dts/final-project/controller"
	"github.com/cperdiansyah/golang-h8-dts/final-project/middleware"
	"github.com/cperdiansyah/golang-h8-dts/final-project/repository"
	"github.com/cperdiansyah/golang-h8-dts/final-project/service"
)

func SocialMediaRouter(router *gin.Engine) {
	db := app.NewDB()

	repo := repository.NewSocialMediaRepository(db)
	srv := service.NewSocialMediaService(repo)
	ctrl := controller.NewSocialMediaController(srv)

	socialMedia := router.Group("/socialmedias", middleware.AuthMiddleware())

	{
		socialMedia.GET("/", ctrl.GetAll)
		socialMedia.POST("/", ctrl.Create)
		{
			socialMedia.PUT("/:id", middleware.SocialMediaMiddleware(srv), ctrl.Update)
			socialMedia.DELETE("/:id", middleware.SocialMediaMiddleware(srv), ctrl.Delete)
		}
	}
}
