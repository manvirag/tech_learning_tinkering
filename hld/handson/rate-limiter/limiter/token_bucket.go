package limiter

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// TokenBucketLimiter implements rate limiting using a token bucket approach
type TokenBucketLimiter struct {
	client     *redis.Client
	capacity   int
	refillRate float64
	refillTime time.Duration
}

// NewTokenBucketLimiter creates a new token bucket rate limiter
func NewTokenBucketLimiter(client *redis.Client, capacity int, refillRate float64, refillTime time.Duration) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		client:     client,
		capacity:   capacity,
		refillRate: refillRate,
		refillTime: refillTime,
	}
}

// IsAllowed checks if a request is allowed under the rate limit
func (l *TokenBucketLimiter) IsAllowed(ctx context.Context, key string) (bool, error) {
	bucketKey := fmt.Sprintf("token_bucket:%s", key)
	fmt.Println("bucketKey Rate limiting key:", bucketKey)
	now := time.Now().Unix()

	// Get current bucket state
	pipe := l.client.Pipeline()
	getTokens := pipe.Get(ctx, bucketKey)
	getLastRefill := pipe.Get(ctx, bucketKey+":last_refill")
	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return false, err
	}
	fmt.Println(getTokens)
	fmt.Println(getLastRefill)
	// Parse current tokens
	var currentTokens float64
	if tokens, err := getTokens.Float64(); err == nil {
		currentTokens = tokens
	} else {
		currentTokens = float64(l.capacity)
	}

	// Parse last refill time
	var lastRefill int64
	if last, err := getLastRefill.Int64(); err == nil {
		lastRefill = last
	} else {
		lastRefill = now
	}

	// Calculate new tokens
	timePassed := float64(now - lastRefill)
	newTokens := currentTokens + (timePassed * l.refillRate)
	if newTokens > float64(l.capacity) {
		newTokens = float64(l.capacity)
	}

	// Check if we have enough tokens
	if newTokens < 1 {
		return false, nil
	}

	// Update bucket state
	pipe = l.client.Pipeline()
	pipe.Set(ctx, bucketKey, newTokens-1, l.refillTime)
	pipe.Set(ctx, bucketKey+":last_refill", now, l.refillTime)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}
