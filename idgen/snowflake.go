package idgen

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

type snowflakeGenerator struct {
	node *snowflake.Node
}

// 默认 Option
func WithWorkerID(id int64) Option {
	return func(c *OptConfig) {
		c.WorkerID = id
	}
}

func newSnowflakeGeneratorWithOpts(opts ...Option) (IDGenerator, error) {
	cfg := &OptConfig{WorkerID: 1}

	for _, opt := range opts {
		opt(cfg)
	}

	// 也可通过环境变量覆盖
	if env := os.Getenv("WORKER_ID"); env != "" {
		if val, err := strconv.ParseInt(env, 10, 64); err == nil {
			cfg.WorkerID = val
		}
	}

	node, err := snowflake.NewNode(cfg.WorkerID)
	if err != nil {
		return nil, fmt.Errorf("failed to init snowflake node: %w", err)
	}

	return &snowflakeGenerator{node: node}, nil
}

func (g *snowflakeGenerator) NextID() string {
	return g.node.Generate().String()
}
