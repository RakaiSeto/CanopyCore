package modules

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"strings"
	"time"
)

func LoadNotificationInLanguageToRedis(db *sql.DB, redisClient *redis.Client, chOnlineLog *amqp.Channel,
	goContext context.Context, transactionId string) bool {
	isSuccess := false

	query := "select lng.word_id, lng.word_description, coalesce(lng.word_default, '') as word_default, coalesce(wni.word_lang_id, '') as word_lang_id, " +
		"coalesce(wni.lang_id, '') as lang_id, coalesce(lang.lang_description, '') as lang_name, coalesce(wni.word_in_lang, '') as word_in_lang " +
		"from global_wording_notification as lng left join global_wording_notification_in_lang as wni on lng.word_id = wni.word_id " +
		"left join global_language as lang on wni.lang_id = lang.lang_id"

	rows, err := db.Query(query)

	if err != nil {
		// Error in query execution
		DoLogNeo(chOnlineLog, transactionId, time.Now(), "DEBUG", Resource(), "",
			"Failed to read data notification language. Error occur.", true, err, true)
		isSuccess = false
	} else {
		// Success query
		defer func(rows *sql.Rows) {
			errC := rows.Close()
			if errC != nil {
				DoLogNeo(chOnlineLog, "", time.Now(), "DEBUG", Resource(), "",
					"Failed to close rows. Error occur.", true, errC, true)
			}
		}(rows)

		for rows.Next() {
			var dataWordId sql.NullString
			var dataWordDescription sql.NullString
			var dataWordDefault sql.NullString
			var dataWordLangId sql.NullString
			var dataLangId sql.NullString
			var dataLangName sql.NullString
			var dataWordInLang sql.NullString

			err = rows.Scan(&dataWordId, &dataWordDescription, &dataWordDefault, &dataWordLangId, &dataLangId, &dataLangName, &dataWordInLang)

			if err != nil {
				// Error scan table countryCodePrefix
				DoLogNeo(chOnlineLog, "", time.Now(), "DEBUG", Resource(), "",
					"Failed to scan data Notification Languages. Error occur.", true, err, true)

				isSuccess = false
			} else {
				theWordId := ConvertSQLNullStringToString(dataWordId)
				theWordDescription := ConvertSQLNullStringToString(dataWordDescription)
				theWordDefault := ConvertSQLNullStringToString(dataWordDefault)
				theWordLangId := ConvertSQLNullStringToString(dataWordLangId)
				theLangId := ConvertSQLNullStringToString(dataLangId)
				theLangName := ConvertSQLNullStringToString(dataLangName)
				theWordInLang := ConvertSQLNullStringToString(dataWordInLang)

				// Load to redis
				redisKey := "lang_" + theWordId

				mapRedis := make(map[string]string)
				mapRedis["wordId"] = theWordId
				mapRedis["wordDescription"] = theWordDescription
				mapRedis["wordDefault"] = theWordDefault
				mapRedis["wordWordLangId"] = theWordLangId
				mapRedis["wordLangId"] = theLangId
				mapRedis["wordLangName"] = theLangName
				mapRedis["wordInLang"] = theWordInLang

				redisVal := ConvertMapStringToJSON(mapRedis)

				errRedis := RedisSet(redisClient, goContext, redisKey, redisVal, -1)

				if errRedis != nil {
					// Failed to load to redis
					DoLogNeo(chOnlineLog, "", time.Now(), "DEBUG", Resource(),
						"", "Failed to load data language notification to redis. WordLangId: "+theWordLangId, true, errRedis, true)

					isSuccess = false
				} else {
					// Success to load to redis
					DoLogNeo(chOnlineLog, "", time.Now(), "DEBUG", Resource(),
						"", "Success to load data language notification to redis. WordLangId: "+theWordLangId, false, nil, true)

					isSuccess = true
				}
			}
		}
	}

	return isSuccess
}

