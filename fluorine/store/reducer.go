package store

import (
	"github.com/DrTeePot/game/fluorine/action"
)

// data is stored as
// [reducer_id][entity_data]

// we need arity so we know what number to skip.
// so to access entity 8, we do
// [component_id][component_arity * 8]

type FloatReducer struct {
	arity   uint32
	reducer func(State_float32, action.Action_float32) State_float32
}

// TODO make a closure function so we can get rid of arity.
func (f FloatReducer) Run(state State_float32, a action.Action_float32) State_float32 {
	return f.reducer(state, a)
}

func NewFloatReducer(
	arity uint32,
	reducer func(State_float32, action.Action_float32) State_float32,
) FloatReducer {
	return FloatReducer{
		arity,
		reducer,
	}
}

func FloatNoOp(state State_float32, _ action.Action_float32) State_float32 {
	return state
}
