package utils

import (
    "golang.org/x/time/rate"
    "sync"
    "time"
    beego "github.com/beego/beego/v2/server/web"
)

type RateLimiter struct {
    ips map[string]*rate.Limiter
    mu  *sync.RWMutex
}

func NewRateLimiter() *RateLimiter {
    return &RateLimiter{
        ips: make(map[string]*rate.Limiter),
        mu:  &sync.RWMutex{},
    }
}

func (r *RateLimiter) GetLimiter(ip string) *rate.Limiter {
    r.mu.Lock()
    defer r.mu.Unlock()

    limiter, exists := r.ips[ip]
    if !exists {
        // Get rate limit configuration
        requests, _ := beego.AppConfig.Int("ratelimit.requests")
        duration, _ := beego.AppConfig.Int("ratelimit.duration")
        
        limiter = rate.NewLimiter(rate.Every(time.Duration(duration)*time.Second), requests)
        r.ips[ip] = limiter
    }

    return limiter
}