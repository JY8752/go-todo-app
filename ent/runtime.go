// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"todo-app/ent/schema"
	"todo-app/ent/todo"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoFields[3].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(time.Time)
	// todoDescUpdatedAt is the schema descriptor for updated_at field.
	todoDescUpdatedAt := todoFields[4].Descriptor()
	// todo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	todo.DefaultUpdatedAt = todoDescUpdatedAt.Default.(time.Time)
}
