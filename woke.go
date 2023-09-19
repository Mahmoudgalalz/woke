package onboardbase

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const content = `
# New Era of Ads built by devs, for devs
[Visit](https://woke.onboardbase.com)`

type wokeView struct {
	viewport viewport.Model
}

func woke() (*wokeView, error) {
	const width = 44

	vp := viewport.New(width, 6)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return nil, err
	}

	str, err := renderer.Render(content)
	if err != nil {
		return nil, err
	}

	vp.SetContent(str)

	return &wokeView{
		viewport: vp,
	}, nil
}

func (e wokeView) Init() tea.Cmd {
	return nil
}

func (e wokeView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return e, nil
}

func (e wokeView) View() string {
	return e.viewport.View()
}

func Wrap(cliFunc func() error) {
	model, err := woke()
	if err != nil {
		fmt.Println("Could not initialize Bubble Tea model:", err)
		os.Exit(1)
	}
	fmt.Println(model.View())

	// Run the CLI function
	if err := cliFunc(); err != nil {
		fmt.Printf("Error running the CLI: %v\n", err)
	}
}
