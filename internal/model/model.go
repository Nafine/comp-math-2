package model

import (
	"comp-math-2/internal/algo"
	"comp-math-2/internal/numeric"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	activeStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	errorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	headerStyle = lipgloss.NewStyle().Background(lipgloss.Color("62")).Foreground(lipgloss.Color("255")).Padding(0, 1).Bold(true)
	resultStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2).BorderForeground(lipgloss.Color("42"))
)

type phase int

const (
	phaseBranch phase = iota
	phaseChoose
	phaseSettings
	phaseSolution
	phaseError
)

const (
	a = iota
	b
	eps
)

type Model struct {
	currentPhase phase
	choiceIndex  int
	isSystem     bool

	equations       []string
	systems         []string
	selectedEqIndex int

	inputs      []textinput.Model
	fieldErrors []string
	focused     int

	solution numeric.Solution
	err      error
}

func InitialModel() Model {
	ins := make([]textinput.Model, 3)
	for i := range ins {
		ins[i] = textinput.New()
	}
	ins[a].Placeholder = "Left border (a)"
	ins[b].Placeholder = "Right border (b)"
	ins[eps].Placeholder = "Margin (eps)"

	return Model{
		currentPhase: phaseBranch,
		equations:    numeric.GetSingleEquations(),
		systems:      numeric.GetSystems(),
		inputs:       ins,
		fieldErrors:  make([]string, 3),
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		m.inputs[0].Focus()
	}
	return m, nil
}

func (m *Model) updateSettings(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if msg.String() == "enter" && m.focused == len(m.inputs)-1 {
		if m.validate() {
			a, _ := strconv.ParseFloat(ReplaceComma(m.inputs[a].Value()), 64)
			b, _ := strconv.ParseFloat(ReplaceComma(m.inputs[b].Value()), 64)
			eps, _ := strconv.ParseFloat(ReplaceComma(m.inputs[eps].Value()), 64)

			eq := numeric.NonlinearEquation{
				F:   numeric.GetEquation(m.equations[m.selectedEqIndex]),
				A:   a,
				B:   b,
				Eps: eps,
			}

			solution, _ := algo.SolveNewton(eq)
			m.solution = solution
			m.currentPhase = phaseSolution
		}
		return m, nil
	}

	if msg.String() == "tab" || msg.String() == "down" || msg.String() == "enter" {
		m.inputs[m.focused].Blur()
		m.focused = (m.focused + 1) % len(m.inputs)
		m.inputs[m.focused].Focus()
	} else if msg.String() == "up" {
		m.inputs[m.focused].Blur()
		m.focused--
		if m.focused < 0 {
			m.focused = len(m.inputs) - 1
		}
		m.inputs[m.focused].Focus()
	}

	var cmd tea.Cmd
	m.inputs[m.focused], cmd = m.inputs[m.focused].Update(msg)
	return m, cmd
}

func (m *Model) validate() bool {
	isValid := true
	for i := range m.fieldErrors {
		m.fieldErrors[i] = ""
	}

	_, errA := strconv.ParseFloat(m.inputs[0].Value(), 64)
	if errA != nil {
		m.fieldErrors[0] = "Invalid number"
		isValid = false
	}

	_, errB := strconv.ParseFloat(m.inputs[1].Value(), 64)
	if errB != nil {
		m.fieldErrors[1] = "Invalid number"
		isValid = false
	}

	eps, errE := strconv.ParseFloat(m.inputs[2].Value(), 64)
	if errE != nil || eps <= 0 {
		m.fieldErrors[2] = "Eps must be > 0"
		isValid = false
	}

	return isValid
}

func (m Model) View() string {
	var s strings.Builder
	s.WriteString(headerStyle.Render(" COMP-MATH Lab #2 ") + "\n\n")

	switch m.currentPhase {
	case phaseBranch:
		s.WriteString("Choose type:\n")
		opts := []string{"Single equation", "System"}
		for i, opt := range opts {
			cursor := "  "
			if m.choiceIndex == i {
				cursor = activeStyle.Render("> ")
			}
			s.WriteString(fmt.Sprintf("%s%s\n", cursor, opt))
		}

	case phaseChoose:
		s.WriteString("Choose equation:\n")
		items := m.equations
		if m.isSystem {
			items = m.systems
		}
		for i, item := range items {
			cursor := "  "
			if m.choiceIndex == i {
				cursor = activeStyle.Render("> ")
			}
			s.WriteString(fmt.Sprintf("%s%s\n", cursor, item))
		}

	case phaseSettings:
		s.WriteString("Parameters:\n\n")
		for i := range m.inputs {
			s.WriteString(m.inputs[i].View() + "\n")
			if m.fieldErrors[i] != "" {
				s.WriteString(errorStyle.Render("  └─ "+m.fieldErrors[i]) + "\n")
			} else {
				s.WriteString("\n")
			}
		}
		s.WriteString("\n[Enter] - Submit")

	case phaseSolution:
		s.WriteString(activeStyle.Render("Computing done:") + "\n")
		output := fmt.Sprintf("Solution: %s\n", m.solution)
		s.WriteString(resultStyle.Render(output) + "\n\n")
		s.WriteString("Press [Enter] for restart")

	case phaseError:
		s.WriteString(headerStyle.Copy().Background(lipgloss.Color("9")).Render(" FILE ERROR ") + "\n\n")
		s.WriteString(m.err.Error() + "\n\n")
		s.WriteString("Check file path and press [Enter]")
	}

	return s.String()
}

func ReplaceComma(s string) string {
	return strings.ReplaceAll(s, ",", ".")
}
