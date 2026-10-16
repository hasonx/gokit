package idgen

import (
	"crypto/rand"
	mrand "math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
)

type ulidGenerator struct {
	entropy *ulid.MonotonicEntropy
	mu      sync.Mutex
}

func newULIDGenerator(opts ...Option) (IDGenerator, error) {
	cfg := &OptConfig{Secure: false}

	for _, opt := range opts {
		opt(cfg)
	}

	var entropy *ulid.MonotonicEntropy
	if cfg.Secure {
		// 加密安全随机数源（适合订单号、公开ID）
		entropy = ulid.Monotonic(rand.Reader, 0)
	} else {
		// 高性能伪随机源（适合内部日志/事件）
		src := mrand.NewSource(time.Now().UnixNano())
		entropy = ulid.Monotonic(mrand.New(src), 0)
	}
	return &ulidGenerator{entropy: entropy}, nil
}

func (g *ulidGenerator) NextID() string {
	g.mu.Lock()
	defer g.mu.Unlock()
	return ulid.MustNew(ulid.Timestamp(time.Now().UTC()), g.entropy).String()
}
