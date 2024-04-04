package controllers

import (
	"encoding/json"
	"log"
	// "fmt"
	models "forum/model"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var Usergmail struct {
	Email string `json:"email"`
}
var (
	oauthConfig      *oauth2.Config
	oauthStateString = "random"
)

func init() {
	oauthConfig = &oauth2.Config{
		ClientID:     "861103099713-agktji6jm06diodorita9o8nnis10o93.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"email"},
		Endpoint:     google.Endpoint,
	}
}

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := oauthConfig.Exchange(r.Context(), code)
	if err != nil {
		log.Println("Error exchanging code: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	client := oauthConfig.Client(r.Context(), token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	println(userInfo)
	if err != nil {
		log.Println("Error getting user info: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer userInfo.Body.Close()
	err = json.NewDecoder(userInfo.Body).Decode(&Usergmail)
	if err != nil {
		log.Println("Error parsing user info: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	exist, uuid := models.GoogleUserCheck(Usergmail.Email)
	if !exist {
		uuid = models.GenerateUUID()
		username := models.CreateUsername()
		models.CreateUser(username, Usergmail.Email, uuid)
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "user",
		Value:  uuid,
		MaxAge: 3600,
	})
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

}
