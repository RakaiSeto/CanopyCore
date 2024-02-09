package modules

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const strSalt = "SayaBisaKarenaKebaikanTuhanSaja"

func generateSignature(incPhoneNumber string, incClientID string, incVariousID string) string {

	_, encPhoneNumber := SimpleStringEncryptWithSalt(incPhoneNumber, incVariousID)
	_, encClientID := SimpleStringEncryptWithSalt(incClientID, incVariousID)

	rawSignature := encPhoneNumber + "||" + encClientID + "||" + incVariousID
	_, strSignature := SimpleStringEncryptWithSalt(rawSignature, strSalt)

	return strSignature
}

func CreateRedisSession(db *sql.DB, redisClient *redis.Client, contextX context.Context, incTraceID string, incClientID string, incPhoneNumber string, incPIN string) (string, bool) {

	strSignature := ""
	status := false

	query := "SELECT user_pin FROM seru " +
		"WHERE client_id = $1 AND user_phone = $2 AND is_active = $3"

	rows, err := db.Query(query, incClientID, incPhoneNumber, true)

	if err != nil {
		// Database Error
		DoLog("ERROR", incTraceID, "RPCModules", "RedisSession",
			"Failed to read database. Error occur.", true, err)
	} else {
		// Success
		for rows.Next() {

			var rawPIN sql.NullString

			errS := rows.Scan(&rawPIN)

			if errS != nil {
				DoLog("ERROR", incTraceID, "RPCModules", "RedisSession",
					"Failed to read database. Error occur.", true, errS)
			} else {
				DoLog("INFO", incTraceID, "RPCModules", "RedisSession",
					"Success to read database", false, nil)

				strPIN := ConvertSQLNullStringToString(rawPIN)

				errX := bcrypt.CompareHashAndPassword([]byte(strPIN), []byte(incPIN))

				if errX == nil {
					DoLog("INFO", incTraceID, "RPCModules", "RedisSession",
						"Credential is Valid", false, nil)

					variousID := GenerateUUID()
					strSignature = generateSignature(incPhoneNumber, incClientID, variousID)

					redisKey := "tokenizer_" + incPhoneNumber

					mapRedis := make(map[string]interface{})
					mapRedis["signature"] = strSignature

					// Save to Destination Redis
					redisValX := ConvertMapInterfaceToJSON(mapRedis)
					RedisSet(redisClient, contextX, redisKey, redisValX, 0)

					status = true

				} else {
					DoLog("ERROR", incTraceID, "RPCModules", "RedisSession",
						"Credential is InValid", false, nil)
				}
			}
		}
	}

	return strSignature, status
}

func expandRefreshRedisSession(db *sql.DB, redisClient *redis.Client, contextX context.Context, incTraceID string, incClientID string, incPhoneNumber string) string {

	strSignature := ""

	variousID := GenerateUUID()
	strSignature = generateSignature(incPhoneNumber, incClientID, variousID)

	redisKey := "tokenizer_" + incPhoneNumber

	mapRedis := make(map[string]interface{})
	mapRedis["signature"] = strSignature

	// Save to Destination Redis
	redisValX := ConvertMapInterfaceToJSON(mapRedis)
	RedisSet(redisClient, contextX, redisKey, redisValX, 0)

	return strSignature
}

func ValidateRedisSession(db *sql.DB, redisClient *redis.Client, contextX context.Context, incTraceID string, incClientID string, incPhoneNumber string, incSessionID string) (string, string) {

	strNewSignature := ""
	redisKey := "tokenizer_" + incPhoneNumber

	redisVal, err := RedisGet(redisClient, contextX, redisKey)

	if err == nil {
		DoLog("INFO", incTraceID, "RPCModules", "Validate",
			"Success to read Rediskey : "+redisKey, false, nil)

		mapRedisVal := ConvertJSONStringToMap("", redisVal)

		// - Read Redis
		strSignature := GetStringFromMapInterface(mapRedisVal, "signature")

		if len(strSignature) > 0 {
			DoLog("INFO", incTraceID, "RPCModules", "Validate",
				"Redis Valid for Rediskey : "+redisKey, false, nil)

			if strSignature == incSessionID {
				// Ada session
				fmt.Println("session Sama")
				strSignatureX := SimpleStringDecryptWithSalt(strSignature, strSalt)
				arrSignature := strings.Split(strSignatureX, "||")
				decPhoneNumber := SimpleStringDecryptWithSalt(arrSignature[0], arrSignature[2])
				decClientiD := SimpleStringDecryptWithSalt(arrSignature[1], arrSignature[2])

				if decPhoneNumber == incPhoneNumber && decClientiD == incClientID {
					strNewSignature = expandRefreshRedisSession(db, redisClient, contextX, incTraceID, incClientID, incPhoneNumber)

					return strNewSignature, "001" // Need to Insert PIN
				} else {
					return "", "003" // Need to input phone number
				}
			} else {
				// Tidak ada session

				query := "SELECT count(*) FROM seru " +
					"WHERE client_id = $1 AND user_phone = $2 AND is_active = $3"

				rows, errZ := db.Query(query, incClientID, incPhoneNumber, true)

				if errZ != nil {
					// Database Error
					DoLog("ERROR", incTraceID, "RPCModules", "RedisSession",
						"Failed to read database. Error occur.", true, errZ)
				} else {
					// Success
					for rows.Next() {

						var rawCount sql.NullFloat64

						errS := rows.Scan(&rawCount)

						if errS != nil {
							DoLog("ERROR", incTraceID, "RPCModules", "RedisSession",
								"Failed to read database. Error occur.", true, errS)
						} else {
							DoLog("INFO", incTraceID, "RPCModules", "RedisSession",
								"Success to read database", false, nil)

							fltCount := ConvertSQLNullFloat64ToFloat64(rawCount)

							if fltCount > 0 {
								// Ada didatabase
								return "", "002" // Need to input pin
							} else {
								// Tidak ada didatabase
								return "", "003" // Need to Login
							}
						}
					}
				}

				fmt.Println("session Beda")
				return "", "003"
			}
		} else {
			DoLog("ERROR", incTraceID, "RPCModules", "Validate",
				"Redis InValid for Rediskey : "+redisKey, false, nil)

			return "", "003"
		}
	} else {
		DoLog("ERROR", incTraceID, "RPCModules", "Validate",
			"Not found redis key : "+redisKey, true, err)

		return "", "003"
	}
}
