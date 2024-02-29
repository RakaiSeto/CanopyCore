package modules

import (
	Config "CanopyCore/Configuration"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

type Receipt struct {
	Receiptno     string         `json:"receiptno"`
	Recipient     string         `json:"recipient"`
	Type          string         `json:"type"`
	Storename     string         `json:"storename"`
	Storeaddress  string         `json:"storeaddress"`
	Storecity     string         `json:"storecity"`
	Storecountry  string         `json:"storecountry"`
	Storecategory string         `json:"storecategory"`
	Datetime      string         `json:"datetime"`
	Order         []ReceiptOrder `json:"order"`
	Total         int            `json:"total"`
	Tax           int            `json:"tax"`
	Service       int            `json:"service"`
	Grandtotal    int            `json:"grandtotal"`
	Payment       string         `json:"payment"`
	Cardnumber    string         `json:"cardnumber"`
	Received      int            `json:"received"`
	Change        int            `json:"change"`
}

type ReceiptOrder struct {
	Name  string `json:"name"`
	Qty   int    `json:"qty"`
	Total int    `json:"total"`
}

func DoOnlineLog(chOnlineLog *amqp.Channel, messageId string, dateTime time.Time, logLevel string, appName string, moduleName string, functionName string,
	serverIPAddress string, logMessage string, isError bool, errMessage string) bool {
	isOK := false

	mapLogMessage := make(map[string]interface{})
	mapLogMessage["messageId"] = messageId
	mapLogMessage["dateTime"] = dateTime.UnixMilli() // Unix millisecond
	mapLogMessage["logLevel"] = logLevel
	mapLogMessage["appName"] = appName
	mapLogMessage["moduleName"] = moduleName
	mapLogMessage["functionName"] = functionName
	mapLogMessage["serverIPAddress"] = serverIPAddress
	mapLogMessage["logMessage"] = logMessage
	mapLogMessage["isError"] = isError
	if isError == false {
		mapLogMessage["errMessage"] = ""
	} else {
		mapLogMessage["errMessage"] = fmt.Sprintf("%+v", errMessage)
	}

	// Convert mapLogMessage to JSON
	jsonLogMessage := ConvertMapInterfaceToJSON(mapLogMessage)

	// Simple encrypt jsonLogMessage
	isEncOK, encLogMessage := SimpleStringEncrypt(jsonLogMessage)

	if isEncOK {
		// Publish to queue online log
		errP := chOnlineLog.Publish(
			"",                             // exchange
			Config.ConstOnlineLogQueueName, // Email queue router
			false,                          // mandatory
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(encLogMessage),
			})

		if errP != nil {
			fmt.Println("Failed to submit to online logging.")

			isOK = false
		} else {
			fmt.Println("Success to submit to online logging.")

			isOK = true
		}
	} else {
		fmt.Println("Failed to encrypt logging message.")

		isOK = false
	}

	return isOK
}

func DoLogNeo(chOnlineLog *amqp.Channel, messageId string, dateTime time.Time, logLevel string, incResource []string,
	serverIPAddress string, logMessage string, isError bool, errMessage error, isOnlineSubmit bool) bool {
	isOK := false

	if logLevel == "DEBUG" {
		message := ""
		if isError == true {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(incResource[0]), 20) + "   " + TrimStringLength(strings.ToUpper(incResource[1]), 20) + " " + logMessage + "." + fmt.Sprintf(" %v", errMessage)
		} else {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(incResource[0]), 20) + "   " + TrimStringLength(strings.ToUpper(incResource[1]), 20) + " " + logMessage
		}

		zapLogger.Debug(message)
	} else if logLevel == "ERROR" {
		message := ""
		if isError == true {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(incResource[0]), 20) + "   " + TrimStringLength(strings.ToUpper(incResource[1]), 20) + "   " + logMessage + "." + fmt.Sprintf(" %v", errMessage)
		} else {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(incResource[0]), 20) + "   " + TrimStringLength(strings.ToUpper(incResource[1]), 20) + "   " + logMessage
		}

		zapLogger.Error(message)
		//zapLogger.Debug(message)
	} else if logLevel == "WARNING" {
		message := ""
		if isError == true {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(incResource[0]), 20) + "   " + TrimStringLength(strings.ToUpper(incResource[1]), 20) + "   " + logMessage + "." + fmt.Sprintf(" %v", errMessage)
		} else {
			message = messageId + "   " + TrimStringLength(strings.ToUpper(incResource[0]), 20) + "   " + TrimStringLength(strings.ToUpper(incResource[1]), 20) + "   " + logMessage
		}

		zapLogger.Warn(message)
		//zapLogger.Debug(message)
	} else if logLevel == "INFO" {
		message := ""
		if isError == true {
			message = messageId + "   " + logMessage + "." + fmt.Sprintf(" %v", errMessage)
		} else {
			message = messageId + "   " + logMessage
		}

		zapLogger.Info(message)
		//zapLogger.Debug(message)
	} else {
		message := ""
		if isError == true {
			message = messageId + "   " + logMessage + "." + fmt.Sprintf(" %v", errMessage)
		} else {
			message = messageId + "   " + logMessage
		}

		//zapLogger.Info(message)
		zapLogger.Debug(message)
	}

	isOK = true
	// Submit online if required
	if isOnlineSubmit == true && chOnlineLog != nil {
		isOK = DoOnlineLog(chOnlineLog, messageId, dateTime, logLevel, Config.ConstAppName, incResource[0], incResource[1], serverIPAddress, logMessage, isError, fmt.Sprintf("%+v", errMessage))
	}

	return isOK
}

func DoSendReceiptToQueue(chOnlineLog *amqp.Channel, messageId string, dateTime time.Time, receipt *Receipt) bool {
	isOK := false

	// Convert mapLogMessage to JSON
	jsonLogMessage, err := json.Marshal(receipt)
	if err != nil {
		Logging(Resource(), messageId, "POS SEND RECEIPT", "SERVER", "Failed to marshal JSON", err)
	}

	// Publish to queue online log
	errP := chOnlineLog.Publish(
		"",                     // exchange
		"TRCV_PAYMENT_RECEIPT", // Email queue router
		false,                  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte(jsonLogMessage),
		})

	if errP != nil {
		Logging(Resource(), messageId, "POS SEND RECEIPT", "SERVER", "Failed to submit to queue", errP)

		isOK = false
	} else {
		fmt.Println("Success to submit to online logging.")

		isOK = true
	}

	return isOK
}
