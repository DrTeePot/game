/**
* Heavy inspiraction from engo.io/ecs
* A Entity-Component-System implementation for Go-lang
 */
package ecs

type System interface {
	Update(delta uint32) // Update takes time since last update
	Remove(b Basic)
}

type Starter interface {
	Start() // starts update loop in new thread
}

type Initializer interface {
	Init() // TODO make this take a world
}

type Prioritizer interface {
	// TODO implement sorting
	Priority() uint32 // determines order of update when present
}
