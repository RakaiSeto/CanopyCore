package RPCHelper

import (
	error_code "CanopyCore/errors"
	rootadminweb "CanopyCore/grpc/rootadminweb"
	"CanopyCore/modules"
	"context"
	"database/sql"
	"github.com/go-redis/redis/v8"
)

func HelperDoLoginRootAdminWeb(ctx context.Context, db *sql.DB, rc *redis.Client, incMapRequest map[string]interface{}) (string, string, []*rootadminweb.DoLoginList) {
	localResponse := new(rootadminweb.DoLoginResponse)
	strStatus := error_code.ErrorServer
	strSession := ""

	//incTimeNow := modules.DoFormatDateTime("YYYY-0M-0D HH:mm:ss.S", time.Now())
	incRequestTraceCode := modules.GetStringFromMapInterface(incMapRequest, "tracecode")
	incRequestRemoteIP := modules.GetStringFromMapInterface(incMapRequest, "remoteip")
	incRequestEmail := modules.GetStringFromMapInterface(incMapRequest, "email")
	incRequestPassword := modules.GetStringFromMapInterface(incMapRequest, "password")

	query := `SELECT u.userid, gr.description, u.email, u.first_name, u.last_name, u.user_password, c.client_id, u.phone, u.roleid FROM web_users u 
    		left join global_role gr on gr.role = u.roleid 
            left join client c on c.client_id = u.client_id
            WHERE u.email = $1 and u.is_active = $2`
	rows, err := db.Query(query, incRequestEmail, true)
	if err != nil {
		modules.Logging(modules.Resource(), incRequestTraceCode, "SERVER", incRequestRemoteIP, "error when query", err)
		strStatus = error_code.ErrorServer
	} else {
		defer rows.Close()
		for rows.Next() {
			var rawUserID sql.NullString
			var rawRole sql.NullString
			var rawEmail sql.NullString
			var rawFirstName sql.NullString
			var rawLastName sql.NullString
			var rawPassword sql.NullString
			var rawClientId sql.NullString
			var rawPhone sql.NullString
			var rawRoleID sql.NullString

			errS := rows.Scan(&rawUserID, &rawRole, &rawEmail, &rawFirstName, &rawLastName, &rawPassword, &rawClientId, &rawPhone, &rawRoleID)
			if errS != nil {
				if errS == sql.ErrNoRows {
					modules.Logging(modules.Resource(), incRequestTraceCode, incRequestEmail, incRequestRemoteIP, "error when scan", errS)
					strStatus = error_code.ErrorBadAccount
				} else {
					modules.Logging(modules.Resource(), incRequestTraceCode, incRequestEmail, incRequestRemoteIP, "error when scan", errS)
					strStatus = error_code.ErrorServer
				}
			} else {
				strUserID := modules.ConvertSQLNullStringToString(rawUserID)
				strRole := modules.ConvertSQLNullStringToString(rawRole)
				strEmail := modules.ConvertSQLNullStringToString(rawEmail)
				strFirstName := modules.ConvertSQLNullStringToString(rawFirstName)
				strLastName := modules.ConvertSQLNullStringToString(rawLastName)
				strPassword := modules.ConvertSQLNullStringToString(rawPassword)
				strClientId := modules.ConvertSQLNullStringToString(rawClientId)
				strPhone := modules.ConvertSQLNullStringToString(rawPhone)
				strRoleID := modules.ConvertSQLNullStringToString(rawRoleID)

				isValidPassword := modules.CheckBCryptHashPassword(incRequestPassword, strPassword)
				if !isValidPassword {
					modules.Logging(modules.Resource(), incRequestTraceCode, incRequestEmail, incRequestRemoteIP, "Invalid Password", nil)
					strStatus = error_code.ErrorBadPassword
				} else {
					redisRequest := modules.WebSessionData{
						Email:     strEmail,
						Role:      strRole,
						ClientId:  strClientId,
						Phone:     strPhone,
						FirstName: strFirstName,
						LastName:  strLastName,
					}

					resultSession, isRedisSuccess := modules.RefreshWebRedisSession(rc, ctx, incRequestTraceCode, incRequestRemoteIP, redisRequest)
					if !isRedisSuccess {
						modules.Logging(modules.Resource(), incRequestTraceCode, incRequestEmail, incRequestRemoteIP, "Failed to create session to redis", nil)
						strStatus = error_code.ErrorServer
					} else {
						strStatus = error_code.ErrorSuccess
						strSession = resultSession
						modules.Logging(modules.Resource(), incRequestTraceCode, incRequestEmail, incRequestRemoteIP, "Login success", nil)
						localResponse.Result = append(localResponse.Result, &rootadminweb.DoLoginList{
							Id:        strUserID,
							Roleid:    strRoleID,
							Email:     strEmail,
							Phone:     strPhone,
							Firstname: strFirstName,
							Lastname:  strLastName,
							Fullname:  strFirstName + " " + strLastName,
							Clientid:  strClientId,
						})
					}
				}
			}
		}
	}

	return strStatus, strSession, localResponse.Result
}
