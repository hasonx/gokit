# gokit
a golang toolkit

# 使用示例

##idgen使用示例

```go
package main

import (
    "fmt"
    "log"
    "github.com/hasonx/gokit/idgen"
)

func main() {
    // 自动初始化所有生成器
    if err := idgen.InitFromConfig("configs/idgen.yaml"); err != nil {
        log.Fatalf("init idgen failed: %v", err)
    }

    // 直接使用
    fmt.Println("user:", idgen.NextID("user"))
    fmt.Println("order:", idgen.NextID("order"))
    fmt.Println("log:", idgen.NextID("log"))
}
```

ulid

```sql
CREATE TABLE event_log (
  id CHAR(26) PRIMARY KEY,
  user_id BIGINT,
  action VARCHAR(64),
  created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

snowflake

```sql
CREATE TABLE orders (
  id BIGINT UNSIGNED PRIMARY KEY,
  user_id BIGINT,
  total DECIMAL(10,2),
  created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3)
);
```

