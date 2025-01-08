package prompts

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/spf13/viper"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#008B8B"))
	noStyle      = lipgloss.NewStyle()
	helpStyle    = blurredStyle.Copy()

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type configData struct {
	repositoryUrl     string
	dotfilesConfigDir string
	nvimConfigDir     string
}

type model struct {
	focusIndex int
	inputs     []Input
	textInputs []textinput.Model
	config     configData
	err        error
}

func ConfigPrompt() model {
	inputs := []Input{
		NewInput("\uf408", "repository (e.g, username/dotfiles)", "pablobfonseca/dotfiles"),
		NewInput("\uebdf", "dotfiles directory (e.g, ~/.dotfiles)", "~/.dotfiles"),
		NewInput("\ue7c5", "config directory (e.g, ~/.config/nvim)", "~/.config/nvim"),
	}

	textInputs := make([]textinput.Model, len(inputs))
	for i, in := range inputs {
		t := textinput.New()
		t.Prompt = in.icon + " "
		t.Placeholder = in.placeholder
		t.Cursor.Style = focusedStyle

		if i == 0 {
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		textInputs[i] = t
	}

	return model{
		inputs:     inputs,
		textInputs: textInputs,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			os.Exit(0)

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs) {
				err := m.persistConfig()
				if err != nil {
					utils.ErrorMessage("Error creating config file", err)
				}
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.textInputs[i].Focus()
					m.textInputs[i].PromptStyle = focusedStyle
					m.textInputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.textInputs[i].Blur()
				m.textInputs[i].PromptStyle = noStyle
				m.textInputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.textInputs))

	for i := range m.textInputs {
		m.textInputs[i], cmds[i] = m.textInputs[i].Update(msg)
		m.inputs[i].value = m.textInputs[i].Value()
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.textInputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}

	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}

func (m *model) persistConfig() error {
	m.config.repositoryUrl = m.inputs[0].Value()
	m.config.dotfilesConfigDir = expandPath(m.inputs[1].Value())
	m.config.nvimConfigDir = expandPath(m.inputs[2].Value())
	m.config.repositoryUrl = m.inputs[3].Value()

	viper.Set("dotfiles.repository", m.config.repositoryUrl)
	viper.Set("dotfiles.default_dir", m.config.dotfilesConfigDir)
	viper.Set("dotfiles.nvim.config_dir", m.config.nvimConfigDir)

	err := viper.SafeWriteConfig()
	if err == nil {
		utils.SuccessMessage("Successfully created your config")
	}

	return err
}

func expandPath(path string) string {
	if !strings.HasPrefix(path, "~") {
		return path
	}

	home, _ := os.UserHomeDir()
	return strings.Replace(path, "~", home, 1)
}
