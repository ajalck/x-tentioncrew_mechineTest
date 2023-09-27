package main

import (
	"log"

	"github.com/ajalck/service_1/pkg/config"
	"github.com/ajalck/service_1/pkg/db"
	"github.com/ajalck/service_1/pkg/handler"
	"github.com/ajalck/service_1/pkg/repository"
	repoInt "github.com/ajalck/service_1/pkg/repository/interfaces"
	"github.com/ajalck/service_1/pkg/usecase"
	usecaseInt "github.com/ajalck/service_1/pkg/usecase/interfaces"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load env; %v", err)
	}
	sDB := db.ConnectPostgres(c)
	db.SyncDB(sDB.DB)
	rDB := db.ConnectRedis(c)

	StartGrpcServer(c,sDB.DB)
	var (
		userRepo    repoInt.UserRepo       = repository.NewUserRepo(sDB.DB, rDB)
		userService usecaseInt.UserUseCase = usecase.NewUserUseCase(userRepo)
		userHandler handler.UserHandler    = *handler.NewUserHandler(userService)
	)
	router.POST("/createuser", userHandler.CreateUser)
	router.GET("/getuserbyid", userHandler.GetUserByID)
	router.PUT("/updateuser", userHandler.UpdateUser)
	router.DELETE("/deleteuser", userHandler.DeleteUser)

	router.Run(c.HttpPort)

}
