package model

import (
	"comp-math-2/internal/algo"
	"comp-math-2/internal/numeric"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) handleUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			m.currentPhase = phaseBranch
			return m, nil
		}

		switch m.currentPhase {
		case phaseBranch:
			return m.updateBranch(msg)
		case phaseChoose:
			return m.updateChoose(msg)
		case phaseSettings:
			return m.updateSettings(msg)
		case phaseSolution, phaseError:
			if msg.String() == "enter" {
				return InitialModel(), nil
			}
		}
	}
	return m, nil
}

func (m *Model) updateBranch(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if msg.String() == "up" && m.choiceIndex > 0 {
		m.choiceIndex--
	}
	if msg.String() == "down" && m.choiceIndex < 1 {
		m.choiceIndex++
	}
	if msg.String() == "enter" {
		m.isSystem = m.choiceIndex == 1
		m.currentPhase = phaseChoose
		m.choiceIndex = 0
	}
	return m, nil
}

func (m *Model) updateChoose(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	items := m.equations
	if m.isSystem {
		items = m.systems
	}
	if msg.String() == "up" && m.choiceIndex > 0 {
		m.choiceIndex--
	}
	if msg.String() == "down" && m.choiceIndex < len(items)-1 {
		m.choiceIndex++
	}
	if msg.String() == "enter" {
		m.selectedEqIndex = m.choiceIndex
		m.currentPhase = phaseSettings
		// Фокус на первое поле в зависимости от типа
		if m.isSystem {
			m.focused = x0
			m.inputs[x0].Focus()
		} else {
			m.focused = a
			m.inputs[a].Focus()
		}
	}
	return m, nil
}

func (m *Model) updateSettings(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	lastField := eps
	if !m.isSystem {
		lastField = eps
	} else {
		lastField = y0
	}

	if msg.String() == "enter" && m.focused == lastField {
		if m.validate() {
			eps, _ := strconv.ParseFloat(replaceComma(m.inputs[eps].Value()), 64)

			if m.isSystem {
				x0, _ := strconv.ParseFloat(replaceComma(m.inputs[x0].Value()), 64)
				y0, _ := strconv.ParseFloat(replaceComma(m.inputs[y0].Value()), 64)

				f1, f2 := numeric.GetSystem(m.systems[m.selectedEqIndex])

				system := numeric.NonlinearSystem{
					F1: f1,
					F2: f2,
					StartCoordinates: numeric.Coordinates{
						X: x0,
						Y: y0,
					},
					Eps: eps,
				}

				solution, err := algo.SolveSystem(system)
				if err != nil {
					m.err = err
					m.currentPhase = phaseError
					return m, nil
				}
				m.systemSolution = solution
			} else {
				a, _ := strconv.ParseFloat(replaceComma(m.inputs[a].Value()), 64)
				b, _ := strconv.ParseFloat(replaceComma(m.inputs[b].Value()), 64)

				eq := numeric.NonlinearEquation{
					F:   numeric.GetEquation(m.equations[m.selectedEqIndex]),
					A:   a,
					B:   b,
					Eps: eps,
				}
				solutions, err := algo.SolveAllSingle(eq)
				if err != nil {
					m.err = err
					m.currentPhase = phaseError
					return m, nil
				}
				m.singleSolutions = solutions
			}
			m.currentPhase = phaseSolution
		}
		return m, nil
	}

	if msg.String() == "tab" || msg.String() == "down" {
		m.inputs[m.focused].Blur()
		m.focused++
		if m.isSystem {
			if m.focused > y0 {
				m.focused = eps
			}
		} else {
			m.focused = (m.focused) % 3
		}
		m.inputs[m.focused].Focus()
	} else if msg.String() == "up" {
		m.inputs[m.focused].Blur()
		m.focused--
		if m.isSystem {
			if m.focused < eps {
				m.focused = y0
			}
		} else {
			if m.focused < a {
				m.focused = eps
			}
		}
		m.inputs[m.focused].Focus()
	}

	var cmd tea.Cmd
	m.inputs[m.focused], cmd = m.inputs[m.focused].Update(msg)
	return m, cmd
}

func replaceComma(s string) string {
	return strings.ReplaceAll(s, ",", ".")
}
