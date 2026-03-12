package prompts

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pablobfonseca/dotfiles/src/installer"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	menuTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFFF")).
			Bold(true).
			Italic(true).
			MarginLeft(2)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#008B8B")).
				Padding(0, 1)
)

type installationStatus int

const (
	statusReady installationStatus = iota
	statusInstalling
	statusDone
	statusError
)

type MenuItem struct {
	title       string
	description string
	installer   func() error
}

func (i MenuItem) Title() string       { return i.title }
func (i MenuItem) Description() string { return i.description }
func (i MenuItem) FilterValue() string { return i.title }

type Model struct {
	list         list.Model
	status       installationStatus
	err          error
	statusMsg    string
	selectedItem string
}

func NewInstallerMenu() Model {
	items := []list.Item{
		MenuItem{
			title:       "Homebrew",
			description: "Install Homebrew and packages from Brewfile",
			installer:   installer.InstallHomebrew,
		},
		MenuItem{
			title:       "Zsh",
			description: "Setup Zsh configuration files",
			installer:   installer.SetupZsh,
		},
		MenuItem{
			title:       "Neovim",
			description: "Install and configure Neovim",
			installer:   installer.InstallNvim,
		},
		MenuItem{
			title:       "Git",
			description: "Setup Git configuration files",
			installer:   installer.SetupGit,
		},
		MenuItem{
			title:       "Tmux",
			description: "Setup Tmux configuration",
			installer:   installer.SetupTmux,
		},
		MenuItem{
			title:       "Starship",
			description: "Setup Starship prompt",
			installer:   installer.SetupStarship,
		},
		MenuItem{
			title:       "Ghostty",
			description: "Setup Ghostty terminal",
			installer:   installer.SetupGhostty,
		},
		MenuItem{
			title:       "Cyberpunk Theme",
			description: "Install cyberpunk color theme",
			installer:   installer.InstallCyberpunkTheme,
		},
		MenuItem{
			title:       "Install All",
			description: "Install all configuration files",
			installer:   installer.InstallConfigFiles,
		},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Dotfiles Installer"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = menuTitleStyle

	return Model{list: l, status: statusReady}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 4) // Make room for status message
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			if m.status == statusReady || m.status == statusDone || m.status == statusError {
				i, ok := m.list.SelectedItem().(MenuItem)
				if ok {
					m.status = statusInstalling
					m.selectedItem = i.title
					m.statusMsg = fmt.Sprintf("Installing %s...", i.title)
					return m, tea.Batch(
						m.list.NewStatusMessage(fmt.Sprintf("Installing %s...", i.title)),
						func() tea.Msg {
							err := i.installer()
							if err != nil {
								return errMsg{err}
							}
							return installedMsg(i.title)
						},
					)
				}
			}
		}

	case errMsg:
		m.err = msg
		m.status = statusError
		m.statusMsg = fmt.Sprintf("Error installing %s: %v", m.selectedItem, m.err)
		return m, m.list.NewStatusMessage(m.statusMsg)

	case installedMsg:
		m.status = statusDone
		m.statusMsg = fmt.Sprintf("Successfully installed %s!", string(msg))
		return m, m.list.NewStatusMessage(m.statusMsg)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	var statusView string

	switch m.status {
	case statusInstalling:
		statusView = statusMessageStyle.Render(fmt.Sprintf("Installing %s...", m.selectedItem))
	case statusDone:
		statusView = statusMessageStyle.Render(m.statusMsg)
	case statusError:
		statusView = statusMessageStyle.
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#FF0000")).
			Render(m.statusMsg)
	default:
		statusView = statusMessageStyle.Render("Press enter to install, q to quit")
	}

	return appStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			m.list.View(),
			"",
			statusView,
		),
	)
}

type errMsg struct{ error }
type installedMsg string

// Launch starts the installer menu
func LaunchInstallerMenu() {
	p := tea.NewProgram(NewInstallerMenu(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running menu:", err)
		os.Exit(1)
	}
}

