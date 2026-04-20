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

	folders, err := collectTargetFolders(rootDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to collect folders: %v\n", err)
		os.Exit(1)
	}

	if len(folders) == 0 {
		fmt.Println("no folders to process")
		return
	}

	outDir := filepath.Join(rootDir, "symbols")
	if stat, err := os.Stat(outDir); err != nil || !stat.IsDir() {
		fmt.Fprintf(os.Stderr, "symbols directory not found: %s\n", outDir)
		os.Exit(1)
	}

	var failed []string
	for _, folder := range folders {
		fmt.Printf("Processing: %s\n", folder)

		pkgPath := modulePrefix + "/" + folder
		cmd := exec.Command("yaegi", "extract", pkgPath)
		cmd.Dir = outDir
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

func findProjectRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		goMod := filepath.Join(wd, "go.mod")
		if _, err := os.Stat(goMod); err == nil {
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

	excluded := map[string]struct{}{
		"tests":   {},
		"symbols": {},
		"doc":     {},
		"cmd":     {},
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
		if _, skip := excluded[name]; skip {
			continue
		}

		folders = append(folders, name)
	}

	sort.Strings(folders)
	return folders, nil
}
