package prompts

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	tableStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00FFFF"))

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FFFF"))

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("10"))

	pendingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("11"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("9"))

	inProgressStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("14"))

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Italic(true).
			Foreground(lipgloss.Color("#00FFFF")).
			MarginLeft(2).
			MarginTop(1).
			MarginBottom(1)
)

type status int

const (
	statusPending status = iota
	statusInProgress
	statusSuccess
	statusFailed
)

type DashboardItem struct {
	Name        string
	Description string
	Status      status
	Error       error
}

type dashboardUpdateMsg struct {
	index  int
	status status
	err    error
}

type completedMsg struct{}

type DashboardModel struct {
	items       []DashboardItem
	table       table.Model
	spinner     spinner.Model
	activeIndex int
	width       int
	height      int
	quitting    bool
}

func NewDashboard(items []DashboardItem) DashboardModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))

	// Create table columns
	columns := []table.Column{
		{Title: "Component", Width: 15},
		{Title: "Description", Width: 40},
		{Title: "Status", Width: 15},
	}

	// Create table rows
	rows := []table.Row{}
	for _, item := range items {
		statusText := pendingStyle.Render("Pending")
		rows = append(rows, table.Row{item.Name, item.Description, statusText})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(len(items)),
	)

	// Style the table
	t.SetStyles(table.Styles{
		Header:   headerStyle,
		Selected: lipgloss.NewStyle(),
		Cell:     lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1),
	})

	return DashboardModel{
		items:       items,
		table:       t,
		spinner:     s,
		activeIndex: 0,
	}
}

func (m DashboardModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.startInstallations,
	)
}

func (m DashboardModel) startInstallations() tea.Msg {
	// Start the first installation
	return dashboardUpdateMsg{
		index:  0,
		status: statusInProgress,
	}
}

func (m DashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quitting = true
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.table.SetWidth(msg.Width - 10)
		return m, nil

	case spinner.TickMsg:
		var spinnerCmd tea.Cmd
		m.spinner, spinnerCmd = m.spinner.Update(msg)
		return m, spinnerCmd

	case dashboardUpdateMsg:
		// Update the status of the current item
		m.items[msg.index].Status = msg.status
		m.items[msg.index].Error = msg.err

		// Update the table row
		rows := m.table.Rows()
		switch msg.status {
		case statusInProgress:
			rows[msg.index][2] = inProgressStyle.Render("Installing...")
		case statusSuccess:
			rows[msg.index][2] = successStyle.Render("✓ Success")
		case statusFailed:
			rows[msg.index][2] = errorStyle.Render(fmt.Sprintf("✗ Failed: %v", msg.err))
		}
		m.table.SetRows(rows)

		// If current item succeeded, move to the next item
		if msg.status == statusSuccess {
			if msg.index+1 < len(m.items) {
				m.activeIndex = msg.index + 1
				return m, tea.Sequence(
					tea.Tick(500*time.Millisecond, func(time.Time) tea.Msg {
						return dashboardUpdateMsg{
							index:  m.activeIndex,
							status: statusInProgress,
						}
					}),
				)
			} else {
				// All done
				return m, func() tea.Msg { return completedMsg{} }
			}
		}

		// If current item failed, stop
		if msg.status == statusFailed {
			return m, nil
		}

		// This is a simulated installation process
		// In real implementation, this would call actual installer methods
		return m, tea.Tick(2*time.Second, func(time.Time) tea.Msg {
			// Simulate successful installation 90% of the time
			if m.items[msg.index].Name != "Error Demo" {
				return dashboardUpdateMsg{
					index:  msg.index,
					status: statusSuccess,
				}
			} else {
				return dashboardUpdateMsg{
					index:  msg.index,
					status: statusFailed,
					err:    fmt.Errorf("simulated error"),
				}
			}
		})

	case completedMsg:
		// All installations completed
		return m, tea.Quit
	}

	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m DashboardModel) View() string {
	if m.quitting {
		return "Installation aborted. Press any key to exit."
	}

	var statusText string
	if m.activeIndex < len(m.items) {
		currentItem := m.items[m.activeIndex]
		switch currentItem.Status {
		case statusPending:
			statusText = pendingStyle.Render("Waiting to install...")
		case statusInProgress:
			statusText = fmt.Sprintf("%s %s", m.spinner.View(), inProgressStyle.Render("Installing "+currentItem.Name+"..."))
		case statusSuccess:
			statusText = successStyle.Render("Installation successful!")
		case statusFailed:
			statusText = errorStyle.Render(fmt.Sprintf("Installation failed: %v", currentItem.Error))
		}
	} else {
		statusText = successStyle.Render("All components installed successfully!")
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		titleStyle.Render("Dotfiles Installer Dashboard"),
		tableStyle.Render(m.table.View()),
		"",
		statusText,
		"",
		"Press q to quit",
	)
}

// LaunchDashboard starts a TUI dashboard for visualizing installation status
func LaunchDashboard() {
	items := []DashboardItem{
		{Name: "Zsh", Description: "Shell configuration", Status: statusPending},
		{Name: "Homebrew", Description: "Package manager", Status: statusPending},
		{Name: "Neovim", Description: "Text editor", Status: statusPending},
		{Name: "Tmux", Description: "Terminal multiplexer", Status: statusPending},
		{Name: "Git", Description: "Version control", Status: statusPending},
		{Name: "Starship", Description: "Shell prompt", Status: statusPending},
		{Name: "Ghostty", Description: "Terminal emulator", Status: statusPending},
		{Name: "Cyberpunk Theme", Description: "Color theme", Status: statusPending},
	}

	p := tea.NewProgram(NewDashboard(items), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running dashboard:", err)
		os.Exit(1)
	}
}
