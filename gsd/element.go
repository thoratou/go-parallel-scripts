package gsd

type Element struct {
	Name       Name        `json:"name"`
	Repetition Repetition  `json:"repetition"`
	Content    *Content    `json:"content,omitempty"`
	Text       *Text       `json:"text,omitempty"`
	Elements   []Element   `json:"element,omitempty"`
	Attributes []Attribute `json:"attribute,omitempty"`
}
