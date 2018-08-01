package web

import (
	"github.com/gin-gonic/gin"
	"github.com/maoltr/blog/web/controllers"
)

func StartServer() {
	router := gin.Default()

	router.Use(controllers.SetUserStatus())

	//db := database.ArticleDatabase()

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
		article.GET("update/:id", controllers.EnsureLoggedIn(), controllers.GetUpdateArticle)
		article.POST("update/:id", controllers.EnsureLoggedIn(), controllers.PostUpdateArticle)
	}

	router.Run("localhost:8080")
}
