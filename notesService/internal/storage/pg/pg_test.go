package pg

import (
	"github.com/joho/godotenv"
	"github.com/norciTest/notesService/internal/config"
	"github.com/norciTest/notesService/internal/model"
	"path/filepath"
	"testing"
)

//без табличных тестов на этот раз :)

func TestPg(t *testing.T) {
	var (
		s   = &Storage{}
		err error
	)

	godotenv.Load(filepath.Join("..", "..", "..", ".env"))

	if err = config.InitConfig(filepath.Join("..", "..", "..", "configs")); err != nil {
		t.Fatal(err)
	}

	if err = s.Init(); err != nil {
		t.Fatal(err)
	}

	n := &model.Note{
		Id:          4,
		Description: "desctiption #4",
	}

	t.Run("add", func(t *testing.T) {
		if err = s.Add(n); err != nil {
			t.Error(err)
		}
	})

	t.Run("get", func(t *testing.T) {
		if n, err = s.Get(4); err != nil {
			t.Error(err)
		}
		t.Log(*n)
	})

	t.Run("delete", func(t *testing.T) {
		if err = s.Delete(4); err != nil {
			t.Error(err)
		}
	})
}
