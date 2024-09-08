package main

import (
	"fmt"
	// "log"

	"github.com/abdul/erp_backend/config"
	branchController "github.com/abdul/erp_backend/controllers/branch"
	countryController "github.com/abdul/erp_backend/controllers/country"
	designationController "github.com/abdul/erp_backend/controllers/designation"
	employeesController "github.com/abdul/erp_backend/controllers/employees"
	featuresController "github.com/abdul/erp_backend/controllers/features"
	"github.com/abdul/erp_backend/controllers/login"
	"github.com/abdul/erp_backend/controllers/middleware"
	regionController "github.com/abdul/erp_backend/controllers/region"
	roleController "github.com/abdul/erp_backend/controllers/role"
	rolefeaturesController "github.com/abdul/erp_backend/controllers/rolefeatures"
	tenantsController "github.com/abdul/erp_backend/controllers/tenants"
	"github.com/abdul/erp_backend/database/dbAdapter"
	"github.com/abdul/erp_backend/database/migration"
	"github.com/abdul/erp_backend/logger"

	// "github.com/abdul/erp_backend/models/organization/branch"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func main() {
	log := logger.Logger
	config.LoadEnv()
	dbAdapter.DbConnect()
	migration.MigrateDb()
	port := fmt.Sprintf(":%d", config.PORT)
	app := fiber.New()
	app.Use(helmet.New())
	app.Use(middleware.AuthHandler)

	//login
	app.Post("/login", login.Login)
	app.Post("/test", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("test")
	})

	//branch routes
	app.Post("/branch/create", branchController.CreateBranch)
	app.Post("/branch/find", branchController.FindBranch)
	app.Post("/branch/update", branchController.UpdateBranch)
	app.Post("/branch/delete", branchController.DeleteBranch)

	//country routes
	app.Post("/country/create", countryController.CreateCountry)
	app.Post("/country/find", countryController.FindCountry)
	app.Post("/country/update", countryController.UpdateCountry)
	app.Post("/country/delete", countryController.DeleteCountry)

	//designation routes
	app.Post("/designation/create", designationController.CreateDesignation)
	app.Post("/designation/find", designationController.FindDesignation)
	app.Post("/designation/update", designationController.UpdateDesignation)
	app.Post("/designation/delete", designationController.DeleteDesignation)

	//employees routes
	app.Post("/employees/create", employeesController.CreateEmployees)
	app.Post("/employees/find", employeesController.FindEmployees)
	app.Post("/employees/update", employeesController.UpdateEmployees)
	app.Post("/employees/delete", employeesController.DeleteEmployees)

	//features routes
	app.Post("/features/create", featuresController.CreateFeatures)
	app.Post("/features/find", featuresController.FindFeatures)
	app.Post("/features/update", featuresController.UpdateFeatures)
	app.Post("/features/delete", featuresController.DeleteFeatures)

	//region routes
	app.Post("/region/create", regionController.CreateRegion)
	app.Post("/region/find", regionController.FindRegion)
	app.Post("/region/update", regionController.UpdateRegion)
	app.Post("/region/delete", regionController.DeleteRegion)

	//role routes
	app.Post("/role/create", roleController.CreateRole)
	app.Post("/role/find", roleController.FindRole)
	app.Post("/role/update", roleController.UpdateRole)
	app.Post("/role/delete", roleController.DeleteRole)

	//rolefeatures routes
	app.Post("/rolefeatures/create", rolefeaturesController.CreateRolefeatures)
	app.Post("/rolefeatures/find", rolefeaturesController.FindRolefeatures)
	app.Post("/rolefeatures/update", rolefeaturesController.UpdateRolefeatures)
	app.Post("/rolefeatures/delete", rolefeaturesController.DeleteRolefeatures)

	//tenants routes
	app.Post("/tenants/create", tenantsController.CreateTenants)
	app.Post("/tenants/find", tenantsController.FindTenants)
	app.Post("/tenants/update", tenantsController.UpdateTenants)
	app.Post("/tenants/delete", tenantsController.DeleteTenants)

	log.Info().Msgf("listening on port %v", port)
	app.Listen(port)
}
