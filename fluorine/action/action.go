package action

// Use globals to define actions like an enum
type Action_float32 struct {
	component   string
	instruction uint32
	entity_id   uint32
	value       []float32
}

// TODO make this New
func Create_float32(
	component string,
	instruction uint32,
	entity_id uint32,
	value []float32,
) Action_float32 {
	return Action_float32{
		component,
		instruction,
		entity_id,
		value,
	}
}

func (a Action_float32) Component() string   { return a.component }
func (a Action_float32) Instruction() uint32 { return a.instruction }
func (a Action_float32) Entity() uint32      { return a.entity_id }
func (a Action_float32) Value() []float32    { return a.value }
