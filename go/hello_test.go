package main

import (
    "testing"
    "regexp"
)

func TestHelloName(t *testing.T) {
    name := "Code.education Rocks!"
    want := regexp.MustCompile(`<b>`+name+`</b>`)
    msg, err := greeting("Code.education Rocks!")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Code.education Rocks!") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

func TestHelloEmpty(t *testing.T) {
    msg, err := greeting("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}