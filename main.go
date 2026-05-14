package main

import (
	"fmt"
	"log"

	"github.com/GustavoMayon/clase5" //importando
)

func main() {
	log.SetPrefix("clase5: ") //prefijo para saber de donde viene el error
	log.SetFlags(0)           //no muestra la fecha ni la hora

	nombres := []string{"lucas", "pedro", "jonas"}
	mensajes, err := clase5.SaludoVarios(nombres)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mensajes)

	mensaje, err := clase5.Saludo("Gustavo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mensaje)
}
