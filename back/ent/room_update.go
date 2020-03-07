// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"studi-guide/ent/door"
	"studi-guide/ent/pathnode"
	"studi-guide/ent/predicate"
	"studi-guide/ent/room"
	"studi-guide/ent/sequence"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RoomUpdate is the builder for updating Room entities.
type RoomUpdate struct {
	config
	Name             *string
	Description      *string
	Floor            *int
	addFloor         *int
	Id               *int
	addId            *int
	doors            map[int]struct{}
	sequences        map[int]struct{}
	pathNodes        map[int]struct{}
	removedDoors     map[int]struct{}
	removedSequences map[int]struct{}
	removedPathNodes map[int]struct{}
	predicates       []predicate.Room
}

// Where adds a new predicate for the builder.
func (ru *RoomUpdate) Where(ps ...predicate.Room) *RoomUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetName sets the Name field.
func (ru *RoomUpdate) SetName(s string) *RoomUpdate {
	ru.Name = &s
	return ru
}

// SetDescription sets the Description field.
func (ru *RoomUpdate) SetDescription(s string) *RoomUpdate {
	ru.Description = &s
	return ru
}

// SetNillableDescription sets the Description field if the given value is not nil.
func (ru *RoomUpdate) SetNillableDescription(s *string) *RoomUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// SetFloor sets the Floor field.
func (ru *RoomUpdate) SetFloor(i int) *RoomUpdate {
	ru.Floor = &i
	ru.addFloor = nil
	return ru
}

// SetNillableFloor sets the Floor field if the given value is not nil.
func (ru *RoomUpdate) SetNillableFloor(i *int) *RoomUpdate {
	if i != nil {
		ru.SetFloor(*i)
	}
	return ru
}

// AddFloor adds i to Floor.
func (ru *RoomUpdate) AddFloor(i int) *RoomUpdate {
	if ru.addFloor == nil {
		ru.addFloor = &i
	} else {
		*ru.addFloor += i
	}
	return ru
}

// SetID sets the Id field.
func (ru *RoomUpdate) SetID(i int) *RoomUpdate {
	ru.Id = &i
	ru.addId = nil
	return ru
}

// AddID adds i to Id.
func (ru *RoomUpdate) AddID(i int) *RoomUpdate {
	if ru.addId == nil {
		ru.addId = &i
	} else {
		*ru.addId += i
	}
	return ru
}

// AddDoorIDs adds the doors edge to Door by ids.
func (ru *RoomUpdate) AddDoorIDs(ids ...int) *RoomUpdate {
	if ru.doors == nil {
		ru.doors = make(map[int]struct{})
	}
	for i := range ids {
		ru.doors[ids[i]] = struct{}{}
	}
	return ru
}

// AddDoors adds the doors edges to Door.
func (ru *RoomUpdate) AddDoors(d ...*Door) *RoomUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ru.AddDoorIDs(ids...)
}

// AddSequenceIDs adds the sequences edge to Sequence by ids.
func (ru *RoomUpdate) AddSequenceIDs(ids ...int) *RoomUpdate {
	if ru.sequences == nil {
		ru.sequences = make(map[int]struct{})
	}
	for i := range ids {
		ru.sequences[ids[i]] = struct{}{}
	}
	return ru
}

// AddSequences adds the sequences edges to Sequence.
func (ru *RoomUpdate) AddSequences(s ...*Sequence) *RoomUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.AddSequenceIDs(ids...)
}

// AddPathNodeIDs adds the pathNodes edge to PathNode by ids.
func (ru *RoomUpdate) AddPathNodeIDs(ids ...int) *RoomUpdate {
	if ru.pathNodes == nil {
		ru.pathNodes = make(map[int]struct{})
	}
	for i := range ids {
		ru.pathNodes[ids[i]] = struct{}{}
	}
	return ru
}

// AddPathNodes adds the pathNodes edges to PathNode.
func (ru *RoomUpdate) AddPathNodes(p ...*PathNode) *RoomUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddPathNodeIDs(ids...)
}

// RemoveDoorIDs removes the doors edge to Door by ids.
func (ru *RoomUpdate) RemoveDoorIDs(ids ...int) *RoomUpdate {
	if ru.removedDoors == nil {
		ru.removedDoors = make(map[int]struct{})
	}
	for i := range ids {
		ru.removedDoors[ids[i]] = struct{}{}
	}
	return ru
}

// RemoveDoors removes doors edges to Door.
func (ru *RoomUpdate) RemoveDoors(d ...*Door) *RoomUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ru.RemoveDoorIDs(ids...)
}

