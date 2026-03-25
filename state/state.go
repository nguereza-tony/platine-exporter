package state

import (
	"encoding/json"
	"os"
)

type State struct {
	Inode  uint64 `json:"inode"`
	Offset int64  `json:"offset"`
}

func Load(path string) *State {
	s := &State{}

	data, err := os.ReadFile(path)
	if err == nil {
		json.Unmarshal(data, s)
	}

	return s
}

func Save(path string, s *State) error {
	data, _ := json.Marshal(s)
	return os.WriteFile(path, data, 0644)
}
