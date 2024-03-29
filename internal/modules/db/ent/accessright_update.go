// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calend/internal/modules/db/ent/accessright"
	"calend/internal/modules/db/ent/invitation"
	"calend/internal/modules/db/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccessRightUpdate is the builder for updating AccessRight entities.
type AccessRightUpdate struct {
	config
	hooks    []Hook
	mutation *AccessRightMutation
}

// Where appends a list predicates to the AccessRightUpdate builder.
func (aru *AccessRightUpdate) Where(ps ...predicate.AccessRight) *AccessRightUpdate {
	aru.mutation.Where(ps...)
	return aru
}

// SetDescription sets the "description" field.
func (aru *AccessRightUpdate) SetDescription(s string) *AccessRightUpdate {
	aru.mutation.SetDescription(s)
	return aru
}

// AddInvitationIDs adds the "invitations" edge to the Invitation entity by IDs.
func (aru *AccessRightUpdate) AddInvitationIDs(ids ...string) *AccessRightUpdate {
	aru.mutation.AddInvitationIDs(ids...)
	return aru
}

// AddInvitations adds the "invitations" edges to the Invitation entity.
func (aru *AccessRightUpdate) AddInvitations(i ...*Invitation) *AccessRightUpdate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return aru.AddInvitationIDs(ids...)
}

// Mutation returns the AccessRightMutation object of the builder.
func (aru *AccessRightUpdate) Mutation() *AccessRightMutation {
	return aru.mutation
}

// ClearInvitations clears all "invitations" edges to the Invitation entity.
func (aru *AccessRightUpdate) ClearInvitations() *AccessRightUpdate {
	aru.mutation.ClearInvitations()
	return aru
}

// RemoveInvitationIDs removes the "invitations" edge to Invitation entities by IDs.
func (aru *AccessRightUpdate) RemoveInvitationIDs(ids ...string) *AccessRightUpdate {
	aru.mutation.RemoveInvitationIDs(ids...)
	return aru
}

// RemoveInvitations removes "invitations" edges to Invitation entities.
func (aru *AccessRightUpdate) RemoveInvitations(i ...*Invitation) *AccessRightUpdate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return aru.RemoveInvitationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *AccessRightUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AccessRightMutation](ctx, aru.sqlSave, aru.mutation, aru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aru *AccessRightUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *AccessRightUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *AccessRightUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aru *AccessRightUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(accessright.Table, accessright.Columns, sqlgraph.NewFieldSpec(accessright.FieldID, field.TypeString))
	if ps := aru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aru.mutation.Description(); ok {
		_spec.SetField(accessright.FieldDescription, field.TypeString, value)
	}
	if aru.mutation.InvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   accessright.InvitationsTable,
			Columns: []string{accessright.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invitation.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.RemovedInvitationsIDs(); len(nodes) > 0 && !aru.mutation.InvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   accessright.InvitationsTable,
			Columns: []string{accessright.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invitation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.InvitationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   accessright.InvitationsTable,
			Columns: []string{accessright.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invitation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accessright.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aru.mutation.done = true
	return n, nil
}

// AccessRightUpdateOne is the builder for updating a single AccessRight entity.
type AccessRightUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccessRightMutation
}

// SetDescription sets the "description" field.
func (aruo *AccessRightUpdateOne) SetDescription(s string) *AccessRightUpdateOne {
	aruo.mutation.SetDescription(s)
	return aruo
}

// AddInvitationIDs adds the "invitations" edge to the Invitation entity by IDs.
func (aruo *AccessRightUpdateOne) AddInvitationIDs(ids ...string) *AccessRightUpdateOne {
	aruo.mutation.AddInvitationIDs(ids...)
	return aruo
}

// AddInvitations adds the "invitations" edges to the Invitation entity.
func (aruo *AccessRightUpdateOne) AddInvitations(i ...*Invitation) *AccessRightUpdateOne {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return aruo.AddInvitationIDs(ids...)
}

// Mutation returns the AccessRightMutation object of the builder.
func (aruo *AccessRightUpdateOne) Mutation() *AccessRightMutation {
	return aruo.mutation
}

// ClearInvitations clears all "invitations" edges to the Invitation entity.
func (aruo *AccessRightUpdateOne) ClearInvitations() *AccessRightUpdateOne {
	aruo.mutation.ClearInvitations()
	return aruo
}

// RemoveInvitationIDs removes the "invitations" edge to Invitation entities by IDs.
func (aruo *AccessRightUpdateOne) RemoveInvitationIDs(ids ...string) *AccessRightUpdateOne {
	aruo.mutation.RemoveInvitationIDs(ids...)
	return aruo
}

// RemoveInvitations removes "invitations" edges to Invitation entities.
func (aruo *AccessRightUpdateOne) RemoveInvitations(i ...*Invitation) *AccessRightUpdateOne {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return aruo.RemoveInvitationIDs(ids...)
}

// Where appends a list predicates to the AccessRightUpdate builder.
func (aruo *AccessRightUpdateOne) Where(ps ...predicate.AccessRight) *AccessRightUpdateOne {
	aruo.mutation.Where(ps...)
	return aruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruo *AccessRightUpdateOne) Select(field string, fields ...string) *AccessRightUpdateOne {
	aruo.fields = append([]string{field}, fields...)
	return aruo
}

// Save executes the query and returns the updated AccessRight entity.
func (aruo *AccessRightUpdateOne) Save(ctx context.Context) (*AccessRight, error) {
	return withHooks[*AccessRight, AccessRightMutation](ctx, aruo.sqlSave, aruo.mutation, aruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aruo *AccessRightUpdateOne) SaveX(ctx context.Context) *AccessRight {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *AccessRightUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *AccessRightUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aruo *AccessRightUpdateOne) sqlSave(ctx context.Context) (_node *AccessRight, err error) {
	_spec := sqlgraph.NewUpdateSpec(accessright.Table, accessright.Columns, sqlgraph.NewFieldSpec(accessright.FieldID, field.TypeString))
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AccessRight.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accessright.FieldID)
		for _, f := range fields {
			if !accessright.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != accessright.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aruo.mutation.Description(); ok {
		_spec.SetField(accessright.FieldDescription, field.TypeString, value)
	}
	if aruo.mutation.InvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   accessright.InvitationsTable,
			Columns: []string{accessright.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invitation.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.RemovedInvitationsIDs(); len(nodes) > 0 && !aruo.mutation.InvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   accessright.InvitationsTable,
			Columns: []string{accessright.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invitation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.InvitationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   accessright.InvitationsTable,
			Columns: []string{accessright.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invitation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AccessRight{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accessright.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aruo.mutation.done = true
	return _node, nil
}
