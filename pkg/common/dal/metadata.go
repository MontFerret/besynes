package dal

import "time"

type Metadata struct {
	UpdateAt time.Time `json:"updated_at"`
}
