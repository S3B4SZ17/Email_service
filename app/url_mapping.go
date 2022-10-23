package app

import "github.com/S3B4SZ17/Email_service/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.GET("/callback-gl", controllers.CallBackFromGoogle)
	router.GET("/login-gl", controllers.HandleGoogleLogin)
}
