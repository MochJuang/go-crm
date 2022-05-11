package lead

import (
	"github.com/MochJuang/go-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func Getleads(c *fiber.Ctx) {
	db := database.DBConn
	var lead []Lead
	db.Find(&lead)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	var lead []Lead

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)

	if lead.Name != "" {
		c.Status(500).Send("Lead not found")
		return
	}
	db.Delete(&lead)
	c.JSON("success deleted lead")
}
