package fluorine

import (
	"time"

	"github.com/go-gl/glfw/v3.2/glfw"

	"github.com/DrTeePot/game/fluorine/store"
)

// type nothing = struct{} // uses 0 space

/*
TODO Fluorine

There are too many structures here, components can be added to fluorine
directly. Current steps:

1. create compnents
2. create systems
3. create registry with components
4. create store with registry
5. create fluorine with system and store

Recognizing that we need to add systems and components to fluorine, it may
make sense to keep the registry. However, the store should be merged into
fluorine.

1. create components
2. create system
3. create registry with components
4. create fluorine wiht system and registry

Ideally we cut out the registry step as well.

*/

const (
	SLEEP_TIME = 1
)

type fluorine struct {
	systems    []store.System
	store      store.Store
	renderTick *time.Ticker
	close      chan (struct{})
}

func New(
	y []store.System,
	s store.Store,
) fluorine {
	dependencies := make(map[string]struct{})

	for _, componentName := range s.RegisteredComponents() {
		dependencies[componentName] = struct{}{}
	}

	for _, system := range y {
		for _, componentName := range system.Dependencies() {
			_, ok := dependencies[componentName]
			if !ok {
				panic("Dependencies of systems are not met")
			}
		}
	}

	return fluorine{
		systems:    y,
		store:      s,
		renderTick: time.NewTicker(time.Duration(time.Second / 60)),
	}
}

func (f fluorine) Start() {
	for true {
		// update state according to dispatched actions
		f.store.Update()

		// TODO this loop should not call systems. Systems should be spun up
		// in their own threads, and this should just pass the new store
		// to them
		for _, system := range f.systems {
			system.Update(f.store)
		}

		// grab input events
		// TODO should be part of an input system
		glfw.PollEvents()

		select {
		// Exit
		case <-f.close:
			break
		default:
			time.Sleep(SLEEP_TIME)
		}
	}

	// **** CLEANUP **** //
	f.store.Close()
}

func (f fluorine) Done() {
	f.close <- struct{}{}
}
