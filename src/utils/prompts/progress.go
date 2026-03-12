package prompts

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	progressBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFFF")).
			Bold(true)

	progressTrackStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#555555"))
)

type ProgressMsg struct{}

type ProgressModel struct {
	Total       int
	Current     int
	Description string
	Done        bool
	Err         error
	startTime   time.Time
}

func NewProgress(total int, description string) ProgressModel {
	return ProgressModel{
		Total:       total,
		Current:     0,
		Description: description,
		startTime:   time.Now(),
	}
}

func (m ProgressModel) Init() tea.Cmd {
	return tick
}

func tick() tea.Msg {
	time.Sleep(100 * time.Millisecond)
	return ProgressMsg{}
}

func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case ProgressMsg:
		if m.Done || m.Err != nil {
			return m, nil
		}

		// Simulate progress for demo purposes
		// In a real implementation, this would be updated based on actual task progress
		if m.Current < m.Total {
			m.Current++
		}

		if m.Current >= m.Total {
			m.Done = true
			return m, nil
		}

		return m, tick
	}

	return m, nil
}

func (m ProgressModel) View() string {
	width := 50
	if m.Total == 0 {
		return fmt.Sprintf("%s\n%s\n%s", m.Description, progressTrackStyle.Render(strings.Repeat("█", width)), "No tasks")
	}
	done := int(float64(m.Current) / float64(m.Total) * float64(width))

	bar := progressBarStyle.Render(strings.Repeat("█", done))
	track := progressTrackStyle.Render(strings.Repeat("█", width-done))

	percent := float64(m.Current) / float64(m.Total) * 100
	status := fmt.Sprintf(" %d/%d (%.0f%%)", m.Current, m.Total, percent)

	elapsed := time.Since(m.startTime).Round(time.Second)
	timeInfo := fmt.Sprintf(" %s elapsed", elapsed)

	if m.Done {
		return fmt.Sprintf("%s\n%s%s%s\n%s",
			m.Description,
			bar, track, status,
			"✓ Done!")
	}

	if m.Err != nil {
		return fmt.Sprintf("%s\n%s%s%s\n%s",
			m.Description,
			bar, track, status,
			fmt.Sprintf("✗ Error: %v", m.Err))
	}

	return fmt.Sprintf("%s\n%s%s%s%s",
		m.Description,
		bar, track, status, timeInfo)
}