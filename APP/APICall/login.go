package apicall

import (
	helper "canopyCore/APP/Helper"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoogleLogin(ctx *gin.Context) {
	url := helper.GoogleOauthLogin.AuthCodeURL(helper.GoogleState)

	ctx.Redirect(http.StatusSeeOther, url)
}

func GoogleLoginCallback(ctx *gin.Context) {
	state := ctx.Query("state")
	code := ctx.Query("code")

	if state != helper.GoogleState {
		errString := "state string is different"
		ctx.JSON(200, gin.H{
			"result": errString,
		})
	}

	data, err := helper.GetGoogleInfo(code)
	if err != nil {
		errString := err.Error()
		ctx.JSON(200, gin.H{
			"result": errString,
		})	
	}

	var x map[string]interface{}
	json.Unmarshal([]byte(data), &x)

	ctx.JSON(200, gin.H{
		"name": x["name"].(string),
		"email": x["email"].(string),
	})
}