package execution

type (
	Query struct {
		Text       string                 `json:"text"`
		Params     map[string]interface{} `json:"params"`
		CDPAddress string                 `json:"cdp"`
	}
)