// RemoveSequenceIDs removes the sequences edge to Sequence by ids.
func (ru *RoomUpdate) RemoveSequenceIDs(ids ...int) *RoomUpdate {
	if ru.removedSequences == nil {
		ru.removedSequences = make(map[int]struct{})
	}
	for i := range ids {
		ru.removedSequences[ids[i]] = struct{}{}
	}
	return ru
}

// RemoveSequences removes sequences edges to Sequence.
func (ru *RoomUpdate) RemoveSequences(s ...*Sequence) *RoomUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.RemoveSequenceIDs(ids...)
}

// RemovePathNodeIDs removes the pathNodes edge to PathNode by ids.
func (ru *RoomUpdate) RemovePathNodeIDs(ids ...int) *RoomUpdate {
	if ru.removedPathNodes == nil {
		ru.removedPathNodes = make(map[int]struct{})
	}
	for i := range ids {
		ru.removedPathNodes[ids[i]] = struct{}{}
	}
	return ru
}

// RemovePathNodes removes pathNodes edges to PathNode.
func (ru *RoomUpdate) RemovePathNodes(p ...*PathNode) *RoomUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemovePathNodeIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *RoomUpdate) Save(ctx context.Context) (int, error) {
	return ru.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoomUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoomUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoomUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RoomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   room.Table,
			Columns: room.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: room.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := ru.Name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: room.FieldName,
		})
	}
	if value := ru.Description; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: room.FieldDescription,
		})
	}
	if value := ru.Floor; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: room.FieldFloor,
		})
	}
	if value := ru.addFloor; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: room.FieldFloor,
		})
	}
	if value := ru.Id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: room.FieldID,
		})
	}
	if value := ru.addId; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: room.FieldID,
		})
	}
	if nodes := ru.removedDoors; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.DoorsTable,
			Columns: []string{room.DoorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: door.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.doors; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.DoorsTable,
			Columns: []string{room.DoorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: door.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := ru.removedSequences; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.SequencesTable,
			Columns: []string{room.SequencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sequence.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.sequences; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.SequencesTable,
			Columns: []string{room.SequencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sequence.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := ru.removedPathNodes; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.PathNodesTable,
			Columns: []string{room.PathNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pathnode.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.pathNodes; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.PathNodesTable,
			Columns: []string{room.PathNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pathnode.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RoomUpdateOne is the builder for updating a single Room entity.
type RoomUpdateOne struct {
	config
	id               int
	Name             *string
	Description      *string
	Floor            *int
	addFloor         *int
	Id               *int
	addId            *int
	doors            map[int]struct{}
	sequences        map[int]struct{}
	pathNodes        map[int]struct{}
	removedDoors     map[int]struct{}
	removedSequences map[int]struct{}
	removedPathNodes map[int]struct{}
}

// SetName sets the Name field.
func (ruo *RoomUpdateOne) SetName(s string) *RoomUpdateOne {
	ruo.Name = &s
	return ruo
}

// SetDescription sets the Description field.
func (ruo *RoomUpdateOne) SetDescription(s string) *RoomUpdateOne {
	ruo.Description = &s
	return ruo
}

// SetNillableDescription sets the Description field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableDescription(s *string) *RoomUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// SetFloor sets the Floor field.
func (ruo *RoomUpdateOne) SetFloor(i int) *RoomUpdateOne {
	ruo.Floor = &i
	ruo.addFloor = nil
	return ruo
}

// SetNillableFloor sets the Floor field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableFloor(i *int) *RoomUpdateOne {
	if i != nil {
		ruo.SetFloor(*i)
	}
	return ruo
}

// AddFloor adds i to Floor.
func (ruo *RoomUpdateOne) AddFloor(i int) *RoomUpdateOne {
	if ruo.addFloor == nil {
		ruo.addFloor = &i
	} else {
		*ruo.addFloor += i
	}
	return ruo
}

// SetID sets the Id field.
func (ruo *RoomUpdateOne) SetID(i int) *RoomUpdateOne {
	ruo.Id = &i
	ruo.addId = nil
	return ruo
}

// AddID adds i to Id.
func (ruo *RoomUpdateOne) AddID(i int) *RoomUpdateOne {
	if ruo.addId == nil {
		ruo.addId = &i
	} else {
		*ruo.addId += i
	}
	return ruo
}

// AddDoorIDs adds the doors edge to Door by ids.
func (ruo *RoomUpdateOne) AddDoorIDs(ids ...int) *RoomUpdateOne {
	if ruo.doors == nil {
		ruo.doors = make(map[int]struct{})
	}
	for i := range ids {
		ruo.doors[ids[i]] = struct{}{}
	}
	return ruo
}

// AddDoors adds the doors edges to Door.
func (ruo *RoomUpdateOne) AddDoors(d ...*Door) *RoomUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ruo.AddDoorIDs(ids...)
}

// AddSequenceIDs adds the sequences edge to Sequence by ids.
func (ruo *RoomUpdateOne) AddSequenceIDs(ids ...int) *RoomUpdateOne {
	if ruo.sequences == nil {
		ruo.sequences = make(map[int]struct{})
	}
	for i := range ids {
		ruo.sequences[ids[i]] = struct{}{}
	}
	return ruo
}

// AddSequences adds the sequences edges to Sequence.
func (ruo *RoomUpdateOne) AddSequences(s ...*Sequence) *RoomUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.AddSequenceIDs(ids...)
}

// AddPathNodeIDs adds the pathNodes edge to PathNode by ids.
func (ruo *RoomUpdateOne) AddPathNodeIDs(ids ...int) *RoomUpdateOne {
	if ruo.pathNodes == nil {
		ruo.pathNodes = make(map[int]struct{})
	}
	for i := range ids {
		ruo.pathNodes[ids[i]] = struct{}{}
	}
	return ruo
}

// AddPathNodes adds the pathNodes edges to PathNode.
func (ruo *RoomUpdateOne) AddPathNodes(p ...*PathNode) *RoomUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddPathNodeIDs(ids...)
}

// RemoveDoorIDs removes the doors edge to Door by ids.
func (ruo *RoomUpdateOne) RemoveDoorIDs(ids ...int) *RoomUpdateOne {
	if ruo.removedDoors == nil {
		ruo.removedDoors = make(map[int]struct{})
	}
	for i := range ids {
		ruo.removedDoors[ids[i]] = struct{}{}
	}
	return ruo
}

// RemoveDoors removes doors edges to Door.
func (ruo *RoomUpdateOne) RemoveDoors(d ...*Door) *RoomUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ruo.RemoveDoorIDs(ids...)
}

// RemoveSequenceIDs removes the sequences edge to Sequence by ids.
func (ruo *RoomUpdateOne) RemoveSequenceIDs(ids ...int) *RoomUpdateOne {
	if ruo.removedSequences == nil {
		ruo.removedSequences = make(map[int]struct{})
	}
	for i := range ids {
		ruo.removedSequences[ids[i]] = struct{}{}
	}
	return ruo
}

// RemoveSequences removes sequences edges to Sequence.
func (ruo *RoomUpdateOne) RemoveSequences(s ...*Sequence) *RoomUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.RemoveSequenceIDs(ids...)
}

