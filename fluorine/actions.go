package fluorine

type StringAction struct {
	instruction  string
	component_id uint32
	entity_id    uint32
	value        []string
}

type FloatAction struct {
	instruction  string
	component_id uint32
	entity_id    uint32
	value        []float32
}

// TODO should these return functions?
func CreateStringAction(
	instruction string,
	component_id uint32,
	entity_id uint32,
	value []string,
) StringAction {
	// TODO do data validations
	return StringAction{instruction, component_id, entity_id, value}
}

func CreateFloatAction(
	instruction string,
	component_id uint32,
	entity_id uint32,
	value []float32,
) FloatAction {
	// TODO do data validations
	return FloatAction{instruction, component_id, entity_id, value}
}

func CreateInputAction(
	instruction string,
	component_id uint32,
	entity_id uint32,
	value []float32,
) FloatAction {
	// TODO do data validations
	return FloatAction{instruction, component_id, entity_id, value}
}
