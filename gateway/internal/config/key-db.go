package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewKeyDBConn(ctx context.Context) (*redis.Client, error) {
	keydb := redis.NewClient(&redis.Options{
		Addr:            GetValue(DbDsn),
		ClientName:      "gateway",
		DB:              0,
		ConnMaxIdleTime: 15 * time.Minute,
	})

	err := keydb.Ping(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return keydb, nil
}
