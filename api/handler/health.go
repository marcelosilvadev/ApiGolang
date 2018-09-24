package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// HealthCheck handler da rota GET /health
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Teste")
	// Setando um header http de resposta
	w.Header().Set("content-type", "application/json")

	// Gerando um objeto customizado à partir de um map, e o convertendo em json
	response, _ := json.Marshal(map[string]interface{}{
		"status": "up",
	})

	// Write escreve o conteúdo do slice de bytes no corpo da resposta
	w.Write(response)

	return
}
