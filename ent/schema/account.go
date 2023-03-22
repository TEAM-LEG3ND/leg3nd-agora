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
		field.String("email"),
		field.String("nickname"),
		field.String("fullname"),
		field.Enum("oauth_provider").
			Values("google"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
