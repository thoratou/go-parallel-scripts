package gsd

type Attribute struct {
	Name       Name       `json:"name"`
	Repetition Repetition `json:"repetition"`
	Type       Type       `json:"type"`
}
