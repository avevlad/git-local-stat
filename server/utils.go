package server

import (
	"path/filepath"
	"os"
	"log"
	"math/rand"
	"time"
)


func GetCurrentDir() string {
	//fmt.Println(os.Args[1])
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir;
}

var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

func srand(size int) string {
	buf := make([]byte, size)
	rand.Seed(time.Now().UnixNano()) // takes the current time in nanoseconds as the seed
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}