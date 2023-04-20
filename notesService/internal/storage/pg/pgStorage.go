package pg

import (
	"github.com/go-pg/pg/v10"
	"github.com/norciTest/notesService/internal/model"
	"github.com/spf13/viper"
	"os"
)

type Storage struct {
	client *pg.DB
}

func (s *Storage) Init() error {
	var (
		options *pg.Options
		err     error
	)
	if options, err = pg.ParseURL("postgres://" +
		viper.GetString("postgres.login") + ":" +
		os.Getenv("POSTGRES_PASSWORD") + "@" +
		viper.GetString("postgres.host") + ":" +
		viper.GetString("postgres.port") + "/" +
		viper.GetString("postgres.name") + "?sslmode=disable"); err != nil {
		return err
	}

	s.client = pg.Connect(options)

	return err
}

func (s *Storage) Get(id int) (*model.Note, error) {
	var (
		n   = &model.Note{}
		err error
	)
	if err = s.client.Model(n).Where("id = ?", id).Select(); err != nil {
		return nil, err
	}
	return n, nil
}

func (s *Storage) Add(n *model.Note) error {
	if _, err := s.client.Model(n).Insert(); err != nil {
		return err
	}
	return nil
}
func (s *Storage) Delete(id int) error {
	if _, err := s.client.Model(&model.Note{}).Where("id = ?", id).Delete(); err != nil {
		return err
	}
	return nil
}