// RemovePathNodeIDs removes the pathNodes edge to PathNode by ids.
func (ruo *RoomUpdateOne) RemovePathNodeIDs(ids ...int) *RoomUpdateOne {
	if ruo.removedPathNodes == nil {
		ruo.removedPathNodes = make(map[int]struct{})
	}
	for i := range ids {
		ruo.removedPathNodes[ids[i]] = struct{}{}
	}
	return ruo
}

// RemovePathNodes removes pathNodes edges to PathNode.
func (ruo *RoomUpdateOne) RemovePathNodes(p ...*PathNode) *RoomUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemovePathNodeIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ruo *RoomUpdateOne) Save(ctx context.Context) (*Room, error) {
	return ruo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoomUpdateOne) SaveX(ctx context.Context) *Room {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *RoomUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoomUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RoomUpdateOne) sqlSave(ctx context.Context) (r *Room, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   room.Table,
			Columns: room.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  ruo.id,
				Type:   field.TypeInt,
				Column: room.FieldID,
			},
		},
	}
	if value := ruo.Name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: room.FieldName,
		})
	}
	if value := ruo.Description; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: room.FieldDescription,
		})
	}
	if value := ruo.Floor; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: room.FieldFloor,
		})
	}
	if value := ruo.addFloor; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: room.FieldFloor,
		})
	}
	if value := ruo.Id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: room.FieldID,
		})
	}
	if value := ruo.addId; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: room.FieldID,
		})
	}
	if nodes := ruo.removedDoors; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.DoorsTable,
			Columns: []string{room.DoorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: door.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.doors; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.DoorsTable,
			Columns: []string{room.DoorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: door.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := ruo.removedSequences; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.SequencesTable,
			Columns: []string{room.SequencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sequence.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.sequences; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.SequencesTable,
			Columns: []string{room.SequencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sequence.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := ruo.removedPathNodes; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.PathNodesTable,
			Columns: []string{room.PathNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pathnode.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.pathNodes; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.PathNodesTable,
			Columns: []string{room.PathNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pathnode.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Room{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}
