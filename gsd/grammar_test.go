package gsd

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestGrammarDecoding(t *testing.T) {

	var grammarStr = `
	{
		"root":
		{
			"name":"component",
			"repetition":"1",
			"attribute":
			[
				{
					"name":"version",
					"repetition":"0..1",
					"type":"uint"
				}
			],
			"element":
			[
				{
					"name":"properties",
					"repetition":"1",
					"element":
					[
						{
							"name":"name",
							"repetition":"1",
							"text":
							{
								"repetition":"1",
								"type":"string"
							}
						},
						{
							"name":"version",
							"repetition":"1",
							"text":
							{
								"repetition":"1",
								"type":"string"
							}
						}
					]
				},
				{
					"name":"target",
					"repetition":"0..n",
					"content":"extension",
					"attribute":
					[
						{
							"name":"name",
							"repetition":"1",
							"type":"string"
						},
						{
							"name":"type",
							"repetition":"1",
							"type":"string"
						}
					]					
				}
			]
		}
	}
	`

	grammar := new(Grammar)

	reader := strings.NewReader(grammarStr)

	//decoding
	if err := json.NewDecoder(reader).Decode(&grammar); err != nil {
		t.Fatal("cannot decode grammar: ", err)
	}

	/****************************
	 * Some data structure test *
	 ****************************/

	if grammar.Root.Name != "component" {
		t.Fatal("grammar.Root.Name found:", grammar.Root.Name, " expected: ", "component")
	}

	if grammar.Root.Repetition != *NewRepetition(1) {
		t.Fatal("grammar.Root.Repetition found:", grammar.Root.Repetition, " expected: ", "1")
	}

	if len(grammar.Root.Attributes) != 1 {
		t.Fatal("found", len(grammar.Root.Attributes), "attribute(s) for ", grammar.Root)
	}

	if grammar.Root.Attributes[0].Name != "version" {
		t.Fatal("grammar.Root.Attributes[0].Name found:", grammar.Root.Attributes[0].Name, " expected: ", "version")
	}

	if grammar.Root.Attributes[0].Repetition != *NewRepetitionMinMax(0, 1) {
		t.Fatal("grammar.Root.Attributes[0].Repetition found:", grammar.Root.Attributes[0].Repetition, " expected: ", "0..1")
	}

	if grammar.Root.Attributes[0].Type != UIntType {
		t.Fatal("grammar.Root.Attributes[0].Type found:", grammar.Root.Attributes[0].Type, " expected: ", "uint")
	}

	if len(grammar.Root.Elements) != 2 {
		t.Fatal("found", len(grammar.Root.Elements), "element(s) for ", grammar.Root)
	}

	if grammar.Root.Elements[0].Name != "properties" {
		t.Fatal("grammar.Root.Elements[0].Name found:", grammar.Root.Elements[0].Name, " expected: ", "properties")
	}

	if grammar.Root.Elements[0].Repetition != *NewRepetition(1) {
		t.Fatal("grammar.Root.Elements[0].Repetition found:", grammar.Root.Elements[0].Repetition, " expected: ", "1")
	}

	if grammar.Root.Elements[0].Content != nil {
		t.Fatal("grammar.Root.Elements[0].Content found:", grammar.Root.Elements[0].Content, " expected empty")
	}

	if grammar.Root.Elements[0].Text != nil {
		t.Fatal("grammar.Root.Elements[0].Text found:", grammar.Root.Elements[0].Text, " expected empty")
	}

	/*********************
	 * Full content test *
	 *********************/

	grammarContent := &Grammar{
		Root: Element{
			Name:       "component",
			Repetition: *NewRepetition(1),
			Elements: []Element{
				{
					Name:       "properties",
					Repetition: *NewRepetition(1),
					Elements: []Element{
						{
							Name:       "name",
							Repetition: *NewRepetition(1),
							Text: &Text{
								Repetition: *NewRepetition(1),
								Type:       StringType,
							},
						},
						{
							Name:       "version",
							Repetition: *NewRepetition(1),
							Text: &Text{
								Repetition: *NewRepetition(1),
								Type:       StringType,
							},
						},
					},
				},
				{
					Name:       "target",
					Repetition: *NewRepetitionInfiniteMax(0),
					Content:    NewContent(Extension),
					Attributes: []Attribute{
						{
							Name:       "name",
							Repetition: *NewRepetition(1),
							Type:       StringType,
						},
						{
							Name:       "type",
							Repetition: *NewRepetition(1),
							Type:       StringType,
						},
					},
				},
			},
			Attributes: []Attribute{
				{
					Name:       "version",
					Repetition: *NewRepetitionMinMax(0, 1),
					Type:       UIntType,
				},
			},
		},
	}

	first, second := "", ""
	fmt.Sprint(grammar, first)
	fmt.Sprint(grammar, second)

	//no equality for arrays :(
	//if grammar != grammarContent {
	if first != second {
		t.Fatal("diff: ", grammar, " / ", grammarContent)
	}

}
