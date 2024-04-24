package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Piero"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Piero")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia %q, nil `, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
