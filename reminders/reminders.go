package reminders

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

const (
	cliCommand = "reminders"
)

type Reminders []Reminder

type Reminder struct {
	DueDate     time.Time `json:"dueDate"`
	ExternalID  string    `json:"externalId"`
	IsCompleted bool      `json:"isCompleted"`
	List        string    `json:"list"`
	Priority    int64     `json:"priority"`
	Title       string    `json:"title"`
}

func GetReminders(listName string) (Reminders, error) {
	var cmd *exec.Cmd
	if listName == "" {
		cmd = exec.Command(cliCommand, "show-all", "-f", "json")
	} else {
		cmd = exec.Command(cliCommand, "show", listName, "-f", "json")
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error getting reminder: %s", string(out))
	}

	var reminders Reminders
	if err := json.Unmarshal(out, &reminders); err != nil {
		return nil, err
	}

	return reminders, nil
}
