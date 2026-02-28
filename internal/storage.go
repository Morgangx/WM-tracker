package wmtracker

import (
	"encoding/json"
	"os"
)

const storageFile = "requests.json"

// LoadRequests reads the persisted list of RequestInput from disk. If the
// storage file doesn\u2019t exist, it returns an empty slice without error.
func LoadRequests() ([]RequestInput, error) {
	if _, err := os.Stat(storageFile); os.IsNotExist(err) {
		return nil, nil
	}

	data, err := os.ReadFile(storageFile)
	if err != nil {
		return nil, err
	}

	var requests []RequestInput
	if err := json.Unmarshal(data, &requests); err != nil {
		return nil, err
	}
	return requests, nil
}

// SaveRequests serializes the provided slice and writes it to disk. It
// overwrites any existing file. Callers should handle any synchronization
// if they're updating the slice from multiple goroutines.
func SaveRequests(requests []RequestInput) error {
	data, err := json.MarshalIndent(requests, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(storageFile, data, 0o644)
}
