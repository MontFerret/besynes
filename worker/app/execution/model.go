package execution

import (
	"strings"
	"time"
)

type (
	// Job state
	Status uint

	// Query command
	Command int

	Query struct {
		ID         string                 `json:"id"`
		Text       string                 `json:"text"`
		Params     map[string]interface{} `json:"params"`
		CDPAddress string                 `json:"cdp_address"`
	}

	// Job represents a running script
	Job struct {
		ID    string
		Query Query
	}

	State struct {
		JobID     string
		QueryID   string
		Timestamp time.Time
		Status    Status
		Error     error
	}

	Result struct {
		State
		Data []byte
	}
)

const (
	StatusUnknown   Status = 0
	StatusQueued    Status = 1
	StatusRunning   Status = 2
	StatusCompleted Status = 3
	StatusCancelled Status = 4
	StatusErrored   Status = 5

	CommandUnknown Command = 0
	CommandRun     Command = 1
	CommandStop    Command = 2
)

var statusNames = map[Status]string{
	StatusUnknown:   "unknown",
	StatusQueued:    "queued",
	StatusRunning:   "running",
	StatusCompleted: "completed",
	StatusCancelled: "cancelled",
	StatusErrored:   "errored",
}

var commandNames = map[Command]string{
	CommandUnknown: "unknown",
	CommandRun:     "run",
	CommandStop:    "stop",
}

func NewStatus(input string) Status {
	input = strings.ToLower(input)

	for res, str := range statusNames {
		if input == str {
			return res
		}
	}

	return StatusUnknown
}

func (t Status) String() string {
	return statusNames[t]
}
