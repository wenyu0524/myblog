package logic

import (
	"context"
	"encoding/json"
	"fmt"
)

// getUserId 从 context 中提取 JWT 解析后的 userId
// go-zero JWT 中间件会将自定义 claim 以 key 为 "userId" 存入 context
// 由于 JSON 解码后数值类型可能为 float64 或 json.Number，需做类型断言兼容
func getUserId(ctx context.Context) (int64, error) {
	uid := ctx.Value("userId")
	if uid == nil {
		return 0, fmt.Errorf("user id not found in context")
	}
	switch v := uid.(type) {
	case int64:
		return v, nil
	case float64:
		return int64(v), nil
	case json.Number:
		return v.Int64()
	default:
		return 0, fmt.Errorf("unexpected type for user id: %T", v)
	}
}
