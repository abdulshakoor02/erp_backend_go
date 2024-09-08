package branchController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/branch"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Branch struct {
	ID         string `json:"id"         gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Name       string `json:"name"       gorm:"type:string;size:100"`
	Mobile     string `json:"mobile"     gorm:"type:string;size:100;unique"`
	Email      string `json:"email"      gorm:"type:string;size:300;unique"`
	RegionId   string `json:"regionId"   gorm:"type:string;size:100"`
	Website    string `json:"website"    gorm:"type:string;size:100"`
	CountryId  string `json:"countryId"  gorm:"type:string;size:100"`
	CreatedBy  string `json:"createdBy"  gorm:"type:string;size:100"`
	ModifiedBy string `json:"modifiedBy" gorm:"type:string;size:100"`
	Status     string `json:"status"     gorm:"type:string;size:100"`
	DeletedAt  gorm.DeletedAt
}

func CreateBranch(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[branch.Branch](c)
}

func FindBranch(c *fiber.Ctx) error {

	return genericHandler.FindHandler[Branch](c)
}

func UpdateBranch(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Branch, branch.Branch](c)
}

func DeleteBranch(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Branch, branch.Branch](c)
}
