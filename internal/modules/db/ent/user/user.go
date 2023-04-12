// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uuid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldLogin holds the string denoting the login field in the database.
	FieldLogin = "login"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// EdgeInvitations holds the string denoting the invitations edge name in mutations.
	EdgeInvitations = "invitations"
	// EdgeCreatedEvents holds the string denoting the created_events edge name in mutations.
	EdgeCreatedEvents = "created_events"
	// Table holds the table name of the user in the database.
	Table = "users"
	// InvitationsTable is the table that holds the invitations relation/edge.
	InvitationsTable = "invitations"
	// InvitationsInverseTable is the table name for the Invitation entity.
	// It exists in this package in order to avoid circular dependency with the "invitation" package.
	InvitationsInverseTable = "invitations"
	// InvitationsColumn is the table column denoting the invitations relation/edge.
	InvitationsColumn = "user_uuid"
	// CreatedEventsTable is the table that holds the created_events relation/edge.
	CreatedEventsTable = "events"
	// CreatedEventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	CreatedEventsInverseTable = "events"
	// CreatedEventsColumn is the table column denoting the created_events relation/edge.
	CreatedEventsColumn = "creator_uuid"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldPhone,
	FieldLogin,
	FieldPasswordHash,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "calend/internal/modules/db/ent/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// Order defines the ordering method for the User queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByLogin orders the results by the login field.
func ByLogin(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldLogin, opts...).ToFunc()
}

// ByPasswordHash orders the results by the password_hash field.
func ByPasswordHash(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPasswordHash, opts...).ToFunc()
}

// ByInvitationsCount orders the results by invitations count.
func ByInvitationsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInvitationsStep(), opts...)
	}
}

// ByInvitations orders the results by invitations terms.
func ByInvitations(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvitationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCreatedEventsCount orders the results by created_events count.
func ByCreatedEventsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreatedEventsStep(), opts...)
	}
}

// ByCreatedEvents orders the results by created_events terms.
func ByCreatedEvents(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedEventsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newInvitationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvitationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, InvitationsTable, InvitationsColumn),
	)
}
func newCreatedEventsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedEventsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CreatedEventsTable, CreatedEventsColumn),
	)
}
