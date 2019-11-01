package dal

import "time"

type Metadata struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  *time.Time `json:"updated_at"`
}
