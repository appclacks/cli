package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/appclacks/cli/client"
	apitypes "github.com/appclacks/go-types"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 2)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Left).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

const (
	defaultTimeout = 10 * time.Second
)

type model struct {
	client                      *client.Client
	healthcheckTable            table.Model
	err                         error
	selectedHealthcheckID       string
	detailsTabs                 []string
	selectedHealthcheck         apitypes.Healthcheck
	selectedHealthcheckFailures []apitypes.HealthcheckResult
	detailsActiveTab            int
}

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func (m *model) Init() tea.Cmd {
	columns := []table.Column{
		{Title: "Name", Width: 20},
		{Title: "Type", Width: 6},
		{Title: "ID", Width: 36},
		{Title: "Failures", Width: 15},
	}

	rows := []table.Row{}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m.healthcheckTable = t

	tabs := []string{"Definition                              ", "Last failure details                              "}
	m.detailsTabs = tabs
	return m.listHealthchecks
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case listHealthcheckMsg:
		rows := []table.Row{}
		for _, c := range msg.healthchecks {
			rows = append(rows, table.Row{c.Name, c.Type, c.ID, fmt.Sprintf("%d", msg.failures[c.ID])})
			m.healthcheckTable.SetRows(rows)
		}
		return m, nil
	case healthcheckDetailsMsg:
		m.selectedHealthcheck = msg.healthcheck
		m.selectedHealthcheckFailures = msg.failures
		if m.detailsActiveTab == 0 {
			m.detailsActiveTab = 1
		}
		return m, nil
	case errMsg:
		m.err = msg
		return m, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.selectedHealthcheckID = ""
			m.detailsActiveTab = 0
			return m, nil
		case "r":
			if m.selectedHealthcheckID != "" {
				return m, m.getHealthchecksDetails
			}
			return m, nil
		case "enter":
			m.selectedHealthcheckID = m.healthcheckTable.SelectedRow()[2]
			return m, m.getHealthchecksDetails
		case "ctrl+c", "q":
			return m, tea.Quit
		case "right", "n":
			if m.selectedHealthcheckID != "" {
				m.detailsActiveTab = min(m.detailsActiveTab+1, len(m.detailsTabs))
			}
			return m, nil
		case "left", "p":
			if m.selectedHealthcheckID != "" {
				m.detailsActiveTab = max(m.detailsActiveTab-1, 1)
			}
			return m, nil
		default:
			m.healthcheckTable, cmd = m.healthcheckTable.Update(msg)

		}
	}
	return m, cmd
}

func (m *model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nError: %v\n\n", m.err)
	}
	if m.detailsActiveTab >= 1 {
		activeTab := m.detailsActiveTab - 1
		doc := strings.Builder{}

		var renderedTabs []string

		for i, t := range m.detailsTabs {
			var style lipgloss.Style
			isFirst, isLast, isActive := i == 0, i == len(m.detailsTabs)-1, i == activeTab
			if isActive {
				style = activeTabStyle.Copy()
			} else {
				style = inactiveTabStyle.Copy()
			}
			border, _, _, _, _ := style.GetBorder()
			if isFirst && isActive {
				border.BottomLeft = "│"
			} else if isFirst && !isActive {
				border.BottomLeft = "├"
			} else if isLast && isActive {
				border.BottomRight = "│"
			} else if isLast && !isActive {
				border.BottomRight = "┤"
			}
			style = style.Border(border)
			renderedTabs = append(renderedTabs, style.Render(t))
		}

		row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
		doc.WriteString(row)
		doc.WriteString("\n")
		if activeTab == 0 {
			doc.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render(detailString(m.selectedHealthcheck, m.selectedHealthcheckFailures)))

		}
		if activeTab == 1 {
			if len(m.selectedHealthcheckFailures) == 0 {
				doc.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render("No failures"))
			} else {
				doc.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render(detailsFailures(m.selectedHealthcheckFailures)))
			}

		}
		return docStyle.Render(doc.String())
	}
	return baseStyle.Render(m.healthcheckTable.View()) + "\n"
}

func initialModel(client *client.Client) (model, error) {

	return model{
		client: client,
	}, nil
}

func Launch(client *client.Client) error {
	initialModel, err := initialModel(client)
	if err != nil {
		return err
	}
	p := tea.NewProgram(&initialModel)
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
