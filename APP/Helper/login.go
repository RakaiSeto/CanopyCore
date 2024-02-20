package helper

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

var GoogleOauthLogin *oauth2.Config
var GoogleState string

func init() {
	GoogleState = uuid.New().String()
}

func GetGoogleInfo(code string) (string, error) {
	token, err := GoogleOauthLogin.Exchange(context.TODO(), code)
	if err != nil {
		return "", fmt.Errorf("code exchange failed: %s", err.Error())
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Response body converted to stringified JSON

	fmt.Println(resp.Status)
	fmt.Println(resp.Body)
	respbody, _ := io.ReadAll(resp.Body)

	return string(respbody), nil
}