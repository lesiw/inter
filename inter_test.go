package main_test

import (
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	inter "lesiw.io/inter"
)

func TestInter(t *testing.T) {
	var (
		reader1  = strings.NewReader("line1\nline2\nline3\n")
		reader2  = strings.NewReader("line2\nline3\nline4\n")
		reader3  = strings.NewReader("line1\nline2\nline3\n")
		result   = inter.Inter([]io.Reader{reader1, reader2, reader3})
		expected = []string{"line2", "line3"}
	)

	if !cmp.Equal(expected, result) {
		t.Errorf(
			"Inter() mismatch: -want +got\n%s", cmp.Diff(expected, result),
		)
	}
}

func TestInterWithEmptyReader(t *testing.T) {
	var (
		reader1  = strings.NewReader("line1\nline2\nline3\n")
		reader2  = strings.NewReader("")
		reader3  = strings.NewReader("line1\nline2\nline3\n")
		result   = inter.Inter([]io.Reader{reader1, reader2, reader3})
		expected = []string{}
	)

	if !cmp.Equal(expected, result) {
		t.Errorf(
			"Inter() mismatch: -want +got\n%s", cmp.Diff(expected, result),
		)
	}
}

func TestInterWithMissingNewline(t *testing.T) {
	var (
		reader1  = strings.NewReader("line1\nline2\nline3\n")
		reader2  = strings.NewReader("line2\nline3")
		reader3  = strings.NewReader("line1\nline2\nline3\n")
		result   = inter.Inter([]io.Reader{reader1, reader2, reader3})
		expected = []string{"line2", "line3"}
	)

	if !cmp.Equal(expected, result) {
		t.Errorf(
			"Inter() mismatch: -want +got\n%s", cmp.Diff(expected, result),
		)
	}
}

func TestInterWithSingleReader(t *testing.T) {
	var (
		reader   = strings.NewReader("line1\nline2\nline3\n")
		result   = inter.Inter([]io.Reader{reader})
		expected = []string{"line1", "line2", "line3"}
	)

	if !cmp.Equal(expected, result) {
		t.Errorf(
			"Inter() mismatch: -want +got\n%s", cmp.Diff(expected, result),
		)
	}
}

func TestInterWithLeadingNewlines(t *testing.T) {
	var (
		reader   = strings.NewReader("\n\nline1\nline2\nline3\n")
		result   = inter.Inter([]io.Reader{reader})
		expected = []string{"line1", "line2", "line3"}
	)

	if !cmp.Equal(expected, result) {
		t.Errorf(
			"Inter() mismatch: -want +got\n%s", cmp.Diff(expected, result),
		)
	}
}

func TestInterWithMultipleBlankLines(t *testing.T) {
	var (
		reader1  = strings.NewReader("\n\nline1\nline2\nline3\n")
		reader2  = strings.NewReader("\n\nline2\nline3\n")
		reader3  = strings.NewReader("\n\nline1\nline2\nline3\n")
		result   = inter.Inter([]io.Reader{reader1, reader2, reader3})
		expected = []string{"line2", "line3"}
	)

	if !cmp.Equal(expected, result) {
		t.Errorf(
			"Inter() mismatch: -want +got\n%s", cmp.Diff(expected, result),
		)
	}
}
