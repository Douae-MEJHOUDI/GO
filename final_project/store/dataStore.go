package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type DataManager struct {
	datadir string
	mu      sync.Mutex
}

func NewDataManager(datadir string) (*DataManager, error) {
	err := os.MkdirAll(datadir, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create data directory", err)
	}

	return &DataManager{
		datadir: datadir,
	}, nil

}

func (dm *DataManager) SaveData(filename string, data interface{}) error {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	filepath := filepath.Join(dm.datadir, filename)
	jsonData, err := json.MarshalIndent(data, " ", "   ")
	if err != nil {
		return fmt.Errorf("failed to marshal")
	}

	err = os.WriteFile(filepath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file")
	}

	return nil
}

func (dm *DataManager) LoadData(filename string, data interface{}) error {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	filepath := filepath.Join(dm.datadir, filename)

	_, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		return nil
	}

	oldData, err := os.ReadFile(filepath)

	if err != nil {
		return fmt.Errorf("failed to read file", err)
	}

	err = json.Unmarshal(oldData, data)
	if err != nil {
		return fmt.Errorf("failed to Unmarshal, ", err)
	}
	return nil
}
