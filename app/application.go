package app

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/S3B4SZ17/Email_service/config"
	"github.com/S3B4SZ17/Email_service/controllers"
	"github.com/S3B4SZ17/Email_service/management"
	"github.com/S3B4SZ17/Email_service/protoServer"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	host         = "localhost"
	gRPCListener = "50051"
	router       *gin.Engine
)

func StartApp(config *config.Config) {

	// Start the gRPC server
	go StartgRPCServer()

	// Start the HTTP server for the application
	StartHTTPServer(config)
}

func StartHTTPServer(config *config.Config) {
	gin_mode := os.Getenv("GIN_MODE")
	fmt.Println(gin_mode)
	if gin_mode == "" {
		gin_mode = "release"
		os.Setenv("GIN_MODE", gin_mode)
		gin.SetMode(gin.ReleaseMode)
	}
	httpPort := config.Http_server.HttpPort
	if httpPort == "" {
		httpPort = "8181"
	}

	management.Log.Info("Starting application\n", zap.String("http_port", httpPort), zap.String("gin_mode", gin_mode))
	router = gin.Default()

	// Initialize Oauth2 Services
	controllers.InitializeOAuthGoogle()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.Http_server.Cors.List_hosts,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	mapUrls()

	router.Run(":" + httpPort)
}

func StartgRPCServer() {
	management.Log.Info("Start gRPCListener", zap.String("http_port", gRPCListener))

	listener, err := net.Listen("tcp", ":"+gRPCListener)
	if err != nil {
		management.Log.Error(err.Error())
	}

	srv := grpc.NewServer()
	protoServer.RegisterServices(srv)

	if e := protoServer.StartServer(srv, listener); e != nil {
		management.Log.Error("An error occurred while serving: " + e.Error())
	}
}
