package fluorine

// data is stored as
// [component_id][entity_data]
//
// we need arity so we know what number to skip.
// so to access entity 8, we do
// [component_id][component_arity * 8]

// this file is why we need generics *eyeroll*

type StringComponent struct {
	name    string
	arity   uint32
	reducer func([]string, StringAction) []string
}

func NewStringComponent(
	name string,
	arity uint32,
	reducer func([]string, StringAction) []string,
) StringComponent {
	return StringComponent{
		name,
		arity,
		reducer,
	}
}

type FloatComponent struct {
	name    string
	arity   uint32
	reducer func([]float32, FloatAction) []float32
}

func NewFloatComponent(
	name string,
	arity uint32,
	reducer func([]float32, FloatAction) []float32,
) FloatComponent {
	return FloatComponent{
		name,
		arity,
		reducer,
	}
}

func FloatNoOp(state []float32, _ FloatAction) []float32 {
	return state
}

func StringNoOp(state []string, _ StringAction) []string {
	return state
}
