package gsd

import (
	"encoding/json"
)

type Type int

const (
	UnknownType Type = iota
	StringType  Type = iota
	BoolType    Type = iota
	IntType     Type = iota
	Int8Type    Type = iota
	Int16Type   Type = iota
	Int32Type   Type = iota
	Int64Type   Type = iota
	UIntType    Type = iota
	UInt8Type   Type = iota
	UInt16Type  Type = iota
	UInt32Type  Type = iota
	UInt64Type  Type = iota
	Float32Type Type = iota
	Float64Type Type = iota
)

func (self Type) String() string {
	value, _ := self.GetString()
	return value
}

func (self Type) GetString() (string, error) {
	switch self {
	case StringType:
		return "string", nil
	case BoolType:
		return "bool", nil
	case IntType:
		return "int", nil
	case Int8Type:
		return "int8", nil
	case Int16Type:
		return "int16", nil
	case Int32Type:
		return "int32", nil
	case Int64Type:
		return "int64", nil
	case UIntType:
		return "uint", nil
	case UInt8Type:
		return "uint8", nil
	case UInt16Type:
		return "uint16", nil
	case UInt32Type:
		return "uint32", nil
	case UInt64Type:
		return "uint64", nil
	case Float32Type:
		return "float32", nil
	case Float64Type:
		return "float64", nil
	case UnknownType:
		return "", nil
	}

	return "", NewError("Unsupported Type value")
}

func (self *Type) SetString(str string) error {
	switch str {
	case "string":
		*self = StringType
		return nil
	case "bool":
		*self = BoolType
		return nil
	case "int":
		*self = IntType
		return nil
	case "int8":
		*self = Int8Type
		return nil
	case "int16":
		*self = Int16Type
		return nil
	case "int32":
		*self = Int32Type
		return nil
	case "int64":
		*self = Int64Type
		return nil
	case "uint":
		*self = UIntType
		return nil
	case "uint8":
		*self = UInt8Type
		return nil
	case "uint16":
		*self = UInt16Type
		return nil
	case "uint32":
		*self = UInt32Type
		return nil
	case "uint64":
		*self = UInt64Type
		return nil
	case "float32":
		*self = Float32Type
		return nil
	case "float64":
		*self = Float64Type
		return nil
	}

	return NewError("Unsupported Type string", str)
}

func (self Type) MarshalJSON() ([]byte, error) {
	value, err := self.GetString()
	if err == nil {
		return json.Marshal(value)
	}
	return []byte{}, err
}

func (self *Type) UnmarshalJSON(buffer []byte) error {
	value := ""
	if err := json.Unmarshal(buffer, &value); err != nil {
		return err
	}
	return self.SetString(value)
}
