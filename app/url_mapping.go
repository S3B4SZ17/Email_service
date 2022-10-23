package app

import (
	"github.com/S3B4SZ17/Email_service/controllers"
	"github.com/S3B4SZ17/Email_service/middleware"
)

func mapUrls() {
	public := router.Group("/api")
	public.GET("/ping", controllers.Ping)
	public.GET("/callback-gl", controllers.CallBackFromGoogle)
	public.GET("/login-gl", controllers.HandleGoogleLogin)

	protected := router.Group("/api/authorized")
	protected.Use(middleware.Oauth2AuthMiddleware())
	protected.GET("/userinfo", controllers.GetUserInfo)

}
