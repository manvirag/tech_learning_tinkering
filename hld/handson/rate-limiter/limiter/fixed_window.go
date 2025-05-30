package limiter

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// FixedWindowLimiter implements rate limiting using a fixed window approach
type FixedWindowLimiter struct {
	client      *redis.Client
	windowSize  time.Duration
	maxRequests int
}

// NewFixedWindowLimiter creates a new fixed window rate limiter
func NewFixedWindowLimiter(client *redis.Client, windowSize time.Duration, maxRequests int) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		client:      client,
		windowSize:  windowSize,
		maxRequests: maxRequests,
	}
}

// IsAllowed checks if a request is allowed under the rate limit
func (l *FixedWindowLimiter) IsAllowed(ctx context.Context, key string) (bool, error) {
	// Get the current window key , dividing so that can get the window like 10 -> 10 divisible -> 20 divisible between divided = 1 after 20 its 2
	windowKey := fmt.Sprintf("fixed_window:%s:%d", key, time.Now().Unix()/int64(l.windowSize.Seconds()))

	// Increment the counter for this window
	count, err := l.client.Incr(ctx, windowKey).Result()
	if err != nil {
		return false, err
	}

	// Set expiry on the key if this is the first request in the window
	if count == 1 {
		err = l.client.Expire(ctx, windowKey, l.windowSize).Err()
		if err != nil {
			return false, err
		}
	}

	// Check if we're under the limit
	return count <= int64(l.maxRequests), nil
}
