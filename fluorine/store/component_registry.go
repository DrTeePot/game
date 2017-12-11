package store

type ComponentRegistry struct {
	components map[string]UniversalComponent_float32
}

func NewRegistry(components []UniversalComponent_float32) ComponentRegistry {
	componentsMap := make(map[string]UniversalComponent_float32)

	for _, component := range components {
		componentsMap[component.name] = component
	}

	return ComponentRegistry{
		components: componentsMap,
	}
}

// TODO this only works for floats
func (r ComponentRegistry) Component(
	name string,
) (component map[uint32][]float32) {
	// TODO this should panic if components[name] doesn't exist
	// TODO should find a way to wrap this to make it readonly?
	// TODO implement a custom iterator
	// 	https://stackoverflow.com/questions/35810674/in-go-is-it-possible-to-iterate-over-a-custom-type
	return r.components[name].state.data
}
