package modules

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func InitiateRedisClient() *redis.Client {
	theDBNo, _ := strconv.Atoi(MapConfig["redisDB"])
	rdb := redis.NewClient(&redis.Options{
		Addr:     MapConfig["redisHost"] + ":" + MapConfig["redisPort"],
		Password: MapConfig["redisPassword"],
		DB:       theDBNo,
	})

	return rdb
}

func InitiateRedisStoreManagement() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     MapConfig["redisHost"] + ":" + MapConfig["redisPort"],
		Password: MapConfig["redisPassword"],
		DB:       1,
	})

	return rdb
}

func InitiateRedisClientWebAdmin() *redis.Client {
	theDBNo, _ := strconv.Atoi(MapConfig["redisAdminDB"])
	rdb := redis.NewClient(&redis.Options{
		Addr:     MapConfig["redisAdminHost"] + ":" + MapConfig["redisAdminPort"],
		Password: MapConfig["redisAdminPass"],
		DB:       theDBNo,
	})

	return rdb
}

func RedisGet(redisClient *redis.Client, goContext context.Context, redisKey string) (string, error) {
	theValue, theError := redisClient.Get(goContext, redisKey).Result()
	if theError == redis.Nil {
		DoLog("DEBUG", "", "RedisNeo", "RedisGet", "key "+redisKey+" does not exist", true, theError)
	} else if theError != nil {
		DoLog("DEBUG", "", "RedisNeo", "RedisGet", fmt.Sprintf("error get value: %+v", redisClient), true, theError)
	}
	return theValue, theError
}

func RedisSet(redisClient *redis.Client, goContext context.Context, redisKey string, redisVal string, expiration time.Duration) error {
	return redisClient.Set(goContext, redisKey, redisVal, expiration).Err()
}

func RedisDel(redisClient *redis.Client, goContext context.Context, redisKey string) error {
	return redisClient.Del(goContext, redisKey).Err()
}

func RedisKeysByPattern(redisClient *redis.Client, goContext context.Context, redisKeyPattern string) (bool, []string) {
	result, err := redisClient.Do(goContext, "keys", redisKeyPattern).Result()

	if err != nil {
		DoLog("DEBUG", "", "RedisNeo", "RedisKeysByPattern", fmt.Sprintf("error get value: %+v", result), true, err)
		return false, nil
	} else {
		arrResult := result.([]interface{})

		finalArrResult := make([]string, len(arrResult))
		for index, value := range arrResult {
			strValue := value.(string)
			//fmt.Printf("%d - %s\n", index, strValue)

			finalArrResult[index] = strValue
		}

		return true, finalArrResult
	}
}

func RedisDeleteKeysByPattern(redisClient *redis.Client, goContext context.Context, redisKeyPattern string) {
	_, arrKeys := RedisKeysByPattern(redisClient, goContext, redisKeyPattern)

	for _, value := range arrKeys {
		err := RedisDel(redisClient, goContext, value)

		if err != nil {
			DoLog("DEBUG", "", "RedisNeo", "RedisDeleteKeysByPattern",
				"Failed to delete redisKey: "+value, false, nil)
		}
	}
}

func RedisRefreshSession(redisClient *redis.Client, goContext context.Context, redisKey string, expiration time.Duration) error {
	return redisClient.Expire(goContext, redisKey, expiration).Err()

}
