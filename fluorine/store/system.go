package store

import ()

type updateFunc = func(Store)

type System struct {
	dependencies []string
	update       updateFunc
}

func NewSystem(dependencies []string, update updateFunc) System {
	return System{
		dependencies,
		update,
	}
}

func (s System) Dependencies() []string { return s.dependencies }
func (s System) Update(store Store)     { s.update(store) }
