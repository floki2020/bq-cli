package pkg

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
)

func PathCreate(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func PathIsExits(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
func GetExecutablePath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	path := filepath.Dir(exe)

	return path, nil
}
func CreateFile(content bytes.Buffer, path string) {
	file, err := os.Create(path)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(content.String())
	if err != nil {
		log.Fatal(err)
	}
}
