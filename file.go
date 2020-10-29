package main

import (
	"io/ioutil"
	"os"
)

// FileX is the interface
type FileX interface {
	EnsureDir(dir string) error
	WriteFile(filename string, data []byte) error
}

type fileX struct {
}

func (f *fileX) EnsureDir(dir string) error {
	err := os.MkdirAll(dir, 0755)
	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}

func (f *fileX) WriteFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0755)
}

// NewFileX is new instance with func
func NewFileX() FileX {
	return &fileX{}
}
