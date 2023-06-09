// Code generated by ent, DO NOT EDIT.

package account

import (
	"fmt"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldOauthProvider holds the string denoting the oauth_provider field in the database.
	FieldOauthProvider = "oauth_provider"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the account in the database.
	Table = "accounts"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldNickname,
	FieldFullName,
	FieldOauthProvider,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OauthProvider defines the type for the "oauth_provider" enum field.
type OauthProvider string

// OauthProvider values.
const (
	OauthProviderGoogle OauthProvider = "google"
)

func (op OauthProvider) String() string {
	return string(op)
}

// OauthProviderValidator is a validator for the "oauth_provider" field enum values. It is called by the builders before save.
func OauthProviderValidator(op OauthProvider) error {
	switch op {
	case OauthProviderGoogle:
		return nil
	default:
		return fmt.Errorf("account: invalid enum value for oauth_provider field: %q", op)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusDraft     Status = "draft"
	StatusOk        Status = "ok"
	StatusSuspended Status = "suspended"
	StatusWithdraw  Status = "withdraw"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusDraft, StatusOk, StatusSuspended, StatusWithdraw:
		return nil
	default:
		return fmt.Errorf("account: invalid enum value for status field: %q", s)
	}
}
