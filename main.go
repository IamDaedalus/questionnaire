package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var qs = []question {
    newQuestion("what is 5 * 10?", "50", [3]string{"510", "15", "50"}),
    newQuestion("what is 4 + 16?", "20", [3]string{"20", "12", "46"}),
    newQuestion("what is 10 / 2", "5", [3]string{"5", "13", "20"}),
    newQuestion("what is 90 - 10", "80", [3]string{"24", "100", "80"}),
    newQuestion("where is Ghana located?", "Africa", [3]string{"Europe", "Asia", "Africa"}),
}

var curQ = 0

type question struct {
	question      string
	correctAns string
	choices       [3]string
}

type model struct {
	qs    []question
	confirmation string
	current      int
}

// create a new question
func newQuestion(q, ans string, choices [3]string) question {
    return question{
        question: q,
        correctAns: ans,
        choices: choices,
    }
}

func main() {
	p := tea.NewProgram(initModel())
	_, err := p.Run()
	if err != nil {
		fmt.Printf("there was an error: %v", err)
		os.Exit(1)
	}
}

func initModel() tea.Model {
	return model{
        qs: qs[:],
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// implements the update method for bubbletea
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.current > 0 {
				m.current--
			} else {
				m.current = len(m.qs[curQ].choices) - 1
			}

		case "down", "j":
			if m.current < len(m.qs[curQ].choices) - 1 {
				m.current++
			} else {
				m.current = 0
			}

		case "enter", " ":
            if m.qs[curQ].choices[m.current] == m.qs[curQ].correctAns{ 
                m.confirmation = "good job!"

                if curQ < len(m.qs) - 1 {
                    curQ++
                    m.current = 0
                } else {
                    return m, tea.Quit
                }
			} else {
				m.confirmation = "try again..."
			}
		}

	}
	return m, nil
}

// implements the view method for bubbletea
func (m model) View() string {
	s := "welcome to the math quiz"
	s += "\n1. " + m.qs[curQ].question + "\n"

	for i := range m.qs[curQ].choices {
		cursor := " "

		if i == m.current {
			cursor = ">"
		}
		s += fmt.Sprintf("%s  %s\n", cursor, m.qs[curQ].choices[i])
	}

	if len(m.confirmation) > 0 {
		s += "\n" + m.confirmation + "\n"
	}

	return s
}
