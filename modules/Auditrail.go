package modules

import (
	"database/sql"
)

func SaveWebActivity(db *sql.DB, inctraceCode string, resource []string, identity string, remoteip string, datetime string, activity string, status string) bool {
	isOK := false

	query := `INSERT INTO auditrail(tracecode, module, identity, remoteip, datetime, activity, status) VALUES($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.Exec(query, inctraceCode, resource[0], identity, remoteip, datetime, activity, status)

	if err != nil {
		isOK = false
		Logging(resource, inctraceCode, identity, remoteip, "error when insert auditrail", err)
	} else {
		isOK = true
		Logging(resource, inctraceCode, identity, remoteip, "success when insert auditrail", err)
	}

	return isOK
}
