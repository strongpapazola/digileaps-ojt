package routes
import (
       "net/http"
       berita "digileaps-ojt/controllers/berita"
       "github.com/gin-gonic/gin"
)
//StartGin function
func StartService() {
       router := gin.Default()
       api := router.Group("/api")
       {
                api.GET("/berita", berita.GetAllBerita)
                api.POST("/berita", berita.CreateBerita)
                api.GET("/berita/:id", berita.GetBerita)
                api.PUT("/berita/:id", berita.UpdateBerita)
                api.DELETE("/berita/:id", berita.DeleteBerita)
       }
       router.NoRoute(func(c *gin.Context) {
              c.AbortWithStatus(http.StatusNotFound)
       })
       router.Run(":8000")
}
