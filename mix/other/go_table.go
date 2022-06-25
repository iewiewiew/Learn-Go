package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

/**
 * @author       weimenghua
 * @time         2024-07-11 13:53
 * @description  表格
 */

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func initialModel() model {
	columns := []table.Column{
		{Title: "Column 1", Width: 20},
		{Title: "Column 2", Width: 20},
		{Title: "Column 3", Width: 20},
		{Title: "Column 4", Width: 20},
		{Title: "Column 5", Width: 20},
	}

	rows := []table.Row{
		{"Data 1", "Data 2", "Data 3", "Data 4", "Data 5"},
		{"Data 6", "Data 7", "Data 8", "Data 9", "Data 10"},
		{"Data 11", "Data 12", "Data 13", "Data 14", "Data 15"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(5),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		/*BorderStyle(lipgloss.Border{
			Top:         "─",
			Bottom:      "─",
			Left:        "│",
			Right:       "│",
			TopLeft:     "┌",
			TopRight:    "┐",
			BottomLeft:  "└",
			BottomRight: "┘",
		}).
		*/
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)

	return model{table: t}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"\n%s\n\n%s",
		m.table.View(),
		"Press q to quit.",
	)

	//return baseStyle.Render(m.table.View()) + "\n"
}

func main() {
	// 创建一个新的 Bubble Tea 程序实例，并传入初始模型 initialModel()
	if err, _ := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
