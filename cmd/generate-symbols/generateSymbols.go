//go:generate go run .

package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

const modulePrefix = "yanlingrpa.com/yanling/protocol"

// excludedDirs are skipped during folder collection.
var excludedDirs = map[string]struct{}{
	".yanling": {},
	"cmd":      {},
	"doc":      {},
	"tests":    {},
	"symbols":  {},
	"schema":   {},
}

func main() {
	rootDir, err := findProjectRoot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to locate project root: %v\n", err)
		os.Exit(1)
	}

	if _, err := exec.LookPath("yaegi"); err != nil {
		fmt.Fprintln(os.Stderr, "yaegi not found in PATH")
		os.Exit(1)
	}

	targetFolders, err := collectTargetFolders(rootDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to collect folders: %v\n", err)
		os.Exit(1)
	}

	if len(targetFolders) == 0 {
		fmt.Println("no folders to process")
		return
	}

	symbolsDir := filepath.Join(rootDir, "symbols")
	if stat, err := os.Stat(symbolsDir); err != nil || !stat.IsDir() {
		fmt.Fprintf(os.Stderr, "symbols directory not found: %s\n", symbolsDir)
		os.Exit(1)
	}

	var failed []string
	for _, folder := range targetFolders {
		pkgPath := modulePrefix + "/" + folder
		fmt.Printf("Processing: %s\n", pkgPath)

		cmd := exec.Command("yaegi", "extract", pkgPath)
		cmd.Dir = symbolsDir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("  Failed: %s (%v)\n", folder, err)
			failed = append(failed, folder)
			continue
		}

		fmt.Printf("  Success: extracted %s\n", folder)
	}

	fmt.Println("\nAll done!")
	if len(failed) > 0 {
		fmt.Fprintf(os.Stderr, "generation failed for %d folder(s): %s\n", len(failed), strings.Join(failed, ", "))
		os.Exit(1)
	}
}

// findProjectRoot walks up from cwd until go.mod is found.
func findProjectRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd, nil
		}
		parent := filepath.Dir(wd)
		if parent == wd {
			return "", errors.New("go.mod not found")
		}
		wd = parent
	}
}
func collectTargetFolders(rootDir string) ([]string, error) {
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}

	folders := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		if _, skip := excludedDirs[name]; skip {
			continue
		}

		folders = append(folders, name)
	}

	sort.Strings(folders)
	return folders, nil
}
