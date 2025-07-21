package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/enescakir/emoji"
)

func ShowBanner() {
	// ASCII Art Banner
	banner := `
    ██████╗  ██████╗ ████████╗███████╗██╗██╗     ███████╗███████╗
    ██╔══██╗██╔═══██╗╚══██╔══╝██╔════╝██║██║     ██╔════╝██╔════╝
    ██║  ██║██║   ██║   ██║   █████╗  ██║██║     █████╗  ███████╗
    ██║  ██║██║   ██║   ██║   ██╔══╝  ██║██║     ██╔══╝  ╚════██║
    ██████╔╝╚██████╔╝   ██║   ██║     ██║███████╗███████╗███████║
    ╚═════╝  ╚═════╝    ╚═╝   ╚═╝     ╚═╝╚══════╝╚══════╝╚══════╝
    `

	// Styles
	bannerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("36")).
		Bold(true).
		Align(lipgloss.Center)

	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("220")).
		Bold(true).
		Align(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Italic(true).
		Align(lipgloss.Center).
		MarginBottom(2)

	// Render
	fmt.Println(bannerStyle.Render(banner))
	fmt.Println(titleStyle.Render(fmt.Sprintf("%s Personal Dotfiles Installer %s", emoji.Rocket, emoji.Sparkles)))
	fmt.Println(subtitleStyle.Render("Automate your development environment setup with style"))
}

func ShowWelcome() {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("36")).
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("36")).
		Padding(1, 2).
		MarginTop(1).
		MarginBottom(1)

	welcome := fmt.Sprintf(
		"%s Welcome to the Dotfiles Installer!\n\n"+
			"This tool will help you set up your development environment\n"+
			"with all your personal configurations and tools.\n\n"+
			"%s Ready to get started?",
		emoji.WavingHand,
		emoji.Rocket,
	)

	fmt.Println(style.Render(welcome))
}

func ShowSuccess(message string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("34")).
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("34")).
		Padding(1, 2).
		MarginTop(1).
		MarginBottom(1)

	fmt.Println(style.Render(fmt.Sprintf("%s %s", emoji.CheckMark, message)))
}

func ShowError(message string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("196")).
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("196")).
		Padding(1, 2).
		MarginTop(1).
		MarginBottom(1)

	fmt.Println(style.Render(fmt.Sprintf("%s %s", emoji.CrossMark, message)))
}

func ShowInfo(message string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("220")).
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("220")).
		Padding(1, 2).
		MarginTop(1).
		MarginBottom(1)

	fmt.Println(style.Render(fmt.Sprintf("%s %s", emoji.Information, message)))
}