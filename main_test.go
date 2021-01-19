package main

import (
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {

	actual := injectSlashDIfneeded("go", "e:\\src")
	expected := "go"

	if !strings.Contains(actual, expected) {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}

func TestSame(t *testing.T) {

	actual := injectSlashDIfneeded("e:\\go", "e:\\src")
	expected := "e:\\go"

	if !strings.Contains(actual, expected) {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}

func TestDifferent(t *testing.T) {

	actual := injectSlashDIfneeded("c:\\go", "e:\\src")
	expected := "/d c:\\go"

	if !strings.Contains(actual, expected) {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}