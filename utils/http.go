package utils

import (
	"context"
)

func GetUserIDFromContext(ctx context.Context) string {
	return ctx.Value("x-user-id").(string)
}
