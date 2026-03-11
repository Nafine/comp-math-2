package model

import (
	"comp-math-2/internal/numeric"

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
	x0
	y0
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

	singleSolutions []numeric.Solution
	systemSolution  numeric.Solution
	err             error
}

func InitialModel() Model {
	ins := make([]textinput.Model, 5)
	for i := range ins {
		ins[i] = textinput.New()
	}
	ins[a].Placeholder = "Left border (a)"
	ins[b].Placeholder = "Right border (b)"
	ins[eps].Placeholder = "Margin (eps)"
	ins[x0].Placeholder = "X0"
	ins[y0].Placeholder = "Y0"

	return Model{
		currentPhase: phaseBranch,
		equations:    numeric.GetFunctionNames(),
		systems:      numeric.GetSystemNames(),
		inputs:       ins,
		fieldErrors:  make([]string, 5),
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.handleUpdate(msg)
}

func (m Model) View() string {
	return m.renderPhase()
}
