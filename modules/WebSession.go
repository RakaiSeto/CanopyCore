package modules

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type WebSessionData struct {
	Email     string
	Role      string
	ClientId  string
	Phone     string
	FirstName string
	LastName  string
}

func RefreshWebRedisSession(rc *redis.Client, ctx context.Context, incTraceCode string, incRemoteIp string, incRequest WebSessionData) (string, bool) {
	isSuccess := false
	strSession := ""

	incRequestEmail := incRequest.Email
	incRequestRole := incRequest.Role
	incRequestClientId := incRequest.ClientId
	incRequestPhone := incRequest.Phone
	incRequestFirstName := incRequest.FirstName
	incRequestLastName := incRequest.LastName
	incRequstExpiry := 24 * time.Hour

	_, encRequestEmail := SimpleStringEncrypt(incRequestEmail)
	_, encRequestRole := SimpleStringEncrypt(incRequestRole)
	_, encRequestClientId := SimpleStringEncrypt(incRequestClientId)
	_, encRequestPhone := SimpleStringEncrypt(incRequestPhone)
	_, encRequestFirstName := SimpleStringEncrypt(incRequestFirstName)
	_, encRequestLastName := SimpleStringEncrypt(incRequestLastName)

	strSignature := encRequestEmail + "||" + encRequestRole + "||" + encRequestClientId + "||" + encRequestPhone + "||" + encRequestFirstName + "||" + encRequestLastName
	_, encSignature := SimpleStringEncrypt(strSignature)

	redisKey := "web_session_" + incRequestEmail
	errRedis := RedisSet(rc, ctx, redisKey, encSignature, incRequstExpiry)
	if errRedis != nil {
		Logging(Resource(), incTraceCode, incRequestEmail, incRemoteIp, "Failed to set redis", errRedis)
		isSuccess = false
	} else {
		isSuccess = true
		strSession = encSignature
	}

	return strSession, isSuccess
}
