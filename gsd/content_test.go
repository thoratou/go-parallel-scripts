package gsd

import (
	"encoding/json"
	"testing"
)

func TestContent(t *testing.T) {
	{
		content := Extension
		value, err := content.GetString()

		if err != nil {
			t.Fatal("GetString on NoContent should be supported, error:", err)
		}

		if value != "extension" {
			t.Fatal("String should be empty for NoContent")
		}
	}

	{
		content := Extension
		err := content.SetString("")

		if err == nil {
			t.Fatal("SetString with empty should nott be supported, error:", err)
		}

		if content != Extension {
			t.Fatal("Content should stay to Extension")
		}
	}

	{
		content := Extension
		err := content.SetString("dummy")

		if err == nil {
			t.Fatal("SetString with extension should be supported, error:", err)
		} else {
			t.Log("returned error: ", err)
		}
	}

}

func TestEncodeContent(t *testing.T) {
	{
		grammarStr := `
		{
			"root":
			{
				"name":"component",
				"repetition":"1",
				"content":null
			}
		}
		`
		grammar := Grammar{}

		if err := json.Unmarshal([]byte(grammarStr), &grammar); err != nil {
			t.Fatal("cannot decode grammar: ", err)
		}

		if grammar.Root.Content != nil {
			t.Fatal("content not decoded")
		}

		result, err := json.Marshal(grammar)

		if err != nil {
			t.Fatal("Unable to encode: ", err)
		}

		s := string(result[:])
		t.Log(s)
	}
	{
		grammarStr := `
		{
			"root":
			{
				"name":"component",
				"repetition":"1"
			}
		}
		`
		grammar := Grammar{}

		if err := json.Unmarshal([]byte(grammarStr), &grammar); err != nil {
			t.Fatal("cannot decode grammar: ", err)
		}

		if grammar.Root.Content != nil {
			t.Fatal("content not decoded")
		}

		result, err := json.Marshal(grammar)

		if err != nil {
			t.Fatal("Unable to encode: ", err)
		}

		s := string(result[:])
		t.Log(s)
	}
	{
		grammarStr := `
		{
			"root":
			{
				"name":"component",
				"repetition":"1",
				"content":"extension"
			}
		}
		`
		grammar := Grammar{}

		if err := json.Unmarshal([]byte(grammarStr), &grammar); err != nil {
			t.Fatal("cannot decode grammar: ", err)
		}

		if grammar.Root.Content == nil {
			t.Fatal("content not decoded")
		}

		result, err := json.Marshal(grammar)

		if err != nil {
			t.Fatal("Unable to encode: ", err)
		}

		s := string(result[:])
		t.Log(s)

	}
	{
		grammarStr := `
		{
			"root":
			{
				"name":"component",
				"repetition":"1",
				"content":"dummy"
			}
		}
		`
		grammar := Grammar{}

		if err := json.Unmarshal([]byte(grammarStr), &grammar); err == nil {
			t.Fatal("Error should have been returned")
		} else {
			t.Log(err)
		}

	}
}
