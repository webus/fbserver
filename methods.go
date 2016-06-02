package main

import (
	"os"
	"io"
	//"log"
	"errors"
	"strings"
	"net/http"
	"encoding/json"
)

// Get user profile picture
func getProfilePicture(client *http.Client) (string, error) {
	resp, err := client.Get(getAPIUrl("/me/picture?type=large"))
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New("Unathorized. Check token.")
	}
	pictureFileExt := "jpg"
	if strings.ToLower(resp.Header.Get("Content-Type")) == "image/png" {
		pictureFileExt = "png"
	}
	pictureFilename := "picture." + pictureFileExt
	picture, err := os.Create(pictureFilename)
	if err != nil {
		return "", err
	}
	defer picture.Close()
	_, err = io.Copy(picture, resp.Body)
	if err != nil {
		return "", err
	}
	return pictureFilename, nil
}

// Get user profile name
func getProfileName(client *http.Client, optional ...string) (string, error) {
	url := getAPIUrl("/me?fields=name")
	if len(optional) == 1 {
		url = optional[0]
	}
	resp, err := client.Get(url)
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var profile FacebookProfile
	err = decoder.Decode(&profile)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New("Unathorized. Check token.")
	}
	return profile.Name, nil
}

// Get full user profile data
func getFullProfile(client *http.Client, optional ...string) (FacebookPublicProfile, error) {
	url := getAPIUrl("/me?fields=name,locale,age_range,gender")
	if len(optional) == 1 {
		url = optional[0]
	}
	resp, err := client.Get(url)
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var profile FacebookPublicProfile
	err = decoder.Decode(&profile)
	if err != nil {
		return FacebookPublicProfile{}, err
	}
	if resp.StatusCode != 200 {
		return FacebookPublicProfile{}, errors.New("Unathorized. Check token.")
	}
	return profile, nil
}

// Get all friends
func getProfileFriendlist(client *http.Client, optional ...string) func() ([]FacebookProfile, error) {
	next := getAPIUrl("/me/friends")
	if len(optional) == 1 {
		next = optional[0]
	}
	return func() ([]FacebookProfile, error) {
		if next == "" {
			return []FacebookProfile{}, errors.New("No data")
		}
		resp, err := client.Get(next)
		defer resp.Body.Close()
		if err != nil {
			return []FacebookProfile{}, err
		}
		if resp.StatusCode != 200 {
			return []FacebookProfile{}, errors.New("Unathorized. Check token.")
		}
		decoder := json.NewDecoder(resp.Body)
		var data FacebookFriends
		err = decoder.Decode(&data)
		if err != nil {
			return []FacebookProfile{}, err
		}
		emptyPaging := FacebookPaging{}
		if data.Paging != emptyPaging {
			next = data.Paging.Next
		} else {
			next = ""
		}
		return data.Data, nil
	}
}
