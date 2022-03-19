package main

import (
	"fmt"
	"github.com/nelsongp/adapter/adapter"
	"log"
)

func main() {
	var t string
	_, err := fmt.Scanln(&t)
	if err != nil {
		log.Fatalf("no se pudo leer el medio de transporte: %v", err)
	}
	transporte := adapter.Factory(t)
	transporte.Mover()
}
