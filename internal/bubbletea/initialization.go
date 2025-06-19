package bubbletea

import (
	"bytes"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/Blaze5333/todo-cli/internal/storage"
	"github.com/Blaze5333/todo-cli/internal/todo"
	"github.com/Blaze5333/todo-cli/internal/user"
	"github.com/Blaze5333/todo-cli/utils"
	"github.com/aquasecurity/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type mode int

const (
	modeView mode = iota
	modeAdd
)

type model struct {
	Task        []todo.Task `json:"task"`
	Cursor      int         `json:"cursor"`
	CurrentMode mode        `json:"current_mode"`
	showDialog  bool
	inputIndex  int
	inputs      []textinput.Model
	Cursor2     int
	UpdateTask  int
}

var strg = storage.NewStorage[[]todo.Todo]("todos.json")

func initialModel() model {
	inputs := make([]textinput.Model, 3)
	t := textinput.New()
	t.Placeholder = "Title"
	t.Focus()
	t.Width = 20
	inputs[0] = t

	d := textinput.New()
	d.Placeholder = "Description"
	d.Width = 20
	inputs[1] = d

	p := textinput.New()
	p.Placeholder = "Priority"
	p.CharLimit = 1
	p.Validate = func(s string) error {
		if s == "" {
			return nil // Allow empty input for priority
		}
		if _, err := strconv.Atoi(s); err != nil {
			return err // Return error if input is not a number
		}
		if s != "1" && s != "2" && s != "3" {
			return errors.New("priority must be 1, 2, or 3") // Ensure priority is one of the valid options
		}
		return nil // Return nil if input is valid
	}

	p.Width = 20
	inputs[2] = p

	return model{
		showDialog: false,
		inputIndex: 0,
		inputs:     inputs,
	}
}

func renderTodos(todos []todo.Task, selected int, selected2 int) string {
	var buf bytes.Buffer
	tbl := table.New(&buf)
	tbl.SetHeaders("ID", "Title", "Description", "Priority", "Completed", "Created At", "Updated At")

	tbl.SetHeaderStyle(table.StyleBold)

	tbl.SetLineStyle(table.StyleMagenta)
	tbl.SetDividers(table.UnicodeRoundedDividers)

	for index, task := range todos {

		completed := "❌"
		if task.Done {
			completed = "✅"
		}
		var colort string
		switch task.Priority {
		case 1:
			colort = "high 🔴"
		case 2:
			colort = "medium 🟠"
		case 3:
			colort = "low 🟡"
		default:
			colort = "⚪️"
		}
		id := strconv.Itoa(index)
		title := task.Title
		description := task.Description
		if index == selected {

			switch selected2 {
			case 0:
				id = "👉 " + id
			case 1:
				title = "👉 " + title
			case 2:
				description = "👉 " + description
			case 3:
				colort = "👉 " + colort
			}

		}

		tbl.AddRow(
			id,
			title,
			description,
			colort,
			completed,
			task.CreatedAt.Format("2006-01-02 15:04:05"),
			task.UpdatedAt.Format("2006-01-02 15:04:05"),
		)

	}
	tbl.Render()
	return buf.String()
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {

	s := "TODO LIST\n\n"
	s += renderTodos(m.Task, m.Cursor, m.Cursor2)
	s += "\n[UP/DOWN] Select  [ENTER] Complete  [DEL] Delete  [A] Add  [U] Update the selected task  [Q] Quit\n"
	if m.showDialog {
		s := "Add New Todo\n\n All fields are required.\n\n"
		for i := range m.inputs {
			s += m.inputs[i].View() + "\n"
		}
		s += "\nFOR Priority, use numbers:\n1. High 🔴\n2. Medium 🟠\n3. Low 🟡\n[Enter] OK   [Tab] Next   [Esc] Cancel"
		return s
	}
	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.showDialog {

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {

			case "down":
				if (m.inputIndex + 1) < len(m.inputs) {
					m.inputIndex++
				} else {
					m.inputIndex = 0
				}

			case "up":
				if m.inputIndex > 0 {
					m.inputIndex--
				} else {
					m.inputIndex = len(m.inputs) - 1
				}
			case "enter":
				if m.inputIndex != len(m.inputs)-1 {
					m.inputIndex++
				}
				if m.inputs[0].Value() == "" || m.inputs[1].Value() == "" || m.inputs[2].Value() == "" {
					break
				}

				if m.UpdateTask != -1 {
					priority, _ := strconv.Atoi(m.inputs[2].Value())
					m.Task[m.UpdateTask].Title = m.inputs[0].Value()
					m.Task[m.UpdateTask].Description = m.inputs[1].Value()
					m.Task[m.UpdateTask].Priority = priority
					m.Task[m.UpdateTask].UpdatedAt = time.Now()
				} else {
					priority, _ := strconv.Atoi(m.inputs[2].Value())
					addedTask := todo.Task{
						Title:       m.inputs[0].Value(),
						Description: m.inputs[1].Value(),
						Priority:    priority, // Default priority, can be changed later
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
						Done:        false,
					}
					m.Task = append(m.Task, addedTask)
				}
				m.UpdateTask = -1
				m.inputs[0].SetValue("")
				m.inputs[1].SetValue("")
				m.inputs[2].SetValue("")
				m.showDialog = false

			case "esc":
				m.showDialog = false
			}
			// Update all input fields' focus
			for i := range m.inputs {
				if i == m.inputIndex {
					m.inputs[i].Focus()
				} else {
					m.inputs[i].Blur()
				}
			}
			// Pass the key to the focused input
			var cmd tea.Cmd
			m.inputs[m.inputIndex], cmd = m.inputs[m.inputIndex].Update(msg)
			return m, cmd
		}
	} else {
		switch msg := msg.(type) {

		case tea.KeyMsg:

			switch msg.String() {

			case "q":
				var currentUserIndex = -1
				for i, todo := range wholeData {
					if todo.Username == name {
						wholeData[i].Task = m.Task
						currentUserIndex = i
						break
					}
				}
				if currentUserIndex == -1 {
					wholeData = append(wholeData, todo.Todo{
						Username: name,
						Task:     m.Task,
					})
				}
				err := strg.Save(wholeData)
				if err != nil {
					utils.ShowErrorMessage("Error Saving Tasks  : " + err.Error())
				}
				return m, tea.Quit
			case "right":
				if m.Cursor2 < 3 {
					m.Cursor2++
				} else {
					m.Cursor2 = 3
				}
			case "left":
				if m.Cursor2 > 1 {
					m.Cursor2--
				} else {
					m.Cursor2 = 0
				}
			case "up":
				if m.Cursor > 0 {
					m.Cursor--
				}
			case "down":
				if m.Cursor < len(m.Task)-1 {
					m.Cursor++
				}
			case "u":
				m.inputs[0].SetValue(m.Task[m.Cursor].Title)
				m.inputs[1].SetValue(m.Task[m.Cursor].Description)
				m.inputs[2].SetValue(strconv.Itoa(m.Task[m.Cursor].Priority))
				m.showDialog = true
				m.inputIndex = 0
				m.inputs[0].Focus()
				m.UpdateTask = m.Cursor
			case "enter":
				m.Task[m.Cursor].Done = !m.Task[m.Cursor].Done

			case "backspace":
				if len(m.Task) > 0 {
					m.Task = append(m.Task[:m.Cursor], m.Task[m.Cursor+1:]...)
					if m.Cursor >= len(m.Task) {
						m.Cursor = len(m.Task) - 1
					}
				}
			case "a":
				m.showDialog = true
				m.inputIndex = 0
				m.inputs[0].Focus()

			}

		}
	}
	return m, nil
}

var wholeData []todo.Todo
var name string

func Start() {
	name = user.CheckSession()

	tasks, _ := todo.GetTasks(name)
	err := strg.Load(&wholeData)

	if err != nil {
		utils.ShowErrorMessage("Error fetching tasks: " + err.Error())
		return
	}
	mod := initialModel()
	m := model{
		Task:        tasks,
		Cursor:      0,
		CurrentMode: modeView,
		showDialog:  mod.showDialog,
		inputIndex:  mod.inputIndex,
		inputs:      mod.inputs,
		Cursor2:     0,
		UpdateTask:  -1,
	}
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		utils.ShowErrorMessage("Error running Bubble Tea program: " + err.Error())
		os.Exit(1)
	}

}
