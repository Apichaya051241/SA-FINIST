// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/Yui/app/ent/drug"
	"github.com/Yui/app/ent/predicate"
	"github.com/Yui/app/ent/registerstore"
	"github.com/Yui/app/ent/store"
	"github.com/Yui/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RegisterstoreUpdate is the builder for updating Registerstore entities.
type RegisterstoreUpdate struct {
	config
	hooks      []Hook
	mutation   *RegisterstoreMutation
	predicates []predicate.Registerstore
}

// Where adds a new predicate for the builder.
func (ru *RegisterstoreUpdate) Where(ps ...predicate.Registerstore) *RegisterstoreUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetAmount sets the amount field.
func (ru *RegisterstoreUpdate) SetAmount(i int) *RegisterstoreUpdate {
	ru.mutation.ResetAmount()
	ru.mutation.SetAmount(i)
	return ru
}

// AddAmount adds i to amount.
func (ru *RegisterstoreUpdate) AddAmount(i int) *RegisterstoreUpdate {
	ru.mutation.AddAmount(i)
	return ru
}

// SetStoreID sets the store edge to Store by id.
func (ru *RegisterstoreUpdate) SetStoreID(id int) *RegisterstoreUpdate {
	ru.mutation.SetStoreID(id)
	return ru
}

// SetNillableStoreID sets the store edge to Store by id if the given value is not nil.
func (ru *RegisterstoreUpdate) SetNillableStoreID(id *int) *RegisterstoreUpdate {
	if id != nil {
		ru = ru.SetStoreID(*id)
	}
	return ru
}

// SetStore sets the store edge to Store.
func (ru *RegisterstoreUpdate) SetStore(s *Store) *RegisterstoreUpdate {
	return ru.SetStoreID(s.ID)
}

