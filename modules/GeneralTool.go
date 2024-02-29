package modules

import (
	helper "CanopyCore/APP/Helper"
	"bytes"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"database/sql"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc/peer"

	"github.com/gomodule/redigo/redis"
	guuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	InfoColor  = Teal
	WarnColor  = Yellow
	ErrorColor = Red
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

type logData struct {
	Datetime     string `bson:"datetime"`
	LogLevel     string `bson:"loglevel"`
	TraceID      string `bson:"traceid"`
	Application  string `bson:"application"`
	Module       string `bson:"module"`
	Function     string `bson:"function"`
	Identity     string `bson:"identity"`
	RemoteIP     string `bson:"remoteip"`
	Message      string `bson:"message"`
	ErrorMessage string `bson:"errormessage"`
	Status       string `bson:"status"`
}

func DoLogDB(db *mongo.Client, ctx context.Context, dbLog string, incLogLevel string, incTraceID string, incApplication string,
	incModule string, incFunction string, incIdentity string, incRemoteIP string, incMessage string, isError bool, incErrorMessage error) {

	//constCollection := ""
	strDatetime := DoFormatDateTime("YYYY-0M-0D HH:mm:ss.S", time.Now())
	colDatetime := "log_" + DoFormatDateTime("YYYY0M0D", time.Now())
	strStatus := "SUCCESS"
	strErrorMessage := ""

	if isError {
		strStatus = "FAILED"
	}

	//if strings.ToUpper(incLogLevel) == "INFO" {
	//	constCollection = "info"
	//} else if strings.ToUpper(incLogLevel) == "ERROR" {
	//	constCollection = "error"
	//} else if strings.ToUpper(incLogLevel) == "DEBUG" {
	//	constCollection = "debug"
	//} else if strings.ToUpper(incLogLevel) == "CRITICAL" {
	//	constCollection = "critical"
	//} else {
	//	constCollection = "info"
	//}

	if incErrorMessage == nil {
		strErrorMessage = ""
	} else {
		strErrorMessage = fmt.Sprintf("%s", incErrorMessage)
	}

	//_, err1 := db.Database(dbLog).Collection(constCollection).InsertOne(ctx, logData{strDatetime, incLogLevel, incTraceID, incApplication,
	//	incModule, incFunction, incMessage, strErrorMessage, strStatus})
	//if err1 != nil {
	//	DoLog("ERROR", "LOG SERVER", "DOLOG", "Save",
	//		"Failed to save Logging. Error occurred.", true, err1)
	//}

	_, err2 := db.Database(dbLog).Collection(colDatetime).InsertOne(ctx, logData{strDatetime, incLogLevel, incTraceID, incApplication,
		incModule, incFunction, incIdentity, incRemoteIP, incMessage, strErrorMessage, strStatus})
	if err2 != nil {
		DoLog("ERROR", "LOG SERVER", "DOLOG", "Save",
			"Failed to save Logging. Error occurred.", true, err2)
	}

	DoLog(incLogLevel, incTraceID, incModule, incFunction, incMessage, isError, incErrorMessage)
}

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

// Logging formatter
type Formatter struct {
	// FieldsOrder - default: fields sorted alphabetically
	FieldsOrder []string

	// TimestampFormat - default: time.StampMilli = "Jan _2 15:04:05.000"
	TimestampFormat string

	// HideKeys - show [fieldValue] instead of [fieldKey:fieldValue]
	HideKeys bool

	// NoColors - disable colors
	NoColors bool

	// NoFieldsColors - apply colors only to the level, default is level + fields
	NoFieldsColors bool

	// ShowFullLevel - show a full level [WARNING] instead of [WARN]
	ShowFullLevel bool

	// TrimMessages - trim whitespaces on messages
	TrimMessages bool

	// CallerFirst - print caller info first
	CallerFirst bool

	// CustomCallerFormatter - set custom formatter for caller info
	CustomCallerFormatter func(*runtime.Frame) string
}

// Global Variable for MapConfig
var MapConfig = make(map[string]string)
var MapConfigScapper = make(map[string]string)

// Global Logger for all
var zapLogger *zap.Logger

// Global RedisPooler for All
var RedisPooler *redis.Pool

func InitiateGlobalVariables() {
	MapConfig = LoadConfig()
	MapConfigScapper = LoadConfigScrapper()
	zapLogger = initiateZapLogger()
	RedisPooler = RedisInitiateRedisPool()
	initiateOauthHandler()
}

func initiateOauthHandler() {
	helper.GoogleOauthLogin = &oauth2.Config{
		RedirectURL:  "https://apicanopy.rakaiseto.com/login/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	fmt.Println(os.Getenv("GOOGLE_CLIENT_ID"))
	fmt.Println(os.Getenv("GOOGLE_CLIENT_SECRET"))
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan 2 15:04:05.000"))
}

func initiateZapLogger() *zap.Logger {
	// THE APPLICATION NEED TO RUN USING SUPERVISOR BECAUSE IT JUST LOGGING TO STDOUT ONLY.
	// SUPERVISOR WILL HANDLE THE LOGGING ---

	//w := zapcore.AddSync(&lumberjack.Logger{
	//	Filename:   MapConfig["loggingPathInfo"],
	//	MaxSize:    100, // megabytes
	//	MaxBackups: 3,
	//	MaxAge:     30, // days
	//})
	//
	//core := zapcore.NewCore(
	//	zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
	//		MessageKey: "message",
	//
	//		LevelKey:    "level",
	//		EncodeLevel: zapcore.CapitalLevelEncoder,
	//
	//		TimeKey:    "time",
	//		EncodeTime: SyslogTimeEncoder,
	//	}),
	//	zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
	//	zap.DebugLevel,
	//)
	//
	//logger := zap.New(core, zap.Development())

	cfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: SyslogTimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, _ := cfg.Build()

	return logger
}

func GenerateUUID() string {
	var theUuid string

	theUuid = guuid.New().String()
	theUuid = strings.Replace(theUuid, "-", "", -1)

	return theUuid
}

func DoLog(logLevel string, messageId string, moduleName string, functionName string, message string, isError bool, theError error) {
	if logLevel == "DEBUG" {
		if isError == true {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(moduleName), 10) + "   " + TrimStringLength(strings.ToUpper(functionName), 15) + "   " + message + "." + fmt.Sprintf(" %v", theError)
		} else {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(moduleName), 10) + "   " + TrimStringLength(strings.ToUpper(functionName), 15) + "   " + message
		}

		zapLogger.Debug(message)
	} else if logLevel == "ERROR" {
		if isError == true {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(moduleName), 10) + "   " + TrimStringLength(strings.ToUpper(functionName), 15) + "   " + message + "." + fmt.Sprintf(" %v", theError)
		} else {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(moduleName), 10) + "   " + TrimStringLength(strings.ToUpper(functionName), 15) + "   " + message
		}

		zapLogger.Error(ErrorColor(message))
		//zapLogger.Debug(message)
	} else if logLevel == "WARNING" {
		if isError == true {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(moduleName), 10) + "   " + TrimStringLength(strings.ToUpper(functionName), 15) + "   " + message + "." + fmt.Sprintf(" %v", theError)
		} else {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(moduleName), 10) + "   " + TrimStringLength(strings.ToUpper(functionName), 15) + "   " + message
		}

		zapLogger.Warn(message)
		//zapLogger.Debug(message)
	} else if logLevel == "INFO" {
		if isError == true {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(moduleName), 10) + "   " + TrimStringLength(strings.ToUpper(functionName), 15) + "   " + message + "." + fmt.Sprintf(" %v", theError)
		} else {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(moduleName), 10) + "   " + TrimStringLength(strings.ToUpper(functionName), 15) + "   " + message
		}

		zapLogger.Info(message)
		//zapLogger.Debug(message)
	} else {
		if isError == true {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(moduleName), 10) + "   " + TrimStringLength(strings.ToUpper(functionName), 15) + "   " + message + "." + fmt.Sprintf(" %v", theError)
		} else {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(moduleName), 10) + "   " + TrimStringLength(strings.ToUpper(functionName), 15) + "   " + message
		}

		//zapLogger.Info(message)
		zapLogger.Debug(message)
	}
}

