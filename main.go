package main

import (
	"context"
	"fmt"
	"log"
	"menu-service/ent"
	menu_handler "menu-service/handler/menu"
	menu_service "menu-service/service/menu"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")

	DB_URL := fmt.Sprintf("host=localhost port=%s user=%s dbname=%s password=%s sslmode=disable", POSTGRES_PORT, POSTGRES_USER, POSTGRES_DB, POSTGRES_PASSWORD)

	client, err := ent.Open("postgres", DB_URL)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	menu_srv := menu_service.NewService(client)
	menu_handler := menu_handler.NewHandler(menu_srv)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/menus/:id", menu_handler.GetMenuById)
	app.Get("/menus", menu_handler.GetAllMenus)
	app.Post("/menus", menu_handler.CreateMenu)
	app.Put("/menus/:id", menu_handler.UpdateMenu)
	app.Delete("/menus/:id", menu_handler.DeleteMenu)

	// Graceful shutdown
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Error when listening to port 3000: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
