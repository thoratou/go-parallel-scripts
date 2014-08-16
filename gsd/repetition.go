package gsd

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

type Repetition struct {
	minimum Minimum
	maximum Maximum
}

type Minimum struct {
	value Value
}

type Maximum struct {
	bound Bound
	value Value
}

type Bound uint

const (
	Infinite Bound = iota
	Finite   Bound = iota
)

type Value uint

func (self Value) String() string {
	return strconv.Itoa(int(self))
}

func NewRepetition(min uint) *Repetition {
	return &Repetition{
		Minimum{Value(min)},
		Maximum{
			Finite,
			Value(min),
		},
	}
}

func NewRepetitionMinMax(min uint, max uint) *Repetition {
	return &Repetition{
		Minimum{Value(min)},
		Maximum{
			Finite,
			Value(max),
		},
	}
}

func NewRepetitionInfiniteMax(min uint) *Repetition {
	return &Repetition{
		Minimum{Value(min)},
		Maximum{
			Infinite,
			Value(0),
		},
	}
}

func (self Repetition) String() string {
	result, _ := self.GetString()
	return result
}

func (self Repetition) GetString() (string, error) {
	if self.maximum.bound == Infinite {
		result := []string{self.minimum.value.String(), "..n"}
		return strings.Join(result, ""), nil
	}

	if self.minimum.value == self.maximum.value {
		return self.minimum.value.String(), nil
	}

	result := []string{self.minimum.value.String(), "..", self.maximum.value.String()}
	return strings.Join(result, ""), nil
}

var regexp_ = regexp.MustCompile(`([0-9]+)(..([0-9]+|[n]))?`)

func (self *Repetition) SetString(str string) error {
	subMatches := regexp_.FindStringSubmatch(str)
	if len(subMatches) != 4 || subMatches[0] != str {
		return NewError("No full match for repetition expression: ", str)
	}

	//the regular expression ensures atoi works
	min, _ := strconv.Atoi(subMatches[1])
	self.minimum.value = Value(min)

	if subMatches[3] == "n" {
		//case 0..n
		self.maximum.bound = Infinite
	} else if subMatches[3] == "" {
		//case 1
		self.maximum.bound = Finite
		self.maximum.value = Value(min)
	} else {
		//the regular expression ensures atoi works
		//case 0..1
		max, _ := strconv.Atoi(subMatches[3])
		self.maximum.bound = Finite
		self.maximum.value = Value(max)
	}

	return nil
}

func (self Repetition) MarshalJSON() ([]byte, error) {
	value, err := self.GetString()
	if err == nil {
		return json.Marshal(value)
	}
	return []byte{}, err
}

func (self *Repetition) UnmarshalJSON(buffer []byte) error {
	value := ""
	if err := json.Unmarshal(buffer, &value); err != nil {
		return err
	}
	return self.SetString(value)
}
