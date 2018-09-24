package handler

import (
	"database/sql"
	"encoding/json"
	"erncliente/api/db"
	"erncliente/model"
	"erncliente/util"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//InsertClient ...
func InsertClient(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var c model.Cliente
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados está down", "")
		return
	}
	db := d.DB
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = c.InsertClient(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Cliente", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, c, 0, 0)
}

//UpdateClient
func UpdateClient(w http.ResponseWriter, r *http.Request) {
	var a model.Cliente
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdateCliente] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()
	a.ID = int64(id)
	if err := a.UpdateClient(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, a, 0, 0)
}

//DeleteClient
func DeleteClient(w http.ResponseWriter, r *http.Request) {
	var c model.Cliente
	var d db.DB
	var t util.App
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/DeleteCliente -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	c.ID = int64(id)
	if err := c.DeleteClient(db); err != nil {
		log.Printf("[handler/DeleteCliente -  Erro ao tentar deletar cliente. Erro: %s", err.Error())
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, c, 0, 0)
}

//GetClient ...
func GetClient(w http.ResponseWriter, r *http.Request) {
	var c model.Cliente
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetClient] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	c.ID = int64(id)
	err = c.GetClient(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetClient -  Não há cliente com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há cliente com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetClient -  Erro ao tentar buscar cliente. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, c, 0, 0)
}

//GetClients ...
func GetClients(w http.ResponseWriter, r *http.Request) {
	var c model.Cliente
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetClient] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	id, _ := strconv.Atoi(r.FormValue("id"))
	nome := r.FormValue("nome")
	dataNascimento := r.FormValue("dataNascimento")

	c.ID = int64(id)
	c.Nome = nome
	c.DataNascimento = dataNascimento

	clientes, err := c.GetClients(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetClient -  Não há cliente com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há clientes cadastrado.", err.Error())
		} else {
			log.Printf("[handler/GetClient -  Erro ao tentar buscar clientes. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, clientes, 0, 0)
}
