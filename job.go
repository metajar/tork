package tork

import (
	"context"
	"time"

	"github.com/runabol/tork/internal/clone"
)

type JobState string

const (
	JobStatePending   JobState = "PENDING"
	JobStateRunning   JobState = "RUNNING"
	JobStateCancelled JobState = "CANCELLED"
	JobStateCompleted JobState = "COMPLETED"
	JobStateFailed    JobState = "FAILED"
	JobStateRestart   JobState = "RESTART"
)

type JobHandler func(context.Context, *Job) error

type Job struct {
	ID          string            `json:"id,omitempty"`
	ParentID    string            `json:"parentId,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	State       JobState          `json:"state,omitempty"`
	CreatedAt   time.Time         `json:"createdAt,omitempty"`
	StartedAt   *time.Time        `json:"startedAt,omitempty"`
	CompletedAt *time.Time        `json:"completedAt,omitempty"`
	FailedAt    *time.Time        `json:"failedAt,omitempty"`
	Tasks       []*Task           `json:"tasks"`
	Execution   []*Task           `json:"execution"`
	Position    int               `json:"position"`
	Inputs      map[string]string `json:"inputs,omitempty"`
	Context     JobContext        `json:"context,omitempty"`
	TaskCount   int               `json:"taskCount,omitempty"`
	Output      string            `json:"output,omitempty"`
	Result      string            `json:"result,omitempty"`
	Error       string            `json:"error,omitempty"`
}

type JobContext struct {
	Inputs map[string]string `json:"inputs,omitempty"`
	Tasks  map[string]string `json:"tasks,omitempty"`
}

func (j *Job) Clone() *Job {
	return &Job{
		ID:          j.ID,
		Name:        j.Name,
		Description: j.Description,
		State:       j.State,
		CreatedAt:   j.CreatedAt,
		StartedAt:   j.StartedAt,
		CompletedAt: j.CompletedAt,
		FailedAt:    j.FailedAt,
		Tasks:       CloneTasks(j.Tasks),
		Execution:   CloneTasks(j.Execution),
		Position:    j.Position,
		Inputs:      j.Inputs,
		Context:     j.Context.Clone(),
		ParentID:    j.ParentID,
		TaskCount:   j.TaskCount,
		Output:      j.Output,
		Result:      j.Result,
		Error:       j.Error,
	}
}

func (c JobContext) Clone() JobContext {
	return JobContext{
		Inputs: clone.CloneStringMap(c.Inputs),
		Tasks:  clone.CloneStringMap(c.Tasks),
	}
}

func (c JobContext) AsMap() map[string]any {
	return map[string]any{
		"inputs": c.Inputs,
		"tasks":  c.Tasks,
	}
}
