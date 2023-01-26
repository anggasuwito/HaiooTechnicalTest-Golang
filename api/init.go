package api

import (
	"HaiooTechnicalTest-Golang/api/routes"
	"HaiooTechnicalTest-Golang/config"
	"HaiooTechnicalTest-Golang/models"
	"github.com/gin-gonic/gin"
	"os"
)

func Init() {
	r := gin.Default()
	conf := &models.Config{DB: config.GormDB()}
	routes.NewCartRoutes(r, conf).Routes()
	r.Run(os.Getenv("ADDRESS"))
}
