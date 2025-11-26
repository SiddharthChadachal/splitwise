package splitwise

import (
	"encoding/json"
	"os"
)

type Storage struct {
	Bills map[string]*Bill `json:"bills"`
}

func LoadStorage(path string) (*Storage, error) {
	data := &Storage{Bills: make(map[string]*Bill)}

	file, err := os.ReadFile(path)
	if err != nil {
		// file not found → first time → return empty storage
		return data, nil
	}

	err = json.Unmarshal(file, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Storage) Save(path string) error {
	bytes, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, bytes, 0644)
}
