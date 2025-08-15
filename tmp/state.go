package tmp

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	SchedSystemd = "systemd"
	SchedAt      = "at"
)

type Record struct {
	Path       string    `json:"path"`
	Scheduler  string    `json:"scheduler"`
	Unit       string    `json:"unit,omitempty"`
	AtJob      int       `json:"at_job,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	Expiration time.Time `json:"expiration"`
}

type stateFile struct {
	Records map[string]Record `json:"records"` // key: absolute path
}

func configDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".utilitaire")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", err
	}
	return dir, nil
}

func statePath() (string, error) {
	dir, err := configDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "schedule.json"), nil
}

func loadState() (*stateFile, error) {
	p, err := statePath()
	if err != nil {
		return nil, err
	}
	b, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return &stateFile{Records: map[string]Record{}}, nil
		}
		return nil, err
	}
	var st stateFile
	if err := json.Unmarshal(b, &st); err != nil {
		return nil, fmt.Errorf("parse state: %w", err)
	}
	if st.recordsNil() {
		st.Records = map[string]Record{}
	}
	return &st, nil
}

func (s *stateFile) recordsNil() bool { return s.Records == nil }

func saveState(s *stateFile) error {
	p, err := statePath()
	if err != nil {
		return err
	}
	tmp := p + ".tmp"
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(tmp, b, 0o600); err != nil {
		return err
	}
	return os.Rename(tmp, p)
}

func addRecord(rec Record) error {
	st, err := loadState()
	if err != nil {
		return err
	}
	st.Records[rec.Path] = rec
	return saveState(st)
}

func delRecord(absPath string) error {
	st, err := loadState()
	if err != nil {
		return err
	}
	delete(st.Records, absPath)
	return saveState(st)
}

func getRecord(absPath string) (Record, bool, error) {
	st, err := loadState()
	if err != nil {
		return Record{}, false, err
	}
	rec, ok := st.Records[absPath]
	return rec, ok, nil
}
