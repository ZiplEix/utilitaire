package tmp

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func TmpDir(directoryPath string, expiration string) error {
	switch {
	case directoryPath == "":
		directoryPath = filepath.Join("/tmp/utilitaire", time.Now().Format("20060102150405"))
	case filepath.IsAbs(directoryPath):
	default:
		if strings.HasPrefix(directoryPath, "./") {
		} else {
			directoryPath = filepath.Join("/tmp/utilitaire", directoryPath)
		}
	}

	absDir, err := filepath.Abs(directoryPath)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(absDir, 0755); err != nil {
		return err
	}

	if err := setExpiration(absDir, expiration); err != nil {
		return err
	}

	return nil
}

func setExpiration(directoryPath string, expiration string) error {
	if expiration == "" {
		expiration = "24h"
	}

	duration, err := parseFlexibleDuration(expiration)
	if err != nil {
		return err
	}
	expAt := time.Now().Add(duration)

	if hasCmd("systemd-run") && isUserSystemdAvailable() {
		label := fmt.Sprintf("tmp-delete-%d", time.Now().UnixNano())
		cmd := exec.Command(
			"systemd-run", "--user",
			"--unit", label,
			"--on-active="+duration.String(),
			"/usr/bin/env", "bash", "-lc", "rm -rf -- "+shellQuote(directoryPath),
		)
		if out, err := cmd.CombinedOutput(); err == nil {
			vLog("Created %s\n", directoryPath)
			vLog("Deletion scheduled via systemd in %s (unit: %s)\n", duration, label)
			fmt.Printf("%s", directoryPath)
			return addRecord(Record{
				Path:       directoryPath,
				Scheduler:  SchedSystemd,
				Unit:       label,
				CreatedAt:  time.Now(),
				Expiration: expAt,
			})
		} else {
			// utile pour debug (par ex. user manager inactif)
			fmt.Fprintf(os.Stderr, "systemd-run failed: %v\nOutput: %s\n", err, string(out))
		}
	}

	// Fallback: try "at"
	if hasCmd("at") {
		deleteCmd := fmt.Sprintf("rm -rf -- %s", shellQuote(directoryPath))
		when := time.Now().Add(duration).Format("15:04 %m/%d/%Y")
		cmd := exec.Command("bash", "-lc", fmt.Sprintf("echo %q | at %q", deleteCmd, when))
		if out, err := cmd.CombinedOutput(); err == nil {
			jobID, _ := parseAtJobID(string(out))
			vLog("Created %s\n", directoryPath)
			vLog("Deletion scheduled via 'at' at %s (job: %d)\n", when, jobID)
			fmt.Printf("%s", directoryPath)
			return addRecord(Record{
				Path:       directoryPath,
				Scheduler:  SchedAt,
				AtJob:      jobID,
				CreatedAt:  time.Now(),
				Expiration: expAt,
			})
		} else {
			return fmt.Errorf("'at' scheduling failed: %v\nOutput: %s", err, string(out))
		}
	}

	// Nothing available
	fmt.Fprintln(os.Stderr, "No system scheduler found (neither systemd-run nor at).")
	fmt.Fprintln(os.Stderr, "Install systemd (user) or 'at', or delete manually later.")

	return nil
}

func Cancel(absPath string) error {
	p, err := filepath.Abs(absPath)
	if err != nil {
		return err
	}
	rec, ok, err := getRecord(p)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("no scheduled deletion found for %s", p)
	}

	switch rec.Scheduler {
	case SchedSystemd:
		// Arrêter le timer (annule la planif) + le service s’il existe
		_ = exec.Command("systemctl", "--user", "stop", rec.Unit+".timer").Run()
		_ = exec.Command("systemctl", "--user", "stop", rec.Unit+".service").Run()
		// (optionnel) nettoyer l'état failed
		_ = exec.Command("systemctl", "--user", "reset-failed", rec.Unit+".service").Run()

	case SchedAt:
		if rec.AtJob == 0 {
			return fmt.Errorf("invalid at job id for %s", p)
		}
		if err := exec.Command("atrm", strconv.Itoa(rec.AtJob)).Run(); err != nil {
			return fmt.Errorf("atrm failed: %w", err)
		}

	default:
		return fmt.Errorf("unknown scheduler: %s", rec.Scheduler)
	}

	if err := delRecord(p); err != nil {
		return err
	}
	return nil
}

func List() ([]Record, error) {
	st, err := loadState()
	if err != nil {
		return nil, err
	}
	out := make([]Record, 0, len(st.Records))
	for _, r := range st.Records {
		out = append(out, r)
	}
	return out, nil
}
