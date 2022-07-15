package util

import "testing"

type model struct {
	S       string `flag:"s"`
	More    bool   `flag:"more"`
	Ignored bool
}

var m model = model{
	S:       "A",
	More:    true,
	Ignored: true,
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

func TestStringifyOptions(t *testing.T) {
	options := GetOptions(m)
	args := StringifyOptions(options)
	expected := []string{"-s", m.S, "--more"}

	if len(args) != len(expected) {
		t.Fail()
	}

	eq := false

	for i := range args {
		eq = args[i] == expected[i]
	}

	if !eq {
		t.Fail()
	}
}