func MapInterfaceHasKey(theMap map[string]interface{}, theKey string) bool {
	_, exist := theMap[theKey]

	return exist
}

// noinspection GoUnusedExportedFunction
func MapStringHasKey(theMap map[string]string, theKey string) bool {
	_, exist := theMap[theKey]

	return exist
}

// noinspection GoUnusedExportedFunction
func DoHashMD5(theStr string) string {
	hasher := md5.New()

	hasher.Write([]byte(theStr))

	return hex.EncodeToString(hasher.Sum(nil))
}

func GetStringFromMapInterface(theMap map[string]interface{}, theKey string) string {
	theValue := ""

	_, exist := theMap[theKey]

	if exist == true && theMap[theKey] != nil {
		theValue = strings.TrimSpace(theMap[theKey].(string))
	}

	return theValue
}

func GetErrorFromMapInterface(theMap map[string]interface{}, theKey string) error {
	theValue := errors.New("nil")

	_, exist := theMap[theKey]

	if exist == true && theMap[theKey] != nil {
		theValue = errors.New(strings.TrimSpace(theMap[theKey].(string)))
	}

	return theValue
}

func GetBoolFromMapInterface(theMap map[string]interface{}, theKey string) bool {
	theValue := false

	theValueString, exist := theMap[theKey]

	if exist == true && theMap[theKey] != nil {
		theValueString = fmt.Sprint(theValueString)
		theValue, _ = strconv.ParseBool(theValueString.(string))
	}
	return theValue
}

