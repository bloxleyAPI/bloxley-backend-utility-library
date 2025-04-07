package utilslibs

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

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
		fmt.Printf("[utilslib] Connecting to Redis at %s\n", config.REDIS_ADDRESS)

		clientInstance = redis.NewClient(&redis.Options{
			Addr:     config.REDIS_ADDRESS + ":" + config.REDIS_PORT,
			Password: config.REDIS_PASSWORD,
			DB:       config.REDIS_DB,
		})

		if clientInstance == nil {
			panic("Failed to create Redis client")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Optionally, you could do a quick ping to verify the connection:
		if err := clientInstance.Ping(ctx).Err(); err != nil {
			panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
		} else {
			fmt.Println("Connected to Redis")
		}
	})

	return clientInstance
}

// RedisValue is a simple struct to hold a string value.
type RedisValue struct {
	Value string `json:"value"`
}

// SetRedisValueToken sets a key-value pair in Redis with the given expiration time.
// GetRedisValueToken retrieves the Redis token value corresponding to the provided key.
// It uses the given Redis client and context to perform the lookup.
//
// If the key does not exist in Redis, it returns an error indicating that the key was not found.
// If any other Redis-related error occurs, it returns that error as well.
// On a successful lookup, it wraps the retrieved value in a RedisValue struct and returns a pointer to it.
func GetRedisValueToken(ctx context.Context, client *redis.Client, key string) (*RedisValue, error) {
	value, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		// Key does not exist
		return nil, errors.New(key + " not found in Redis")
	} else if err != nil {
		// Some other Redis-related error
		return nil, err
	}

	// Return the token wrapped in an
	return &RedisValue{Value: value}, nil
}
