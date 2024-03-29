package main

import (
	"strings"
	"testing"
)

func TestProcessArgs(t *testing.T) {
	param := []string{"Program", "Files", "(x86)"}
	actual := processArgs(param)
	expected := "\"Program Files (x86)\""
	if actual != expected {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}

func TestCodePage(t *testing.T) {
	actual := GetConsoleCP()
	expected := uint32(1252)
	if actual != expected {
		t.Errorf("Expected: %d, but was: %d", expected, actual)
	}
}

func TestSimple(t *testing.T) {
	actual := addSwitchIfNeeded("go", "e:\\src")
	expected := "go"
	compare(actual, expected, t)
}

func TestWithSpace(t *testing.T) {
	actual := addSwitchIfNeeded("c:\\Program Files", "e:\\src")
	expected := "c:\\Program Files"
	compare(actual, expected, t)
}

func TestSameDrive(t *testing.T) {
	actual := addSwitchIfNeeded("e:\\go", "e:\\src")
	expected := "e:\\go"
	compare(actual, expected, t)
}

func TestDifferentDrive(t *testing.T) {
	actual := addSwitchIfNeeded("c:\\go", "e:\\src")
	expected := "/d \"c:\\go\""
	compare(actual, expected, t)
}

func compare(actual, expected string, t *testing.T) {
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}
