package main

import (
	"reflect"
	"testing"
)

func TestLastCommitArgs(t *testing.T) {
	got := lastCommitArgs()
	want := []string{"reset", "--hard", "HEAD~1"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("lastCommitArgs() = %v, want %v", got, want)
	}
}

func TestFileArgs(t *testing.T) {
	got := fileArgs("main.go")
	want := []string{"checkout", "--", "main.go"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("fileArgs() = %v, want %v", got, want)
	}
}

func TestUnstageArgs(t *testing.T) {
	got := unstageArgs("main.go")
	want := []string{"reset", "HEAD", "main.go"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("unstageArgs() = %v, want %v", got, want)
	}
}

func TestRevertArgs(t *testing.T) {
	got := revertArgs("abc123")
	want := []string{"revert", "abc123"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("revertArgs() = %v, want %v", got, want)
	}
}

func TestLogArgs(t *testing.T) {
	got := logArgs()
	want := []string{"log", "--oneline", "-n", "10"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("logArgs() = %v, want %v", got, want)
	}
}