package woke

import (
	"fmt"
	"log"
	"os"

	ad "github.com/mahmoudgalalz/woke/data"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type wokeView struct {
	viewport viewport.Model
}

func woke() (*wokeView, error) {
	const width = 75
	var content = `**Show WK:** `

	vp := viewport.New(width, 5)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#3C3C3C"))
		// Inline(true).
		// Align(lipgloss.Left)
		// Reverse(true).

	ads, err := ad.RandomAd()

	if err != nil {
		log.Fatal(err)
	}

	content += fmt.Sprintf("[%s](%s)", ads.Ad, ads.Link)

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

	if err := cliFunc(); err != nil {
		fmt.Printf("Error running the CLI: %v\n", err)
	}
}
