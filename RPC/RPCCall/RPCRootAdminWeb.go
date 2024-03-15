package RPCCall

import (
	"CanopyCore/RPC/RPCHelper"
	error_code "CanopyCore/errors"
	rootadminweb "CanopyCore/grpc/rootadminweb"
	"CanopyCore/modules"
	"context"
	"database/sql"
	"github.com/go-redis/redis/v8"
	"time"
)

func CallFunctionDoLoginRootAdminWeb(ctx context.Context, db *sql.DB, rc *redis.Client, request *rootadminweb.DoLoginRequest) (*rootadminweb.DoLoginResponse, error) {
	localResponse := new(rootadminweb.DoLoginResponse)
	var responseResult []*rootadminweb.DoLoginList

	incTraceCode := modules.GenerateUUID()

	var mapRequest = make(map[string]interface{})

	responseStatus := error_code.ErrorServer
	responseDescription := ""
	responseSession := ""

	strRemoteIP := modules.GetGRPCRemoteIP(ctx)

	incRequestEmail := request.Email
	incRequestPassword := request.Password

	// log marking the start of process
	modules.Logging(modules.Resource(), incTraceCode, incRequestEmail, strRemoteIP, "Start process Do Login", nil)

	if len(incRequestEmail) > 0 && len(incRequestPassword) > 0 {
		modules.Logging(modules.Resource(), incTraceCode, incRequestEmail, strRemoteIP, "Parameter is valid", nil)

		mapRequest["tracecode"] = incTraceCode
		mapRequest["remoteip"] = strRemoteIP
		mapRequest["email"] = incRequestEmail
		mapRequest["password"] = incRequestPassword

		responseStatus, responseSession, responseResult = RPCHelper.HelperDoLoginRootAdminWeb(ctx, db, rc, mapRequest)
	} else {
		responseStatus = error_code.ErrorInvalidParameter
		modules.Logging(modules.Resource(), incTraceCode, incRequestEmail, strRemoteIP, "Invalid Parameter", nil)
	}

	// if responseStatus == error_code.ErrorSuccess then log success, else log failed
	if responseStatus == error_code.ErrorSuccess {
		modules.Logging(modules.Resource(), incTraceCode, incRequestEmail, strRemoteIP, "Do Login success", nil)
	} else {
		modules.Logging(modules.Resource(), incTraceCode, incRequestEmail, strRemoteIP, "Do Login failed", nil)
	}

	var getError error
	responseDescription, getError = error_code.GetErrorStatus(db, responseStatus)
	if getError != nil {
		modules.Logging(modules.Resource(), incTraceCode, incRequestEmail, strRemoteIP, "error when get error status", getError)
		responseStatus = error_code.ErrorServer
	} else {
		localResponse.Description = responseDescription
	}

	localResponse.Result = responseResult
	localResponse.Session = responseSession
	localResponse.Statuscode = responseStatus

	status := ""
	if responseStatus == error_code.ErrorSuccess {
		status = "SUCCESS"
	} else {
		status = "FAILED"
	}
	modules.SaveWebActivity(db, incTraceCode, modules.Resource(), incRequestEmail, strRemoteIP, modules.DoFormatDateTime("YYYY-0M-0D HH:mm:ss.S", time.Now()), "DO LOGIN ROOT ADMIN WEB", status)

	return localResponse, nil
}

func CallFunctionDoLogoutRootAdminWeb(ctx context.Context, db *sql.DB, rc *redis.Client, request *rootadminweb.EmptyRequest) (*rootadminweb.DoLogoutResponse, error) {
	localResponse := new(rootadminweb.DoLogoutResponse)

	incTraceCode := modules.GenerateUUID()

	//var mapRequest = make(map[string]interface{})

	responseStatus := error_code.ErrorServer
	responseDescription := ""

	strRemoteIP := modules.GetGRPCRemoteIP(ctx)

	var getError error
	responseDescription, getError = error_code.GetErrorStatus(db, responseStatus)
	if getError != nil {
		modules.Logging(modules.Resource(), incTraceCode, "SERVER", strRemoteIP, "error when get error status", getError)
		responseStatus = error_code.ErrorServer
	} else {
		localResponse.Description = responseDescription
	}

	localResponse.Statuscode = responseStatus

	//status := ""
	//if responseStatus == error_code.ErrorSuccess {
	//	status = "SUCCESS"
	//} else {
	//	status = "FAILED"
	//}
	//modules.SaveWebActivity(db, incTraceCode, modules.Resource(), incRequestEmail, strRemoteIP, modules.DoFormatDateTime("YYYY-0M-0D HH:mm:ss.S", time.Now()), "DO LOGOUT ROOT ADMIN WEB", status)

	return localResponse, nil
}
