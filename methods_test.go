package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProfileName(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		fmt.Fprintln(w,`{"id":"1","name":"Test User"}`)
	}))
	defer ts.Close()
	data, err := getProfileName(&http.Client{}, ts.URL)
	if err != nil {
		t.Error(err)
	}
	if data != "Test User" {
		t.Error("Wrong user name")
	}
}

func TestGetProfileFriendlist(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		fmt.Fprintln(w,`{"data":[{"name":"A","id":"1"},{"name":"B","id":"2"}],"summary":{"total_count":2}}`)
	}))
	defer ts.Close()
	friends := getProfileFriendlist(&http.Client{}, ts.URL)
	var profilesAll []FacebookProfile
	profiles, err := friends()
	log.Println(profiles)
	log.Println(err)
	appendProfiles := func(profiles []FacebookProfile) {
		for _, profile := range profiles {
			profilesAll = append(profilesAll, profile)
		}
	}
	appendProfiles(profiles)
	for err == nil {
		profiles, err = friends()
		appendProfiles(profiles)
	}
	log.Println(len(profilesAll))
	if len(profilesAll) != 2 {
		t.Error("Wrong profiles count")
	}
}
