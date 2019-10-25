package execution

import "time"

type (
	Query struct {
		Text       string                 `json:"text"`
		Params     map[string]interface{} `json:"params"`
		CDPAddress string                 `json:"cdp"`
	}

	Statistics struct {
		Compilation time.Duration
		Runtime     time.Duration
	}

	Result struct {
		Data  []byte
		Error error
		Stats Statistics
	}
)
