package main

import (
	"log"

	"github.com/albarogarzon/twittor/bd"
	"github.com/albarogarzon/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la bd")
		return
	}
	handlers.Manejadores()
}
