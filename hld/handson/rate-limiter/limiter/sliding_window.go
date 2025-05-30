package limiter

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// SlidingWindowLimiter implements rate limiting using a sliding window approach
type SlidingWindowLimiter struct {
	client      *redis.Client
	windowSize  time.Duration
	maxRequests int
}

// NewSlidingWindowLimiter creates a new sliding window rate limiter
func NewSlidingWindowLimiter(client *redis.Client, windowSize time.Duration, maxRequests int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		client:      client,
		windowSize:  windowSize,
		maxRequests: maxRequests,
	}
}

// IsAllowed checks if a request is allowed under the rate limit
func (l *SlidingWindowLimiter) IsAllowed(ctx context.Context, key string) (bool, error) {
	now := time.Now().Unix()
	windowKey := fmt.Sprintf("sliding_window:%s", key)

	// Add current timestamp to the sorted set
	err := l.client.ZAdd(ctx, windowKey, redis.Z{
		Score:  float64(now),
		Member: now,
	}).Err()
	if err != nil {
		return false, err
	}

	// Remove old entries outside the window
	err = l.client.ZRemRangeByScore(ctx, windowKey, "0", fmt.Sprintf("%d", now-int64(l.windowSize.Seconds()))).Err()
	if err != nil {
		return false, err
	}

	// Count requests in the current window
	count, err := l.client.ZCard(ctx, windowKey).Result()
	if err != nil {
		return false, err
	}

	// Set expiry on the key
	err = l.client.Expire(ctx, windowKey, l.windowSize).Err()
	if err != nil {
		return false, err
	}

	return count <= int64(l.maxRequests), nil
}
