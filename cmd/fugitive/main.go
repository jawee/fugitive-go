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
    // choices  []string           // items on the to-do list
    cursor   int                // which to-do list item our cursor is pointing at
    current string
    currentType string

    footer string
}

func initialModel() model {
    state, err := git.GetStatus()
    if err != nil {
        os.Exit(1)
    }
	return model{
        state: *state,
        len: 6,
		// Our to-do list is a grocery list
		// choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// // A map which indicates which choices are selected. We're using
		// // the  map like a mathematical set. The keys refer to the indexes
		// // of the `choices` slice, above.
		// selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            fmt.Printf("%s %s\n", m.current, m.currentType)
            return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursor < m.len-1 {
                m.cursor++
            }

        // The "enter" key and the spacebar (a literal space) toggle
        // the selected state for the item that the cursor is pointing at.
        case "enter", "s":
            m.footer = fmt.Sprintf("Selected: %s %s\n", m.current, m.currentType)
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() string {
    // The header
    s := "Fugitive\n\n"

    // Iterate over our choices
    i := 0
    s += fmt.Sprintf("Staged\n")
    for _, choice := range m.state.Staged {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
            m.current = choice
            m.currentType = "staged"
        }

        // Is this choice selected?
        // checked := " " // not selected
        // if _, ok := m.selected[i]; ok {
        //     checked = "x" // selected!
        // }

        // Render the row
        s += fmt.Sprintf("%s %s\n", cursor, choice)
        i++
    }

    s += fmt.Sprintf("Unstaged\n")
    for _, choice := range m.state.Unstaged {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
            m.current = choice
            m.currentType = "Unstaged"
        }

        // Is this choice selected?
        // checked := " " // not selected
        // if _, ok := m.selected[i]; ok {
        //     checked = "x" // selected!
        // }

        // Render the row
        s += fmt.Sprintf("%s %s\n", cursor, choice)
        i++
    }

    s += fmt.Sprintf("Untracked\n")
    for _, choice := range m.state.Untracked {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
            m.current = choice
            m.currentType = "Untracked"
        }

        // Is this choice selected?
        // checked := " " // not selected
        // if _, ok := m.selected[i]; ok {
        //     checked = "x" // selected!
        // }

        // Render the row
        s += fmt.Sprintf("%s %s\n", cursor, choice)
        i++
    }
    // The footer
    s += "\nPress q to quit.\n"
    s += fmt.Sprintf("%s\n", m.footer)

    // Send the UI for rendering
    return s
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
