package migration

import (
	"fmt"

	"github.com/abdul/erp_backend/database/dbAdapter"
	"github.com/abdul/erp_backend/models/organization/additionalInfo"
	"github.com/abdul/erp_backend/models/organization/branch"
	"github.com/abdul/erp_backend/models/organization/country"
	"github.com/abdul/erp_backend/models/organization/designation"
	"github.com/abdul/erp_backend/models/organization/employees"
	"github.com/abdul/erp_backend/models/organization/features"
	"github.com/abdul/erp_backend/models/organization/invoice"
	"github.com/abdul/erp_backend/models/organization/leadCategory"
	"github.com/abdul/erp_backend/models/organization/leads"
	"github.com/abdul/erp_backend/models/organization/orderedProduct"
	"github.com/abdul/erp_backend/models/organization/products"
	"github.com/abdul/erp_backend/models/organization/reciepts"
	"github.com/abdul/erp_backend/models/organization/region"
	"github.com/abdul/erp_backend/models/organization/role"
	"github.com/abdul/erp_backend/models/organization/rolefeatures"
	tenant "github.com/abdul/erp_backend/models/organization/tenants"
)

func MigrateDb() {
	if err := dbAdapter.DB.AutoMigrate(
		&country.Country{},
		&tenant.Tenant{},
		&region.Region{},
		&branch.Branch{},
		&designation.Designation{},
		&employees.Employees{},
		&features.Features{},
		&role.Role{},
		&rolefeatures.RoleFeatures{},
		&products.Products{},
		&leadCategory.LeadCategory{},
		&leads.Leads{},
		&additionalInfo.AdditionalInfo{},
		&invoice.Invoice{},
		&reciepts.Reciepts{},
		&orderedProduct.OrderedProduct{},
	); err != nil {
		fmt.Printf("fialed to migrate %v \n", err)
	}
	// dbAdapter.DB.Migrator().CreateConstraint(&branch.Branch{}, "Country")
	// dbAdapter.DB.Migrator().CreateConstraint(&branch.Branch{}, "Region")
	// dbAdapter.DB.Migrator().CreateConstraint(&designation.Designation{}, "Tenant")
	// dbAdapter.DB.Migrator().CreateConstraint(&employees.Employees{}, "Country")
	// dbAdapter.DB.Migrator().CreateConstraint(&employees.Employees{}, "Tenant")
	// dbAdapter.DB.Migrator().CreateConstraint(&region.Region{}, "Tenant")
	// dbAdapter.DB.Migrator().CreateConstraint(&rolefeatures.RoleFeatures{}, "Tenant")
	// dbAdapter.DB.Migrator().CreateConstraint(&rolefeatures.RoleFeatures{}, "Features")
	// dbAdapter.DB.Migrator().CreateConstraint(&rolefeatures.RoleFeatures{}, "Role")
}
