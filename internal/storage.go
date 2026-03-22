package wmtracker

import (
	"encoding/json"
	"os"
)

const storageFile = "requests.json"

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
f
func SaveRequests(requests []RequestInput) error {
	data, err := json.MarshalIndent(requests, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(storageFile, data, 0o644)
}
