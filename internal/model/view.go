package model

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderPhase() string {
	var b strings.Builder

	b.WriteString(headerStyle.Render(" COMP-MATH Lab #2 ") + "\n\n")

	switch m.currentPhase {
	case phaseBranch:
		m.viewBranch(&b)
	case phaseChoose:
		m.viewChoose(&b)
	case phaseSettings:
		m.viewSettings(&b)
	case phaseSolution:
		m.viewSolution(&b)
	case phaseError:
		m.viewError(&b)
	}

	return b.String()
}

func (m Model) viewBranch(b *strings.Builder) {
	b.WriteString("Choose type:\n")
	opts := []string{"Single equation", "System"}
	for i, opt := range opts {
		cursor := "  "
		if m.choiceIndex == i {
			cursor = activeStyle.Render("> ")
		}
		b.WriteString(fmt.Sprintf("%s%s\n", cursor, opt))
	}
}

func (m Model) viewChoose(b *strings.Builder) {
	b.WriteString("Choose equation:\n")
	items := m.equations
	if m.isSystem {
		items = m.systems
	}
	for i, item := range items {
		cursor := "  "
		if m.choiceIndex == i {
			cursor = activeStyle.Render("> ")
		}
		b.WriteString(fmt.Sprintf("%s%s\n", cursor, item))
	}
}

func (m Model) viewSettings(b *strings.Builder) {
	b.WriteString("Parameters:\n\n")

	if m.isSystem {
		m.renderSystemSettings(b)
	} else {
		m.renderSingleSettings(b)
	}

	b.WriteString("\n[Enter] - Submit")
}

func (m Model) renderSingleSettings(b *strings.Builder) {
	for i := a; i <= eps; i++ {
		b.WriteString(m.inputs[i].View() + "\n")
		if m.fieldErrors[i] != "" {
			b.WriteString(errorStyle.Render("  └─ "+m.fieldErrors[i]) + "\n")
		} else {
			b.WriteString("\n")
		}
	}
}

func (m Model) renderSystemSettings(b *strings.Builder) {
	for i := eps; i <= y0; i++ {
		b.WriteString(m.inputs[i].View() + "\n")
		if m.fieldErrors[i] != "" {
			b.WriteString(errorStyle.Render("  └─ "+m.fieldErrors[i]) + "\n")
		} else {
			b.WriteString("\n")
		}
	}
}

func (m Model) viewSolution(b *strings.Builder) {
	b.WriteString(activeStyle.Render("Computing done:") + "\n")

	if m.isSystem {
		m.renderSystemSolution(b)
	} else {
		m.renderSingleSolutions(b)
	}
	b.WriteString("Press [Enter] for restart")
}

func (m Model) renderSingleSolutions(b *strings.Builder) {
	for _, solution := range m.singleSolutions {
		output := fmt.Sprintf("Solution: %s\n", solution)
		b.WriteString(output)
	}
}

func (m Model) renderSystemSolution(b *strings.Builder) {
	output := fmt.Sprintf("Solution: %s\n", m.systemSolution)
	b.WriteString(resultStyle.Render(output) + "\n\n")
}

func (m Model) viewError(b *strings.Builder) {
	b.WriteString(headerStyle.Background(lipgloss.Color("9")).Render(" FILE ERROR ") + "\n\n")
	b.WriteString(m.err.Error() + "\n\n")
	b.WriteString("Check file path and press [Enter]")
}
