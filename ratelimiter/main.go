package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type RateLimiter struct {
	limit    int
	window   time.Duration
	requests map[string][]time.Time
	mu       sync.Mutex
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:    limit,
		window:   window,
		requests: make(map[string][]time.Time),
	}
}

func (r *RateLimiter) AllowRequest(ip string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()

	windowStart := now.Add(-r.window)

	timestamps := r.requests[ip]

	idx := 0
	for idx < len(timestamps) {
		if timestamps[idx].Before(windowStart) {
			idx++
		} else {
			break
		}
	}

	//newStamps := make([]time.Time, len(timestamps)-idx)
	newStamps := timestamps[idx:]
	copy(newStamps, timestamps)

	if len(timestamps) >= r.limit {
		return false
	}

	timestamps = append(timestamps, now)
	r.requests[ip] = timestamps
	return true

}

func main() {
	ips := []string{"1", "1", "1", "1", "1", "1", "2"}

	rl := NewRateLimiter(3, 5*time.Second)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("HeapAlloc:", m.HeapAlloc)
	fmt.Println("StackInuse:", m.StackInuse)

	for _, ip := range ips {
		fmt.Print(rl.AllowRequest(ip), "\n")
	}

}
