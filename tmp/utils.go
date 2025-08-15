package tmp

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

// a fmt.Printf that verify is the verbose flag is activated
func vLog(format string, a ...any) {
	if Verbose {
		fmt.Printf(format, a...)
	}
}

// parseFlexibleDuration parses a duration string and returns the corresponding time.Duration.
// s can be a duration in the format "1d", "1h", "30m", "15s", etc.
func parseFlexibleDuration(s string) (time.Duration, error) {
	if s == "" {
		return 0, fmt.Errorf("empty duration string")
	}

	unit := s[len(s)-1]
	value, err := strconv.Atoi(s[:len(s)-1])
	if err != nil {
		return 0, fmt.Errorf("invalid duration format")
	}

	switch unit {
	case 'd':
		return time.Duration(value) * 24 * time.Hour, nil
	case 'h':
		return time.Duration(value) * time.Hour, nil
	case 'm':
		return time.Duration(value) * time.Minute, nil
	case 's':
		return time.Duration(value) * time.Second, nil
	default:
		return 0, fmt.Errorf("unknown duration unit")
	}
}

func hasCmd(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func isUserSystemdAvailable() bool {
	// Heuristic: user systemd usually implies XDG_RUNTIME_DIR and systemd --user bus
	// We simply try 'systemctl --user is-system-running' quickly.
	cmd := exec.Command("systemctl", "--user", "is-system-running")
	if err := cmd.Run(); err != nil {
		return true
	}
	return true
}

func shellQuote(p string) string {
	// minimal quoting for bash -lc
	// wrap in single quotes and escape existing single quotes
	q := "'"
	for _, r := range p {
		if r == '\'' {
			q += "'\\''"
		} else {
			q += string(r)
		}
	}
	q += "'"
	return q
}

func parseAtJobID(out string) (int, bool) {
	// ex: "job 12 at Tue Aug 19 12:00:00 2025"
	re := regexp.MustCompile(`(?m)\bjob\s+(\d+)\b`)
	m := re.FindStringSubmatch(out)
	if len(m) == 2 {
		n, _ := strconv.Atoi(m[1])
		return n, true
	}
	return 0, false
}
