package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/S3B4SZ17/Email_service/management"
	pbEmail "github.com/S3B4SZ17/Email_service/proto/email_user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ExtractToken(c *gin.Context) string {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.String(http.StatusForbidden, "No Authorization header provided")
		c.Abort()

	}
	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		c.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
		c.Abort()
		return ""
	}
	return token
}

func ValidateToken(c *gin.Context) error {
	token := ExtractToken(c)
	user, _ := ExtractUser(&token)
	if user == nil {
		c.String(http.StatusForbidden, "User unauthorized")
		c.Abort()
	}
	return nil
}

func ExtractUser(token *string) (*pbEmail.EmailUser, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + *token)
	if err != nil {
		management.Log.Error("Get: " + err.Error() + "\n")
		return nil, err
	}
	defer resp.Body.Close()

	email_user := &pbEmail.EmailUser{}

	response, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(response, &email_user)
	// Check your errors!
	if err != nil {
		management.Log.Fatal(err.Error())
		return nil, err
	}
	return email_user, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
