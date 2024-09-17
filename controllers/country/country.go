package countryController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/country"
	"github.com/gofiber/fiber/v2"
)

type Country struct {
	ID             string `json:"id"              gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Name           string `json:"name"`
	Code           string `json:"code"`
	Currency       string `json:"currency"`
	CurrencyName   string `json:"currency_name"`
	CurrencySymbol string `json:"currency_symbol"`
}

func CreateCountry(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[country.Country](c)
}

func FindCountry(c *fiber.Ctx) error {

	return genericHandler.FindHandler[Country](c)
}

func UpdateCountry(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Country, country.Country](c)
}

func DeleteCountry(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Country, country.Country](c)
}
