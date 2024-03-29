// Code generated by ent, DO NOT EDIT.

package accessright

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the accessright type in the database.
	Label = "access_right"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "code"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeInvitations holds the string denoting the invitations edge name in mutations.
	EdgeInvitations = "invitations"
	// InvitationFieldID holds the string denoting the ID field of the Invitation.
	InvitationFieldID = "uuid"
	// Table holds the table name of the accessright in the database.
	Table = "access_rights"
	// InvitationsTable is the table that holds the invitations relation/edge.
	InvitationsTable = "invitations"
	// InvitationsInverseTable is the table name for the Invitation entity.
	// It exists in this package in order to avoid circular dependency with the "invitation" package.
	InvitationsInverseTable = "invitations"
	// InvitationsColumn is the table column denoting the invitations relation/edge.
	InvitationsColumn = "access_right_code"
)

// Columns holds all SQL columns for accessright fields.
var Columns = []string{
	FieldID,
	FieldDescription,
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

// Order defines the ordering method for the AccessRight queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
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
func newInvitationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvitationsInverseTable, InvitationFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, InvitationsTable, InvitationsColumn),
	)
}
