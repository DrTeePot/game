package ecs

// TODO this isn't threadsafe
var next uint32

func Init() {
	// our first create will incriment by one
	next := -1
}

type Basic struct {
	ID uint32
}

func Create() {
	next = next + 1
	return &Basic{
		ID: next,
	}
}

func (c Basic) ID() uint32 {
	return c.ID
}
