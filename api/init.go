package api

import (
	"HaiooTechnicalTest-Golang/config"
	"HaiooTechnicalTest-Golang/models"
	"github.com/gin-gonic/gin"
	"os"
)

func Init() {
	r := gin.Default()
	conf := &models.Config{DB: config.GormDB()}
	r.Run(os.Getenv("ADDRESS"))
}
