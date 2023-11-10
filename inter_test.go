package main_test

import (
	"io"
	"reflect"
	"strings"
	"testing"

	inter "lesiw.io/inter"
)

func TestInter(t *testing.T) {
	reader1 := strings.NewReader("line1\nline2\nline3\n")
	reader2 := strings.NewReader("line2\nline3\nline4\n")
	reader3 := strings.NewReader("line1\nline2\nline3\n")

	result := inter.Inter([]io.Reader{reader1, reader2, reader3})

	expected := []string{"line2", "line3"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestInterWithEmptyReader(t *testing.T) {
	reader1 := strings.NewReader("line1\nline2\nline3\n")
	reader2 := strings.NewReader("")
	reader3 := strings.NewReader("line1\nline2\nline3\n")

	result := inter.Inter([]io.Reader{reader1, reader2, reader3})

	expected := []string{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestInterWithMissingNewline(t *testing.T) {
	reader1 := strings.NewReader("line1\nline2\nline3\n")
	reader2 := strings.NewReader("line2\nline3")
	reader3 := strings.NewReader("line1\nline2\nline3\n")

	result := inter.Inter([]io.Reader{reader1, reader2, reader3})

	expected := []string{"line2", "line3"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestInterWithSingleReader(t *testing.T) {
	reader := strings.NewReader("line1\nline2\nline3\n")

	result := inter.Inter([]io.Reader{reader})

	expected := []string{"line1", "line2", "line3"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestInterWithLeadingNewlines(t *testing.T) {
	reader := strings.NewReader("\n\nline1\nline2\nline3\n")

	result := inter.Inter([]io.Reader{reader})

	expected := []string{"line1", "line2", "line3"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestInterWithMultipleBlankLines(t *testing.T) {
	reader1 := strings.NewReader("\n\nline1\nline2\nline3\n")
	reader2 := strings.NewReader("\n\nline2\nline3\n")
	reader3 := strings.NewReader("\n\nline1\nline2\nline3\n")

	result := inter.Inter([]io.Reader{reader1, reader2, reader3})

	expected := []string{"line2", "line3"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
