package model

import (
	"errors"
	"github.com/abel1105/go_gin_boilerplate/db"
)

type Test struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name"`
}

func (Test) GetById(id string)(*Test, error) {
	DB := db.GetDB()
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	row := DB.QueryRow("SELECT * FROM tests WHERE id = ?;", id)
	Test := &Test{}
	err := row.Scan(&Test.ID, &Test.Name)
	if err != nil {
		return nil, err
	}
	return Test, nil
}
