package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nathannlu/microservice/pkg/auth"
	"github.com/nathannlu/microservice/pkg/config"
)

func main() {
	c, err := config.loadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	authSvc := *auth.RegisterRoutes(r, &c)


	r.Run(c.Port)
}
