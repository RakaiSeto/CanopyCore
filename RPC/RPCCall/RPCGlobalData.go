package RPCCall

import (
	error_code "CanopyCore/errors"
	globaldata "CanopyCore/grpc/globaldata"
	"CanopyCore/modules"
	"context"
	"database/sql"
	"time"
)

func CallFunctionGetAllCountry(ctx context.Context, db *sql.DB) (*globaldata.AllCountryResponse, error) {
	localResponse := new(globaldata.AllCountryResponse)

	incTraceCode := modules.GenerateUUID()

	//var mapRequest = make(map[string]interface{})

	responseStatus := error_code.ErrorServer
	responseDescription := ""

	strRemoteIP := modules.GetGRPCRemoteIP(ctx)

	//log marking the start of process
	modules.Logging(modules.Resource(), incTraceCode, "SERVER", strRemoteIP, "Start process Get All Country", nil)

	query := `SELECT iso3, nicename FROM global_country ORDER BY nicename ASC`
	rows, err := db.Query(query)
	if err != nil {
		modules.Logging(modules.Resource(), incTraceCode, "SERVER", strRemoteIP, "error when query", err)
		responseStatus = error_code.ErrorServer
	} else {
		defer rows.Close()
		for rows.Next() {
			var rawIso3 sql.NullString
			var rawNicename sql.NullString
			err := rows.Scan(&rawIso3, &rawNicename)
			if err != nil {
				modules.Logging(modules.Resource(), incTraceCode, "SERVER", strRemoteIP, "error when scan", err)
				responseStatus = error_code.ErrorServer
				break
			} else {
				localResponse.Result = append(localResponse.Result, &globaldata.AllCountryList{
					Code: modules.ConvertSQLNullStringToString(rawIso3),
					Name: modules.ConvertSQLNullStringToString(rawNicename),
				})
				responseStatus = error_code.ErrorSuccess
			}
		}

	}

	//if response status is success, then log
	if responseStatus == error_code.ErrorSuccess {
		modules.Logging(modules.Resource(), incTraceCode, "SERVER", strRemoteIP, "Get All Country success", nil)
	} else {
		modules.Logging(modules.Resource(), incTraceCode, "SERVER", strRemoteIP, "Get All Country failed", nil)
	}

	var getError error
	responseDescription, getError = error_code.GetErrorStatus(db, responseStatus)
	if getError != nil {
		modules.Logging(modules.Resource(), incTraceCode, "SERVER", strRemoteIP, "error when get error status", getError)
		responseStatus = error_code.ErrorServer
	} else {
		localResponse.Description = responseDescription
	}

	localResponse.Statuscode = responseStatus

	status := ""
	if responseStatus == error_code.ErrorSuccess {
		status = "SUCCESS"
	} else {
		status = "FAILED"
	}
	modules.SaveWebActivity(db, incTraceCode, modules.Resource(), "SERVER", strRemoteIP, modules.DoFormatDateTime("YYYY-0M-0D HH:mm:ss.S", time.Now()), "DO GET ALL COUNTRY", status)

	return localResponse, nil
}

func CallFunctionGetAllRoles(ctx context.Context, db *sql.DB) (*globaldata.AllRoleResponse, error) {
	localResponse := new(globaldata.AllRoleResponse)

	incTraceCode := modules.GenerateUUID()

	//var mapRequest = make(map[string]interface{})

	responseStatus := error_code.ErrorServer
	responseDescription := ""

	strRemoteIP := modules.GetGRPCRemoteIP(ctx)

	//log marking the start of process
	modules.Logging(modules.Resource(), incTraceCode, "SERVER", strRemoteIP, "Start process Get All Roles", nil)

	query := `SELECT role, description FROM global_role ORDER BY role ASC`
	rows, err := db.Query(query)
	if err != nil {
		modules.Logging(modules.Resource(), incTraceCode, "SERVER", strRemoteIP, "error when query", err)
		responseStatus = error_code.ErrorServer
	} else {
		defer rows.Close()
		for rows.Next() {
			var rawRole sql.NullString
			var rawDescription sql.NullString
			err := rows.Scan(&rawRole, &rawDescription)
			if err != nil {
				modules.Logging(modules.Resource(), incTraceCode, "SERVER", strRemoteIP, "error when scan", err)
				responseStatus = error_code.ErrorServer
				break
			} else {
				localResponse.Result = append(localResponse.Result, &globaldata.AllRoleList{
					Code: modules.ConvertSQLNullStringToString(rawRole),
					Name: modules.ConvertSQLNullStringToString(rawDescription),
				})
				responseStatus = error_code.ErrorSuccess
			}
		}

	}

	//if response status is success, then log
	if responseStatus == error_code.ErrorSuccess {
		modules.Logging(modules.Resource(), incTraceCode, "SERVER", strRemoteIP, "Get All Roles success", nil)
	} else {
		modules.Logging(modules.Resource(), incTraceCode, "SERVER", strRemoteIP, "Get All Roles failed", nil)
	}

	var getError error
	responseDescription, getError = error_code.GetErrorStatus(db, responseStatus)
	if getError != nil {
		modules.Logging(modules.Resource(), incTraceCode, "SERVER", strRemoteIP, "error when get error status", getError)
		responseStatus = error_code.ErrorServer
	} else {
		localResponse.Description = responseDescription
	}

	localResponse.Statuscode = responseStatus

	status := ""
	if responseStatus == error_code.ErrorSuccess {
		status = "SUCCESS"
	} else {
		status = "FAILED"
	}
	modules.SaveWebActivity(db, incTraceCode, modules.Resource(), "SERVER", strRemoteIP, modules.DoFormatDateTime("YYYY-0M-0D HH:mm:ss.S", time.Now()), "DO GET ALL ROLES", status)

	return localResponse, nil
}
