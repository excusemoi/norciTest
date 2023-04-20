package storage

import "github.com/norciTest/notesService/internal/model"

type Storage interface {
	Get(id int) (*model.Note, error)
	Add(*model.Note) error
	Delete(id int) error
}
