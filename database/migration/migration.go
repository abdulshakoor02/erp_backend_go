package migration

import (
	"github.com/abdul/erp_backend/database/dbAdapter"
	"github.com/abdul/erp_backend/models/organization/branch"
	"github.com/abdul/erp_backend/models/organization/country"
	"github.com/abdul/erp_backend/models/organization/designation"
	"github.com/abdul/erp_backend/models/organization/employees"
	"github.com/abdul/erp_backend/models/organization/features"
	"github.com/abdul/erp_backend/models/organization/region"
	"github.com/abdul/erp_backend/models/organization/role"
	"github.com/abdul/erp_backend/models/organization/rolefeatures"
	"github.com/abdul/erp_backend/models/organization/tenants"
)

func MigrateDb() {
	dbAdapter.DB.AutoMigrate(&country.Country{}, &tenant.Tenant{}, &region.Region{}, &branch.Branch{}, &designation.Designation{}, &employees.Employees{}, &features.Features{}, &role.Role{}, &rolefeatures.RoleFeatures{})
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
