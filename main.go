package main

import (
	"fmt"

	"github.com/abdul/erp_backend/config"
	addtionalInfoController "github.com/abdul/erp_backend/controllers/addtionalInfo"
	appointmentsController "github.com/abdul/erp_backend/controllers/appointments"
	branchController "github.com/abdul/erp_backend/controllers/branch"
	countryController "github.com/abdul/erp_backend/controllers/country"
	designationController "github.com/abdul/erp_backend/controllers/designation"
	employeesController "github.com/abdul/erp_backend/controllers/employees"
	featuresController "github.com/abdul/erp_backend/controllers/features"
	"github.com/abdul/erp_backend/controllers/fileUpload"
	invoiceController "github.com/abdul/erp_backend/controllers/invoice"
	leadCategoryController "github.com/abdul/erp_backend/controllers/leadCategory"
	leadViewController "github.com/abdul/erp_backend/controllers/leadView"
	leadsController "github.com/abdul/erp_backend/controllers/leads"
	"github.com/abdul/erp_backend/controllers/login"
	"github.com/abdul/erp_backend/controllers/middleware"
	productsController "github.com/abdul/erp_backend/controllers/products"
	regionController "github.com/abdul/erp_backend/controllers/region"
	roleController "github.com/abdul/erp_backend/controllers/role"
	rolefeaturesController "github.com/abdul/erp_backend/controllers/rolefeatures"
	tenantsController "github.com/abdul/erp_backend/controllers/tenants"
	"github.com/abdul/erp_backend/database/dbAdapter"
	"github.com/abdul/erp_backend/database/migration"
	"github.com/abdul/erp_backend/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New())
	app.Use(middleware.AuthHandler)

	//login
	app.Post("/login", login.Login)
	app.Post("/test", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("test")
	})

	//branch routes
	app.Post("/branch/create", branchController.CreateBranch)
	app.Post("/branch/find", branchController.FindBranch)
	app.Post("/branch/findAssociated", branchController.FindBranchAssociated)
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
	app.Post("/employees/findAssociated", employeesController.FindEmployeesAssociated)
	app.Post("/employees/update", employeesController.UpdateEmployees)
	app.Post("/employees/delete", employeesController.DeleteEmployees)

	//features routes
	app.Post("/features/create", featuresController.CreateFeatures)
	app.Post("/features/upsert", featuresController.UpsertFeatures)
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
	app.Post("/rolefeatures/findAssociated", rolefeaturesController.FindRolefeaturesAssociated)
	app.Post("/rolefeatures/update", rolefeaturesController.UpdateRolefeatures)
	app.Post("/rolefeatures/delete", rolefeaturesController.DeleteRolefeatures)

	//tenants routes
	app.Post("/tenants/create", tenantsController.CreateTenants)
	app.Post("/tenants/find", tenantsController.FindTenants)
	app.Post("/tenants/findAssociated", tenantsController.FindAssociatedTenants)
	app.Post("/tenants/update", tenantsController.UpdateTenants)
	app.Post("/tenants/delete", tenantsController.DeleteTenants)

	//leads routes
	app.Post("/lead/create", leadsController.Create)
	app.Post("/lead/find", leadsController.Find)
	app.Post("/lead/update", leadsController.Update)
	app.Post("/lead/delete", leadsController.Delete)
	app.Post("/leadView/find", leadViewController.Find)

	//products routes
	app.Post("/products/create", productsController.Create)
	app.Post("/products/find", productsController.Find)
	app.Post("/products/update", productsController.Update)
	app.Post("/products/delete", productsController.Delete)

	//leadCategory routes
	app.Post("/leadCategory/create", leadCategoryController.Create)
	app.Post("/leadCategory/find", leadCategoryController.Find)
	app.Post("/leadCategory/update", leadCategoryController.Update)
	app.Post("/leadCategory/delete", leadCategoryController.Delete)

	//addtionalInfo routes
	app.Post("/additionalInfo/create", addtionalInfoController.Create)
	app.Post("/additionalInfo/find", addtionalInfoController.Find)
	app.Post("/additionalInfo/update", addtionalInfoController.Update)
	app.Post("/additionalInfo/delete", addtionalInfoController.Delete)

	//appointment routes
	app.Post("/appointment/create", appointmentsController.Create)
	app.Post("/appointment/find", appointmentsController.Find)
	app.Post("/appointment/update", appointmentsController.Update)
	app.Post("/appointment/delete", appointmentsController.Delete)

	//fileUpload
	app.Post("/fileUpload", fileUpload.UploadHandler)
	app.Post("/fileDownload", fileUpload.ImagePostHandler)

	//invoice
	app.Post("/invoice/create", invoiceController.Create)
	app.Post("/reciept/create", invoiceController.GenerateReciept)
	app.Post("/invoice/findone", invoiceController.FindOne)
	app.Post("/invoice/find", invoiceController.FindInvoices)
	app.Post("/reciepts/find", invoiceController.FindReciepts)

	log.Info().Msgf("listening on port %v", port)
	app.Listen(port)
}
