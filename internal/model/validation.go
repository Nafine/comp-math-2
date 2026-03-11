package model

import "strconv"

func (m *Model) validate() bool {
	valid := true

	// Валидация eps (общее для всех)
	epsVal, err := strconv.ParseFloat(replaceComma(m.inputs[eps].Value()), 64)
	if err != nil || epsVal <= 0 {
		m.fieldErrors[eps] = "must be positive number"
		valid = false
	} else {
		m.fieldErrors[eps] = ""
	}

	if m.isSystem {
		// Валидация x0
		_, err := strconv.ParseFloat(replaceComma(m.inputs[x0].Value()), 64)
		if err != nil {
			m.fieldErrors[x0] = "invalid number"
			valid = false
		} else {
			m.fieldErrors[x0] = ""
		}

		// Валидация y0
		_, err = strconv.ParseFloat(replaceComma(m.inputs[y0].Value()), 64)
		if err != nil {
			m.fieldErrors[y0] = "invalid number"
			valid = false
		} else {
			m.fieldErrors[y0] = ""
		}
	} else {
		// Валидация a
		aVal, err := strconv.ParseFloat(replaceComma(m.inputs[a].Value()), 64)
		if err != nil {
			m.fieldErrors[a] = "invalid number"
			valid = false
		} else {
			m.fieldErrors[a] = ""
		}

		// Валидация b
		bVal, err := strconv.ParseFloat(replaceComma(m.inputs[b].Value()), 64)
		if err != nil {
			m.fieldErrors[b] = "invalid number"
			valid = false
		} else {
			m.fieldErrors[b] = ""
		}

		// Проверка a < b
		if m.fieldErrors[a] == "" && m.fieldErrors[b] == "" && aVal >= bVal {
			m.fieldErrors[a] = "a must be < b"
			m.fieldErrors[b] = "a must be < b"
			valid = false
		}
	}

	return valid
}
