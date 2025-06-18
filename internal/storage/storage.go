package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *storage[T] {
	return &storage[T]{FileName: fileName}
}
func (s *storage[T]) Save(data T) error {
	if _, err := os.Stat(s.FileName); os.IsNotExist(err) {
		if _, err := os.Create(s.FileName); err != nil {
			return fmt.Errorf("error creating file: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("error checking file: %v", err)
	}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(s.FileName, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
func (s *storage[T]) Load(data *T) error {
	//what is the file does not exist?
	if _, err := os.Stat(s.FileName); os.IsNotExist(err) {

		if _, err := os.Create(s.FileName); err != nil {
			return fmt.Errorf("error creating file: %v", err)
		}

	} else if err != nil {
		return fmt.Errorf("error checking file: %v", err)
	}
	jsonData, err := os.ReadFile(s.FileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	if len(jsonData) == 0 {
		return nil
		// If the file is empty, initialize data to an empty slice
	}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return err
	}
	return nil
}
