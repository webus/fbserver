package main

import (
	"os"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

// ClientID - Facebook Client ID credentials
var ClientID string
// ClientSecret - Facebook secret key
var ClientSecret string
// APIVersion - Facebook API version number
const APIVersion = "v2.5"

// api url builder
func getAPIUrl(method string) string {
	return "https://graph.facebook.com/" + APIVersion + method
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var fb FacebookToken
		err := decoder.Decode(&fb)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(fb.Token)
		workOAuth(fb.Token)
	}
}

// support twelve-factor app idea
func main() {
	ClientID = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")
	Token := os.Getenv("TOKEN")
	if Token != "" {
		workOAuth(Token)
	} else {
		log.Println("Listening: 0.0.0.0:3000")
		http.HandleFunc("/", handler)
		http.ListenAndServe(":3000", nil)
	}
}

// get all data from Facebook API
func workOAuth(token string) {
	conf := &oauth2.Config{
		ClientID: ClientID,
		ClientSecret: ClientSecret,
		Scopes: []string{"public_profile","email","user_friends"},
		Endpoint: facebook.Endpoint,
	}
	tok := &oauth2.Token{
		AccessToken: token,
	}
	fmt.Println(" === ")
	client := conf.Client(oauth2.NoContext, tok)
	var err error
	var profileName string
	if profileName, err = getProfileName(client); err != nil {
		log.Println(err)
		return
	}
	fmt.Println("User name:", profileName)
	fmt.Println(" === ")
	var profilePicture string
	if profilePicture, err = getProfilePicture(client); err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Profile picture saved to: ", profilePicture)
	fmt.Println(" === ")
	friends := getProfileFriendlist(client)
	fmt.Println("Friends:")
	profiles, err := friends()
	for err == nil {
		for _, user := range profiles {
			fmt.Println(user.Name)
		}
		profiles, err = friends()
	}
	fmt.Println(" === ")
	var data FacebookPublicProfile
	if data, err = getFullProfile(client); err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Locale : ", data.Locale)
	fmt.Println("Age range : ", data.AgeRange.Min)
	fmt.Println("Gender : ", data.Gender)
}
