package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"rate-limiter/limiter"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient    *redis.Client
	fixedLimiter   *limiter.FixedWindowLimiter
	slidingLimiter *limiter.SlidingWindowLimiter
	tokenLimiter   *limiter.TokenBucketLimiter
)

func init() {
	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Test Redis connection
	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	// Initialize rate limiters
	fixedLimiter = limiter.NewFixedWindowLimiter(redisClient, time.Minute, 10)           // 10 requests per minute
	slidingLimiter = limiter.NewSlidingWindowLimiter(redisClient, time.Minute, 10)       // 10 requests per minute
	tokenLimiter = limiter.NewTokenBucketLimiter(redisClient, 10, 1.0/60.0, time.Minute) // 1 token per minute, max 10 tokens
}

func rateLimitMiddleware(limiter interface {
	IsAllowed(context.Context, string) (bool, error)
}) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP() // Use client IP as the rate limit key
		fmt.Println("Rate limiting key:", key)
		allowed, err := limiter.IsAllowed(c.Request.Context(), key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Rate limit error"})
			c.Abort()
			return
		}

		if !allowed {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Fixed window rate limiter endpoint
	r.GET("/fixed", rateLimitMiddleware(fixedLimiter), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Fixed window rate limit passed"})
	})

	// Sliding window rate limiter endpoint
	r.GET("/sliding", rateLimitMiddleware(slidingLimiter), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Sliding window rate limit passed"})
	})

	// Token bucket rate limiter endpoint
	r.GET("/token", rateLimitMiddleware(tokenLimiter), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Token bucket rate limit passed"})
	})

	log.Fatal(r.Run(":8080"))
}
