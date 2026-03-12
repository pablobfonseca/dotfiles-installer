package prompts

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pablobfonseca/dotfiles/src/installer"
	"github.com/pablobfonseca/dotfiles/src/utils"
)

var (
	tableStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00FFFF"))

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FFFF"))

	selectedRowStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#000000")).
				Background(lipgloss.Color("#00FFFF"))

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("10"))

	pendingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("11"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("9"))

	inProgressStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("14"))

	skippedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

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
	statusSelected
	statusInProgress
	statusSuccess
	statusFailed
	statusSkipped
)

type phase int

const (
	phaseSelecting phase = iota
	phaseInstalling
	phaseDone
)

// DashboardItem represents a single installable component in the dashboard.
type DashboardItem struct {
	Name        string
	Description string
	Status      status
	Error       error
	Installer   func() error
}

// dashboardUpdateMsg is sent when an installation completes or fails.
type dashboardUpdateMsg struct {
	index  int
	status status
	err    error
}

// doneMsg is sent after a brief delay once all installations finish.
type doneMsg struct{}

// DashboardModel is the Bubble Tea model for the two-phase installation dashboard.
type DashboardModel struct {
	items    []DashboardItem
	table    table.Model
	spinner  spinner.Model
	selected map[int]bool
	phase    phase
	current  int   // position within queue currently being installed
	queue    []int // indices of selected items in insertion order
	width    int
	height   int
	quitting bool
}

// NewDashboard constructs a DashboardModel ready for the selection phase.
func NewDashboard(items []DashboardItem) DashboardModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))

	columns := []table.Column{
		{Title: "Component", Width: 15},
		{Title: "Description", Width: 40},
		{Title: "Status", Width: 18},
	}

	rows := make([]table.Row, 0, len(items))
	for _, item := range items {
		rows = append(rows, table.Row{item.Name, item.Description, pendingStyle.Render("Pending")})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(items)),
	)

	t.SetStyles(table.Styles{
		Header:   headerStyle,
		Selected: selectedRowStyle,
		Cell:     lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1),
	})

	return DashboardModel{
		items:    items,
		table:    t,
		spinner:  s,
		selected: make(map[int]bool),
		phase:    phaseSelecting,
	}
}

// Init starts the spinner tick; installation does not begin until the user confirms.
func (m DashboardModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m DashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}

		switch m.phase {
		case phaseSelecting:
			return m.updateSelecting(msg)
		case phaseInstalling:
			// Ignore most keys during installation; q/ctrl+c handled above.
			return m, nil
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
		return m.updateInstalling(msg)

	case doneMsg:
		return m, tea.Quit
	}

	// Forward remaining messages to the table only during selection (navigation).
	if m.phase == phaseSelecting {
		var cmd tea.Cmd
		m.table, cmd = m.table.Update(msg)
		return m, cmd
	}

	return m, nil
}

// updateSelecting handles all key events during the selection phase.
func (m DashboardModel) updateSelecting(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case " ":
		cursor := m.table.Cursor()
		if cursor < 0 || cursor >= len(m.items) {
			return m, nil
		}
		m.selected[cursor] = !m.selected[cursor]
		m.refreshRowStatus(cursor)
		return m, nil

	case "a":
		for i := range m.items {
			m.selected[i] = true
			m.refreshRowStatus(i)
		}
		return m, nil

	case "n":
		for i := range m.items {
			m.selected[i] = false
			m.refreshRowStatus(i)
		}
		return m, nil

	case "enter":
		// Build the install queue from selected indices (preserving order).
		m.queue = m.queue[:0]
		for i := range m.items {
			if m.selected[i] {
				m.queue = append(m.queue, i)
			}
		}
		if len(m.queue) == 0 {
			return m, nil
		}

		m.phase = phaseInstalling
		m.current = 0

		// Mark all non-selected items as skipped.
		rows := m.table.Rows()
		for i := range m.items {
			if !m.selected[i] {
				m.items[i].Status = statusSkipped
				rows[i][2] = skippedStyle.Render("Skipped")
			}
		}
		m.table.SetRows(rows)

		return m, m.startNextInstall()

	default:
		// Forward navigation keys (↑/↓, etc.) to the table.
		var cmd tea.Cmd
		m.table, cmd = m.table.Update(msg)
		return m, cmd
	}
}

