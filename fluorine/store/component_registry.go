package store

type ComponentRegistry struct {
	components map[string]UniversalComponent_float32
	names      []string
}

func NewRegistry(components []UniversalComponent_float32) ComponentRegistry {
	componentsMap := make(map[string]UniversalComponent_float32)

	var names []string
	for _, component := range components {
		componentsMap[component.name] = component
		names = append(names, component.name)
	}

	return ComponentRegistry{
		components: componentsMap,
		names:      names,
	}
}

// TODO this only works for floats
func (r ComponentRegistry) Component(
	name string,
) (component UniversalComponent_float32) {
	// TODO this should panic if components[name] doesn't exist
	// TODO should find a way to wrap this to make it readonly?
	// TODO implement a custom iterator
	// 	https://stackoverflow.com/questions/35810674/in-go-is-it-possible-to-iterate-over-a-custom-type
	return r.components[name]
}
func (r ComponentRegistry) RegisteredNames() []string {
	return r.names
}
