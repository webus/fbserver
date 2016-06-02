package main

// FacebookToken - Facebook Token type
type FacebookToken struct {
	Token string
}

// FacebookProfile - Facebook profile type
type FacebookProfile struct {
	Id string
	Name string
}

// FacebookPublicProfile - A person's name, profile picture, locale, age range and gender are included by default with the public_profile permission.
type FacebookPublicProfile struct {
	Id string
	Name string
	Locale string
	AgeRange FacebookAgeRange `json:"age_range"`
	Gender string
}

// FacebookAgeRange - Facebook age range type
type FacebookAgeRange struct {
	Min int
}

// FacebookPaging - Facebook paging type
type FacebookPaging struct {
	Next string
}

// FacebookSummary - Facebook summary info type
type FacebookSummary struct {
	TotalCount int `json:"total_count,omitempty"`
}

// FacebookFriends - friends list
type FacebookFriends struct {
	Data []FacebookProfile
	Paging FacebookPaging `json:"paging,omitempty"`
	Summary FacebookSummary
}
