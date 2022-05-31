package cli

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	Mode   string // mine or wallet mode
	Cursor int
	Wallet string // current wallet
}

func initialModel() model {
	return model{
		Mode:   "wallet",
		Cursor: 0,
	}
}

func (m *model) Init() *tea.Cmd {
	m.Mode = "wallet"
	m.Cursor = 0
}
