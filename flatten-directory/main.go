package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	dryRun            bool
	directory         string
	includeParentDir  bool
}

func log(message string) {
	fmt.Printf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}

func moveFile(path string, config Config) error {
	// Skip hidden files
	base := filepath.Base(path)
	if strings.HasPrefix(base, ".") {
		return nil
	}

	// Skip if already in root
	dir := filepath.Dir(path)
	if dir == config.directory {
		return nil
	}

	// Create new filename, optionally including parent directory name
	newBase := base
	if config.includeParentDir {
		parentDir := filepath.Base(filepath.Dir(path))
		if parentDir != config.directory {
			ext := filepath.Ext(base)
			name := strings.TrimSuffix(base, ext)
			newBase = fmt.Sprintf("%s_%s%s", name, parentDir, ext)
		}
	}

	// Create new filename with counter if needed
	newPath := filepath.Join(config.directory, newBase)
	counter := 1
	for {
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			break
		}
		ext := filepath.Ext(newBase)
		name := strings.TrimSuffix(newBase, ext)
		newPath = filepath.Join(config.directory, fmt.Sprintf("%s_%d%s", name, counter, ext))
		counter++
	}

	if config.dryRun {
		log(fmt.Sprintf("Would move: %s -> %s", path, newPath))
		return nil
	}

	if err := os.Rename(path, newPath); err != nil {
		return fmt.Errorf("error moving file: %w", err)
	}
	log(fmt.Sprintf("Moved: %s -> %s", path, newPath))
	return nil
}

func main() {
	config := Config{}
	flag.BoolVar(&config.dryRun, "d", false, "Dry run mode")
	flag.BoolVar(&config.includeParentDir, "p", false, "Include parent directory name in filename")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Error: Directory argument is required")
		fmt.Printf("Usage: %s [-d] [-p] <directory>\n", os.Args[0])
		fmt.Println("  -d    Dry run mode")
		fmt.Println("  -p    Include parent directory name in filename")
		os.Exit(1)
	}

	config.directory = args[0]

	// Validate directory
	if info, err := os.Stat(config.directory); err != nil || !info.IsDir() {
		fmt.Printf("Error: '%s' is not a directory or does not exist\n", config.directory)
		os.Exit(1)
	}

	log(fmt.Sprintf("Starting file reorganization in directory: %s", config.directory))

	err := filepath.WalkDir(config.directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && d.Name() != ".DS_Store" {
			if err := moveFile(path, config); err != nil {
				log(fmt.Sprintf("Failed to process: %s - %v", path, err))
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}

	log("Finished file reorganization")
}
