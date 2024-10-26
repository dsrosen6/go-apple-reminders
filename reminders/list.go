package reminders

import (
	"encoding/json"
	"os/exec"
)

type Lists []string

func GetLists() (Lists, error) {
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
