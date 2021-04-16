package main

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Response struct {
	Token string `json:"token"`
}

func main() {
	shouldRunWithPrefork := os.Getenv("PRE_FORK")
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               shouldRunWithPrefork == "true",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		var (
			id    = uuid.New().String()
			token string
			err   error
		)

		if token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uuid": id,
			"nbf":  time.Now().Unix(),
		}).SigningString(); err != nil {
			return err
		}

		return c.JSON(Response{token})
	})

	app.Listen(":80")
}
