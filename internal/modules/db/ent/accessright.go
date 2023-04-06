// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calend/internal/modules/db/ent/accessright"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AccessRight is the model entity for the AccessRight schema.
type AccessRight struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccessRightQuery when eager-loading is set.
	Edges AccessRightEdges `json:"edges"`
}

// AccessRightEdges holds the relations/edges for other nodes in the graph.
type AccessRightEdges struct {
	// Invitations holds the value of the invitations edge.
	Invitations []*Invitation `json:"invitations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// InvitationsOrErr returns the Invitations value or an error if the edge
// was not loaded in eager-loading.
func (e AccessRightEdges) InvitationsOrErr() ([]*Invitation, error) {
	if e.loadedTypes[0] {
		return e.Invitations, nil
	}
	return nil, &NotLoadedError{edge: "invitations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AccessRight) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case accessright.FieldID, accessright.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AccessRight", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AccessRight fields.
func (ar *AccessRight) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case accessright.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ar.ID = value.String
			}
		case accessright.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ar.Description = value.String
			}
		}
	}
	return nil
}

// QueryInvitations queries the "invitations" edge of the AccessRight entity.
func (ar *AccessRight) QueryInvitations() *InvitationQuery {
	return NewAccessRightClient(ar.config).QueryInvitations(ar)
}

// Update returns a builder for updating this AccessRight.
// Note that you need to call AccessRight.Unwrap() before calling this method if this AccessRight
// was returned from a transaction, and the transaction was committed or rolled back.
func (ar *AccessRight) Update() *AccessRightUpdateOne {
	return NewAccessRightClient(ar.config).UpdateOne(ar)
}

// Unwrap unwraps the AccessRight entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ar *AccessRight) Unwrap() *AccessRight {
	_tx, ok := ar.config.driver.(*txDriver)
	if !ok {
		panic("ent: AccessRight is not a transactional entity")
	}
	ar.config.driver = _tx.drv
	return ar
}

// String implements the fmt.Stringer.
func (ar *AccessRight) String() string {
	var builder strings.Builder
	builder.WriteString("AccessRight(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ar.ID))
	builder.WriteString("description=")
	builder.WriteString(ar.Description)
	builder.WriteByte(')')
	return builder.String()
}

// AccessRights is a parsable slice of AccessRight.
type AccessRights []*AccessRight
