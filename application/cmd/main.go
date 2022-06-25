package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/piTch-time/pitch-backend/application/controller"
	"github.com/piTch-time/pitch-backend/application/route"
	"github.com/piTch-time/pitch-backend/docs"
	"github.com/piTch-time/pitch-backend/domain/service"
	"github.com/piTch-time/pitch-backend/infrastructure"
	"github.com/piTch-time/pitch-backend/infrastructure/configs"
	"github.com/piTch-time/pitch-backend/infrastructure/logger"
	"github.com/piTch-time/pitch-backend/infrastructure/persistence"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

const (
	//
	versionPrefix = "/v1"
	defaultPhase  = "dev"
	configPath    = "./infrastructure/configs"
)

var (
	phase string
	conf  configs.Config
)

func main() {
	var err error
	flag.StringVar(&phase, "phase", defaultPhase, "name of configuration file with no extension")
	flag.Parse()
	viper.SetDefault("PHASE", phase)

	conf, err = configs.Load(configPath)
	if err != nil {
		panic("Failed to load config file: " + err.Error())
	}

	server := bootstrap()
	server.Run(":8080")
	shutdown()
}

func bootstrap() *gin.Engine {
	// init db
	db := infrastructure.ConnectDatabase(phase)
	infrastructure.Migrate(db)

	// set DI
	taskRepository := persistence.NewTaskRepository(db)
	roomRepository := persistence.NewRoomRepository(db)

	taskService := service.NewTaskService(taskRepository)
	roomService := service.NewRoomService(roomRepository)

	taskController := controller.NewTaskController(taskService)
	roomController := controller.NewRoomController(roomService)

	// init server
	server := gin.New()

	swagger(server)

	// zap middlewares
	server.Use(ginzap.Ginzap(logger.Log, time.RFC3339, true))

	// init routes
	v1 := server.Group(versionPrefix)
	route.RoomRoutes(v1, roomController)
	route.TaskRoutes(v1, taskController)
	return server
}

func swagger(server *gin.Engine) {
	docs.SwaggerInfo.BasePath = versionPrefix
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func shutdown() {
	// Wait for termination signals.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	osSignal := <-c
	logger.Info("Application terminates", zap.Any("Signal", osSignal))
}
