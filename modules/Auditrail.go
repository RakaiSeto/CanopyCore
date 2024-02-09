package modules

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	Config "canopyCore/Configuration"
)

type TheWebActivity struct {
	Id              primitive.ObjectID `bson:"_id, omitempty"`
	MessageId       string             `bson:"messageId"`
	SessionId       string             `bson:"sessionId"`
	ResellerId      string             `bson:"resellerId"`
	ClientId        string             `bson:"clientId"`
	UserName        string             `bson:"username"`
	Privilege       string             `bson:"privilege"`
	RemoteIPAddress string             `bson:"remoteIPAddress"`
	Menu            string             `bson:"menu"`
	Activity        string             `bson:"activity"`
	WebURL          string             `bson:"webURL"`
	IsSuccess       bool               `bson:"isSuccess"`
	RespStatus      string             `bson:"respStatus"`
	ResultNote      string             `bson:"resultNote"`
}

func SaveWebActivity(mongoClient *mongo.Client, goContext context.Context, messageId string, sessionId string,
	resellerId string, username string, privilege string, remoteIPAddress string, menu string, activity string, webURL string,
	isSuccess bool, respStatus string, resultNote string) bool {
	isOK := false

	collection := mongoClient.Database(Config.ConstMongoAuditrailDB).Collection(Config.ConstAuditrailMongoColl)

	theWebActivityData := TheWebActivity{
		Id:              primitive.NewObjectID(),
		MessageId:       messageId,
		SessionId:       sessionId,
		ResellerId:      resellerId,
		UserName:        username,
		Privilege:       privilege,
		RemoteIPAddress: remoteIPAddress,
		Menu:            menu,
		Activity:        activity,
		WebURL:          webURL,
		IsSuccess:       isSuccess,
		RespStatus:      respStatus,
		ResultNote:      resultNote,
	}

	insertResult, err := collection.InsertOne(goContext, theWebActivityData)

	if err != nil {
		isOK = false
		fmt.Printf("%+v", err)
	} else {
		isOK = true
		fmt.Printf("insertResult: %+v\n", insertResult)
	}

	return isOK
}
