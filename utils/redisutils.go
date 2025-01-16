package utilslib

import (
	"bloxley-backend-utility-library/constants"
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
)

// RedisConfig holds the connection details for Redis.
// The calling application is responsible for populating this struct.
type RedisConfig struct {
	Address  string // "host:port" e.g. "localhost:6379" or "10.0.0.5:6379"
	Password string // Redis password (can be empty if none)
	DB       int    // Redis DB index (e.g. 0, 1, 2, ...)
}

var (
	clientInstance *redis.Client
	once           sync.Once
)

// InitializeRedis initializes and returns a singleton Redis client instance
// using the given RedisConfig.
//
// Example usage from the calling service:
//
//	cfg := utilslib.RedisConfig{
//	    Address:  "10.0.0.5:6379",
//	    Password: "my-secret-pwd",
//	    DB:       1,
//	}
//	client := utils.InitializeRedis(cfg)
//	err := client.Set(context.Background(), "key", "value", 0).Err()
func InitializeRedis(config RedisConfig) *redis.Client {
	once.Do(func() {
		fmt.Printf("[utilslib] Connecting to Redis at %s (DB %d)\n", config.Address, config.DB)

		clientInstance = redis.NewClient(&redis.Options{
			Addr:     config.Address,
			Password: config.Password,
			DB:       config.DB,
		})

		// Optionally, you could do a quick ping to verify the connection:
		if err := clientInstance.Ping(context.Background()).Err(); err != nil {
			panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
		}
	})

	return clientInstance
}

// GetMBanqToken fetches a JWT token stored under the "mbanq_token" key in Redis.
// Returns the token, or an error if it isn't found or if a Redis error occurs.
func GetMBanqToken(ctx context.Context, client *redis.Client) (string, error) {
	token, err := client.Get(ctx, constants.MbanqTokenKey).Result()
	if err == redis.Nil {
		// Key does not exist
		return "", errors.New("mbanq_token not found in Redis")
	} else if err != nil {
		// Some other Redis-related error
		return "", err
	}

	// Return the token value
	return token, nil
}
