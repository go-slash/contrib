// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/zero/ent/pet"
	"entgo.io/contrib/entproto/internal/zero/ent/predicate"
	"entgo.io/contrib/entproto/internal/zero/ent/user"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUserName sets the "user_name" field.
func (uu *UserUpdate) SetUserName(s string) *UserUpdate {
	uu.mutation.SetUserName(s)
	return uu
}

// SetPoints sets the "points" field.
func (uu *UserUpdate) SetPoints(u uint) *UserUpdate {
	uu.mutation.ResetPoints()
	uu.mutation.SetPoints(u)
	return uu
}

// AddPoints adds u to the "points" field.
func (uu *UserUpdate) AddPoints(u int) *UserUpdate {
	uu.mutation.AddPoints(u)
	return uu
}

// SetExp sets the "exp" field.
func (uu *UserUpdate) SetExp(u uint64) *UserUpdate {
	uu.mutation.ResetExp()
	uu.mutation.SetExp(u)
	return uu
}

// AddExp adds u to the "exp" field.
func (uu *UserUpdate) AddExp(u int64) *UserUpdate {
	uu.mutation.AddExp(u)
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(u user.Status) *UserUpdate {
	uu.mutation.SetStatus(u)
	return uu
}

// SetPetID sets the "pet" edge to the Pet entity by ID.
func (uu *UserUpdate) SetPetID(id int) *UserUpdate {
	uu.mutation.SetPetID(id)
	return uu
}

// SetNillablePetID sets the "pet" edge to the Pet entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillablePetID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetPetID(*id)
	}
	return uu
}

// SetPet sets the "pet" edge to the Pet entity.
func (uu *UserUpdate) SetPet(p *Pet) *UserUpdate {
	return uu.SetPetID(p.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearPet clears the "pet" edge to the Pet entity.
func (uu *UserUpdate) ClearPet() *UserUpdate {
	uu.mutation.ClearPet()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Points(); ok {
		_spec.SetField(user.FieldPoints, field.TypeUint, value)
	}
	if value, ok := uu.mutation.AddedPoints(); ok {
		_spec.AddField(user.FieldPoints, field.TypeUint, value)
	}
	if value, ok := uu.mutation.Exp(); ok {
		_spec.SetField(user.FieldExp, field.TypeUint64, value)
	}
	if value, ok := uu.mutation.AddedExp(); ok {
		_spec.AddField(user.FieldExp, field.TypeUint64, value)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeEnum, value)
	}
	if uu.mutation.PetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PetTable,
			Columns: []string{user.PetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PetTable,
			Columns: []string{user.PetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUserName sets the "user_name" field.
func (uuo *UserUpdateOne) SetUserName(s string) *UserUpdateOne {
	uuo.mutation.SetUserName(s)
	return uuo
}

// SetPoints sets the "points" field.
func (uuo *UserUpdateOne) SetPoints(u uint) *UserUpdateOne {
	uuo.mutation.ResetPoints()
	uuo.mutation.SetPoints(u)
	return uuo
}

// AddPoints adds u to the "points" field.
func (uuo *UserUpdateOne) AddPoints(u int) *UserUpdateOne {
	uuo.mutation.AddPoints(u)
	return uuo
}

// SetExp sets the "exp" field.
func (uuo *UserUpdateOne) SetExp(u uint64) *UserUpdateOne {
	uuo.mutation.ResetExp()
	uuo.mutation.SetExp(u)
	return uuo
}

// AddExp adds u to the "exp" field.
func (uuo *UserUpdateOne) AddExp(u int64) *UserUpdateOne {
	uuo.mutation.AddExp(u)
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(u user.Status) *UserUpdateOne {
	uuo.mutation.SetStatus(u)
	return uuo
}

// SetPetID sets the "pet" edge to the Pet entity by ID.
func (uuo *UserUpdateOne) SetPetID(id int) *UserUpdateOne {
	uuo.mutation.SetPetID(id)
	return uuo
}

// SetNillablePetID sets the "pet" edge to the Pet entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePetID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetPetID(*id)
	}
	return uuo
}

// SetPet sets the "pet" edge to the Pet entity.
func (uuo *UserUpdateOne) SetPet(p *Pet) *UserUpdateOne {
	return uuo.SetPetID(p.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearPet clears the "pet" edge to the Pet entity.
func (uuo *UserUpdateOne) ClearPet() *UserUpdateOne {
	uuo.mutation.ClearPet()
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Points(); ok {
		_spec.SetField(user.FieldPoints, field.TypeUint, value)
	}
	if value, ok := uuo.mutation.AddedPoints(); ok {
		_spec.AddField(user.FieldPoints, field.TypeUint, value)
	}
	if value, ok := uuo.mutation.Exp(); ok {
		_spec.SetField(user.FieldExp, field.TypeUint64, value)
	}
	if value, ok := uuo.mutation.AddedExp(); ok {
		_spec.AddField(user.FieldExp, field.TypeUint64, value)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeEnum, value)
	}
	if uuo.mutation.PetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PetTable,
			Columns: []string{user.PetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PetTable,
			Columns: []string{user.PetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
