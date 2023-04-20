package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/norciTest/notesService/internal/model"
	"net/http"
	"strconv"
)

func (s *Service) Get(c *gin.Context) {
	var (
		id  int
		err error
	)
	id, err = strconv.Atoi(c.Query("id"))
	n, err := s.Db.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, n)
}

func (s *Service) Delete(c *gin.Context) {
	var (
		id  int
		err error
	)
	if id, err = strconv.Atoi(c.Query("id")); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err = s.Db.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("note with id = %d deleted", id))
}

func (s *Service) Add(c *gin.Context) {
	var (
		n   = &model.Note{}
		err error
	)
	if err = c.BindJSON(n); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err = s.Db.Add(n); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("note with id = %d added", n.Id))
}
