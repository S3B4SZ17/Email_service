package services

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/S3B4SZ17/Email_service/management"
	"go.uber.org/zap"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// func ConfigureEmailService() *http.Client {

// 	b, err := os.ReadFile("credentials.json")
// 	if err != nil {
// 		management.Log.Error("Unable to read client secret file:", zap.String("credentials_file", err.Error()))
// 		os.Exit(1)
// 	}

// 	// If modifying these scopes, delete your previously saved token.json.
// 	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
// 	if err != nil {
// 		management.Log.Error("Unable to parse client secret file to config:", zap.String("err_message", err.Error()))
// 		os.Exit(1)
// 	}

// 	client := getClient(config)

// 	return client

// }

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
// func getClient(config *oauth2.Config) *http.Client {
// 	// The file token.json stores the user's access and refresh tokens, and is
// 	// created automatically when the authorization flow completes for the first
// 	// time.
// 	tokFile := "token.json"
// 	tok, err := tokenFromFile(tokFile)
// 	if err != nil {
// 		tok = GetTokenFromWeb(config)
// 		saveToken(tokFile, tok)
// 	}
// 	return config.Client(context.Background(), tok)
// }

// // Request a token from the web, then returns the retrieved token.
// func GetTokenFromWeb(config *oauth2.Config) *oauth2.Token {
// 	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
// 	fmt.Printf("Go to the following link in your browser then type the "+
// 		"authorization code: \n%v\n", authURL)

// 	// openbrowser(authURL)

// 	var urlresponse string
// 	if _, err := fmt.Scan(&urlresponse); err != nil {
// 		management.Log.Error("Unable to read URI response:", zap.String("err_message", err.Error()))
// 	}
// 	// parseURI(urlresponse)

// 	var authCode string
// 	if _, err := fmt.Scan(&authCode); err != nil {
// 		management.Log.Error("Unable to read authorization code:", zap.String("err_message", err.Error()))
// 	}

// 	tok, err := config.Exchange(context.TODO(), authCode)
// 	if err != nil {
// 		management.Log.Error("Unable to retrieve token from web:", zap.String("err_message", err.Error()))
// 	}
// 	return tok
// }

// func openbrowser(url_gmail string) {
// 	var err error

// 	switch runtime.GOOS {
// 	case "linux":
// 		err = exec.Command("xdg-open", url_gmail).Start()
// 	case "windows":
// 		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url_gmail).Start()
// 	case "darwin":
// 		err = exec.Command("open", url_gmail).Start()
// 	default:
// 		err = fmt.Errorf("unsupported platform")
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// // parse URL response
// func parseURI(url_string string) {
// 	u, err := url.Parse(url_string)

// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	fmt.Println("Scheme: ", u.Scheme)
// 	fmt.Println("Host: ", u.Host)

// 	queries := u.Query()
// 	fmt.Println("Query Strings: ")
// 	for key, value := range queries {
// 		fmt.Printf("  %v = %v\n", key, value)
// 	}
// 	fmt.Println("Path: ", u.Path)
// }

// // Retrieves a token from a local file.
// func tokenFromFile(file string) (*oauth2.Token, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()
// 	tok := &oauth2.Token{}
// 	err = json.NewDecoder(f).Decode(tok)
// 	return tok, err
// }

// // Saves a token to a file path.
// func saveToken(path string, token *oauth2.Token) {
// 	fmt.Printf("Saving credential file to: %s\n", path)
// 	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
// 	if err != nil {
// 		management.Log.Error("Unable to cache oauth token:", zap.String("err_message", err.Error()))
// 	}
// 	defer f.Close()
// 	json.NewEncoder(f).Encode(token)
// }
