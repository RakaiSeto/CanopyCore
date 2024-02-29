package apicall

import (
	helper "CanopyCore/APP/Helper"
	"CanopyCore/modules"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoogleLogin(ctx *gin.Context) {
	url := helper.GoogleOauthLogin.AuthCodeURL(helper.GoogleState)

	ctx.Redirect(http.StatusSeeOther, url)
}

func GoogleLoginCallback(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		incTraceCode := modules.GenerateUUID()
		incClientIP := ctx.ClientIP()
		state := ctx.Query("state")
		code := ctx.Query("code")

		if state != helper.GoogleState {
			modules.Logging(modules.Resource(), incTraceCode, "GUESTLOGIN", incClientIP, "Google Oauth state mismatch", nil)
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

		// query := `SELECT `

		ctx.JSON(200, gin.H{
			"name":  x["name"].(string),
			"email": x["email"].(string),
		})
	}
}
