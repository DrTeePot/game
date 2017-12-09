package store

// TODO reimplement as a real immutable data structure
type State_float32 struct {
	data map[uint32][]float32
}

func NewState_float32() State_float32 {
	return State_float32{
		data: make(map[uint32][]float32),
	}
}

func (s State_float32) GetEntity(id uint32) []float32 {
	return s.data[id]
}

func (s State_float32) Assign(id uint32, value []float32) State_float32 {
	s.data[id] = value
	return s
}
