package main

import (
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	actual := addSwitchIfNeeded("go", "e:\\src")
	expected := "go"
	compare(actual, expected, t)
}

func TestSame(t *testing.T) {
	actual := addSwitchIfNeeded("e:\\go", "e:\\src")
	expected := "e:\\go"
	compare(actual, expected, t)
}

func TestDifferent(t *testing.T) {
	actual := addSwitchIfNeeded("c:\\go", "e:\\src")
	expected := "/d c:\\go"
	compare(actual, expected, t)
}

func TestQuote(t *testing.T) {
	actual := quote("c:\\Program Files")
	expected := "\"c:\\Program Files\""
	compare(actual, expected, t)
}

func TestNoQuote(t *testing.T) {
	actual := quote("c:\\go")
	expected := "c:\\go"
	compare(actual, expected, t)
}

func compare(actual, expected string, t *testing.T) {
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}
