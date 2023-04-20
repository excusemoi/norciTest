package service

import (
	"github.com/gin-gonic/gin"
	"github.com/norciTest/notesService/internal/storage"
)

type Service struct {
	Db storage.Storage
	S  *gin.Engine
}

func New(db storage.Storage) *Service {
	service := &Service{Db: db, S: gin.Default()}
	service.initRoutes()
	return service
}
