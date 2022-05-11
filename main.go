package main

import (
	"fmt"

	"github.com/MochJuang/go-crm/database"
	"github.com/MochJuang/go-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.Getleads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "lead.db")

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}
