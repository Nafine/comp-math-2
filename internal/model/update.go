package model

func (m *Model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *Model) prevInput() {
	m.focused--

	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
