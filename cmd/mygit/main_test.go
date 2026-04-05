package main

import (
	"os"
	"path/filepath"
	"testing"
)

// TestInitRepo verifies that initRepo creates the expected directory structure
// and writes the correct HEAD file in a temporary working directory.
func TestInitRepo(t *testing.T) {
	// Run the test in an isolated temporary directory so it does not touch the
	// real repository's .git folder.
	tmpDir := t.TempDir()
	origDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("chdir to tmpDir: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(origDir) })

	initRepo()

	// Verify that all required directories were created.
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		info, err := os.Stat(filepath.Join(tmpDir, dir))
		if err != nil {
			t.Errorf("expected directory %q to exist: %v", dir, err)
			continue
		}
		if !info.IsDir() {
			t.Errorf("expected %q to be a directory", dir)
		}
	}

	// Verify the HEAD file contents.
	head, err := os.ReadFile(filepath.Join(tmpDir, ".git", "HEAD"))
	if err != nil {
		t.Fatalf("reading .git/HEAD: %v", err)
	}
	want := "ref: refs/heads/main\n"
	if string(head) != want {
		t.Errorf(".git/HEAD = %q, want %q", string(head), want)
	}
}

// TestInitRepoIdempotent ensures that calling initRepo twice does not produce
// an error (the directories already exist on the second call).
func TestInitRepoIdempotent(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("chdir to tmpDir: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(origDir) })

	// Calling twice should not panic or overwrite with wrong data.
	initRepo()
	initRepo()

	head, err := os.ReadFile(filepath.Join(tmpDir, ".git", "HEAD"))
	if err != nil {
		t.Fatalf("reading .git/HEAD after second init: %v", err)
	}
	want := "ref: refs/heads/main\n"
	if string(head) != want {
		t.Errorf(".git/HEAD = %q, want %q", string(head), want)
	}
}
