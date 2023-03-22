// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString},
		{Name: "fullname", Type: field.TypeString},
		{Name: "oauth_provider", Type: field.TypeEnum, Enums: []string{"google"}},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
	}
)

func init() {
}
