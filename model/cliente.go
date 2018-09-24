package model

import (
	"database/sql"
	"strings"
)

//Cliente struct
type Cliente struct {
	ID             int64  `json:"id"`
	Nome           string `json:"nome"`
	DataNascimento string `json:"dataNascimento"`
}

//InsertClient ...
func (c *Cliente) InsertClient(db *sql.DB) error {
	statement, err := db.Prepare(`insert into clientes (nome, data_nascimento)
								values
								(?, ?)`)
	if err != nil {
		return err
	}
	res, err := statement.Exec(c.Nome, c.DataNascimento)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	c.ID = id
	return nil
}

//UpdateClient ...
func (c *Cliente) UpdateClient(db *sql.DB) error {
	statement, err := db.Prepare(`update clientes set nome = ?, data_nascimento = ? where id = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(c.Nome, c.DataNascimento, c.ID)

	return err
}

//GetClient ...
func (c *Cliente) GetClient(db *sql.DB) error {
	err := db.QueryRow(`select id, nome, data_nascimento
					from clientes
					where id =  ?`, c.ID).Scan(&c.ID, &c.Nome, &c.DataNascimento)
	if err != nil {
		return err
	}

	return err
}

//GetClients ...
func (c *Cliente) GetClients(db *sql.DB) ([]Cliente, error) {
	var values []interface{}
	var where []string

	if c.ID != 0 {
		where = append(where, "id = ?")
		values = append(values, c.ID)
	}

	if c.DataNascimento != "" {
		where = append(where, "data_nascimento = ?")
		values = append(values, c.DataNascimento)
	}
	if c.Nome != "" {
		where = append(where, "nome = ?")
		values = append(values, c.Nome)
	}

	rows, err := db.Query(`select id, nome, data_nascimento
					from clientes
					where 1=1 `+strings.Join(where, " AND "), values...)
	if err != nil {
		return nil, err
	}

	clientes := []Cliente{}
	defer rows.Close()
	for rows.Next() {
		var cl Cliente
		if err = rows.Scan(&cl.ID, &cl.Nome, &cl.DataNascimento); err != nil {
			return nil, err
		}
		clientes = append(clientes, cl)
	}
	return clientes, nil
}

//DeleteClient ...
func (c *Cliente) DeleteClient(db *sql.DB) error {
	statement, err := db.Prepare(`delete from clientes where id = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(c.ID)

	return err
}
