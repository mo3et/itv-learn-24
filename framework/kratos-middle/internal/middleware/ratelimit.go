package middleware

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"golang.org/x/time/rate"
	"google.golang.org/grpc/status"
)

type Limiter struct {
	rateLimiter *rate.Limiter
	maxDelay    time.Duration
}

func (l *Limiter) Handle(next middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if err := l.rateLimiter.Wait(ctx); err != nil {
			return nil, err
		}
		start := time.Now()
		reply, err := next(ctx, req)
		duration := time.Since(start)

		if err != nil {
			// 判断是否超最大延迟时间
			if duration > l.maxDelay {
				l.rateLimiter.SetLimit(l.rateLimiter.Limit() / 2)
			} else {
				st, _ := status.FromError(err)
				// 检查gRPC错误码，根据不同错误码进行限流调整
				if st.Code() != 0 {
					l.rateLimiter.SetLimit(l.rateLimiter.Limit() * 2)
				}
			}
		}
		return reply, err
	}
}
