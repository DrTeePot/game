package fluorine

type store struct {
	floatComponents []FloatComponent
	floatData       [][]float32
	dispatchFloat   chan FloatAction
	outboundFloat   []chan []float32

	stringComponents []StringComponent
	stringData       [][]string
	dispatchString   chan StringAction
	outboundString   []chan []string

	inputStream chan FloatAction
}

func CreateStore(
	floatComponents []FloatComponent,
	stringComponents []StringComponent,
) store {

	floatLength := len(floatComponents)
	stringLength := len(stringComponents)

	return store{
		floatComponents: floatComponents,
		floatData:       [][]float32{},
		dispatchFloat:   make(chan FloatAction),
		outboundFloat:   make([]chan []float32, floatLength),

		stringComponents: stringComponents,
		stringData:       [][]string{},
		dispatchString:   make(chan StringAction),
		outboundString:   make([]chan []string, stringLength),

		inputStream: make(chan FloatAction),
	}
}

func (s *store) update() {
	// grab input events every frame
	select {
	case in := <-s.inputStream:
		id := in.component_id
		component := s.floatComponents[id]
		data := s.floatData[id]

		s.floatData[id] = component.reducer(data, in)
		s.outboundFloat[id] <- s.floatData[id]
	default:
		// no input events this update, carry on
	}

	// grab other component events
	select {
	case a := <-s.dispatchString:
		id := a.component_id
		component := s.stringComponents[id]
		data := s.stringData[id]

		s.stringData[id] = component.reducer(data, a)
		s.outboundString[id] <- s.stringData[id]

	case a := <-s.dispatchFloat:
		id := a.component_id
		component := s.floatComponents[id]
		data := s.floatData[id]

		s.floatData[id] = component.reducer(data, a)
		s.outboundFloat[id] <- s.floatData[id]
	default:
		// no actions this update, carry on
	}
}

func (s store) DispatchString(action StringAction) {
	s.dispatchString <- action
}

func (s store) DispatchFloat(action FloatAction) {
	s.dispatchFloat <- action
}

func (s store) DispatchInput(action FloatAction) {
	s.inputStream <- action
}

func (s store) SubscribeString(
	component_id uint32,
	watcher func([]string),
) {
	outbound := s.outboundString[component_id]

	go func(outbound chan []string, f func([]string)) {
		for updated := range outbound {
			f(updated)
		}
	}(outbound, watcher)
}

func (s store) SubscribeFloat(
	component_id uint32,
	watcher func([]float32),
) {
	outbound := s.outboundFloat[component_id]

	go func(outbound chan []float32, f func([]float32)) {
		for updated := range outbound {
			f(updated)
		}
	}(outbound, watcher)

}

func (s store) Close() {
	close(s.dispatchFloat)
	close(s.dispatchString)
	close(s.inputStream)

	for i := range s.floatComponents {
		close(s.outboundFloat[i])
	}
	for i := range s.stringComponents {
		close(s.outboundString[i])
	}
}
