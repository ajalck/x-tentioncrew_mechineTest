package main

import (
	"log"

	"github.com/ajalck/service_2/pkg/client"
	"github.com/ajalck/service_2/pkg/config"
	"github.com/ajalck/service_2/pkg/methods"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}
	client:=client.InitService1Client(c.Service_1_port)
	m:=&methods.Client{
		C: client.Client,
	}
	engine := gin.New()
	engine.Use(gin.Logger())

	engine.GET("/getusers",m.ClientMethod)

	engine.Run(c.Port)

}
