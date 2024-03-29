// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lixin9311/vac-bot/ent/predicate"
	"github.com/lixin9311/vac-bot/ent/vacuser"
)

// VacUserDelete is the builder for deleting a VacUser entity.
type VacUserDelete struct {
	config
	hooks    []Hook
	mutation *VacUserMutation
}

// Where adds a new predicate to the VacUserDelete builder.
func (vud *VacUserDelete) Where(ps ...predicate.VacUser) *VacUserDelete {
	vud.mutation.predicates = append(vud.mutation.predicates, ps...)
	return vud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vud *VacUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vud.hooks) == 0 {
		affected, err = vud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VacUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vud.mutation = mutation
			affected, err = vud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vud.hooks) - 1; i >= 0; i-- {
			mut = vud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vud *VacUserDelete) ExecX(ctx context.Context) int {
	n, err := vud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vud *VacUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: vacuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vacuser.FieldID,
			},
		},
	}
	if ps := vud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vud.driver, _spec)
}

// VacUserDeleteOne is the builder for deleting a single VacUser entity.
type VacUserDeleteOne struct {
	vud *VacUserDelete
}

// Exec executes the deletion query.
func (vudo *VacUserDeleteOne) Exec(ctx context.Context) error {
	n, err := vudo.vud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vacuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vudo *VacUserDeleteOne) ExecX(ctx context.Context) {
	vudo.vud.ExecX(ctx)
}
