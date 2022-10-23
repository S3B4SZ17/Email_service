package services

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/S3B4SZ17/Email_service/management"
	pbEmail "github.com/S3B4SZ17/Email_service/proto/email_user"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func SendEmail(message *pbEmail.EmailMessage) (res *pbEmail.EmailResponse, err error) {

	return res, err
}

// func GetAuthenticatedUser() (res *pbEmail.EmailUser, err error) {

// 	return res, err
// }

func AuthenticateUser(c *gin.Context, oauthConfGl *oauth2.Config, oauthStateStringGl *string) {
	compareState(c, oauthConfGl, oauthStateStringGl)
	checkCode(c, oauthConfGl, oauthStateStringGl)

}

func checkCode(c *gin.Context, oauthConfGl *oauth2.Config, oauthStateStringGl *string) *oauth2.Token {
	management.Log.Info("Getting code from Oauth2")
	code := c.Request.FormValue("code")

	if code == "" {
		management.Log.Warn("Code not found..")
		reason := c.Request.FormValue("error_reason")
		if reason == "user_denied" {
			management.Log.Error("User has denied Permission..")
			front_end_url := url.URL{Path: viper.GetString("front_end_url") + "/loginerror"}
			c.Redirect(http.StatusUnauthorized, front_end_url.RequestURI())
		}

	} else {
		token, err := oauthConfGl.Exchange(c, code)
		if err != nil {
			management.Log.Error("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
			front_end_url := url.URL{Path: viper.GetString("front_end_url") + "/loginerror"}
			c.Redirect(http.StatusTemporaryRedirect, front_end_url.RequestURI())
		}

		c.SetCookie("gmail_token", url.QueryEscape(token.AccessToken), 1000, "/authorized", c.Request.URL.Hostname(), false, false)
		front_end_url := url.URL{Path: viper.GetString("front_end_url") + "/authorized/user"}
		c.Redirect(http.StatusTemporaryRedirect, front_end_url.RequestURI())
		return token
	}
	return nil
}

func emailSrv(c *gin.Context, oauthConfGl *oauth2.Config, oauthStateStringGl *string, token *oauth2.Token) {
	gmail_client := oauthConfGl.Client(c, token)
	gmail, _ := GetGmailService(c, gmail_client)
	GetGmailLabels(gmail)
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		management.Log.Error("Get: " + err.Error() + "\n")
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	defer resp.Body.Close()

	email_user := &pbEmail.EmailUser{}

	response, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(response, &email_user)
	// Check your errors!
	if err != nil {
		management.Log.Fatal(err.Error())
	}
}

func compareState(c *gin.Context, oauthConfGl *oauth2.Config, oauthStateStringGl *string) {
	state := c.Request.FormValue("state")
	management.Log.Info(state)
	if state != *oauthStateStringGl {
		management.Log.Error("invalid oauth state, expected " + *oauthStateStringGl + ", got " + state + "\n")
		front_end_url := url.URL{Path: viper.GetString("front_end_url") + "/error"}
		c.Redirect(http.StatusUnauthorized, front_end_url.RequestURI())
	}
}
