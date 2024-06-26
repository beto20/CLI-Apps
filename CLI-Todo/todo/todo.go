package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/beto20/cli/colors"
)

type Item struct {
	Task        string
	Description string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type ItemTemp struct {
	Title       string
	Description string
}

type Todos []Item

func Plus(x, y int) int {
	return x + y
}

func (t *Todos) Addv2(temp ItemTemp) {
	todo := Item{
		Task:        temp.Title,
		Description: temp.Description,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Add(task string) {
	todo := Item{
		Task:        task,
		Description: "description",
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{
				Align: simpletable.AlignCenter,
				Text:  "#",
			},
			{
				Align: simpletable.AlignCenter,
				Text:  "Task",
			},
			{
				Align: simpletable.AlignCenter,
				Text:  "Desc",
			},
			{
				Align: simpletable.AlignCenter,
				Text:  "Done",
			},
			{
				Align: simpletable.AlignRight,
				Text:  "Created at",
			},
			{
				Align: simpletable.AlignRight,
				Text:  "Completed at",
			},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := colors.Blue(item.Task)
		done := colors.Blue("no")
		desc := colors.Blue(item.Description)

		if item.Done {
			task = colors.Green(fmt.Sprintf("\u2705 %s", item.Task))
			done = colors.Green("yes")
			desc = colors.Green(item.Description)
		}

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: desc},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 6, Text: colors.Red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}
