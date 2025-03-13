package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"russian-roulette/internal/config"
	"russian-roulette/internal/entities/custom_errors"
	"time"
)

type CacheService interface {
	SetExpire(key string, value string, expiration time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
	PushToQueue(key string, value string) error
	PushArrayToQueue(key string, values []string) error
	PopFromQueue(key string) (string, error)
	PeekQueue(key string) (string, error)
	QueueLength(key string) (int64, error)
	GetQueueValues(key string) ([]string, error)
	SetMap(key string, value map[string]interface{}, expiration time.Duration) error
	GetMap(key string) (map[string]interface{}, error)
	DeleteFromMap(key string, field string, expiration time.Duration) error
	RemoveFromQueue(key, value string) error
}

type RedisCache struct {
	client *redis.Client

	cfg    *config.Redis
	logger *zap.Logger
}

func NewRedisCache(cfg *config.Redis, logger *zap.Logger) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Url,
		Password: cfg.Password,
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(errors.Wrap(err, "redis ping failed"))
	}

	return &RedisCache{
		client: client,
		logger: logger,
	}
}

// SetExpire устанавливает значение с заданным временем истечения.
// key: ключ для хранения значения.
// value: значение, которое нужно сохранить.
// expiration: длительность времени истечения.
func (r *RedisCache) SetExpire(key string, value string, expiration time.Duration) error {
	err := r.client.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return custom_errors.ErrRedisFailedSetValue(value, err)
	}
	return nil
}

// Get получает значение по ключу.
// key: ключ для извлечения значения.
// Возвращает значение и ошибку, если есть.
func (r *RedisCache) Get(key string) (string, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if errors.Is(err, redis.Nil) {
		return "", custom_errors.ErrRedisKeyNotFound(key, err)
	} else if err != nil {
		return "", custom_errors.ErrRedisFailedGetValue(key, err)
	}
	return val, nil
}

// Delete удаляет значение по ключу.
// key: ключ для удаления значения.
// Возвращает ошибку, если есть.
func (r *RedisCache) Delete(key string) error {
	err := r.client.Del(context.Background(), key).Err()
	if err != nil {
		return custom_errors.ErrRedisFailedDeleteValue(key, err)
	}
	return nil
}

// PushToQueue добавляет элемент в конец очереди.
// key: имя очереди.
// value: значение, которое нужно добавить в очередь.
func (r *RedisCache) PushToQueue(key string, value string) error {
	err := r.client.RPush(context.Background(), key, value).Err()
	if err != nil {
		return custom_errors.ErrRedisFailedSetValue(value, err)
	}
	return nil
}

// PushArrayToQueue добавляет массив значений в конец очереди.
// key: имя очереди.
// values: массив значений, которые нужно добавить в очередь.
func (r *RedisCache) PushArrayToQueue(key string, values []string) error {
	err := r.client.RPush(context.Background(), key, values).Err()
	if err != nil {
		return custom_errors.ErrRedisFailedSetValue(fmt.Sprintf("key: %s, values: %v", key, values), err)
	}
	return nil
}

// PopFromQueue удаляет элемент из начала очереди.
// key: имя очереди.
// Возвращает значение элемента и ошибку, если есть.
func (r *RedisCache) PopFromQueue(key string) (string, error) {
	val, err := r.client.LPop(context.Background(), key).Result()
	if errors.Is(err, redis.Nil) {
		return "", custom_errors.ErrRedisKeyNotFound(key, err)
	} else if err != nil {
		return "", custom_errors.ErrRedisFailedGetValue(key, err)
	}
	return val, nil
}

// QueueLength возвращает количество элементов в очереди.
// key: имя очереди.
func (r *RedisCache) QueueLength(key string) (int64, error) {
	length, err := r.client.LLen(context.Background(), key).Result()
	if err != nil {
		return 0, custom_errors.ErrRedisFailedGetValue(key, err)
	}
	return length, nil
}

// GetQueueValues возвращает все значения из очереди.
// key: имя очереди.
// Возвращает срез строк и ошибку, если есть.
func (r *RedisCache) GetQueueValues(key string) ([]string, error) {
	values, err := r.client.LRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		return nil, custom_errors.ErrRedisFailedGetValue(key, err)
	}
	return values, nil
}

// PeekQueue возвращает текущее значение из начала очереди без удаления.
// key: имя очереди.
// Возвращает строку и ошибку, если есть.
func (r *RedisCache) PeekQueue(key string) (string, error) {
	val, err := r.client.LIndex(context.Background(), key, 0).Result()
	if errors.Is(err, redis.Nil) {
		return "", custom_errors.ErrRedisKeyNotFound(key, err)
	} else if err != nil {
		return "", custom_errors.ErrRedisFailedGetValue(key, err)
	}
	return val, nil
}

// SetMap сохраняет карту (map) в Redis с заданным временем истечения.
// key: ключ для хранения карты.
// value: карта, которую нужно сохранить.
// expiration: длительность времени истечения.
func (r *RedisCache) SetMap(key string, value map[string]interface{}, expiration time.Duration) error {
	// Сериализуем map в JSON
	jsonData, err := json.Marshal(value)
	if err != nil {
		return errors.Wrap(err, "failed to marshal map to JSON")
	}

	// Сохраняем JSON как строку
	err = r.client.Set(context.Background(), key, string(jsonData), expiration).Err()
	if err != nil {
		return custom_errors.ErrRedisFailedSetValue(key, err)
	}
	return nil
}

// DeleteFromMap удаляет элемент из карты (map), хранящейся в Redis.
// key: ключ карты в Redis.
// field: ключ элемента, который нужно удалить.
// expiration: обновленное время истечения (если нужно сохранить).
func (r *RedisCache) DeleteFromMap(key string, field string, expiration time.Duration) error {
	// Извлекаем карту из Redis
	currentMap, err := r.GetMap(key)
	if err != nil {
		return errors.Wrap(err, "failed to get map from Redis")
	}

	// Удаляем элемент из карты
	delete(currentMap, field)

	// Сохраняем обновленную карту обратно в Redis
	err = r.SetMap(key, currentMap, expiration)
	if err != nil {
		return errors.Wrap(err, "failed to update map in Redis")
	}

	return nil
}

// GetMap извлекает карту (map) из Redis.
// key: ключ для извлечения карты.
// Возвращает карту и ошибку, если есть.
func (r *RedisCache) GetMap(key string) (map[string]interface{}, error) {
	// Получаем JSON строку
	jsonData, err := r.client.Get(context.Background(), key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, custom_errors.ErrRedisKeyNotFound(key, err)
	} else if err != nil {
		return nil, custom_errors.ErrRedisFailedGetValue(key, err)
	}

	// Десериализуем JSON обратно в map
	var result map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal JSON to map")
	}

	return result, nil
}

func (r *RedisCache) RemoveFromQueue(key, value string) error {
	return r.client.LRem(context.Background(), key, 0, value).Err()
}
