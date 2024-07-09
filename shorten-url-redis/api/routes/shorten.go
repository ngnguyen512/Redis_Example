package routes

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/ngnguyen/shorten-url-redis/api/database"
	"github.com/ngnguyen/shorten-url-redis/api/helpers"
)

type request struct {
	URL         string
	CustomShort string
	Expiry      time.Duration
}

type response struct {
	URL             string
	CustomShort     string
	Expiry          time.Duration
	XRateRemaining  int
	XRateLimitReset time.Duration
}

func ShortenURL(c *gin.Context) {
	body := new(request)

	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot parse JSON",
		})
		return
	}

	// Rate limiting
	r2 := database.CreateClient(1)
	defer r2.Close()
	val, err := r2.Get(database.Ctx, c.ClientIP()).Result()
	if err == redis.Nil {
		r2.Set(database.Ctx, c.ClientIP(), os.Getenv("API_QUOTA"), 30*time.Minute)
	} else if err == nil {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.Get(database.Ctx, c.ClientIP()).Result()
			limitDuration, _ := time.ParseDuration(limit)
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error":            "Rate limit exceeded",
				"rate_limit_reset": limitDuration / time.Minute,
			})
			return
		}
	}

	// Validate URL
	if !govalidator.IsURL(body.URL) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid URL",
		})
		return
	}

	// Domain error checking
	if !helpers.RemoveDomainError(body.URL) {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "nice try",
		})
		return
	}

	// Enforce HTTPS and SSL
	body.URL = helpers.EnforceHTTP(body.URL)
	var id string

	if body.CustomShort == "" {
		id = uuid.NewString()[:6]
	} else {
		id = body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	existingURL, _ := r.Get(database.Ctx, id).Result()
	// check if the user provided short is already in use
	if existingURL != "" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "URL short already in use",
		})
		return
	}
	if body.Expiry == 0 {
		body.Expiry = 24 // default expiry of 24 hours
	}
	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to connect to server",
		})
		return
	}
	// respond with the url, short, expiry in hours, calls remaining and time to reset
	resp := response{
		URL:             body.URL,
		CustomShort:     os.Getenv("DOMAIN") + "/" + id,
		Expiry:          body.Expiry,
		XRateRemaining:  10, // This should be dynamic based on actual remaining requests allowed
		XRateLimitReset: 30, // This should be set dynamically based on TTL from Redis
	}

	remainingCalls, _ := strconv.Atoi(val)
	ttl, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
	resp.XRateRemaining = remainingCalls
	resp.XRateLimitReset = ttl / time.Minute

	c.JSON(http.StatusOK, resp)
}
