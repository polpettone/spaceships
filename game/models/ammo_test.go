package models

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	for n := 0; n < 100; n++ {
		x := rand.Intn(2)
		fmt.Printf("Result %d \n", x)
	}

}
