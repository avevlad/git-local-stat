package server

import (
	"path/filepath"
	"os"
	"log"
)

func GetCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir;
}
