package dal

type Entity struct {
	Metadata
	ID string `json:"id"`
}

func (e Entity) IsEmpty() bool {
	return e.ID == ""
}
