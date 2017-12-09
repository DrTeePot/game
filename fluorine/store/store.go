package store

import (
	"github.com/DrTeePot/game/fluorine/action"
)

type Store struct {
	dispatchFloat chan action.Action_float32
	inputStream   chan action.Action_float32

	registry ComponentRegistry
}

func CreateStore(
	registry ComponentRegistry,
) Store {

	return Store{
		registry: registry,

		dispatchFloat: make(chan action.Action_float32),
		inputStream:   make(chan action.Action_float32),
	}
}

func (s *Store) Update() {
	// grab input events every frame
	select {
	case in := <-s.inputStream:
		id := in.Component()
		component := s.registry.components[id]
		component.update(in)
	default:
		// no input events this update, carry on
	}

	// grab other component events
	select {
	case a := <-s.dispatchFloat:
		id := a.Component()
		component := s.registry.components[id]
		component.update(a)
	default:
		// no actions this update, carry on
	}
}

func (s Store) DispatchFloat(action action.Action_float32) {
	s.dispatchFloat <- action
}

func (s Store) DispatchInput(action action.Action_float32) {
	s.inputStream <- action
}

func (s Store) Close() {
	close(s.dispatchFloat)
	close(s.inputStream)
	for _, c := range s.registry.components {
		c.Close()
	}
}
