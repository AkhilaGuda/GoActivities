package main

import "fmt"

type Storage interface {
	Save(data string) error
}

func SaveData(storage Storage, data string) error {
	if err := storage.Save(data); err != nil {
		return fmt.Errorf("failed to save: %w", err)
	}
	return nil
}
