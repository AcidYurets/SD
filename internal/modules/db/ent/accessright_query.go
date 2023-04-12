// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calend/internal/models/access"
	"calend/internal/modules/db/ent/accessright"
	"calend/internal/modules/db/ent/invitation"
	"calend/internal/modules/db/ent/predicate"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccessRightQuery is the builder for querying AccessRight entities.
type AccessRightQuery struct {
	config
	ctx             *QueryContext
	order           []accessright.Order
	inters          []Interceptor
	predicates      []predicate.AccessRight
	withInvitations *InvitationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccessRightQuery builder.
func (arq *AccessRightQuery) Where(ps ...predicate.AccessRight) *AccessRightQuery {
	arq.predicates = append(arq.predicates, ps...)
	return arq
}

// Limit the number of records to be returned by this query.
func (arq *AccessRightQuery) Limit(limit int) *AccessRightQuery {
	arq.ctx.Limit = &limit
	return arq
}

// Offset to start from.
func (arq *AccessRightQuery) Offset(offset int) *AccessRightQuery {
	arq.ctx.Offset = &offset
	return arq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (arq *AccessRightQuery) Unique(unique bool) *AccessRightQuery {
	arq.ctx.Unique = &unique
	return arq
}

// Order specifies how the records should be ordered.
func (arq *AccessRightQuery) Order(o ...accessright.Order) *AccessRightQuery {
	arq.order = append(arq.order, o...)
	return arq
}

// QueryInvitations chains the current query on the "invitations" edge.
func (arq *AccessRightQuery) QueryInvitations() *InvitationQuery {
	query := (&InvitationClient{config: arq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := arq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accessright.Table, accessright.FieldID, selector),
			sqlgraph.To(invitation.Table, invitation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, accessright.InvitationsTable, accessright.InvitationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(arq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AccessRight entity from the query.
// Returns a *NotFoundError when no AccessRight was found.
func (arq *AccessRightQuery) First(ctx context.Context) (*AccessRight, error) {
	nodes, err := arq.Limit(1).All(setContextOp(ctx, arq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{accessright.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (arq *AccessRightQuery) FirstX(ctx context.Context) *AccessRight {
	node, err := arq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AccessRight ID from the query.
// Returns a *NotFoundError when no AccessRight ID was found.
func (arq *AccessRightQuery) FirstID(ctx context.Context) (id access.Type, err error) {
	var ids []access.Type
	if ids, err = arq.Limit(1).IDs(setContextOp(ctx, arq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{accessright.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (arq *AccessRightQuery) FirstIDX(ctx context.Context) access.Type {
	id, err := arq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AccessRight entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AccessRight entity is found.
// Returns a *NotFoundError when no AccessRight entities are found.
func (arq *AccessRightQuery) Only(ctx context.Context) (*AccessRight, error) {
	nodes, err := arq.Limit(2).All(setContextOp(ctx, arq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{accessright.Label}
	default:
		return nil, &NotSingularError{accessright.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (arq *AccessRightQuery) OnlyX(ctx context.Context) *AccessRight {
	node, err := arq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AccessRight ID in the query.
// Returns a *NotSingularError when more than one AccessRight ID is found.
// Returns a *NotFoundError when no entities are found.
func (arq *AccessRightQuery) OnlyID(ctx context.Context) (id access.Type, err error) {
	var ids []access.Type
	if ids, err = arq.Limit(2).IDs(setContextOp(ctx, arq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{accessright.Label}
	default:
		err = &NotSingularError{accessright.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (arq *AccessRightQuery) OnlyIDX(ctx context.Context) access.Type {
	id, err := arq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AccessRights.
func (arq *AccessRightQuery) All(ctx context.Context) ([]*AccessRight, error) {
	ctx = setContextOp(ctx, arq.ctx, "All")
	if err := arq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AccessRight, *AccessRightQuery]()
	return withInterceptors[[]*AccessRight](ctx, arq, qr, arq.inters)
}

// AllX is like All, but panics if an error occurs.
func (arq *AccessRightQuery) AllX(ctx context.Context) []*AccessRight {
	nodes, err := arq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AccessRight IDs.
func (arq *AccessRightQuery) IDs(ctx context.Context) (ids []access.Type, err error) {
	if arq.ctx.Unique == nil && arq.path != nil {
		arq.Unique(true)
	}
	ctx = setContextOp(ctx, arq.ctx, "IDs")
	if err = arq.Select(accessright.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (arq *AccessRightQuery) IDsX(ctx context.Context) []access.Type {
	ids, err := arq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (arq *AccessRightQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, arq.ctx, "Count")
	if err := arq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, arq, querierCount[*AccessRightQuery](), arq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (arq *AccessRightQuery) CountX(ctx context.Context) int {
	count, err := arq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (arq *AccessRightQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, arq.ctx, "Exist")
	switch _, err := arq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (arq *AccessRightQuery) ExistX(ctx context.Context) bool {
	exist, err := arq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccessRightQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (arq *AccessRightQuery) Clone() *AccessRightQuery {
	if arq == nil {
		return nil
	}
	return &AccessRightQuery{
		config:          arq.config,
		ctx:             arq.ctx.Clone(),
		order:           append([]accessright.Order{}, arq.order...),
		inters:          append([]Interceptor{}, arq.inters...),
		predicates:      append([]predicate.AccessRight{}, arq.predicates...),
		withInvitations: arq.withInvitations.Clone(),
		// clone intermediate query.
		sql:  arq.sql.Clone(),
		path: arq.path,
	}
}

// WithInvitations tells the query-builder to eager-load the nodes that are connected to
// the "invitations" edge. The optional arguments are used to configure the query builder of the edge.
func (arq *AccessRightQuery) WithInvitations(opts ...func(*InvitationQuery)) *AccessRightQuery {
	query := (&InvitationClient{config: arq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	arq.withInvitations = query
	return arq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AccessRight.Query().
//		GroupBy(accessright.FieldDescription).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (arq *AccessRightQuery) GroupBy(field string, fields ...string) *AccessRightGroupBy {
	arq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AccessRightGroupBy{build: arq}
	grbuild.flds = &arq.ctx.Fields
	grbuild.label = accessright.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//	}
//
//	client.AccessRight.Query().
//		Select(accessright.FieldDescription).
//		Scan(ctx, &v)
func (arq *AccessRightQuery) Select(fields ...string) *AccessRightSelect {
	arq.ctx.Fields = append(arq.ctx.Fields, fields...)
	sbuild := &AccessRightSelect{AccessRightQuery: arq}
	sbuild.label = accessright.Label
	sbuild.flds, sbuild.scan = &arq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AccessRightSelect configured with the given aggregations.
func (arq *AccessRightQuery) Aggregate(fns ...AggregateFunc) *AccessRightSelect {
	return arq.Select().Aggregate(fns...)
}

func (arq *AccessRightQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range arq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, arq); err != nil {
				return err
			}
		}
	}
	for _, f := range arq.ctx.Fields {
		if !accessright.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if arq.path != nil {
		prev, err := arq.path(ctx)
		if err != nil {
			return err
		}
		arq.sql = prev
	}
	return nil
}

func (arq *AccessRightQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AccessRight, error) {
	var (
		nodes       = []*AccessRight{}
		_spec       = arq.querySpec()
		loadedTypes = [1]bool{
			arq.withInvitations != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AccessRight).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AccessRight{config: arq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, arq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := arq.withInvitations; query != nil {
		if err := arq.loadInvitations(ctx, query, nodes,
			func(n *AccessRight) { n.Edges.Invitations = []*Invitation{} },
			func(n *AccessRight, e *Invitation) { n.Edges.Invitations = append(n.Edges.Invitations, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (arq *AccessRightQuery) loadInvitations(ctx context.Context, query *InvitationQuery, nodes []*AccessRight, init func(*AccessRight), assign func(*AccessRight, *Invitation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[access.Type]*AccessRight)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Invitation(func(s *sql.Selector) {
		s.Where(sql.InValues(accessright.InvitationsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AccessRightCode
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "access_right_code" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (arq *AccessRightQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := arq.querySpec()
	_spec.Node.Columns = arq.ctx.Fields
	if len(arq.ctx.Fields) > 0 {
		_spec.Unique = arq.ctx.Unique != nil && *arq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, arq.driver, _spec)
}

func (arq *AccessRightQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(accessright.Table, accessright.Columns, sqlgraph.NewFieldSpec(accessright.FieldID, field.TypeString))
	_spec.From = arq.sql
	if unique := arq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if arq.path != nil {
		_spec.Unique = true
	}
	if fields := arq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accessright.FieldID)
		for i := range fields {
			if fields[i] != accessright.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := arq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := arq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := arq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := arq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (arq *AccessRightQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(arq.driver.Dialect())
	t1 := builder.Table(accessright.Table)
	columns := arq.ctx.Fields
	if len(columns) == 0 {
		columns = accessright.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if arq.sql != nil {
		selector = arq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if arq.ctx.Unique != nil && *arq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range arq.predicates {
		p(selector)
	}
	for _, p := range arq.order {
		p(selector)
	}
	if offset := arq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := arq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AccessRightGroupBy is the group-by builder for AccessRight entities.
type AccessRightGroupBy struct {
	selector
	build *AccessRightQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (argb *AccessRightGroupBy) Aggregate(fns ...AggregateFunc) *AccessRightGroupBy {
	argb.fns = append(argb.fns, fns...)
	return argb
}

// Scan applies the selector query and scans the result into the given value.
func (argb *AccessRightGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, argb.build.ctx, "GroupBy")
	if err := argb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccessRightQuery, *AccessRightGroupBy](ctx, argb.build, argb, argb.build.inters, v)
}

func (argb *AccessRightGroupBy) sqlScan(ctx context.Context, root *AccessRightQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(argb.fns))
	for _, fn := range argb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*argb.flds)+len(argb.fns))
		for _, f := range *argb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*argb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := argb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AccessRightSelect is the builder for selecting fields of AccessRight entities.
type AccessRightSelect struct {
	*AccessRightQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ars *AccessRightSelect) Aggregate(fns ...AggregateFunc) *AccessRightSelect {
	ars.fns = append(ars.fns, fns...)
	return ars
}

// Scan applies the selector query and scans the result into the given value.
func (ars *AccessRightSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ars.ctx, "Select")
	if err := ars.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccessRightQuery, *AccessRightSelect](ctx, ars.AccessRightQuery, ars, ars.inters, v)
}

func (ars *AccessRightSelect) sqlScan(ctx context.Context, root *AccessRightQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ars.fns))
	for _, fn := range ars.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ars.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
