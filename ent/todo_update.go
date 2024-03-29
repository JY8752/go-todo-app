// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todo-app/ent/predicate"
	"todo-app/ent/todo"
	"todo-app/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoUpdate is the builder for updating Todo entities.
type TodoUpdate struct {
	config
	hooks    []Hook
	mutation *TodoMutation
}

// Where appends a list predicates to the TodoUpdate builder.
func (tu *TodoUpdate) Where(ps ...predicate.Todo) *TodoUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUserID sets the "user_id" field.
func (tu *TodoUpdate) SetUserID(i int) *TodoUpdate {
	tu.mutation.SetUserID(i)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TodoUpdate) SetTitle(s string) *TodoUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetDetail sets the "detail" field.
func (tu *TodoUpdate) SetDetail(s string) *TodoUpdate {
	tu.mutation.SetDetail(s)
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TodoUpdate) SetCreatedAt(t time.Time) *TodoUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableCreatedAt(t *time.Time) *TodoUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TodoUpdate) SetUpdatedAt(t time.Time) *TodoUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableUpdatedAt(t *time.Time) *TodoUpdate {
	if t != nil {
		tu.SetUpdatedAt(*t)
	}
	return tu
}

// SetCompletedAt sets the "completed_at" field.
func (tu *TodoUpdate) SetCompletedAt(t time.Time) *TodoUpdate {
	tu.mutation.SetCompletedAt(t)
	return tu
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableCompletedAt(t *time.Time) *TodoUpdate {
	if t != nil {
		tu.SetCompletedAt(*t)
	}
	return tu
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (tu *TodoUpdate) ClearCompletedAt() *TodoUpdate {
	tu.mutation.ClearCompletedAt()
	return tu
}

// SetUser sets the "user" edge to the User entity.
func (tu *TodoUpdate) SetUser(u *User) *TodoUpdate {
	return tu.SetUserID(u.ID)
}

// Mutation returns the TodoMutation object of the builder.
func (tu *TodoUpdate) Mutation() *TodoMutation {
	return tu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (tu *TodoUpdate) ClearUser() *TodoUpdate {
	tu.mutation.ClearUser()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TodoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TodoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TodoUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TodoUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TodoUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TodoUpdate) check() error {
	if _, ok := tu.mutation.UserID(); tu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Todo.user"`)
	}
	return nil
}

func (tu *TodoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   todo.Table,
			Columns: todo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: todo.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todo.FieldTitle,
		})
	}
	if value, ok := tu.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todo.FieldDetail,
		})
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldCreatedAt,
		})
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldUpdatedAt,
		})
	}
	if value, ok := tu.mutation.CompletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldCompletedAt,
		})
	}
	if tu.mutation.CompletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: todo.FieldCompletedAt,
		})
	}
	if tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.UserTable,
			Columns: []string{todo.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.UserTable,
			Columns: []string{todo.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TodoUpdateOne is the builder for updating a single Todo entity.
type TodoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TodoMutation
}

// SetUserID sets the "user_id" field.
func (tuo *TodoUpdateOne) SetUserID(i int) *TodoUpdateOne {
	tuo.mutation.SetUserID(i)
	return tuo
}

// SetTitle sets the "title" field.
func (tuo *TodoUpdateOne) SetTitle(s string) *TodoUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetDetail sets the "detail" field.
func (tuo *TodoUpdateOne) SetDetail(s string) *TodoUpdateOne {
	tuo.mutation.SetDetail(s)
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TodoUpdateOne) SetCreatedAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableCreatedAt(t *time.Time) *TodoUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TodoUpdateOne) SetUpdatedAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableUpdatedAt(t *time.Time) *TodoUpdateOne {
	if t != nil {
		tuo.SetUpdatedAt(*t)
	}
	return tuo
}

// SetCompletedAt sets the "completed_at" field.
func (tuo *TodoUpdateOne) SetCompletedAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetCompletedAt(t)
	return tuo
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableCompletedAt(t *time.Time) *TodoUpdateOne {
	if t != nil {
		tuo.SetCompletedAt(*t)
	}
	return tuo
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (tuo *TodoUpdateOne) ClearCompletedAt() *TodoUpdateOne {
	tuo.mutation.ClearCompletedAt()
	return tuo
}

// SetUser sets the "user" edge to the User entity.
func (tuo *TodoUpdateOne) SetUser(u *User) *TodoUpdateOne {
	return tuo.SetUserID(u.ID)
}

// Mutation returns the TodoMutation object of the builder.
func (tuo *TodoUpdateOne) Mutation() *TodoMutation {
	return tuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (tuo *TodoUpdateOne) ClearUser() *TodoUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TodoUpdateOne) Select(field string, fields ...string) *TodoUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Todo entity.
func (tuo *TodoUpdateOne) Save(ctx context.Context) (*Todo, error) {
	var (
		err  error
		node *Todo
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TodoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Todo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TodoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TodoUpdateOne) SaveX(ctx context.Context) *Todo {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TodoUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TodoUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TodoUpdateOne) check() error {
	if _, ok := tuo.mutation.UserID(); tuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Todo.user"`)
	}
	return nil
}

func (tuo *TodoUpdateOne) sqlSave(ctx context.Context) (_node *Todo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   todo.Table,
			Columns: todo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: todo.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Todo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todo.FieldID)
		for _, f := range fields {
			if !todo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != todo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todo.FieldTitle,
		})
	}
	if value, ok := tuo.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todo.FieldDetail,
		})
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldCreatedAt,
		})
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldUpdatedAt,
		})
	}
	if value, ok := tuo.mutation.CompletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldCompletedAt,
		})
	}
	if tuo.mutation.CompletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: todo.FieldCompletedAt,
		})
	}
	if tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.UserTable,
			Columns: []string{todo.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.UserTable,
			Columns: []string{todo.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Todo{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
