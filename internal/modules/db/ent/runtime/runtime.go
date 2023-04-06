// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"calend/internal/modules/db/ent/event"
	"calend/internal/modules/db/ent/invitation"
	"calend/internal/modules/db/ent/tag"
	"calend/internal/modules/db/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	eventMixin := schema.Event{}.Mixin()
	eventMixinHooks0 := eventMixin[0].Hooks()
	event.Hooks[0] = eventMixinHooks0[0]
	eventMixinInters0 := eventMixin[0].Interceptors()
	event.Interceptors[0] = eventMixinInters0[0]
	eventMixinFields0 := eventMixin[0].Fields()
	_ = eventMixinFields0
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescCreatedAt is the schema descriptor for created_at field.
	eventDescCreatedAt := eventMixinFields0[1].Descriptor()
	// event.DefaultCreatedAt holds the default value on creation for the created_at field.
	event.DefaultCreatedAt = eventDescCreatedAt.Default.(func() time.Time)
	// eventDescUpdatedAt is the schema descriptor for updated_at field.
	eventDescUpdatedAt := eventMixinFields0[2].Descriptor()
	// event.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	event.DefaultUpdatedAt = eventDescUpdatedAt.Default.(func() time.Time)
	// event.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	event.UpdateDefaultUpdatedAt = eventDescUpdatedAt.UpdateDefault.(func() time.Time)
	// eventDescID is the schema descriptor for id field.
	eventDescID := eventMixinFields0[0].Descriptor()
	// event.DefaultID holds the default value on creation for the id field.
	event.DefaultID = eventDescID.Default.(func() string)
	invitationMixin := schema.Invitation{}.Mixin()
	invitationMixinFields0 := invitationMixin[0].Fields()
	_ = invitationMixinFields0
	invitationFields := schema.Invitation{}.Fields()
	_ = invitationFields
	// invitationDescID is the schema descriptor for id field.
	invitationDescID := invitationMixinFields0[0].Descriptor()
	// invitation.DefaultID holds the default value on creation for the id field.
	invitation.DefaultID = invitationDescID.Default.(func() string)
	tagMixin := schema.Tag{}.Mixin()
	tagMixinHooks0 := tagMixin[0].Hooks()
	tag.Hooks[0] = tagMixinHooks0[0]
	tagMixinInters0 := tagMixin[0].Interceptors()
	tag.Interceptors[0] = tagMixinInters0[0]
	tagMixinFields0 := tagMixin[0].Fields()
	_ = tagMixinFields0
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescCreatedAt is the schema descriptor for created_at field.
	tagDescCreatedAt := tagMixinFields0[1].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the created_at field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	// tagDescUpdatedAt is the schema descriptor for updated_at field.
	tagDescUpdatedAt := tagMixinFields0[2].Descriptor()
	// tag.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tag.DefaultUpdatedAt = tagDescUpdatedAt.Default.(func() time.Time)
	// tag.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	tag.UpdateDefaultUpdatedAt = tagDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tagDescID is the schema descriptor for id field.
	tagDescID := tagMixinFields0[0].Descriptor()
	// tag.DefaultID holds the default value on creation for the id field.
	tag.DefaultID = tagDescID.Default.(func() string)
}

const (
	Version = "v0.11.10"                                        // Version of ent codegen.
	Sum     = "h1:iqn32ybY5HRW3xSAyMNdNKpZhKgMf1Zunsej9yPKUI8=" // Sum of ent codegen.
)