// updateInstalling handles dashboardUpdateMsg events during the installation phase.
func (m DashboardModel) updateInstalling(msg dashboardUpdateMsg) (tea.Model, tea.Cmd) {
	m.items[msg.index].Status = msg.status
	m.items[msg.index].Error = msg.err

	rows := m.table.Rows()
	switch msg.status {
	case statusSuccess:
		rows[msg.index][2] = successStyle.Render("✓ Success")
	case statusFailed:
		rows[msg.index][2] = errorStyle.Render(fmt.Sprintf("✗ Failed: %v", msg.err))
	}
	m.table.SetRows(rows)

	m.current++
	return m, m.startNextInstall()
}

// startNextInstall advances to the next item in the queue and launches its installer,
// or transitions to phaseDone when the queue is exhausted.
func (m *DashboardModel) startNextInstall() tea.Cmd {
	if m.current >= len(m.queue) {
		m.phase = phaseDone
		return tea.Tick(2*time.Second, func(time.Time) tea.Msg {
			return doneMsg{}
		})
	}

	idx := m.queue[m.current]
	m.items[idx].Status = statusInProgress

	rows := m.table.Rows()
	rows[idx][2] = inProgressStyle.Render("Installing...")
	m.table.SetRows(rows)

	installerFn := m.items[idx].Installer
	return func() tea.Msg {
		err := installerFn()
		if err != nil {
			return dashboardUpdateMsg{index: idx, status: statusFailed, err: err}
		}
		return dashboardUpdateMsg{index: idx, status: statusSuccess}
	}
}

// refreshRowStatus updates the Status cell for a single row to reflect its selection state.
func (m *DashboardModel) refreshRowStatus(i int) {
	rows := m.table.Rows()
	if m.selected[i] {
		rows[i][2] = successStyle.Render("☑ Selected")
	} else {
		rows[i][2] = pendingStyle.Render("Pending")
	}
	m.table.SetRows(rows)
}

func (m DashboardModel) View() string {
	if m.quitting {
		return "Installer exited.\n"
	}

	header := titleStyle.Render("Dotfiles Installer Dashboard")
	tbl := tableStyle.Render(m.table.View())

	switch m.phase {
	case phaseSelecting:
		selectedCount := len(m.selected)
		actualCount := 0
		for _, v := range m.selected {
			if v {
				actualCount++
			}
		}
		hint := fmt.Sprintf("↑/↓: navigate • space: toggle • a: all • n: none • enter: install • q: quit   (%d/%d selected)",
			actualCount, selectedCount)
		return lipgloss.JoinVertical(lipgloss.Left, header, tbl, "", hint)

	case phaseInstalling:
		var currentName string
		if m.current < len(m.queue) {
			currentName = m.items[m.queue[m.current]].Name
		}
		installing := fmt.Sprintf("%s %s", m.spinner.View(), inProgressStyle.Render("Installing "+currentName+"..."))
		return lipgloss.JoinVertical(lipgloss.Left, header, tbl, "", installing, "", "q: quit")

	case phaseDone:
		succeeded, failed, skipped := 0, 0, 0
		for _, item := range m.items {
			switch item.Status {
			case statusSuccess:
				succeeded++
			case statusFailed:
				failed++
			case statusSkipped:
				skipped++
			}
		}
		summary := successStyle.Render(fmt.Sprintf(
			"Done! %d succeeded • %d failed • %d skipped — closing...",
			succeeded, failed, skipped,
		))
		return lipgloss.JoinVertical(lipgloss.Left, header, tbl, "", summary)
	}

	return ""
}

// LaunchDashboard starts the interactive TUI dashboard.
// Non-interactive mode is set so installer prompts auto-confirm inside the TUI.
func LaunchDashboard() {
	utils.SetNonInteractiveMode(true)

	items := []DashboardItem{
		{Name: "Homebrew", Description: "Package manager", Installer: installer.InstallHomebrew},
		{Name: "Zsh", Description: "Shell configuration", Installer: installer.SetupZsh},
		{Name: "Neovim", Description: "Text editor", Installer: installer.InstallNvim},
		{Name: "Git", Description: "Version control", Installer: installer.SetupGit},
		{Name: "Tmux", Description: "Terminal multiplexer", Installer: installer.SetupTmux},
		{Name: "Starship", Description: "Shell prompt", Installer: installer.SetupStarship},
		{Name: "Ghostty", Description: "Terminal emulator", Installer: installer.SetupGhostty},
		{Name: "Karabiner", Description: "Keyboard customization", Installer: func() error {
			if err := installer.InstallKarabiner(); err != nil {
				return err
			}
			return installer.SetupKarabiner()
		}},
		{Name: "Cyberpunk Theme", Description: "Color theme", Installer: installer.InstallCyberpunkTheme},
	}

	p := tea.NewProgram(NewDashboard(items), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running dashboard:", err)
		os.Exit(1)
	}
}
