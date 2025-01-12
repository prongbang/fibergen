//go:build !test
// +build !test

package filex

import (
	"os"
)

// FileX is the interface
type FileX interface {
	EnsureDir(dir string) error
	WriteFile(filename string, data []byte) error
	ReadFile(filename string) string
	Getwd() (string, error)
	Chdir(dir string) error
	IsExist(filename string) bool
	IsDirExist(dir string) bool
}

type fileX struct {
}

func (f *fileX) IsDirExist(dir string) bool {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func (f *fileX) IsExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (f *fileX) EnsureDir(dir string) error {
	err := os.MkdirAll(dir, 0755)
	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}

func (f *fileX) WriteFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0755)
}

func (f *fileX) ReadFile(filename string) string {
	if text, err := os.ReadFile(filename); err == nil {
		return string(text)
	}
	return ""
}

func (f *fileX) Getwd() (string, error) {
	return os.Getwd()
}

func (f *fileX) Chdir(dir string) error {
	return os.Chdir(dir)
}

// NewFileX is new instance with func
func NewFileX() FileX {
	return &fileX{}
}
