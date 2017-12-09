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
