package router

import (
	"github.com/calendarproject/common"
	"github.com/calendarproject/service"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.Use(common.CORSMiddleware(), common.CORSMiddleware())

	postRoutes := r.Group("/posts")

	// postRoutes.Use(middleware.AuthMiddleware())

	postController := service.NewPostController()

	postRoutes.POST("", postController.Create)

	postRoutes.DELETE(":id", postController.Delete)

	postRoutes.PUT(":id", postController.Put)

	postRoutes.GET(":id", postController.SelectByID)

	postRoutes.POST("page/list", postController.PageList)

	return r

}
