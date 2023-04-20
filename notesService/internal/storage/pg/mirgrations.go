package pg

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/norciTest/notesService/internal/model"
)

func (s *Storage) CreateNoteTable() error {
	return s.client.Model(&model.Note{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
}
