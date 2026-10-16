package idgen

import (
	"fmt"
	"sync"
)

// generatorManager 统一管理所有 ID 生成器
type generatorManager struct {
	mu         sync.RWMutex
	generators map[string]IDGenerator
}

var (
	managerOnce sync.Once
	manager     *generatorManager
)

type Option func(*OptConfig)

type OptConfig struct {
	WorkerID int64
	Secure   bool
}

// 获取全局单例
func getManager() *generatorManager {
	managerOnce.Do(func() {
		manager = &generatorManager{
			generators: make(map[string]IDGenerator),
		}
	})
	return manager
}

// Register 注册一个新的 ID 生成器
func Register(name string, idType IDType, opts ...Option) error {
	m := getManager()
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.generators[name]; exists {
		return fmt.Errorf("id generator '%s' already registered", name)
	}

	var gen IDGenerator
	var err error

	switch idType {
	case IDTypeULID:
		gen, err = newULIDGenerator(opts...)
	case IDTypeSnowflake:
		gen, err = newSnowflakeGeneratorWithOpts(opts...)
	case ID_UUID:
		gen = &UUIDGen{}
	default:
		return fmt.Errorf("unsupported id type: %s", idType)
	}

	if err != nil {
		return err
	}

	m.generators[name] = gen
	return nil
}

// Get 获取指定生成器
func Get(name string) (IDGenerator, error) {
	m := getManager()
	m.mu.RLock()
	defer m.mu.RUnlock()

	gen, ok := m.generators[name]
	if !ok {
		return nil, fmt.Errorf("id generator '%s' not found", name)
	}
	return gen, nil
}

// MustGet 获取指定生成器（panic 版本）
func MustGet(name string) IDGenerator {
	gen, err := Get(name)
	if err != nil {
		panic(err)
	}
	return gen
}

// NextID 获取指定生成器生成的 ID
func NextID(name string) string {
	return MustGet(name).NextID()
}
