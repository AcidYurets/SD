// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calend/internal/modules/db/ent/event"
	"calend/internal/modules/db/ent/invitation"
	"calend/internal/modules/db/ent/predicate"
	"calend/internal/modules/db/ent/user"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InvitationQuery is the builder for querying Invitation entities.
type InvitationQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Invitation
	withEvent  *EventQuery
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InvitationQuery builder.
func (iq *InvitationQuery) Where(ps ...predicate.Invitation) *InvitationQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *InvitationQuery) Limit(limit int) *InvitationQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *InvitationQuery) Offset(offset int) *InvitationQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *InvitationQuery) Unique(unique bool) *InvitationQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *InvitationQuery) Order(o ...OrderFunc) *InvitationQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryEvent chains the current query on the "event" edge.
func (iq *InvitationQuery) QueryEvent() *EventQuery {
	query := (&EventClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invitation.Table, invitation.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, invitation.EventTable, invitation.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (iq *InvitationQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invitation.Table, invitation.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, invitation.UserTable, invitation.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Invitation entity from the query.
// Returns a *NotFoundError when no Invitation was found.
func (iq *InvitationQuery) First(ctx context.Context) (*Invitation, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{invitation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InvitationQuery) FirstX(ctx context.Context) *Invitation {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Invitation ID from the query.
// Returns a *NotFoundError when no Invitation ID was found.
func (iq *InvitationQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{invitation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InvitationQuery) FirstIDX(ctx context.Context) string {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Invitation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Invitation entity is found.
// Returns a *NotFoundError when no Invitation entities are found.
func (iq *InvitationQuery) Only(ctx context.Context) (*Invitation, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{invitation.Label}
	default:
		return nil, &NotSingularError{invitation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InvitationQuery) OnlyX(ctx context.Context) *Invitation {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Invitation ID in the query.
// Returns a *NotSingularError when more than one Invitation ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *InvitationQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{invitation.Label}
	default:
		err = &NotSingularError{invitation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InvitationQuery) OnlyIDX(ctx context.Context) string {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Invitations.
func (iq *InvitationQuery) All(ctx context.Context) ([]*Invitation, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Invitation, *InvitationQuery]()
	return withInterceptors[[]*Invitation](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *InvitationQuery) AllX(ctx context.Context) []*Invitation {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Invitation IDs.
func (iq *InvitationQuery) IDs(ctx context.Context) (ids []string, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(invitation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InvitationQuery) IDsX(ctx context.Context) []string {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InvitationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*InvitationQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InvitationQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InvitationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InvitationQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InvitationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InvitationQuery) Clone() *InvitationQuery {
	if iq == nil {
		return nil
	}
	return &InvitationQuery{
		config:     iq.config,
		ctx:        iq.ctx.Clone(),
		order:      append([]OrderFunc{}, iq.order...),
		inters:     append([]Interceptor{}, iq.inters...),
		predicates: append([]predicate.Invitation{}, iq.predicates...),
		withEvent:  iq.withEvent.Clone(),
		withUser:   iq.withUser.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InvitationQuery) WithEvent(opts ...func(*EventQuery)) *InvitationQuery {
	query := (&EventClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withEvent = query
	return iq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InvitationQuery) WithUser(opts ...func(*UserQuery)) *InvitationQuery {
	query := (&UserClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withUser = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (iq *InvitationQuery) GroupBy(field string, fields ...string) *InvitationGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InvitationGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = invitation.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (iq *InvitationQuery) Select(fields ...string) *InvitationSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &InvitationSelect{InvitationQuery: iq}
	sbuild.label = invitation.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InvitationSelect configured with the given aggregations.
func (iq *InvitationQuery) Aggregate(fns ...AggregateFunc) *InvitationSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *InvitationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !invitation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InvitationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Invitation, error) {
	var (
		nodes       = []*Invitation{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [2]bool{
			iq.withEvent != nil,
			iq.withUser != nil,
		}
	)
	if iq.withEvent != nil || iq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, invitation.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Invitation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Invitation{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withEvent; query != nil {
		if err := iq.loadEvent(ctx, query, nodes, nil,
			func(n *Invitation, e *Event) { n.Edges.Event = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withUser; query != nil {
		if err := iq.loadUser(ctx, query, nodes, nil,
			func(n *Invitation, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *InvitationQuery) loadEvent(ctx context.Context, query *EventQuery, nodes []*Invitation, init func(*Invitation), assign func(*Invitation, *Event)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Invitation)
	for i := range nodes {
		if nodes[i].event_uuid == nil {
			continue
		}
		fk := *nodes[i].event_uuid
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(event.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "event_uuid" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (iq *InvitationQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Invitation, init func(*Invitation), assign func(*Invitation, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Invitation)
	for i := range nodes {
		if nodes[i].user_uuid == nil {
			continue
		}
		fk := *nodes[i].user_uuid
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_uuid" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (iq *InvitationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InvitationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(invitation.Table, invitation.Columns, sqlgraph.NewFieldSpec(invitation.FieldID, field.TypeString))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invitation.FieldID)
		for i := range fields {
			if fields[i] != invitation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InvitationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(invitation.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = invitation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InvitationGroupBy is the group-by builder for Invitation entities.
type InvitationGroupBy struct {
	selector
	build *InvitationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InvitationGroupBy) Aggregate(fns ...AggregateFunc) *InvitationGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *InvitationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvitationQuery, *InvitationGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *InvitationGroupBy) sqlScan(ctx context.Context, root *InvitationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InvitationSelect is the builder for selecting fields of Invitation entities.
type InvitationSelect struct {
	*InvitationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *InvitationSelect) Aggregate(fns ...AggregateFunc) *InvitationSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *InvitationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvitationQuery, *InvitationSelect](ctx, is.InvitationQuery, is, is.inters, v)
}

func (is *InvitationSelect) sqlScan(ctx context.Context, root *InvitationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
