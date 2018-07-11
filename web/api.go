package web

import (
	"blog/web/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	router.Use(controllers.SetUserStatus())

	router.GET("/", controllers.GetAllArticles)
	router.LoadHTMLGlob("./resources/templates/*.html")

	u := router.Group("/u/")
	{
		u.GET("login/", controllers.EnsureNotLoggedIn(), controllers.ShowLoginPage)
		u.POST("login/", controllers.EnsureNotLoggedIn(), controllers.Login)
		u.GET("logout/", controllers.EnsureLoggedIn(), controllers.LogOut)
		u.GET("register/", controllers.EnsureNotLoggedIn(), controllers.ShowRegisterPage)
		u.POST("register/", controllers.EnsureNotLoggedIn(), controllers.Register)
	}

	article := router.Group("/article/")
	{
		article.GET("create/", controllers.EnsureLoggedIn(), controllers.CreateArticle)
		article.POST("create/", controllers.EnsureLoggedIn(), controllers.PostArticle)
		article.GET("view/:id", controllers.GetArticle)
		article.POST("delete/:id", controllers.EnsureLoggedIn(), controllers.DeleteArticle)
		article.GET("manage/", controllers.EnsureLoggedIn(), controllers.ManageArticles)
	}

	v1 := router.Group("/api/v1/")
	{
		v1.POST("product/", controllers.CreateProduct)
		v1.GET("product/", controllers.FetchAllProduct)
		v1.GET("product/:id", controllers.FetchSingleProduct)
		v1.PUT("product/:id", controllers.UpdateProduct)
		v1.DELETE("product/:id", controllers.DeleteProduct)
	}

	router.Run("localhost:8080")
}
