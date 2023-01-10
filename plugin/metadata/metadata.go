// Package metadata is a way of defining message headers
package metadata

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type metaKey struct{}

// Metadata is our way of representing request headers internally.
// They're used at the RPC level and translate back and forth
// from Transport headers.
type Metadata map[string]string

// Copy makes a copy of the metadata
func Copy(md Metadata) Metadata {
	cmd := make(Metadata)
	for k, v := range md {
		cmd[k] = v
	}
	return cmd
}

// Get returns a single value from metadata in the context
func Get(ctx context.Context, key string) (string, bool) {
	md, ok := FromContext(ctx)
	if !ok {
		return "", ok
	}
	val, ok := md[key]
	return val, ok
}

// FromContext returns metadata from the given context
func FromContext(ctx context.Context) (Metadata, bool) {
	md, ok := ctx.Value(metaKey{}).(Metadata)
	return md, ok
}

// NewContext creates a new context with the given metadata
func NewContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, metaKey{}, md)
}

// MergeContext merges metadata to existing metadata, overwriting if specified
func MergeContext(ctx context.Context, patchMd Metadata, overwrite bool) context.Context {
	md, _ := ctx.Value(metaKey{}).(Metadata)
	cmd := make(Metadata)
	for k, v := range md {
		cmd[k] = v
	}
	for k, v := range patchMd {
		if _, ok := cmd[k]; ok && !overwrite {
			// skip
		} else {
			cmd[k] = v
		}
	}
	return context.WithValue(ctx, metaKey{}, cmd)

}

// GetContext 获取context 上线
func GetContext(c *gin.Context) context.Context {
	if ctx, ok := c.Get("Context"); ok {
		return ctx.(context.Context)
	} else {
		return context.Background()
	}
}

// GetDiffDays 计算日期相差天数
func GetDiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

// GetDiffDaysBySecond 计算日期相差天数-参数为时间戳时
func GetDiffDaysBySecond(t1, t2 int64) int {
	time1 := time.Unix(t1, 0)
	time2 := time.Unix(t2, 0)

	// 调用上面的函数
	return GetDiffDays(time1, time2)
}
