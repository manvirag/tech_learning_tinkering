# Rate Limiter Implementation

This project implements three different rate limiting algorithms using Redis:
1. Fixed Window Rate Limiter
2. Sliding Window Rate Limiter
3. Token Bucket Rate Limiter

## Prerequisites

- Go 1.21 or higher
- Redis server running on localhost:6379

## Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod tidy
```

## Running the Application

1. Start Redis server
2. Run the application:
```bash
go run main.go
```

The server will start on http://localhost:8080

## API Endpoints

### 1. Fixed Window Rate Limiter
- Endpoint: `GET /fixed`
- Rate Limit: 10 requests per minute
- Window: Fixed 1-minute window

### 2. Sliding Window Rate Limiter
- Endpoint: `GET /sliding`
- Rate Limit: 10 requests per minute
- Window: Sliding 1-minute window

### 3. Token Bucket Rate Limiter
- Endpoint: `GET /token`
- Rate Limit: 1 token per second
- Bucket Capacity: 10 tokens

## Testing the Rate Limiters

You can test the rate limiters using curl:

```bash
# Test fixed window
for i in {1..15}; do
  curl http://localhost:8080/fixed
  sleep 1
done

# Test sliding window
for i in {1..15}; do
  curl http://localhost:8080/sliding
  sleep 1
done

# Test token bucket
for i in {1..15}; do
  curl http://localhost:8080/token
  sleep 1
done
```

## Rate Limiter Algorithms

### 1. Fixed Window
- Simple to implement
- Less accurate at window boundaries
- Can allow burst of requests at window start
- Uses Redis INCR command

### 2. Sliding Window
- More accurate than fixed window
- Smoother rate limiting
- Uses Redis Sorted Sets
- Better for distributed systems

### 3. Token Bucket
- Allows burst of traffic
- Smooths out traffic
- Good for APIs with varying load
- Uses Redis for token storage

## Error Responses

- 429 Too Many Requests: Rate limit exceeded
- 500 Internal Server Error: Rate limiter error

## Configuration

You can modify the rate limits in `main.go`:

```go
fixedLimiter = limiter.NewFixedWindowLimiter(redisClient, time.Minute, 10)    // 10 requests per minute
slidingLimiter = limiter.NewSlidingWindowLimiter(redisClient, time.Minute, 10) // 10 requests per minute
tokenLimiter = limiter.NewTokenBucketLimiter(redisClient, 10, 1.0, time.Minute) // 1 token per second, max 10 tokens
``` 