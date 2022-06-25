package main

/**
 * @author       weimenghua
 * @time         2024-07-10 10:03
 * @description  表格
 */

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "An example CLI application with a styled table",
	Run: func(cmd *cobra.Command, args []string) {
		columns := []table.Column{
			{Title: "ID", Width: 5},
			{Title: "Name", Width: 20},
			{Title: "Age", Width: 5},
		}

		rows := []table.Row{
			{"1", "Alice", "30"},
			{"2", "Bob", "25"},
			{"3", "Charlie", "35"},
		}

		styles := table.DefaultStyles()
		styles.Header = styles.Header.BorderLeft(true).BorderRight(true).BorderTop(true).BorderBottom(true)
		styles.Cell = styles.Cell.BorderLeft(true).BorderRight(true).BorderTop(true).BorderBottom(true)

		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithFocused(true),
			table.WithStyles(styles),
		)

		p := tea.NewProgram(model{t: t})

		if err, _ := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error running program: %v", err)
			os.Exit(1)
		}
	},
}

type model struct {
	t table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.t.Focused() {
				m.t.Blur()
			} else {
				m.t.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.t, cmd = m.t.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.t.View()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing root command: %v", err)
		os.Exit(1)
	}
}
