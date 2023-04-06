// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calend/internal/modules/db/ent/event"
	"calend/internal/modules/db/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Event is the model entity for the Event schema.
type Event struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// IsWholeDay holds the value of the "is_whole_day" field.
	IsWholeDay bool `json:"is_whole_day,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventQuery when eager-loading is set.
	Edges        EventEdges `json:"edges"`
	creator_uuid *string
}

// EventEdges holds the relations/edges for other nodes in the graph.
type EventEdges struct {
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// Invitations holds the value of the invitations edge.
	Invitations []*Invitation `json:"invitations,omitempty"`
	// Creator holds the value of the creator edge.
	Creator *User `json:"creator,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e EventEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// InvitationsOrErr returns the Invitations value or an error if the edge
// was not loaded in eager-loading.
func (e EventEdges) InvitationsOrErr() ([]*Invitation, error) {
	if e.loadedTypes[1] {
		return e.Invitations, nil
	}
	return nil, &NotLoadedError{edge: "invitations"}
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EventEdges) CreatorOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.Creator == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Event) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case event.FieldIsWholeDay:
			values[i] = new(sql.NullBool)
		case event.FieldID, event.FieldName, event.FieldDescription, event.FieldType:
			values[i] = new(sql.NullString)
		case event.FieldCreatedAt, event.FieldUpdatedAt, event.FieldDeletedAt, event.FieldTimestamp:
			values[i] = new(sql.NullTime)
		case event.ForeignKeys[0]: // creator_uuid
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Event", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Event fields.
func (e *Event) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case event.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				e.ID = value.String
			}
		case event.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case event.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		case event.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				e.DeletedAt = value.Time
			}
		case event.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				e.Timestamp = value.Time
			}
		case event.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case event.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				e.Description = new(string)
				*e.Description = value.String
			}
		case event.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				e.Type = value.String
			}
		case event.FieldIsWholeDay:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_whole_day", values[i])
			} else if value.Valid {
				e.IsWholeDay = value.Bool
			}
		case event.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator_uuid", values[i])
			} else if value.Valid {
				e.creator_uuid = new(string)
				*e.creator_uuid = value.String
			}
		}
	}
	return nil
}

// QueryTags queries the "tags" edge of the Event entity.
func (e *Event) QueryTags() *TagQuery {
	return NewEventClient(e.config).QueryTags(e)
}

// QueryInvitations queries the "invitations" edge of the Event entity.
func (e *Event) QueryInvitations() *InvitationQuery {
	return NewEventClient(e.config).QueryInvitations(e)
}

// QueryCreator queries the "creator" edge of the Event entity.
func (e *Event) QueryCreator() *UserQuery {
	return NewEventClient(e.config).QueryCreator(e)
}

// Update returns a builder for updating this Event.
// Note that you need to call Event.Unwrap() before calling this method if this Event
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Event) Update() *EventUpdateOne {
	return NewEventClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Event entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Event) Unwrap() *Event {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Event is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Event) String() string {
	var builder strings.Builder
	builder.WriteString("Event(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(e.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(e.Timestamp.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	if v := e.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(e.Type)
	builder.WriteString(", ")
	builder.WriteString("is_whole_day=")
	builder.WriteString(fmt.Sprintf("%v", e.IsWholeDay))
	builder.WriteByte(')')
	return builder.String()
}

// Events is a parsable slice of Event.
type Events []*Event
