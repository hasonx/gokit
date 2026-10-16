package idgen

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config 文件结构
type Config struct {
	Generators map[string]GeneratorConfig `yaml:"generators"`
}

type GeneratorConfig struct {
	Type     IDType `yaml:"type"`
	WorkerID int64  `yaml:"worker_id,omitempty"`
	Secure   bool   `yaml:"secure,omitempty"`
}

// LoadConfig 从文件加载配置
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	return &cfg, nil
}

// InitFromConfig 初始化全部生成器
func InitFromConfig(path string) error {
	cfg, err := LoadConfig(path)
	if err != nil {
		return err
	}

	for name, g := range cfg.Generators {
		var opts []Option
		if g.WorkerID > 0 {
			opts = append(opts, WithWorkerID(g.WorkerID))
		}

		if err := Register(name, g.Type, opts...); err != nil {
			return fmt.Errorf("failed to register generator '%s': %w", name, err)
		}
	}
	return nil
}
