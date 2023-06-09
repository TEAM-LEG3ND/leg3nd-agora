// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"leg3nd-agora/ent/account"
	"leg3nd-agora/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetEmail sets the "email" field.
func (au *AccountUpdate) SetEmail(s string) *AccountUpdate {
	au.mutation.SetEmail(s)
	return au
}

// SetNickname sets the "nickname" field.
func (au *AccountUpdate) SetNickname(s string) *AccountUpdate {
	au.mutation.SetNickname(s)
	return au
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (au *AccountUpdate) SetNillableNickname(s *string) *AccountUpdate {
	if s != nil {
		au.SetNickname(*s)
	}
	return au
}

// ClearNickname clears the value of the "nickname" field.
func (au *AccountUpdate) ClearNickname() *AccountUpdate {
	au.mutation.ClearNickname()
	return au
}

// SetFullName sets the "full_name" field.
func (au *AccountUpdate) SetFullName(s string) *AccountUpdate {
	au.mutation.SetFullName(s)
	return au
}

// SetOauthProvider sets the "oauth_provider" field.
func (au *AccountUpdate) SetOauthProvider(ap account.OauthProvider) *AccountUpdate {
	au.mutation.SetOauthProvider(ap)
	return au
}

// SetStatus sets the "status" field.
func (au *AccountUpdate) SetStatus(a account.Status) *AccountUpdate {
	au.mutation.SetStatus(a)
	return au
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AccountMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.OauthProvider(); ok {
		if err := account.OauthProviderValidator(v); err != nil {
			return &ValidationError{Name: "oauth_provider", err: fmt.Errorf(`ent: validator failed for field "Account.oauth_provider": %w`, err)}
		}
	}
	if v, ok := au.mutation.Status(); ok {
		if err := account.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Account.status": %w`, err)}
		}
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Email(); ok {
		_spec.SetField(account.FieldEmail, field.TypeString, value)
	}
	if value, ok := au.mutation.Nickname(); ok {
		_spec.SetField(account.FieldNickname, field.TypeString, value)
	}
	if au.mutation.NicknameCleared() {
		_spec.ClearField(account.FieldNickname, field.TypeString)
	}
	if value, ok := au.mutation.FullName(); ok {
		_spec.SetField(account.FieldFullName, field.TypeString, value)
	}
	if value, ok := au.mutation.OauthProvider(); ok {
		_spec.SetField(account.FieldOauthProvider, field.TypeEnum, value)
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.SetField(account.FieldStatus, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetEmail sets the "email" field.
func (auo *AccountUpdateOne) SetEmail(s string) *AccountUpdateOne {
	auo.mutation.SetEmail(s)
	return auo
}

// SetNickname sets the "nickname" field.
func (auo *AccountUpdateOne) SetNickname(s string) *AccountUpdateOne {
	auo.mutation.SetNickname(s)
	return auo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableNickname(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetNickname(*s)
	}
	return auo
}

// ClearNickname clears the value of the "nickname" field.
func (auo *AccountUpdateOne) ClearNickname() *AccountUpdateOne {
	auo.mutation.ClearNickname()
	return auo
}

// SetFullName sets the "full_name" field.
func (auo *AccountUpdateOne) SetFullName(s string) *AccountUpdateOne {
	auo.mutation.SetFullName(s)
	return auo
}

// SetOauthProvider sets the "oauth_provider" field.
func (auo *AccountUpdateOne) SetOauthProvider(ap account.OauthProvider) *AccountUpdateOne {
	auo.mutation.SetOauthProvider(ap)
	return auo
}

// SetStatus sets the "status" field.
func (auo *AccountUpdateOne) SetStatus(a account.Status) *AccountUpdateOne {
	auo.mutation.SetStatus(a)
	return auo
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (auo *AccountUpdateOne) Where(ps ...predicate.Account) *AccountUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	return withHooks[*Account, AccountMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.OauthProvider(); ok {
		if err := account.OauthProviderValidator(v); err != nil {
			return &ValidationError{Name: "oauth_provider", err: fmt.Errorf(`ent: validator failed for field "Account.oauth_provider": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Status(); ok {
		if err := account.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Account.status": %w`, err)}
		}
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Email(); ok {
		_spec.SetField(account.FieldEmail, field.TypeString, value)
	}
	if value, ok := auo.mutation.Nickname(); ok {
		_spec.SetField(account.FieldNickname, field.TypeString, value)
	}
	if auo.mutation.NicknameCleared() {
		_spec.ClearField(account.FieldNickname, field.TypeString)
	}
	if value, ok := auo.mutation.FullName(); ok {
		_spec.SetField(account.FieldFullName, field.TypeString, value)
	}
	if value, ok := auo.mutation.OauthProvider(); ok {
		_spec.SetField(account.FieldOauthProvider, field.TypeEnum, value)
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.SetField(account.FieldStatus, field.TypeEnum, value)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
