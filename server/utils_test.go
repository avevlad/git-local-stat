package server

import (
	"testing"
	"fmt"
)

func TestSrand(t *testing.T) {
	fmt.Print(srand(10))
}