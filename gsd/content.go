package gsd

import (
	"encoding/json"
)

type Content int

const (
	Extension Content = iota
)

func NewContent(content Content) *Content {
	result := new(Content)
	*result = content
	return result
}

func (self Content) String() string {
	value, _ := self.GetString()
	return value
}

func (self Content) GetString() (string, error) {
	switch self {
	case Extension:
		return "extension", nil
	default:
		return "", NewError("unsupported content type:", self)
	}
}

func (self *Content) SetString(value string) error {
	switch value {
	case "extension":
		*self = Extension
		return nil
	default:
		return NewError("unsupported content value:", value)
	}
}

func (self Content) MarshalJSON() ([]byte, error) {
	value, err := self.GetString()
	if err == nil {
		return json.Marshal(value)
	}
	return []byte{}, err
}

func (self *Content) UnmarshalJSON(buffer []byte) error {
	value := ""
	if err := json.Unmarshal(buffer, &value); err != nil {
		return err
	}
	return self.SetString(value)
}
