package concat

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var OutputFile string

func processFile(file string, outputFile *os.File) error {
	_, err := fmt.Fprintf(outputFile, "### %s ###\n", file)
	if err != nil {
		return err
	}

	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(outputFile, "%s\n", content)
	return err
}

func Concat(globPattern []string) error {
	outputFile, err := os.OpenFile(OutputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open output file %s: %w", OutputFile, err)
	}
	defer outputFile.Close()

	for _, pattern := range globPattern {
		files, err := filepath.Glob(pattern)
		if err != nil {
			return fmt.Errorf("failed to glob files with pattern %s: %w", pattern, err)
		}

		if len(files) == 0 {
			return fmt.Errorf("no files found for pattern %s", pattern)
		}

		for _, file := range files {
			info, err := os.Stat(file)
			if err != nil {
				return fmt.Errorf("failed to stat file %s: %w", file, err)
			}

			if info.IsDir() {
				err := filepath.Walk(file, func(path string, info fs.FileInfo, err error) error {
					if err != nil {
						return fmt.Errorf("error walking path %s: %w", path, err)
					}

					if !info.IsDir() {
						err := processFile(path, outputFile)
						if err != nil {
							return fmt.Errorf("failed to process file %s: %w", path, err)
						}
					}

					return nil
				})
				if err != nil {
					return fmt.Errorf("error walking directory %s: %w", file, err)
				}
			} else {
				err := processFile(file, outputFile)
				if err != nil {
					return fmt.Errorf("failed to process file %s: %w", file, err)
				}
			}
		}
	}

	return nil
}
