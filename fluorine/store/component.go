package store

import (
	"github.com/DrTeePot/game/fluorine/action"
)

/*
components have a state and a reducer. the store is made up of various
components.

components can be chunked or universal. the state in a universal component
is always loaded in memory, whereas chunked components will only
load data in a certain range of chunks
*/

type UniversalComponent_float32 struct {
	name    string
	state   State_float32
	reducer FloatReducer // TODO FloatReducer should probably just be a
	// function  alias
	subscribe chan State_float32
}

func NewUniversalComponent_float32(
	name string,
	reducer FloatReducer,
) UniversalComponent_float32 {
	state := NewState_float32()
	return UniversalComponent_float32{
		name:      name,
		state:     state,
		reducer:   reducer,
		subscribe: make(chan State_float32),
	}
}

func (u UniversalComponent_float32) State() map[uint32][]float32 {
	return u.state.data
}

func (u UniversalComponent_float32) Name() string { return u.name }

func (u UniversalComponent_float32) Subscribe(watcher func(State_float32)) {
	outbound := u.subscribe

	go func(outbound chan State_float32, f func(State_float32)) {
		for updated := range outbound {
			f(updated)
		}
	}(outbound, watcher)
}

func (u UniversalComponent_float32) Close() {
	close(u.subscribe)
}

func (u *UniversalComponent_float32) update(a action.Action_float32) {
	data := u.state
	u.state = u.reducer.Run(data, a)
	select {
	case u.subscribe <- u.state:
	default:
	}
}
