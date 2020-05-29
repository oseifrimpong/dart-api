package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"

	"dart-api/api/config"
	"dart-api/api/model"
)

type dserver struct {
	router *gin.Engine
	cont   *dig.Container
}

func NewServer(e *gin.Engine, c *dig.Container) *dserver {
	return &dserver{
		router: e,
		cont:   c,
	}
}

func (ds *dserver) SetupDB() error {
	var db *gorm.DB

	if err := ds.cont.Invoke(func(d *gorm.DB) { db = d }); err != nil {
		return err
	}
	db.AutoMigrate(&model.User{})
	fmt.Println("Database connection successful.")
	return nil
}
func (ds *dserver) Start() error {
	var cfg *config.Config
	if err := ds.cont.Invoke(func(c *config.Config) { cfg = c }); err != nil {
		return err
	}
	return ds.router.Run(fmt.Sprintf(":%s", cfg.Port))
}
