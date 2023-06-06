package routes

import (
	"gin/common"
	"gin/controller"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	category := r.Group("")
	{
		category.GET("/category", controller.ChooseCategory)
	}

	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.GET("/info", middleware.AuthMiddleware(), controller.Info)

	home := r.Group("home")
	{
		home.GET("/goods", controller.GetGoods)
		home.GET("/banner", controller.GetBanner)
		home.GET("/new", controller.RecentIdle)

	}
	member := r.Group("member")
	{
		member.POST("/order", middleware.AuthMiddleware(), controller.CreateOrder)
		member.GET("/order/:id", middleware.AuthMiddleware(), controller.GetOrder)
	}

	goods := r.Group("")
	{
		goods.GET("/goods", controller.GetOneGood)
		goods.GET("/goods/relevant", middleware.AuthMiddleware(), controller.RecommendGoods)
	}

	chatList := r.Group("chat")
	{
		chatList.GET("/getmsg", middleware.AuthMiddleware(), controller.GetMsg)
		chatList.POST("/sendmsg", middleware.AuthMiddleware(), controller.SendMsg)
		chatList.POST("/addchat", middleware.AuthMiddleware(), controller.AddChat)
	}
<<<<<<< HEAD
	//member路由完善后可以将下面这个路由整合
	CartGroup := r.Group("member/cart")
	{
		CartGroup.POST("/add", middleware.AuthMiddleware(), controller.CartIn)
		CartGroup.GET("/pull", middleware.AuthMiddleware(), controller.CartOut)
		CartGroup.DELETE("/del", middleware.AuthMiddleware(), controller.CartDel)
	}
=======

	db := common.GetDB()
	imageController := controller.NewImageController(db)

	imageRoutes := r.Group("/images")
	{
		imageRoutes.POST("", imageController.UploadImage)
		imageRoutes.GET("/:id", imageController.GetImage)
	}

>>>>>>> 2f7999f1a791b3d5437a6c64f80f409d0cf26d0c
	return r
}
