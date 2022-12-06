package util

import "testing"

type model struct {
	S        string `flag:"s"`
	More     bool   `flag:"more"`
	Ignored  bool
	Argument string `arg:"the value is blue"`
}

var m model = model{
	S:        "A",
	More:     true,
	Ignored:  true,
	Argument: "blue",
}

func TestGetOptions(t *testing.T) {
	options := GetOptions(m)

	if _, ok := options["s"]; !ok {
		t.Fatalf("Flag `s` not found.")
	}

	if _, ok := options["more"]; !ok {
		t.Fatalf("Flag `more` not found.")
	}

	if v, ok := options["ignored"]; ok {
		t.Fatalf("Found a key that should be ignored with value %s", v)
	}
}

func TestGetArguments(t *testing.T) {
	arguments := GetArguments(m)
	expected := []string{"blue"}

	t.Log(arguments, expected)

	eq := false

	if len(arguments) != 1 {
		t.Fail()
	}

	for i := range arguments {
		eq = arguments[i] == expected[i]
	}

	if !eq {
		t.Fail()
	}
}

func TestStringifyOptions(t *testing.T) {
	options := GetOptions(m)
	args := StringifyOptions(options)
	expected := []string{"-s", m.S, "--more"}

	if len(args) != len(expected) {
		t.Logf("data: %v", args)
		t.Fatalf("Invalid length received, expected %d received %d", len(expected), len(args))
	}

	eq := false

	for i := range args {
		eq = args[i] == expected[i]
	}

	if !eq {
		t.Fail()
	}
}
