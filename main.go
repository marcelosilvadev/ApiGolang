package main

import (
	"erncliente/api"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) //Informa o local do erro
	api := api.App{}
	api.StartServer()
}
