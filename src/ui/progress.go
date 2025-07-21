package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/enescakir/emoji"
)

type progressMsg struct {
	step        int
	total       int
	description string
	completed   bool
	err         error
}

type taskCompleteMsg struct {
	step int
	desc string
}

type ProgressModel struct {
	progress    progress.Model
	spinner     spinner.Model
	steps       []string
	currentStep int
	completed   []bool
	errors      []error
	width       int
	quitting    bool
}

func NewProgressModel(steps []string) ProgressModel {
	p := progress.New(progress.WithDefaultGradient())
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return ProgressModel{
		progress:  p,
		spinner:   s,
		steps:     steps,
		completed: make([]bool, len(steps)),
		errors:    make([]error, len(steps)),
		width:     80,
	}
}

func (m ProgressModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.progress.Init())
}

func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.progress.Width = msg.Width - 4
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case progressMsg:
		if msg.err != nil {
			m.errors[msg.step] = msg.err
		}
		m.completed[msg.step] = msg.completed
		
		if msg.completed {
			m.currentStep = msg.step + 1
		}

		progressPercent := float64(m.currentStep) / float64(len(m.steps))
		return m, m.progress.SetPercent(progressPercent)

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

func (m ProgressModel) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder

	// Header
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("36")).
		MarginBottom(1)
	
	b.WriteString(headerStyle.Render("🚀 Dotfiles Installation Progress"))
	b.WriteString("\n\n")

	// Progress bar
	b.WriteString(m.progress.View())
	b.WriteString("\n\n")

	// Steps
	for i, step := range m.steps {
		var icon, color string
		var style lipgloss.Style

		if i < m.currentStep || m.completed[i] {
			if m.errors[i] != nil {
				icon = emoji.CrossMark.String()
				color = "196"
				style = lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Strikethrough(true)
			} else {
				icon = emoji.CheckMark.String()
				color = "34"
				style = lipgloss.NewStyle().Foreground(lipgloss.Color(color))
			}
		} else if i == m.currentStep {
			icon = m.spinner.View()
			color = "220"
			style = lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Bold(true)
		} else {
			icon = emoji.HollowRedCircle.String()
			color = "240"
			style = lipgloss.NewStyle().Foreground(lipgloss.Color(color))
		}

		b.WriteString(fmt.Sprintf("  %s %s", icon, style.Render(step)))
		
		// Show error if exists
		if m.errors[i] != nil {
			errorStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color("196")).
				Italic(true).
				MarginLeft(4)
			b.WriteString(errorStyle.Render(fmt.Sprintf("\n    Error: %v", m.errors[i])))
		}
		
		b.WriteString("\n")
	}

	// Status
	completed := 0
	for _, c := range m.completed {
		if c {
			completed++
		}
	}

	statusStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Italic(true).
		MarginTop(1)

	if m.currentStep >= len(m.steps) {
		b.WriteString(statusStyle.Render(fmt.Sprintf("\n%v Installation completed! (%d/%d steps)", emoji.CheckMark, completed, len(m.steps))))
	} else {
		b.WriteString(statusStyle.Render(fmt.Sprintf("\n%v Working... (%d/%d steps completed)", emoji.Gear, completed, len(m.steps))))
	}

	return b.String()
}

func SendProgress(step int, total int, description string, completed bool, err error) tea.Cmd {
	return func() tea.Msg {
		return progressMsg{
			step:        step,
			total:       total,
			description: description,
			completed:   completed,
			err:         err,
		}
	}
}

func RunWithProgress(steps []string, tasks []func() error) error {
	if len(steps) != len(tasks) {
		return fmt.Errorf("steps and tasks length mismatch")
	}

	m := NewProgressModel(steps)
	p := tea.NewProgram(m)

	// Run tasks in background
	go func() {
		for i, task := range tasks {
			time.Sleep(200 * time.Millisecond) // Small delay for UX
			
			err := task()
			
			p.Send(progressMsg{
				step:      i,
				completed: true,
				err:       err,
			})
			
			if err != nil {
				time.Sleep(1 * time.Second) // Show error for a moment
			}
		}
		
		// Wait a bit then quit
		time.Sleep(2 * time.Second)
		p.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	}()

	_, err := p.Run()
	return err
}