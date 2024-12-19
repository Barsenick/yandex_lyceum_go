package tasks

import "time"

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (t Task) IsOverdue() (is bool) {
	return time.Now().After(t.deadline)
}

func (t Task) IsTopPriority() (is bool) {
	return t.priority >= 4
}
