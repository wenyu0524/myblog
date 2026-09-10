package userAdmin

import (
	"context"
	"encoding/json"
	"errors"
)

func userIdFromCtx(ctx context.Context) (int64, error) {
	switch v := ctx.Value("userId").(type) {
	case json.Number:
		return v.Int64()
	case float64:
		return int64(v), nil
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	default:
		return 0, errors.New("未认证")
	}
}
