package reminders

import (
	"testing"
	"time"
)

const (
	testList = "Test List"
)

func TestAddReminderMinimum(t *testing.T) {
	n := NewReminder{
		List:  testList,
		Title: "Test Reminder Minimum",
	}

	_, err := AddReminder(n)
	if err != nil {
		t.Errorf("error adding reminder: %s", err)
	}
}

func TestAddReminderWithDueDateNL(t *testing.T) {
	n := NewReminder{
		List:    testList,
		Title:   "Test Reminder with Due Date (Natural Language, No Time)",
		DueDate: "tomorrow",
	}

	_, err := AddReminder(n)
	if err != nil {
		t.Errorf("error adding reminder: %s", err)
	}
}

func TestAddReminderWithDueDateNLTime(t *testing.T) {
	n := NewReminder{
		List:    testList,
		Title:   "Test Reminder with Due Date (Natural Language With Time)",
		DueDate: "tomorrow at 2:33 pm",
	}

	_, err := AddReminder(n)
	if err != nil {
		t.Errorf("error adding reminder: %s", err)
	}
}

func TestAddReminderWithDueDateFormal(t *testing.T) {
	n := NewReminder{
		List:    testList,
		Title:   "Test Reminder with Due Date (Formatted Date, No Time)",
		DueDate: "12/25",
	}

	_, err := AddReminder(n)
	if err != nil {
		t.Errorf("error adding reminder: %s", err)
	}
}

func TestAddReminderWithDueDateFormalTime(t *testing.T) {
	n := NewReminder{
		List:    testList,
		Title:   "Test Reminder with Due Date (Formatted Date, With Time)",
		DueDate: "12/25 2:12 PM",
	}

	_, err := AddReminder(n)
	if err != nil {
		t.Errorf("error adding reminder: %s", err)
	}
}

func TestAddReminderWithPriority(t *testing.T) {
	n := NewReminder{
		List:     testList,
		Title:    "Test Reminder with Priority",
		Priority: PriorityMedium,
	}

	_, err := AddReminder(n)
	if err != nil {
		t.Errorf("error adding reminder: %s", err)
	}
}

func TestAddReminderWithNotes(t *testing.T) {
	n := NewReminder{
		List:  testList,
		Title: "Test Reminder with Notes",
		Notes: "This is a test reminder with notes.",
	}

	_, err := AddReminder(n)
	if err != nil {
		t.Errorf("error adding reminder: %s", err)
	}
}

func TestAddReminderWithAll(t *testing.T) {
	n := NewReminder{
		List:     testList,
		Title:    "Test Reminder with All Options",
		DueDate:  "tomorrow",
		Priority: PriorityHigh,
		Notes:    "This is a test reminder with all options.",
	}

	_, err := AddReminder(n)
	if err != nil {
		t.Errorf("error adding reminder: %s", err)
	}
}

func TestCompleteReminder(t *testing.T) {
	n := NewReminder{
		List:  testList,
		Title: "Test Reminder to be Completed",
	}

	r, err := AddReminder(n)
	if err != nil {
		t.Errorf("error adding reminder: %s", err)
	}

	if err := r.Complete(); err != nil {
		t.Errorf("error completing reminder: %s", err)
	}
}

func TestUncompleteReminder(t *testing.T) {
	n := NewReminder{
		List:  testList,
		Title: "Test Reminder to be Uncompleted",
	}

	r, err := AddReminder(n)
	if err != nil {
		t.Errorf("error adding reminder: %s", err)
	}

	time.Sleep(1 * time.Second) // just need time to see it appear
	if err := r.Complete(); err != nil {
		t.Errorf("error completing reminder: %s", err)
	}

	time.Sleep(1 * time.Second) // just need time to see it appear as completed
	if err := r.Uncomplete(); err != nil {
		t.Errorf("error uncompleting reminder: %s", err)
	}
}

func TestDeleteReminder(t *testing.T) {
	n := NewReminder{
		List:  testList,
		Title: "Test Reminder to be Deleted",
	}

	r, err := AddReminder(n)
	if err != nil {
		t.Errorf("error adding reminder: %s", err)
	}

	time.Sleep(1 * time.Second) // just need time to see it appear
	if err := r.Delete(); err != nil {
		t.Errorf("error deleting reminder: %s", err)
	}
}

func TestEditReminder(t *testing.T) {
	n := NewReminder{
		List:  testList,
		Title: "Test Reminder to be Edited",
	}

	r, err := AddReminder(n)
	if err != nil {
		t.Errorf("error adding reminder: %s", err)
	}

	time.Sleep(1 * time.Second) // just need time to see it appear

	e := EditArgs{
		Title: "Test Reminder Edited",
		Notes: "This is a test reminder that has been edited.",
	}

	if err := r.Edit(e); err != nil {
		t.Errorf("error editing reminder: %s", err)
	}
}

// TODO: ShowReminders test
