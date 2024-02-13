package modules

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

//noinspection GoUnusedExportedFunction
func RedisInitiateRedisPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 10,
		// max number of connections
		MaxActive: 1000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			redisStrConn := "redis://:" + MapConfig["redisPass"] + "@" + MapConfig["redisHost"] + ":" + MapConfig["redisPort"] + "/" + MapConfig["redisDB"]

			fmt.Println(" ------- " + redisStrConn + " -------- ")
			//c, err := redis.DialURL("redis://user:Eliandri3@localhost:6379/0")
			c, err := redis.DialURL(redisStrConn)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

//noinspection GoUnusedExportedFunction
func RedisSetDataRedis(traceCode string, redisKey string, redisVal string) bool {
	result := true

	// RedisPooler need to be initiated in caller function
	redisConn := RedisPooler.Get()

	_, err := redisConn.Do("SET", redisKey, redisVal)

	if err != nil {
		DoLog("DEBUG", traceCode, "APIGetToken", "SetDataRedis",
			"Set data redis key: "+redisKey+" is failed. Error occured.", true, err)

		result = false
	} else {
		result = true
	}

	// Close redisConn
	_ = redisConn.Close()

	DoLog("DEBUG", traceCode, "APIGetToken", "SetDataRedis",
		"Set data redis key: "+redisKey+", result: "+fmt.Sprintf("%t", result), false, nil)

	return result
}

//noinspection GoUnusedExportedFunction
func RedisSetDataRedisWithExpiry(traceCode string, redisKey string, redisVal string, expiry int32) bool { // expiry in second
	result := true

	// RedisPooler need to be initiated in caller function
	redisConn := RedisPooler.Get()

	_, err := redisConn.Do("SET", redisKey, redisVal, "EX", expiry)

	if err != nil {
		result = false
	} else {
		result = true
	}

	// Close redisConn
	_ = redisConn.Close()

	DoLog("DEBUG", traceCode, "APIGetToken", "SetDataRedis",
		"Set data redis key: "+redisKey+" with expiry: "+strconv.Itoa(int(expiry))+" seconds , result: "+
			fmt.Sprintf("%t", result), false, nil)

	return result
}

//noinspection GoUnusedExportedFunction
func RedisGetDataRedis(traceCode string, redisKey string) string {
	//fmt.Println("GET redisKey: " + redisKey)

	// RedisPooler need to be initiated in caller function
	redisConn := RedisPooler.Get()

	hasilGet, err := redis.String(redisConn.Do("GET", redisKey))

	//fmt.Println(fmt.Sprintf("hasilGet: %v", hasilGet))

	result := ""
	if err != nil {
		result = ""
	} else {
		result = hasilGet
	}

	// Close redisConn
	_ = redisConn.Close()

	//DoLog("DEBUG", traceCode, "APIGetToken", "GetDataRedis",
	//	"Get data redis key: " + redisKey + ", redis val: " + result,false, nil)

	return result
}

func RedisDelDataRedis(messageId string, redisKey string) bool {
	result := true

	// RedisPooler need to be initiated in caller function
	redisConn := RedisPooler.Get()

	_, err := redisConn.Do("DEL", redisKey)

	if err != nil {
		result = false
	} else {
		result = true
	}

	// Close redisConn
	_ = redisConn.Close()

	DoLog("DEBUG", messageId, "APIGetToken", "DelDataRedis",
		"Del data redis - redisKey: "+redisKey+", result: "+
			fmt.Sprintf("%t", result), false, nil)

	return result
}

//noinspection GoUnusedExportedFunction
func RedisDelKeysWithPatternRedis(messageId string, redisKeyPattern string) bool {
	// RedisPooler need to be initiated in caller function
	redisConn := RedisPooler.Get()

	keys, err := redis.Strings(redisConn.Do("KEYS", redisKeyPattern))

	_ = redisConn.Close()

	if err != nil {
		return false
	} else {
		for _, key := range keys {
			fmt.Println("Deleting key: " + key)
			RedisDelDataRedis(messageId, key)
		}

		return true
	}
}

func RedisGetKeysWithPatternRedis(messageId string, redisKeyPattern string) []string {
	var arrResult []string
	redisConn := RedisPooler.Get()

	keys, _ := redis.Strings(redisConn.Do("KEYS", redisKeyPattern))

	_ = redisConn.Close()

	for _, key := range keys {
		arrResult = append(arrResult, key)
	}

	return arrResult
}
