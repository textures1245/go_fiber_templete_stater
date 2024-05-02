package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/textures1245/go-template/router"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	log.Info("-= Start Simple Service =-")
	router.SetupRoutes(app)

	app.Listen(":" + viper.GetString("SERVER_PORT"))

}
