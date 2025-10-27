package main

import (
	"log"
	"os"
	"path/filepath"
)

type File struct {
	FilePath   string
	OutputFile *os.File
}

func NewFile() *File {
	return &File{}
}
func (f *File) SetFile(fileName, path string) error {
	f.FilePath = filepath.Join(path, fileName)
	log.Printf("Uploads folder full path: %s\n", filepath.Join(path, fileName))

	file, err := os.Create(f.FilePath)
	if err != nil {
		return err
	}
	f.OutputFile = file
	return nil
}
func (f *File) Write(chunk []byte) error {
	if f.OutputFile == nil {
		return nil
	}
	_, err := f.OutputFile.Write(chunk)
	return err
}
func (f *File) Close() error {
	return f.OutputFile.Close()
}
