package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/enescakir/emoji"
)

type Tool struct {
	Name        string
	Desc        string
	Installed   bool
	Selected    bool
}

func (t Tool) FilterValue() string { return t.Name }
func (t Tool) Title() string       { 
	icon := emoji.Package.String()
	if t.Installed {
		icon = emoji.CheckMark.String()
	}
	return fmt.Sprintf("%s %s", icon, t.Name) 
}
func (t Tool) Description() string { return t.Desc }

type InstallerModel struct {
	list     list.Model
	tools    []Tool
	selected map[int]bool
	width    int
	height   int
	quitting bool
}

func NewInstallerModel(tools []Tool) InstallerModel {
	items := make([]list.Item, len(tools))
	for i, tool := range tools {
		items[i] = tool
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = fmt.Sprintf("%s Select Tools to Install", emoji.Gear)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("36")).
		Bold(true).
		Padding(0, 1)

	return InstallerModel{
		list:     l,
		tools:    tools,
		selected: make(map[int]bool),
	}
}

func (m InstallerModel) Init() tea.Cmd {
	return nil
}

func (m InstallerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 4)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case " ":
			// Toggle selection
			idx := m.list.Index()
			m.selected[idx] = !m.selected[idx]
			return m, nil

		case "a":
			// Select all
			for i := range m.tools {
				m.selected[i] = true
			}
			return m, nil

		case "n":
			// Select none
			m.selected = make(map[int]bool)
			return m, nil

		case "enter":
			// Start installation
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m InstallerModel) View() string {
	if m.quitting {
		return ""
	}

	// Header
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("36")).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("36")).
		Padding(1, 2).
		MarginBottom(1)

	header := headerStyle.Render("🚀 Dotfiles Installer")

	// List with custom rendering for selection
	listView := m.renderCustomList()

	// Instructions
	instructionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Italic(true)

	instructions := instructionStyle.Render(
		"• space: toggle selection  • a: select all  • n: select none  • enter: install  • q: quit",
	)

	// Selected count
	selectedCount := 0
	for _, selected := range m.selected {
		if selected {
			selectedCount++
		}
	}

	statusStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("220")).
		Bold(true)

	status := statusStyle.Render(fmt.Sprintf("Selected: %d tools", selectedCount))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		listView,
		"",
		status,
		instructions,
	)
}

func (m InstallerModel) renderCustomList() string {
	var b strings.Builder
	
	start, end := m.list.Paginator.GetSliceBounds(len(m.tools))
	for i, tool := range m.tools[start:end] {
		realIndex := start + i
		isSelected := realIndex == m.list.Index()
		isChecked := m.selected[realIndex]
		
		// Style for current item
		var style lipgloss.Style
		if isSelected {
			style = lipgloss.NewStyle().
				Foreground(lipgloss.Color("36")).
				Bold(true).
				Background(lipgloss.Color("240"))
		} else {
			style = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
		}
		
		// Checkbox
		checkbox := "☐"
		if isChecked {
			checkbox = emoji.CheckBoxWithCheck.String()
		}
		
		// Status icon
		statusIcon := emoji.Package.String()
		if tool.Installed {
			statusIcon = emoji.CheckMark.String()
		}
		
		// Cursor
		cursor := " "
		if isSelected {
			cursor = "❯"
		}
		
		line := fmt.Sprintf("  %s %s %s %s", 
			cursor,
			checkbox, 
			statusIcon, 
			tool.Name,
		)
		
		if isSelected {
			// Add description for selected item
			line += fmt.Sprintf("\n    %s", 
				lipgloss.NewStyle().
					Foreground(lipgloss.Color("240")).
					Italic(true).
					Render(tool.Desc),
			)
		}
		
		b.WriteString(style.Render(line))
		b.WriteString("\n")
	}
	
	return b.String()
}

func (m InstallerModel) GetSelectedTools() []Tool {
	var selected []Tool
	for i, tool := range m.tools {
		if m.selected[i] {
			selected = append(selected, tool)
		}
	}
	return selected
}

func RunToolSelector(tools []Tool) ([]Tool, error) {
	m := NewInstallerModel(tools)
	p := tea.NewProgram(m, tea.WithAltScreen())
	
	finalModel, err := p.Run()
	if err != nil {
		return nil, err
	}
	
	if finalModel, ok := finalModel.(InstallerModel); ok {
		return finalModel.GetSelectedTools(), nil
	}
	
	return []Tool{}, nil
}