// SetUserID sets the user edge to User by id.
func (ru *RegisterstoreUpdate) SetUserID(id int) *RegisterstoreUpdate {
	ru.mutation.SetUserID(id)
	return ru
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (ru *RegisterstoreUpdate) SetNillableUserID(id *int) *RegisterstoreUpdate {
	if id != nil {
		ru = ru.SetUserID(*id)
	}
	return ru
}

// SetUser sets the user edge to User.
func (ru *RegisterstoreUpdate) SetUser(u *User) *RegisterstoreUpdate {
	return ru.SetUserID(u.ID)
}

// SetDrugID sets the drug edge to Drug by id.
func (ru *RegisterstoreUpdate) SetDrugID(id int) *RegisterstoreUpdate {
	ru.mutation.SetDrugID(id)
	return ru
}

// SetNillableDrugID sets the drug edge to Drug by id if the given value is not nil.
func (ru *RegisterstoreUpdate) SetNillableDrugID(id *int) *RegisterstoreUpdate {
	if id != nil {
		ru = ru.SetDrugID(*id)
	}
	return ru
}

// SetDrug sets the drug edge to Drug.
func (ru *RegisterstoreUpdate) SetDrug(d *Drug) *RegisterstoreUpdate {
	return ru.SetDrugID(d.ID)
}

// Mutation returns the RegisterstoreMutation object of the builder.
func (ru *RegisterstoreUpdate) Mutation() *RegisterstoreMutation {
	return ru.mutation
}

// ClearStore clears the store edge to Store.
func (ru *RegisterstoreUpdate) ClearStore() *RegisterstoreUpdate {
	ru.mutation.ClearStore()
	return ru
}

// ClearUser clears the user edge to User.
func (ru *RegisterstoreUpdate) ClearUser() *RegisterstoreUpdate {
	ru.mutation.ClearUser()
	return ru
}

// ClearDrug clears the drug edge to Drug.
func (ru *RegisterstoreUpdate) ClearDrug() *RegisterstoreUpdate {
	ru.mutation.ClearDrug()
	return ru
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *RegisterstoreUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ru.mutation.Amount(); ok {
		if err := registerstore.AmountValidator(v); err != nil {
			return 0, &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RegisterstoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RegisterstoreUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RegisterstoreUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RegisterstoreUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RegisterstoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   registerstore.Table,
			Columns: registerstore.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: registerstore.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: registerstore.FieldAmount,
		})
	}
	if value, ok := ru.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: registerstore.FieldAmount,
		})
	}
	if ru.mutation.StoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   registerstore.StoreTable,
			Columns: []string{registerstore.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: store.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   registerstore.StoreTable,
			Columns: []string{registerstore.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: store.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   registerstore.UserTable,
			Columns: []string{registerstore.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   registerstore.UserTable,
			Columns: []string{registerstore.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.DrugCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   registerstore.DrugTable,
			Columns: []string{registerstore.DrugColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drug.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.DrugIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   registerstore.DrugTable,
			Columns: []string{registerstore.DrugColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drug.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{registerstore.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RegisterstoreUpdateOne is the builder for updating a single Registerstore entity.
type RegisterstoreUpdateOne struct {
	config
	hooks    []Hook
	mutation *RegisterstoreMutation
}

// SetAmount sets the amount field.
func (ruo *RegisterstoreUpdateOne) SetAmount(i int) *RegisterstoreUpdateOne {
	ruo.mutation.ResetAmount()
	ruo.mutation.SetAmount(i)
	return ruo
}

// AddAmount adds i to amount.
func (ruo *RegisterstoreUpdateOne) AddAmount(i int) *RegisterstoreUpdateOne {
	ruo.mutation.AddAmount(i)
	return ruo
}

// SetStoreID sets the store edge to Store by id.
func (ruo *RegisterstoreUpdateOne) SetStoreID(id int) *RegisterstoreUpdateOne {
	ruo.mutation.SetStoreID(id)
	return ruo
}

// SetNillableStoreID sets the store edge to Store by id if the given value is not nil.
func (ruo *RegisterstoreUpdateOne) SetNillableStoreID(id *int) *RegisterstoreUpdateOne {
	if id != nil {
		ruo = ruo.SetStoreID(*id)
	}
	return ruo
}

// SetStore sets the store edge to Store.
func (ruo *RegisterstoreUpdateOne) SetStore(s *Store) *RegisterstoreUpdateOne {
	return ruo.SetStoreID(s.ID)
}

// SetUserID sets the user edge to User by id.
func (ruo *RegisterstoreUpdateOne) SetUserID(id int) *RegisterstoreUpdateOne {
	ruo.mutation.SetUserID(id)
	return ruo
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (ruo *RegisterstoreUpdateOne) SetNillableUserID(id *int) *RegisterstoreUpdateOne {
	if id != nil {
		ruo = ruo.SetUserID(*id)
	}
	return ruo
}

// SetUser sets the user edge to User.
func (ruo *RegisterstoreUpdateOne) SetUser(u *User) *RegisterstoreUpdateOne {
	return ruo.SetUserID(u.ID)
}

// SetDrugID sets the drug edge to Drug by id.
func (ruo *RegisterstoreUpdateOne) SetDrugID(id int) *RegisterstoreUpdateOne {
	ruo.mutation.SetDrugID(id)
	return ruo
}

// SetNillableDrugID sets the drug edge to Drug by id if the given value is not nil.
func (ruo *RegisterstoreUpdateOne) SetNillableDrugID(id *int) *RegisterstoreUpdateOne {
	if id != nil {
		ruo = ruo.SetDrugID(*id)
	}
	return ruo
}

// SetDrug sets the drug edge to Drug.
func (ruo *RegisterstoreUpdateOne) SetDrug(d *Drug) *RegisterstoreUpdateOne {
	return ruo.SetDrugID(d.ID)
}

// Mutation returns the RegisterstoreMutation object of the builder.
func (ruo *RegisterstoreUpdateOne) Mutation() *RegisterstoreMutation {
	return ruo.mutation
}

// ClearStore clears the store edge to Store.
func (ruo *RegisterstoreUpdateOne) ClearStore() *RegisterstoreUpdateOne {
	ruo.mutation.ClearStore()
	return ruo
}

// ClearUser clears the user edge to User.
func (ruo *RegisterstoreUpdateOne) ClearUser() *RegisterstoreUpdateOne {
	ruo.mutation.ClearUser()
	return ruo
}

// ClearDrug clears the drug edge to Drug.
func (ruo *RegisterstoreUpdateOne) ClearDrug() *RegisterstoreUpdateOne {
	ruo.mutation.ClearDrug()
	return ruo
}

// Save executes the query and returns the updated entity.
func (ruo *RegisterstoreUpdateOne) Save(ctx context.Context) (*Registerstore, error) {
	if v, ok := ruo.mutation.Amount(); ok {
		if err := registerstore.AmountValidator(v); err != nil {
			return nil, &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}

	var (
		err  error
		node *Registerstore
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RegisterstoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RegisterstoreUpdateOne) SaveX(ctx context.Context) *Registerstore {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *RegisterstoreUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RegisterstoreUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RegisterstoreUpdateOne) sqlSave(ctx context.Context) (r *Registerstore, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   registerstore.Table,
			Columns: registerstore.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: registerstore.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Registerstore.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: registerstore.FieldAmount,
		})
	}
	if value, ok := ruo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: registerstore.FieldAmount,
		})
	}
	if ruo.mutation.StoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   registerstore.StoreTable,
			Columns: []string{registerstore.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: store.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   registerstore.StoreTable,
			Columns: []string{registerstore.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: store.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   registerstore.UserTable,
			Columns: []string{registerstore.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   registerstore.UserTable,
			Columns: []string{registerstore.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.DrugCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   registerstore.DrugTable,
			Columns: []string{registerstore.DrugColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drug.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.DrugIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   registerstore.DrugTable,
			Columns: []string{registerstore.DrugColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drug.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Registerstore{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{registerstore.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}
