package routes

import "github.com/gin-gonic/gin"

func CartRoutes(r *gin.Engine) {
	cartAPI := r.Group("/cart")
	cartAPI.POST("")
	cartAPI.DELETE("")
	cartAPI.GET("")
}
