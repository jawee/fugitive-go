package main

import (
	"fmt"
	"os"

	"github.com/jawee/fugitive-go/internal/git"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    state git.GitStatus
    len int
    cursor int                
    current string
    currentType string

    footer string
}

func initialModel() *model {
    state, err := git.GetStatus()
    if err != nil {
        os.Exit(1)
    }
	return &model{
        state: *state,
        len: 6,
	}
}

func (m *model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyMsg:

        switch msg.String() {

        case "ctrl+c", "q":
            fmt.Printf("%s %s\n", m.current, m.currentType)
            return m, tea.Quit

        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        case "down", "j":
            if m.cursor < m.len-1 {
                m.cursor++
            }

        case "enter", "s":
            m.footer = fmt.Sprintf("Selected: %s %s\n", m.current, m.currentType)
        }
    }

    return m, nil
}

func (m *model) View() string {
    s := "Fugitive\n\n"

    i := 0
    s += fmt.Sprintf("Staged\n")
    for _, choice := range m.state.Staged {

        cursor := " " 
        if m.cursor == i {
            cursor = ">" 
            m.current = choice
            m.currentType = "staged"
        }

        s += fmt.Sprintf("%s %s\n", cursor, choice)
        i++
    }

    s += fmt.Sprintf("Unstaged\n")
    for _, choice := range m.state.Unstaged {

        cursor := " " 
        if m.cursor == i {
            cursor = ">" 
            m.current = choice
            m.currentType = "Unstaged"
        }

        s += fmt.Sprintf("%s %s\n", cursor, choice)
        i++
    }

    s += fmt.Sprintf("Untracked\n")
    for _, choice := range m.state.Untracked {

        cursor := " " 
        if m.cursor == i {
            cursor = ">" 
            m.current = choice
            m.currentType = "Untracked"
        }

        s += fmt.Sprintf("%s %s\n", cursor, choice)
        i++
    }

    s += "\nPress q to quit.\n"
    s += fmt.Sprintf("%s\n", m.footer)

    return s
}

func main() {
    p := tea.NewProgram(initialModel(), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
