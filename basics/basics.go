package basics

import (
	"fmt"
	"log"
	"time"
)

func Basics() {
	n := time.Now()

	log.Println("Hello, world!")
	log.Println("The time is", n)
	fmt.Println("The time is", n)
}
