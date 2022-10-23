package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"

	"github.com/S3B4SZ17/Email_service/management"
	pbEmail "github.com/S3B4SZ17/Email_service/proto/email_user"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func GetAuthenticatedUser(c *gin.Context, oauthConfGl *oauth2.Config, oauthStateStringGl *string) {
	state := c.Request.FormValue("state")
	management.Log.Info(state)
	if state != *oauthStateStringGl {
		management.Log.Error("invalid oauth state, expected " + *oauthStateStringGl + ", got " + state + "\n")
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	code := c.Request.FormValue("code")
	management.Log.Info(code)

	if code == "" {
		management.Log.Warn("Code not found..")
		c.Writer.Write([]byte("Code Not Found to provide AccessToken..\n"))
		reason := c.Request.FormValue("error_reason")
		if reason == "user_denied" {
			management.Log.Error("User has denied Permission..")
		}

	} else {
		token, err := oauthConfGl.Exchange(c, code)
		if err != nil {
			management.Log.Error("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
		}

		management.Log.Info("TOKEN>> AccessToken>> " + token.AccessToken)
		management.Log.Info("TOKEN>> Expiration Time>> " + token.Expiry.String())
		management.Log.Info("TOKEN>> RefreshToken>> " + token.RefreshToken)

		// gmail_client := oauthConfGl.Client(c, token)
		// gmail, _ := GetGmailService(c, gmail_client)
		// GetGmailLabels(gmail)

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		// if resp.StatusCode == 401 {

		// }
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

		c.SetCookie("gmail_token", url.QueryEscape(token.AccessToken), 1000, "/user", c.Request.URL.Hostname(), false, false)
		front_end_url := url.URL{Path: viper.GetString("front_end_url") + "/user"}
		c.Redirect(http.StatusTemporaryRedirect, front_end_url.RequestURI())

		// management.Log.Info("parseResponseBody: " + string(response) + "\n")

		// c.Writer.Write([]byte("Hello, I'm protected\n"))
		// c.Writer.Write([]byte(string(response)))

		// return
	}

}

func ConfigureEmailService() *http.Client {

	b, err := os.ReadFile("credentials.json")
	if err != nil {
		management.Log.Error("Unable to read client secret file:", zap.String("credentials_file", err.Error()))
		os.Exit(1)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		management.Log.Error("Unable to parse client secret file to config:", zap.String("err_message", err.Error()))
		os.Exit(1)
	}

	client := getClient(config)

	return client

}

func GetGmailService(ctx context.Context, client *http.Client) (*gmail.Service, error) {

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		management.Log.Error("Unable to retrieve Gmail client:", zap.String("err_message", err.Error()))
		os.Exit(1)
	}

	return srv, err
}

func GetGmailLabels(mail *gmail.Service) {

	// Hardcoding me since it represents the context of the user that is authenticated.
	// Reference this thread https://stackoverflow.com/questions/26135310/gmail-api-returns-403-error-code-and-delegation-denied-for-user-email
	user := "me"
	r, err := mail.Users.Labels.List(user).Do()
	if err != nil {
		management.Log.Error("Unable to retrieve labels:", zap.String("err_message", err.Error()))
		os.Exit(1)
	}
	if len(r.Labels) == 0 {
		fmt.Println("No labels found.")
		return
	}
	fmt.Println("Labels:")
	for _, l := range r.Labels {
		fmt.Printf("- %s\n", l.Name)
	}
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = GetTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func GetTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	openbrowser(authURL)

	var urlresponse string
	if _, err := fmt.Scan(&urlresponse); err != nil {
		management.Log.Error("Unable to read URI response:", zap.String("err_message", err.Error()))
	}
	parseURI(urlresponse)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		management.Log.Error("Unable to read authorization code:", zap.String("err_message", err.Error()))
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		management.Log.Error("Unable to retrieve token from web:", zap.String("err_message", err.Error()))
	}
	return tok
}

func openbrowser(url_gmail string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url_gmail).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url_gmail).Start()
	case "darwin":
		err = exec.Command("open", url_gmail).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

// parse URL response
func parseURI(url_string string) {
	u, err := url.Parse(url_string)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Scheme: ", u.Scheme)
	fmt.Println("Host: ", u.Host)

	queries := u.Query()
	fmt.Println("Query Strings: ")
	for key, value := range queries {
		fmt.Printf("  %v = %v\n", key, value)
	}
	fmt.Println("Path: ", u.Path)
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		management.Log.Error("Unable to cache oauth token:", zap.String("err_message", err.Error()))
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
