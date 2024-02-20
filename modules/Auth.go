package modules

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v8"
)

func reloadGuestSessionToRedis(db *sql.DB, rc *redis.Client, cx context.Context, incUserID string) (string, bool) {
	isSuccess := false
	strSession := ""
	
	return strSession, isSuccess
}