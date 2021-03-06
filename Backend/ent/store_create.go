// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Yui/app/ent/registerstore"
	"github.com/Yui/app/ent/store"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// StoreCreate is the builder for creating a Store entity.
type StoreCreate struct {
	config
	mutation *StoreMutation
	hooks    []Hook
}

// SetName sets the name field.
func (sc *StoreCreate) SetName(s string) *StoreCreate {
	sc.mutation.SetName(s)
	return sc
}

// AddStoreIDs adds the stores edge to Registerstore by ids.
func (sc *StoreCreate) AddStoreIDs(ids ...int) *StoreCreate {
	sc.mutation.AddStoreIDs(ids...)
	return sc
}

// AddStores adds the stores edges to Registerstore.
func (sc *StoreCreate) AddStores(r ...*Registerstore) *StoreCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return sc.AddStoreIDs(ids...)
}

// Mutation returns the StoreMutation object of the builder.
func (sc *StoreCreate) Mutation() *StoreMutation {
	return sc.mutation
}

// Save creates the Store in the database.
func (sc *StoreCreate) Save(ctx context.Context) (*Store, error) {
	if _, ok := sc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := store.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	var (
		err  error
		node *Store
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StoreCreate) SaveX(ctx context.Context) *Store {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *StoreCreate) sqlSave(ctx context.Context) (*Store, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *StoreCreate) createSpec() (*Store, *sqlgraph.CreateSpec) {
	var (
		s     = &Store{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: store.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: store.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: store.FieldName,
		})
		s.Name = value
	}
	if nodes := sc.mutation.StoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.StoresTable,
			Columns: []string{store.StoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: registerstore.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}
