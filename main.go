package main

import (
	"EmuTwitter/bd"
	"EmuTwitter/handlers"
	"log"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}