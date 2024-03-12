package RPCCall

import (
	error_code "CanopyCore/errors"
	globaldata "CanopyCore/grpc/globaldata"
	"CanopyCore/modules"
	"context"
	"database/sql"
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

	localResponse.Code = responseStatus

	return localResponse, nil
}
