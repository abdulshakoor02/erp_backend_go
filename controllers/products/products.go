package productsController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/products"
	"github.com/gofiber/fiber/v2"
)

type Products struct {
	ID    string  `json:"id"    gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Name  string  `json:"name"  gorm:"type:string;size:100;unique"`
	Desc  string  `json:"desc"       gorm:"type:text"`
	Price float64 `json:"price" gorm:"type:float"`
}

func Create(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[products.Products](c)
}

func Find(c *fiber.Ctx) error {

	return genericHandler.FindHandler[products.Products](c)
}

func Update(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Products, products.Products](c)
}

func Delete(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Products, products.Products](c)
}
