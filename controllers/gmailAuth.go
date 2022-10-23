package controllers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/S3B4SZ17/Email_service/management"
	"github.com/S3B4SZ17/Email_service/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

var (
	oauthConfGl        = &oauth2.Config{}
	oauthStateStringGl = ""
	host               = "localhost"
	gRPCListener       = "50051"
)

/*
InitializeOAuthGoogle Function
*/
func InitializeOAuthGoogle() {
	oauthConfGl.ClientID = viper.GetString("google.client_id")
	oauthConfGl.ClientSecret = viper.GetString("google.client_secret")
	oauthConfGl.RedirectURL = viper.GetString("google.redirect_uri")
	oauthConfGl.Scopes = []string{gmail.GmailReadonlyScope, "https://www.googleapis.com/auth/userinfo.email"}
	oauthConfGl.Endpoint = google.Endpoint
	oauthStateStringGl = viper.GetString("oauthStateString")
	management.Log.Info("Oauth Config", zap.String("redirect_uri", oauthConfGl.RedirectURL))
}

/*
HandleGoogleLogin Function
*/
func HandleGoogleLogin(c *gin.Context) {
	HandleLogin(c, oauthConfGl, oauthStateStringGl)
}

/*
CallBackFromGoogle Function
*/
func CallBackFromGoogle(c *gin.Context) {
	management.Log.Info("Callback-gl..")
	// // Set up a connection to the AddTwoNumbers server.
	// conn, err := grpc.Dial(host+":"+gRPCListener, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	management.Log.Error("Did not connect: " + err.Error())
	// }
	// defer conn.Close()
	// client := pbEmail.NewGetAuthenticatedUserClient(conn)
	// in := &emptypb.Empty{}

	services.AuthenticateUser(c, oauthConfGl, &oauthStateStringGl)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	c.JSON(415, gin.H{"errcode": 415, "description": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, res)
}

/*
HandleLogin Function
*/
func HandleLogin(c *gin.Context, oauthConf *oauth2.Config, oauthStateString string) {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)

	if err != nil {
		management.Log.Error("Parse: " + err.Error())
	}
	management.Log.Info(URL.String())

	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	management.Log.Info(url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}
