package reminders

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

// TODO: Description
const (
	cliCommand = "reminders"
)

// TODO: Description
type NewReminder struct {
	List     string
	Title    string
	DueDate  string // uses natural language
	Priority Priority
	Notes    string
}

// TODO: Description
type Priority string

// TODO: Description
const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

// TODO: Description
type Lists []string

// TODO: Description
type Reminders []Reminder

// TODO: Description
type Reminder struct {
	List        string    `json:"list"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"isCompleted"`
	DueDate     time.Time `json:"dueDate,omitempty"`
	ExternalID  string    `json:"externalId,omitempty"`
	Priority    int64     `json:"priority,omitempty"`
}

// TODO: Description
type EditArgs struct {
	Title string // optional new title
	Notes string // optional new notes
}

// TODO: Description
type ShowRemindersArgs struct {
	List             string
	OnlyCompleted    bool
	IncludeCompleted bool
	Sort             SortType  // one of SortCreationDate or SortDueDate - leave empty for no sort
	SortOrder        SortOrder // one of SortAscending or SortDescending - default is SortAscending
	DueDate          string    // natural language
}

// TODO: Description
type (
	SortType  string
	SortOrder string
)

const (
	SortCreationDate SortType  = "creation-date"
	SortDueDate      SortType  = "due-date"
	SortAscending    SortOrder = "ascending"
	SortDescending   SortOrder = "descending"
)

// TODO: Description
func AddReminder(newReminder NewReminder) (*Reminder, error) {
	// USAGE: reminders add <list-name> <reminder> ... [--due-date <due-date>] [--priority <priority>] [--format <format>] [--notes <notes>]
	args := []string{"add", newReminder.List, newReminder.Title}
	if newReminder.DueDate != "" {
		args = append(args, "--due-date", newReminder.DueDate)
	}

	if newReminder.Priority != "" {
		args = append(args, "--priority", string(newReminder.Priority))
	}

	if newReminder.Notes != "" {
		args = append(args, "--notes", newReminder.Notes)
	}

	// add the format flag to get the reminder back as JSON
	args = append(args, "--format", "json")

	cmd := exec.Command(cliCommand, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error adding reminder: %s", string(out))
	}

	var reminder Reminder
	if err := json.Unmarshal(out, &reminder); err != nil {
		return nil, fmt.Errorf("error unmarshaling reminder: %s", string(out))
	}

	return &reminder, nil
}

// TODO: Description
func (r Reminder) Complete() error {
	// USAGE: reminders complete <list-name> <index>
	cmd := exec.Command(cliCommand, "complete", r.List, r.ExternalID)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error completing reminder: %s", string(out))
	}

	return nil
}

// TODO: Description
func (r Reminder) Uncomplete() error {
	// USAGE: reminders uncomplete <list-name> <index>
	cmd := exec.Command(cliCommand, "uncomplete", r.List, r.ExternalID)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error uncompleting reminder: %s", string(out))
	}

	return nil
}

// TODO: Description
func (r Reminder) Delete() error {
	// USAGE: reminders delete <list-name> <index>
	cmd := exec.Command(cliCommand, "delete", r.List, r.ExternalID)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error deleting reminder: %s", string(out))
	}

	return nil
}

// TODO: Description
func (r Reminder) Edit(e EditArgs) error {
	// USAGE: reminders edit <list-name> <index> [--notes <notes>] [<reminder> ...]
	args := []string{"edit", r.List, r.ExternalID}
	if e.Title != "" {
		args = append(args, e.Title)
	}

	if e.Notes != "" {
		args = append(args, "--notes", e.Notes)
	}

	cmd := exec.Command(cliCommand, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error editing reminder: %s", string(out))
	}

	return nil
}

// TODO: Description
func ShowReminders(s ShowRemindersArgs) (Reminders, error) {
	// USAGE: reminders show <list-name> [--only-completed] [--include-completed] [--sort <sort>] [--sort-order <sort-order>] [--due-date <due-date>] [--format <format>]
	args := []string{"show"}
	if s.List != "" {
		args = append(args, s.List)
	} else {
		return nil, fmt.Errorf("list name is required")
	}

	if s.OnlyCompleted {
		args = append(args, "--completed")
	}

	if s.IncludeCompleted {
		args = append(args, "--include-completed")
	}

	if s.Sort != "" {
		args = append(args, "--sort", string(s.Sort))
	}

	if s.SortOrder != "" {
		args = append(args, "--sort-order", string(s.SortOrder))
	}

	if s.DueDate != "" {
		args = append(args, "--due-date", s.DueDate)
	}

	args = append(args, "--format", "json")

	cmd := exec.Command(cliCommand, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error showing reminders: %s", string(out))
	}

	var reminders Reminders
	if err := json.Unmarshal(out, &reminders); err != nil {
		return nil, err
	}

	return reminders, nil
}

// TODO: Description
func ShowLists() (Lists, error) {
	// USAGE: reminders show-lists [--format <format>]
	cmd := exec.Command(cliCommand, "show-lists", "-f", "json")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var lists Lists
	if err := json.Unmarshal(out, &lists); err != nil {
		return nil, err
	}

	return lists, nil
}

// TODO: NewList()
// TODO: ShowAllReminders()
