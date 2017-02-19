package main

import (
	"fmt"
	"time"

	"github.com/ssoudan/hugo/input"
	"github.com/ssoudan/hugo/logging"
)

var log = logging.Log("hugotune")

func main() {

	p, err := input.New()
	if err != nil {
		log.Fatal("Failed to open channel:", err)
	}

	fmt.Println("Starting")

	for {
		v, err := p.Read(0)
		if err != nil {
			log.Fatal("Failed to read value: ", err)
		}
		fmt.Printf("p(0)=%f\n", v)

		v, err = p.Read(1)
		if err != nil {
			log.Fatal("Failed to read value: ", err)
		}
		fmt.Printf("p(1)=%f\n", v)

		v, err = p.Read(2)
		if err != nil {
			log.Fatal("Failed to read value: ", err)
		}
		fmt.Printf("p(2)=%f\n", v)

		v, err = p.Read(3)
		if err != nil {
			log.Fatal("Failed to read value: ", err)
		}
		fmt.Printf("p(3)=%f\n", v)

		time.Sleep(500 * time.Millisecond)
	}
}
