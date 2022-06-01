package cli

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type CommandLine struct{}

type model struct {
	Mode    string // wallet or mining mode
	Cursor  int
	Address string
	Balance float32
}

func InitialModel() model {
	return model{
		Mode:    "",
		Cursor:  0,
		Address: "",
		Balance: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "w":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return ""
}

func New() *CommandLine {
	return &CommandLine{}
}

func (cli *CommandLine) Run() {
	p := tea.NewProgram(InitialModel())

	if err := p.Start(); err != nil {
		os.Exit(1)
	}
}
