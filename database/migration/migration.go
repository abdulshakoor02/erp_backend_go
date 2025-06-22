package migration

import (
	"fmt"

	"github.com/abdul/erp_backend/database/dbAdapter"
	"github.com/abdul/erp_backend/models/organization/additionalInfo"
	"github.com/abdul/erp_backend/models/organization/appointments"
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
		&appointments.Appointments{},
	); err != nil {
		fmt.Printf("fialed to migrate %v \n", err)
	}
	// Drop view if exists (optional)
	dbAdapter.DB.Exec("DROP VIEW IF EXISTS lead_views")
	dbAdapter.DB.Exec("DROP VIEW IF EXISTS reciepts_view")
	dbAdapter.DB.Exec("DROP VIEW IF EXISTS ordered_products_view")

	// Create view
	createViewSQL := `
        CREATE VIEW lead_views AS
select  l.*,c."name" as country_name,c.currency_symbol,c.currency,
b."name" as branch_name,b.mobile as branch_mobile,b.email as branch_email ,b.website ,b.address as branch_address,b.tax,
lc."name" as category_name , e.first_name as employee_name  from leads l left join branches b on b.id = l.branch_id 
left join countries c on c.id = l.country_id left join lead_categories lc on lc.id = l.lead_category_id 
left join employees e on e.id = l.employee_id;
    `
	if err := dbAdapter.DB.Exec(createViewSQL).Error; err != nil {
		fmt.Printf("failed to create lead view %v \n", err)
	} else {
		fmt.Println("lead view successfully created")
	}

	// Create reciepts view
	createRecieptView := `
        CREATE VIEW reciepts_view AS
select r.id as id,r.invoice_id as invoice_id,r.amount_paid as amount_paid,i.discount as discount,r.reciept_no as reciept_no,
r.created_at as created_at,i.lead_id as lead_id,i.total as total,i.pending_amount as pending_amount,
i.tenant_id as tenant_id,l."name" as lead_name,l.email as lead_email,l.mobile as lead_mobile,
b."name" as branch_name,b.address as branch_address,b.mobile as branch_mobile,c.currency as currency,b.tax as tax
from reciepts r inner join invoices i on r.invoice_id = i.id inner join leads l on i.lead_id = l.id inner join branches b  on l.branch_id = b.id
inner join countries c on b.country_id = c.id;
    `
	if err := dbAdapter.DB.Exec(createRecieptView).Error; err != nil {
		fmt.Printf("failed to create reciepts view %v \n", err)
	} else {
		fmt.Println("reciepts view successfully created")
	}

	// Create reciepts view
	createOrderProdView := `
        CREATE VIEW ordered_products_view AS
select op.quantity as quantity,op.invoice_id as invoice_id,op.product_id as product_id,op.id as id,p."desc" as product_desc,
op.created_at as created_at,p.price as product_price,p."name" as product_name
from ordered_products op inner join products p on op.product_id = p.id;
    `
	if err := dbAdapter.DB.Exec(createOrderProdView).Error; err != nil {
		fmt.Printf("failed to create orderd products view %v \n", err)
	} else {
		fmt.Println("orderd products view successfully created")
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
