package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("email"),
		field.String("nickname"),
		field.String("full_name"),
		field.Enum("oauth_provider").
			Values("google"),
		field.Enum("status").
			Values("draft", "ok", "suspended", "withdraw"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
