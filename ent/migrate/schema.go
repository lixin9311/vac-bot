// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// VacDepartmentsColumns holds the columns for the "vac_departments" table.
	VacDepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "partition", Type: field.TypeString},
		{Name: "department_id", Type: field.TypeInt},
		{Name: "data", Type: field.TypeJSON},
	}
	// VacDepartmentsTable holds the schema information for the "vac_departments" table.
	VacDepartmentsTable = &schema.Table{
		Name:        "vac_departments",
		Columns:     VacDepartmentsColumns,
		PrimaryKey:  []*schema.Column{VacDepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "vacdepartment_partition_department_id",
				Unique:  false,
				Columns: []*schema.Column{VacDepartmentsColumns[3], VacDepartmentsColumns[4]},
			},
		},
	}
	// VacUsersColumns holds the columns for the "vac_users" table.
	VacUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "sns_id", Type: field.TypeString, Unique: true},
		{Name: "partition", Type: field.TypeString, Nullable: true},
		{Name: "range_key", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "token", Type: field.TypeString, Nullable: true},
		{Name: "watcher_enabled", Type: field.TypeBool, Nullable: true},
		{Name: "reserve_config", Type: field.TypeJSON, Nullable: true},
		{Name: "reserve_enabled", Type: field.TypeBool, Nullable: true},
		{Name: "reservations", Type: field.TypeJSON, Nullable: true},
	}
	// VacUsersTable holds the schema information for the "vac_users" table.
	VacUsersTable = &schema.Table{
		Name:        "vac_users",
		Columns:     VacUsersColumns,
		PrimaryKey:  []*schema.Column{VacUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "vacuser_sns_id_partition",
				Unique:  false,
				Columns: []*schema.Column{VacUsersColumns[3], VacUsersColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		VacDepartmentsTable,
		VacUsersTable,
	}
)

func init() {
}