func GetNotificationInLanguageFromDB(db *sql.DB, chOnlineLog *amqp.Channel, transactionId string, wordingId string, langId string) (bool, string) {
	isSuccess := false
	theNotificationInLang := ""
	theLogResource := make([]string, 2)
	theLogResource[0] = "LANG_TOOL"
	theLogResource[1] = "GET_NOTIF_IN_LANG"

	query := "select lng.word_id, lng.word_description, coalesce(lng.word_default, '') as word_default, coalesce(wni.word_lang_id, '') as word_lang_id, " +
		"coalesce(wni.lang_id, '') as lang_id, coalesce(lang.lang_description, '') as lang_name, coalesce(wni.word_in_lang, '') as word_in_lang " +
		"from global_wording_notification as lng left join global_wording_notification_in_lang as wni on lng.word_id = wni.word_id " +
		"left join global_language as lang on wni.lang_id = lang.lang_id where lng.word_id = $1"

	rows, err := db.Query(query, wordingId)

	if err != nil {
		// Error in query execution
		DoLogNeo(chOnlineLog, transactionId, time.Now(), "DEBUG", theLogResource, "",
			"Failed to read data notification language. Error occur.", true, err, true)
		isSuccess = false
	} else {
		// Success query
		defer func(rows *sql.Rows) {
			errC := rows.Close()
			if errC != nil {
				DoLogNeo(chOnlineLog, "", time.Now(), "DEBUG", theLogResource, "",
					"Failed to close rows. Error occur.", true, errC, true)
			}
		}(rows)

		for rows.Next() {
			var dataWordId sql.NullString
			var dataWordDescription sql.NullString
			var dataWordDefault sql.NullString
			var dataWordLangId sql.NullString
			var dataLangId sql.NullString
			var dataLangName sql.NullString
			var dataWordInLang sql.NullString

			err = rows.Scan(&dataWordId, &dataWordDescription, &dataWordDefault, &dataWordLangId, &dataLangId, &dataLangName, &dataWordInLang)

			if err != nil {
				// Error scan table countryCodePrefix
				DoLogNeo(chOnlineLog, "", time.Now(), "DEBUG", theLogResource, "",
					"Failed to scan data Notification Languages. Error occur.", true, err, true)

				isSuccess = false
			} else {
				theWordId := ConvertSQLNullStringToString(dataWordId)
				strWordId := fmt.Sprintf("%d", theWordId)
				//theWordDescription := ConvertSQLNullStringToString(dataWordDescription)
				theWordDefault := ConvertSQLNullStringToString(dataWordDefault)
				theWordLangId := ConvertSQLNullStringToString(dataWordLangId)
				//theLangId := ConvertSQLNullStringToString(dataLangId)
				//theLangName := ConvertSQLNullStringToString(dataLangName)
				theWordInLang := ConvertSQLNullStringToString(dataWordInLang)

				if len(strings.TrimSpace(theWordLangId)) > 0 {
					if theWordLangId == strWordId+"_"+strings.TrimSpace(strings.ToUpper(langId)) {
						theNotificationInLang = theWordInLang
					}
				} else {
					if theWordId == wordingId {
						theNotificationInLang = theWordDefault
					}
				}
			}
		}
	}

	return isSuccess, theNotificationInLang
}

func GetErrorWordInLangFromDB(db *sql.DB, chOnlineLog *amqp.Channel, transactionId string, errCode string, langId string) (bool, string) {
	isSuccess := false
	theNotifInLang := ""
	theLogResource := make([]string, 2)
	theLogResource[0] = "LANG_TOOL"
	theLogResource[1] = "GET_ERRCODE_IN_LANG"

	query := "select gec.error_code, gec.error_description, gec.default_notif_word_id, gwn.word_description, gwn.word_default, gwl.word_lang_id, " +
		"gwl.lang_id, gwl.word_in_lang from global_error_code as gec left join global_wording_notification as gwn on gec.default_notif_word_id = gwn.word_id " +
		"left join global_wording_notification_in_lang as gwl on gec.default_notif_word_id = gwl.word_id where gec.error_code = $1"

	rows, err := db.Query(query, errCode)

	if err != nil {
		// Error in query execution
		DoLogNeo(chOnlineLog, transactionId, time.Now(), "DEBUG", theLogResource, "",
			"Failed to read data error code notification language. Error occur.", true, err, true)
		isSuccess = false
	} else {
		// Success query
		defer func(rows *sql.Rows) {
			errC := rows.Close()
			if errC != nil {
				DoLogNeo(chOnlineLog, "", time.Now(), "DEBUG", theLogResource, "",
					"Failed to close rows. Error occur.", true, errC, true)
			}
		}(rows)

		counterMatchingData := 0
		defaultWording := ""
		for rows.Next() {
			var dataErrorCode sql.NullString
			var dataErrorDescription sql.NullString
			var dataDefaultNotifWordId sql.NullString
			var dataWordDescription sql.NullString
			var dataWordDefault sql.NullString
			var dataWordLangId sql.NullString
			var dataLangId sql.NullString
			var dataWordInLangId sql.NullString

			err = rows.Scan(&dataErrorCode, &dataErrorDescription, &dataDefaultNotifWordId, &dataWordDescription, &dataWordDefault,
				&dataWordLangId, &dataLangId, &dataWordInLangId)

			if err != nil {
				// Error scan table countryCodePrefix
				DoLogNeo(chOnlineLog, "", time.Now(), "DEBUG", theLogResource, "",
					"Failed to scan data Error Notification Languages. Error occur.", true, err, true)

				isSuccess = false
			} else {
				//strErrorCode := ConvertSQLNullStringToString(dataErrorCode)
				//strErrorDescription := ConvertSQLNullStringToString(dataErrorDescription)
				//strDefaultNotifWordId := ConvertSQLNullStringToString(dataDefaultNotifWordId)
				//strWordDescription := ConvertSQLNullStringToString(dataWordDescription)
				strWordDefault := ConvertSQLNullStringToString(dataWordDescription)
				//strWordLangId := ConvertSQLNullStringToString(dataWordLangId)
				strLangId := ConvertSQLNullStringToString(dataLangId)
				strWordInLangId := ConvertSQLNullStringToString(dataWordInLangId)

				if strLangId == langId {
					// This is it
					isSuccess = true
					theNotifInLang = strWordInLangId
					counterMatchingData = counterMatchingData + 1
				} else {
					// Just in case not found, use default.
					defaultWording = strWordDefault
				}
			}
		}

		if counterMatchingData == 0 {
			// Matching lang and error not found, use default
			theNotifInLang = defaultWording

			if len(theNotifInLang) > 0 {
				isSuccess = true
			} else {
				isSuccess = false
			}
		} else {
			isSuccess = true
		}
	}

	return isSuccess, theNotifInLang
}
