package gitign

import (
	"fmt"
	"strings"
)

func optimizeGitignore(gitignore *string) error {
	// Split the content by lines
	lines := strings.Split(string(*gitignore), "\n")
	seen := make(map[string]bool)
	var optimizedLines []string

	// Iterate over lines, only keep the first occurrence of each line
	for _, line := range lines {
		// Skip empty lines and comments
		trimedLine := strings.TrimSpace(line)
		if seen[line] && !strings.HasPrefix(trimedLine, "#") {
			println("Duplicate rule found:", line)
			continue
		}
		optimizedLines = append(optimizedLines, line)
		if !strings.HasPrefix(trimedLine, "#") && trimedLine != "" {
			seen[line] = true
		}
	}

	// Join optimized lines and write back to the file
	*gitignore = strings.Join(optimizedLines, "\n")
	// err := os.WriteFile(".gitignore", []byte(optimizedContent), 0644)
	// if err != nil {
	// 	return fmt.Errorf("failed to write optimized .gitignore: %w", err)
	// }

	fmt.Println("Optimization complete.")
	return nil
}
