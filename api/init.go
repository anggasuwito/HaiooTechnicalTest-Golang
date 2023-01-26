package api

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Init() {
	r := gin.Default()

	r.Run(os.Getenv("ADDRESS"))
}
