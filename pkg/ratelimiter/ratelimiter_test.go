package ratelimiter

import (
	"testing"

	"github.com/9d77v/band/pkg/stores/redis"
)

func TestRateLimiter_AllowByDayRange(t *testing.T) {
	cli, err := redis.NewRedis(redis.FromEnv())
	if err != nil {
		panic(err)
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		r       *RateLimiter
		args    args
		want    bool
		wantErr bool
	}{
		{"test", NewRateLimiter(cli, 3), args{"TEST"}, true, false},
		{"test", NewRateLimiter(cli, 3), args{"TEST"}, true, false},
		{"test", NewRateLimiter(cli, 3), args{"TEST"}, true, false},
		{"test", NewRateLimiter(cli, 3), args{"TEST"}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.AllowByDayRange(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RateLimiter.AllowByDayRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RateLimiter.AllowByDayRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
