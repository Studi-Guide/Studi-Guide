// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"studi-guide/ent/door"
	"studi-guide/ent/pathnode"
	"studi-guide/ent/room"
	"studi-guide/ent/sequence"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RoomCreate is the builder for creating a Room entity.
type RoomCreate struct {
	config
	Name        *string
	Description *string
	Floor       *int
	doors       map[int]struct{}
	sequences   map[int]struct{}
	pathNodes   map[int]struct{}
}

// SetName sets the Name field.
func (rc *RoomCreate) SetName(s string) *RoomCreate {
	rc.Name = &s
	return rc
}

// SetDescription sets the Description field.
func (rc *RoomCreate) SetDescription(s string) *RoomCreate {
	rc.Description = &s
	return rc
}

// SetNillableDescription sets the Description field if the given value is not nil.
func (rc *RoomCreate) SetNillableDescription(s *string) *RoomCreate {
	if s != nil {
		rc.SetDescription(*s)
	}
	return rc
}

// SetFloor sets the Floor field.
func (rc *RoomCreate) SetFloor(i int) *RoomCreate {
	rc.Floor = &i
	return rc
}

// SetNillableFloor sets the Floor field if the given value is not nil.
func (rc *RoomCreate) SetNillableFloor(i *int) *RoomCreate {
	if i != nil {
		rc.SetFloor(*i)
	}
	return rc
}

// AddDoorIDs adds the doors edge to Door by ids.
func (rc *RoomCreate) AddDoorIDs(ids ...int) *RoomCreate {
	if rc.doors == nil {
		rc.doors = make(map[int]struct{})
	}
	for i := range ids {
		rc.doors[ids[i]] = struct{}{}
	}
	return rc
}

// AddDoors adds the doors edges to Door.
func (rc *RoomCreate) AddDoors(d ...*Door) *RoomCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return rc.AddDoorIDs(ids...)
}

// AddSequenceIDs adds the sequences edge to Sequence by ids.
func (rc *RoomCreate) AddSequenceIDs(ids ...int) *RoomCreate {
	if rc.sequences == nil {
		rc.sequences = make(map[int]struct{})
	}
	for i := range ids {
		rc.sequences[ids[i]] = struct{}{}
	}
	return rc
}

// AddSequences adds the sequences edges to Sequence.
func (rc *RoomCreate) AddSequences(s ...*Sequence) *RoomCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return rc.AddSequenceIDs(ids...)
}

// AddPathNodeIDs adds the pathNodes edge to PathNode by ids.
func (rc *RoomCreate) AddPathNodeIDs(ids ...int) *RoomCreate {
	if rc.pathNodes == nil {
		rc.pathNodes = make(map[int]struct{})
	}
	for i := range ids {
		rc.pathNodes[ids[i]] = struct{}{}
	}
	return rc
}

// AddPathNodes adds the pathNodes edges to PathNode.
func (rc *RoomCreate) AddPathNodes(p ...*PathNode) *RoomCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPathNodeIDs(ids...)
}

// Save creates the Room in the database.
func (rc *RoomCreate) Save(ctx context.Context) (*Room, error) {
	if rc.Name == nil {
		return nil, errors.New("ent: missing required field \"Name\"")
	}
	if rc.Description == nil {
		v := room.DefaultDescription
		rc.Description = &v
	}
	if rc.Floor == nil {
		v := room.DefaultFloor
		rc.Floor = &v
	}
	return rc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoomCreate) SaveX(ctx context.Context) *Room {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *RoomCreate) sqlSave(ctx context.Context) (*Room, error) {
	var (
		r     = &Room{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: room.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: room.FieldID,
			},
		}
	)
	if value := rc.Name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: room.FieldName,
		})
		r.Name = *value
	}
	if value := rc.Description; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: room.FieldDescription,
		})
		r.Description = *value
	}
	if value := rc.Floor; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: room.FieldFloor,
		})
		r.Floor = *value
	}
	if nodes := rc.doors; len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.sequences; len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.pathNodes; len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	r.ID = int(id)
	return r, nil
}
