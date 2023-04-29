package main

type News struct {
	News []NewsItem `json:"news"`
}

type NewsItem struct {
	Title   string `json:"title"`
	Perex   string `json:"perex"`
	Link    string `json:"link"`
	Server  string `json:"server"`
	PubDate string `json:"pub_date"`
}
