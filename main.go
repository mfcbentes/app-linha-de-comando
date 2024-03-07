package main

import (
	"log"
	"mfcbentes/linha-de-comando/app"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
