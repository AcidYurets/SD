// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calend/internal/models/access"
	"calend/internal/modules/db/ent/accessright"
	"calend/internal/modules/db/ent/event"
	"calend/internal/modules/db/ent/invitation"
	"calend/internal/modules/db/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Invitation is the model entity for the Invitation schema.
type Invitation struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UserUUID holds the value of the "user_uuid" field.
	UserUUID string `json:"user_uuid,omitempty"`
	// EventUUID holds the value of the "event_uuid" field.
	EventUUID string `json:"event_uuid,omitempty"`
	// AccessRightCode holds the value of the "access_right_code" field.
	AccessRightCode access.Type `json:"access_right_code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InvitationQuery when eager-loading is set.
	Edges InvitationEdges `json:"edges"`
}

// InvitationEdges holds the relations/edges for other nodes in the graph.
type InvitationEdges struct {
	// Event holds the value of the event edge.
	Event *Event `json:"event,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// AccessRight holds the value of the access_right edge.
	AccessRight *AccessRight `json:"access_right,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvitationEdges) EventOrErr() (*Event, error) {
	if e.loadedTypes[0] {
		if e.Event == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: event.Label}
		}
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvitationEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// AccessRightOrErr returns the AccessRight value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvitationEdges) AccessRightOrErr() (*AccessRight, error) {
	if e.loadedTypes[2] {
		if e.AccessRight == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: accessright.Label}
		}
		return e.AccessRight, nil
	}
	return nil, &NotLoadedError{edge: "access_right"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Invitation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case invitation.FieldID, invitation.FieldUserUUID, invitation.FieldEventUUID, invitation.FieldAccessRightCode:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Invitation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Invitation fields.
func (i *Invitation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case invitation.FieldID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value.Valid {
				i.ID = value.String
			}
		case invitation.FieldUserUUID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_uuid", values[j])
			} else if value.Valid {
				i.UserUUID = value.String
			}
		case invitation.FieldEventUUID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_uuid", values[j])
			} else if value.Valid {
				i.EventUUID = value.String
			}
		case invitation.FieldAccessRightCode:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_right_code", values[j])
			} else if value.Valid {
				i.AccessRightCode = access.Type(value.String)
			}
		}
	}
	return nil
}

// QueryEvent queries the "event" edge of the Invitation entity.
func (i *Invitation) QueryEvent() *EventQuery {
	return NewInvitationClient(i.config).QueryEvent(i)
}

// QueryUser queries the "user" edge of the Invitation entity.
func (i *Invitation) QueryUser() *UserQuery {
	return NewInvitationClient(i.config).QueryUser(i)
}

// QueryAccessRight queries the "access_right" edge of the Invitation entity.
func (i *Invitation) QueryAccessRight() *AccessRightQuery {
	return NewInvitationClient(i.config).QueryAccessRight(i)
}

// Update returns a builder for updating this Invitation.
// Note that you need to call Invitation.Unwrap() before calling this method if this Invitation
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Invitation) Update() *InvitationUpdateOne {
	return NewInvitationClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Invitation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Invitation) Unwrap() *Invitation {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Invitation is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Invitation) String() string {
	var builder strings.Builder
	builder.WriteString("Invitation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("user_uuid=")
	builder.WriteString(i.UserUUID)
	builder.WriteString(", ")
	builder.WriteString("event_uuid=")
	builder.WriteString(i.EventUUID)
	builder.WriteString(", ")
	builder.WriteString("access_right_code=")
	builder.WriteString(fmt.Sprintf("%v", i.AccessRightCode))
	builder.WriteByte(')')
	return builder.String()
}

// Invitations is a parsable slice of Invitation.
type Invitations []*Invitation