func GetInt64FromMapInterface(theMap map[string]interface{}, theKey string) int64 {
	theValue := int64(0)

	i64Val, ok := theMap[theKey].(int64)

	if !ok {
		// Bukan int64 Check apakah string
		stringVal, strOk := theMap[theKey].(string)
		if !strOk {
			// Bukan string, apakah int64
			floatVal, floatOk := theMap[theKey].(float64)

			if !floatOk {
				// Bukan float, apa dong ya?
				theValue = 0
			} else {
				// float64, convert it to int64
				theValue = int64(floatVal)
			}
		} else {
			// String, convert it to float64
			theValuei64, errStr := strconv.ParseInt(stringVal, 10, 64)

			if errStr == nil {
				theValue = theValuei64
			}
		}
	} else {
		theValue = i64Val
	}

	return theValue
}

func GetFloatFromMapInterface(theMap map[string]interface{}, theKey string) float64 {
	theValue := 0.00000000

	floatVal, ok := theMap[theKey].(float64)

	if !ok {
		// Bukan float64 Check apakah string
		stringVal, strOk := theMap[theKey].(string)
		if !strOk {
			// Bukan string, apakah int64
			intVal, intOk := theMap[theKey].(int64)

			if !intOk {
				// Bukan integer, apa dong ya?
				theValue = 0.00000000
			} else {
				// Integer64, convert it to float64
				theValue = float64(intVal)
			}
		} else {
			// String, convert it to float64
			theValueF, errStr := strconv.ParseFloat(stringVal, 64)

			if errStr == nil {
				theValue = theValueF
			}
		}
	} else {
		theValue = floatVal
	}

	return theValue
}

// noinspection GoUnusedExportedFunction
func GetStringFromMapString(theMap map[string]string, theKey string) string {
	theValue := ""

	_, exist := theMap[theKey]

	if exist == true {
		theValue = strings.TrimSpace(theMap[theKey])
	}

	return theValue
}

func ConvertMapStringToJSON(theMap map[string]string) string {
	jsonString, err := json.Marshal(theMap)

	if err != nil {
		return "{}"
	}

	return string(jsonString)
}

func ConvertMapInterfaceToJSON(theMap map[string]interface{}) string {
	jsonString, err := json.Marshal(theMap)

	if err != nil {
		return "{}"
	}

	return string(jsonString)
}

func ConvertJSONStringToMap(messageId string, theJSON string) map[string]interface{} {
	resultMap := make(map[string]interface{})

	err := json.Unmarshal([]byte(theJSON), &resultMap)

	if err != nil {
		log.Debugln(messageId + " ." + fmt.Sprintf("Failed to convert json to map for json content: %s", theJSON))
		resultMap = nil
	} else {
		log.Debugln(fmt.Sprintf("Success converting json %s to hashmap: %+v", theJSON, resultMap))
	}

	return resultMap
}

func ConvertMapStringToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

// noinspection GoUnusedExportedFunction
func GenerateRandomNumericString(strLength int) string {
	var letters = []rune("1234567890")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, strLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func GenerateRandomLongNumericString(strLength int) string {
	var letters = []rune("456781451234567890987654321")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, strLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// noinspection GoUnusedExportedFunction
func GenerateRandomAlphaNumericString(strLength int) string {
	var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, strLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// noinspection GoUnusedExportedFunction
func GenerateRandomAlphabeticalString(strLength int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, strLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// noinspection GoUnusedExportedFunction
func GenerateRandomCapitalAlphabeticalString(strLength int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXZABCDEFGHIJKLMNOPQRSTUVWXZABCDEFGHIJKLMNOPQRSTUVWXZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, strLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

//func GenerateTrxId(trxIDType string, length int) string {
//	trxID		:= ""
//
//	if trxIDType == "ALPHANUMERIC" {
//		trxID 	= GenerateRandomAlphaNumericString(length)
//	} else if trxIDType == "ALPHABETICAL" {
//		trxID 	= GenerateRandomAlphabeticalString(length)
//	} else {
//		trxID	= GenerateRandomNumericString(length)
//	}
//
//	return trxID
//}

func DoFormatDateTime(dateTimeFormat string, theTime time.Time) string {
	const (
		stdLongMonth   = "January"
		stdMonth       = "Jan"
		stdNumMonth    = "1"
		stdZeroMonth   = "01"
		stdLongWeekDay = "Monday"
		stdWeekDay     = "Mon"
		stdDay         = "2"
		//stdUnderDay       = "_2"
		stdZeroDay    = "02"
		stdHour       = "15"
		stdHour12     = "3"
		stdZeroHour12 = "03"
		stdMinute     = "4"
		stdZeroMinute = "04"
		stdSecond     = "5"
		stdZeroSecond = "05"
		stdLongYear   = "2006"
		stdYear       = "06"
		//stdPM             = "PM"
		//stdpm             = "pm"
		//stdTZ             = "MST"
		//stdISO8601TZ      = "Z0700"  // prints Z for UTC
		//stdISO8601ColonTZ = "Z07:00" // prints Z for UTC
		//stdNumTZ          = "-0700"  // always numeric
		//stdNumShortTZ     = "-07"    // always numeric
		//stdNumColonTZ     = "-07:00" // always numeric
		//stdMilliSecond	  = "000"
	)

	theFormat := dateTimeFormat

	// replace YYYY with stdLongYear
	theFormat = strings.Replace(theFormat, "YYYY", stdLongYear, -1)

	// replace YY with stdYear
	theFormat = strings.Replace(theFormat, "YY", stdYear, -1)

	// replace MMMM with stdLongMonth - January
	theFormat = strings.Replace(theFormat, "MMMM", stdLongMonth, -1)

	// replace MM with stdMonth - Jan
	theFormat = strings.Replace(theFormat, "MM", stdMonth, -1)

	// replace 0M with zeroMonth - 01
	theFormat = strings.Replace(theFormat, "0M", stdZeroMonth, -1)

	// replace M with oneMonth - 1
	theFormat = strings.Replace(theFormat, "M", stdNumMonth, -1)

	// replace DDDD with stdLongWeekDay - Monday
	theFormat = strings.Replace(theFormat, "DDDD", stdLongWeekDay, -1)

	// replace DD with stdWeekDay - Mon
	theFormat = strings.Replace(theFormat, "DD", stdWeekDay, -1)

	// replace 0D with stdZeroDay - 01
	theFormat = strings.Replace(theFormat, "0D", stdZeroDay, -1)

	// replace D with stdNumDate - 1
	theFormat = strings.Replace(theFormat, "D", stdDay, -1)

	// replace HH with 24 hours hour
	theFormat = strings.Replace(theFormat, "HH", stdHour, -1)

	// replace 0H with 2 digits 12 hour
	theFormat = strings.Replace(theFormat, "0H", stdZeroHour12, -1)

	// replace H with num hour 12
	theFormat = strings.Replace(theFormat, "H", stdHour12, -1)

	// repalce mm with 2 digits minute start with 0
	theFormat = strings.Replace(theFormat, "mm", stdZeroMinute, -1)

	// replace m with number minute
	theFormat = strings.Replace(theFormat, "m", stdMinute, -1)

	// replace ss with 2 digits seconds
	theFormat = strings.Replace(theFormat, "ss", stdZeroSecond, -1)

	// replace s with number second
	theFormat = strings.Replace(theFormat, "s", stdSecond, -1)

	// replace S with millisecond after theFormat is implemented
	theReturn := theTime.Format(theFormat)
	//fmt.Println("theReturn before milisecond: " + theReturn)

	// Create the milist of current theTime
	milis := (theTime.UnixNano()) / 1000000

	nowMilis := milis - (theTime.Unix() * 1000)

	strMilis := ConvertInt64ToStringFixLength("0", "LEFT", 3, nowMilis)

	// Kalo masih 4 digit, cut jadi 3 digit
	if len(strMilis) > 3 {
		theRune := []rune(strMilis)
		strMilis = string(theRune[0:2])
	}

	// Replace S with millisecond
	theReturn = strings.Replace(theReturn, "S", strMilis, -1)

	// Replce x with nano second
	theFormat = strings.Replace(theFormat, "x", strconv.FormatInt(theTime.UnixNano(), 10), -1)

	return theReturn
}

// noinspection GoUnusedExportedFunction
func GetStringInBetweenInsideBoundary(str string, start string, end string) (result string) {
	fmt.Println("GetStringInBetweenInsideBoundary - start: " + start + ", end: " + end + ", from string: " + str)
	s := strings.Index(str, start)
	if s == -1 {
		return ""
	}
	s += len(start)

	withoutStartStr := str[s:]
	fmt.Println("GetStringInBetweenInsideBoundary - withoutStartStr: " + withoutStartStr)
	e := strings.Index(withoutStartStr, end) // ambil batas akhir paling dalam
	fmt.Printf("GetStringInBetweenInsideBoundary = e: " + fmt.Sprintf("%d", e))
	if e == -1 {
		return ""
	}
	return withoutStartStr[0:e]
}

// noinspection GoUnusedExportedFunction
func GetStringInBetweenOutsideBoundary(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return ""
	}
	s += len(start)
	e := strings.LastIndex(str, end) // ambil batas akhir paling luar
	return str[s:e]
}

// noinspection GoUnusedExportedFunction
func DoReverseString(theString string) string {
	runes := []rune(theString)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// noinspection GoUnusedExportedFunction
func DoMD5(theString string) string {
	hash := md5.Sum([]byte(theString))
	return hex.EncodeToString(hash[:])
}

// noinspection GoUnusedExportedFunction
func DoSHA1(theString string) string {
	h := sha1.New()
	h.Write([]byte(theString))
	sha1Hash := hex.EncodeToString(h.Sum(nil))

	return sha1Hash
}

// noinspection GoUnusedExportedFunction
func ConvertMapInterfaceToMapString(mapInterface map[string]interface{}) map[string]string {
	mapHasil := make(map[string]string)

	//for k := range mapInterface {
	//	if fmt.Sprintf("%T", reflect.TypeOf(mapInterface[k])) == "bool" {
	//		// variable is bool
	//		mapHasil[k] = strconv.FormatBool(mapInterface[k].(bool))
	//	} else if fmt.Sprintf("%T", reflect.TypeOf(mapInterface[k])) == "string" {
	//		// variable is string
	//		mapHasil[k] = mapInterface[k].(string)
	//	} else if fmt.Sprintf("%T", reflect.TypeOf(mapInterface[k])) == "int" {
	//		// variable is int
	//		mapHasil[k] = strconv.Itoa(mapInterface[k].(int))
	//	} else if fmt.Sprintf("%T", reflect.TypeOf(mapInterface[k])) == "float64" {
	//		// variable is float64
	//		mapHasil[k] = fmt.Sprintf("%.2f", mapInterface[k].(float64))
	//	} else if fmt.Sprintf("%T", reflect.TypeOf(mapInterface[k])) == "float32" {
	//		// variable is float64
	//		mapHasil[k] = fmt.Sprintf("%.2f", mapInterface[k].(float32))
	//	}
	//}

	for key, value := range mapInterface {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		mapHasil[strKey] = strValue
	}

	return mapHasil
}

// noinspection GoUnusedExportedFunction
func ConvertIntToStringFixLength(emptySpaceFiller string, fillerPosition string, theExpectedLength int, theInteger int) string {
	// fillerPosition = "RIGHT" or "LEFT"
	response := ""

	// Convert theInteger to str
	theStrInteger := strconv.Itoa(theInteger)

	// Check if original length of theInteger > expectedLength
	if len(theStrInteger) >= theExpectedLength {
		response = theStrInteger
	} else {
		theFiller := emptySpaceFiller
		for i := 0; i < theExpectedLength-len(theStrInteger); i++ {
			theFiller = theFiller + emptySpaceFiller
		}

		if fillerPosition == "RIGHT" {
			response = theStrInteger + theFiller
		} else {
			response = theFiller + theStrInteger
		}
	}

	return response
}

func ConvertInt64ToStringFixLength(emptySpaceFiller string, fillerPosition string, theExpectedLength int, theInteger int64) string {
	// fillerPosition = "RIGHT" or "LEFT"
	response := ""

	// Convert theInteger to str
	theStrInteger := strconv.FormatInt(theInteger, 10)

	// Check if original length of theInteger > expectedLength
	if len(theStrInteger) >= theExpectedLength {
		response = theStrInteger
	} else {
		theFiller := emptySpaceFiller
		for i := 0; i < theExpectedLength-len(theStrInteger); i++ {
			theFiller = theFiller + emptySpaceFiller
		}

		if fillerPosition == "RIGHT" {
			response = theStrInteger + theFiller
		} else {
			response = theFiller + theStrInteger
		}
	}

	return response
}

// noinspection GoUnusedExportedFunction
func IsStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// noinspection GoUnusedExportedFunction
func DoAppendStringToSlice(slice []string, data ...string) []string {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]string, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

//func SubstrUTF8(input string, start int, length int) string {
//	asRunes := []rune(input)
//
//	if start >= len(asRunes) {
//		return ""
//	}
//
//	if start+length > len(asRunes) {
//		length = len(asRunes) - start
//	}
//
//	return string(asRunes[start : start+length])
//}
//
//func RemoveStringStartToIndex(theStr string, startIndex int, endIndex int) string {
//	theResult := theStr
//
//	// temp char to replace then to be deleted is ;
//	theResult = theResult[:startIndex] + "" + theResult[endIndex:]
//	//for i:=startIndex; i<endIndex; i++ {
//	//}
//	fmt.Println("temp theResult: " + theResult)
//
//	// Remove all ;
//	theResult = strings.Replace(theResult, ";", "", -1)
//	fmt.Println("final theResult: " + theResult)
//
//	return theResult
//}

//func GetStringInBetween(str string, start string, end string) string {
//	result := ""
//	theString := str
//
//	// Get pos start
//	posStart := strings.Index(theString, start)
//
//	// Remove theString from index 0 to posStart + len(start)
//	result = RemoveStringStartToIndex(theString, 0, posStart + len(start))
//	fmt.Println("result - clean depan: " + result)
//
//	// Remove dari posEnd sampe end of string
//	// Get pos end
//	posEndDrResult := strings.Index(result, end)
//
//	result = RemoveStringStartToIndex(result, posEndDrResult, len(result))
//	fmt.Println("result - clean all: " + result)
//
//	return result
//}

func TrimStringLength(theString string, length int) string {
	hasil := theString

	if len(theString) > length {
		hasil = fmt.Sprintf("%."+strconv.Itoa(length)+"s", theString)
	} else {
		hasil = theString

		for x := 0; x < length-len(theString); x++ {
			hasil = hasil + " "
		}
	}

	return hasil
}

func ConvertSQLNullStringToString(sqlNullString sql.NullString) string {
	if sqlNullString.Valid == true {
		return sqlNullString.String
	} else {
		return ""
	}
}

func ConvertSQLNullFloat64ToFloat64(sqlNullFloat64 sql.NullFloat64) float64 {
	if sqlNullFloat64.Valid == true {
		return sqlNullFloat64.Float64
	} else {
		return 0.00
	}
}

func ConvertSQLNullBoolToBool(sqlNullBool sql.NullBool) bool {
	if sqlNullBool.Valid == true {
		return sqlNullBool.Bool
	} else {
		return false
	}
}

func ConvertSQLNullInt64ToInt64(sqlNullInt64 sql.NullInt64) int64 {
	if sqlNullInt64.Valid == true {
		return sqlNullInt64.Int64
	} else {
		return 0
	}
}

func EncodeBase64(toEncode string) string {
	hash64 := b64.URLEncoding.EncodeToString([]byte(toEncode))

	return hash64
}

func SHA256(toHash string) string {
	hasher := sha256.New()
	hasher.Write([]byte(toHash))

	result := hex.EncodeToString(hasher.Sum(nil))

	return result
}

func Base64(toHash string) string {
	return b64.StdEncoding.EncodeToString([]byte(toHash))
}

func IsEmailValid(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func Resource() []string {

	arrResource := make([]string, 2)

	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, _ := f.FileLine(pc[0])

	// Get Module Name
	arrModule := strings.Split(file, "/")
	moduleName := arrModule[len(arrModule)-1]
	//strTrace := moduleName[:len(moduleName)-3] + ":" + fmt.Sprintf("%v", line)
	strModuleName := moduleName[:len(moduleName)-3]

	// Get Function Name
	arrFunc := strings.Split(f.Name(), ".")
	strFunctionName := arrFunc[len(arrFunc)-1]

	arrResource[0] = strModuleName
	arrResource[1] = strFunctionName

	return arrResource
}

func GetGRPCRemoteIP(cx context.Context) string {

	p, _ := peer.FromContext(cx)
	hostStr := p.Addr.String()
	arrIP := strings.Split(hostStr, ":")
	incRemoteIP := arrIP[0]

	return incRemoteIP
}

func ResourceX() []string {

	arrResource := make([]string, 2)

	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, _ := f.FileLine(pc[0])

	// Get Module Name
	arrModule := strings.Split(file, "/")
	moduleName := arrModule[len(arrModule)-1]
	//strTrace := moduleName[:len(moduleName)-3] + ":" + fmt.Sprintf("%v", line)
	strModuleName := moduleName[:len(moduleName)-3]

	// Get Function Name
	arrFunc := strings.Split(f.Name(), ".")
	strFunctionName := arrFunc[len(arrFunc)-1]

	arrResource[0] = strModuleName
	arrResource[1] = strFunctionName

	return arrResource
}

func Logging(incResource []string, incTraceCode string, incIdentity string, incRemoteIP string, incMessage string, incError error) {
	isConsole := false
	//strURL := "http://172.31.27.36:55555/log"
	strURL := MapConfig["logEndpoint"]

	if strings.ToUpper(MapConfig["state"]) == "PRODUCTION" {
		isConsole = true
	} else if strings.ToUpper(MapConfig["state"]) == "STAGING" {
		isConsole = false
	}

	strError := fmt.Sprintf("%v", incError)

	mapRequest := make(map[string]interface{})
	mapRequest["tracecode"] = incTraceCode
	mapRequest["application"] = MapConfig["appname"]
	mapRequest["module"] = incResource[0]
	mapRequest["function"] = incResource[1]
	mapRequest["identity"] = incIdentity
	mapRequest["remoteip"] = incRemoteIP
	mapRequest["message"] = incMessage
	mapRequest["error"] = strError

	jsonRequest := ConvertMapInterfaceToJSON(mapRequest)
	var bodyByte = []byte(jsonRequest)
	_, err := http.Post(strURL, "application/json", bytes.NewBuffer(bodyByte))

	if err != nil {
		log.Fatal(err)
	} else {
		if !isConsole {
			if len(strError) == 0 || strings.Contains(strError, "nil") {
				DoLog("INFO", incTraceCode, incResource[0], incResource[1], incMessage, false, nil)
			} else {
				DoLog("ERROR", incTraceCode, incResource[0], incResource[1], incMessage, true, incError)
			}
		}
	}
}

func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}
