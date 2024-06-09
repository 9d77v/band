package ratelimiter

import (
	"context"

	"fmt"

	"time"

	"github.com/9d77v/band/pkg/stores/redis"
)

var ctx = context.Background()

type RateLimiter struct {
	client *redis.Redis
	limit  int
}

func NewRateLimiter(client *redis.Redis, limit int) *RateLimiter {
	return &RateLimiter{
		client: client,
		limit:  limit,
	}
}

// AllowByDayRange 实现限制一天内的访问次数的逻辑
func (r *RateLimiter) AllowByDayRange(key string) (bool, error) {
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location()).Unix()
	return r.AllowByTimeRange(key, startTime, endTime)
}

// AllowByDayRange 实现限制一月内的访问次数的逻辑
func (r *RateLimiter) AllowByMonthRange(key string) (bool, error) {
	now := time.Now()
	next := now.AddDate(0, 1, 0)
	startTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Unix()
	endTime := time.Date(next.Year(), next.Month(), 1, 0, 0, 0, 0, now.Location()).Unix() - 1
	return r.AllowByTimeRange(key, startTime, endTime)
}

// AllowByTimeRange 实现限制一个时间段内的访问次数的逻辑
func (r *RateLimiter) AllowByTimeRange(key string, startTime, endTime int64) (bool, error) {
	redisKey := fmt.Sprintf("%s:%d", key, startTime)
	// 使用INCR和EXPIRE来增加计数并设置过期时间
	result, err := r.client.Incr(ctx, redisKey).Result()
	if err != nil {
		return false, err
	}

	if result == 1 {
		// 如果是第一次设置这个键，设置过期时间
		r.client.ExpireAt(ctx, redisKey, time.Unix(endTime, 0))

	}
	return result <= int64(r.limit), nil
}

// AllowByDuration 实现限制一段时间内的访问次数的逻辑
func (r *RateLimiter) AllowByDuration(key string, windowSec int) (bool, error) {
	// 使用INCR和EXPIRE来增加计数并设置过期时间
	result, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if result == 1 {
		// 如果是第一次设置这个键，设置过期时间
		r.client.Expire(ctx, key, time.Duration(windowSec)*time.Second)
	}
	return result <= int64(r.limit), nil
}
