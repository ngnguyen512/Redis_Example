package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/ngnguyen/shorten-url-redis/api/database"
	"net/http"
)

func ResolveURL(c *gin.Context) {
	url := c.Param("url")
	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "short not found in the database",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot connect to DB",
		})
		return
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	c.Redirect(http.StatusMovedPermanently, value)
}
