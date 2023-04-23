// Code generated by ent, DO NOT EDIT.

package account

import (
	"leg3nd-agora/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldID, id))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldEmail, v))
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldNickname, v))
}

// FullName applies equality check predicate on the "full_name" field. It's identical to FullNameEQ.
func FullName(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldFullName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldEmail, v))
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldNickname, v))
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldNickname, v))
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldNickname, vs...))
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldNickname, vs...))
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldNickname, v))
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldNickname, v))
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldNickname, v))
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldNickname, v))
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldNickname, v))
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldNickname, v))
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldNickname, v))
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldNickname, v))
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldNickname, v))
}

// FullNameEQ applies the EQ predicate on the "full_name" field.
func FullNameEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldFullName, v))
}

// FullNameNEQ applies the NEQ predicate on the "full_name" field.
func FullNameNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldFullName, v))
}

// FullNameIn applies the In predicate on the "full_name" field.
func FullNameIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldFullName, vs...))
}

// FullNameNotIn applies the NotIn predicate on the "full_name" field.
func FullNameNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldFullName, vs...))
}

// FullNameGT applies the GT predicate on the "full_name" field.
func FullNameGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldFullName, v))
}

// FullNameGTE applies the GTE predicate on the "full_name" field.
func FullNameGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldFullName, v))
}

// FullNameLT applies the LT predicate on the "full_name" field.
func FullNameLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldFullName, v))
}

// FullNameLTE applies the LTE predicate on the "full_name" field.
func FullNameLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldFullName, v))
}

// FullNameContains applies the Contains predicate on the "full_name" field.
func FullNameContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldFullName, v))
}

// FullNameHasPrefix applies the HasPrefix predicate on the "full_name" field.
func FullNameHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldFullName, v))
}

// FullNameHasSuffix applies the HasSuffix predicate on the "full_name" field.
func FullNameHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldFullName, v))
}

// FullNameEqualFold applies the EqualFold predicate on the "full_name" field.
func FullNameEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldFullName, v))
}

// FullNameContainsFold applies the ContainsFold predicate on the "full_name" field.
func FullNameContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldFullName, v))
}

// OauthProviderEQ applies the EQ predicate on the "oauth_provider" field.
func OauthProviderEQ(v OauthProvider) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldOauthProvider, v))
}

// OauthProviderNEQ applies the NEQ predicate on the "oauth_provider" field.
func OauthProviderNEQ(v OauthProvider) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldOauthProvider, v))
}

// OauthProviderIn applies the In predicate on the "oauth_provider" field.
func OauthProviderIn(vs ...OauthProvider) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldOauthProvider, vs...))
}

// OauthProviderNotIn applies the NotIn predicate on the "oauth_provider" field.
func OauthProviderNotIn(vs ...OauthProvider) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldOauthProvider, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldStatus, vs...))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		p(s.Not())
	})
}
