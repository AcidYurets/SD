// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"calend/internal/modules/db/ent"
	"calend/internal/modules/db/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...ent.OrderFunc)
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The AccessRightFunc type is an adapter to allow the use of ordinary function as a Querier.
type AccessRightFunc func(context.Context, *ent.AccessRightQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f AccessRightFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.AccessRightQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.AccessRightQuery", q)
}

// The TraverseAccessRight type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAccessRight func(context.Context, *ent.AccessRightQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAccessRight) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAccessRight) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AccessRightQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.AccessRightQuery", q)
}

// The EventFunc type is an adapter to allow the use of ordinary function as a Querier.
type EventFunc func(context.Context, *ent.EventQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f EventFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.EventQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.EventQuery", q)
}

// The TraverseEvent type is an adapter to allow the use of ordinary function as Traverser.
type TraverseEvent func(context.Context, *ent.EventQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseEvent) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseEvent) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EventQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.EventQuery", q)
}

// The InvitationFunc type is an adapter to allow the use of ordinary function as a Querier.
type InvitationFunc func(context.Context, *ent.InvitationQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f InvitationFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.InvitationQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.InvitationQuery", q)
}

// The TraverseInvitation type is an adapter to allow the use of ordinary function as Traverser.
type TraverseInvitation func(context.Context, *ent.InvitationQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseInvitation) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseInvitation) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.InvitationQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.InvitationQuery", q)
}

// The TagFunc type is an adapter to allow the use of ordinary function as a Querier.
type TagFunc func(context.Context, *ent.TagQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f TagFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.TagQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.TagQuery", q)
}

// The TraverseTag type is an adapter to allow the use of ordinary function as Traverser.
type TraverseTag func(context.Context, *ent.TagQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseTag) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseTag) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TagQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.TagQuery", q)
}

// The UserFunc type is an adapter to allow the use of ordinary function as a Querier.
type UserFunc func(context.Context, *ent.UserQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f UserFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.UserQuery", q)
}

// The TraverseUser type is an adapter to allow the use of ordinary function as Traverser.
type TraverseUser func(context.Context, *ent.UserQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseUser) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseUser) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.UserQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.AccessRightQuery:
		return &query[*ent.AccessRightQuery, predicate.AccessRight]{typ: ent.TypeAccessRight, tq: q}, nil
	case *ent.EventQuery:
		return &query[*ent.EventQuery, predicate.Event]{typ: ent.TypeEvent, tq: q}, nil
	case *ent.InvitationQuery:
		return &query[*ent.InvitationQuery, predicate.Invitation]{typ: ent.TypeInvitation, tq: q}, nil
	case *ent.TagQuery:
		return &query[*ent.TagQuery, predicate.Tag]{typ: ent.TypeTag, tq: q}, nil
	case *ent.UserQuery:
		return &query[*ent.UserQuery, predicate.User]{typ: ent.TypeUser, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...ent.OrderFunc) T
		Where(...P) T
	}
}

func (q query[T, P]) Type() string {
	return q.typ
}

func (q query[T, P]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P]) Order(orders ...ent.OrderFunc) {
	q.tq.Order(orders...)
}

func (q query[T, P]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}