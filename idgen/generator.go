package idgen

type IDType string

const (
	IDTypeULID      IDType = "ulid"
	IDTypeSnowflake IDType = "snowflake"
	ID_UUID         IDType = "uuid"
)

// IDGenerator 定义生成器通用接口
type IDGenerator interface {
	NextID() string
}
