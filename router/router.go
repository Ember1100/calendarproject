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
	//添加提醒
	postRoutes.POST("", postController.Create)
	//删除提醒
	postRoutes.DELETE(":id", postController.Delete)
	//更改提醒
	postRoutes.PUT(":id", postController.Put)

	postRoutes.GET(":id", postController.SelectByID)
	//用户提醒列表
	postRoutes.POST("page/list", postController.PageList)

	return r

}
