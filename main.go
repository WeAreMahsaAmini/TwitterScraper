package main

import (
	"fmt"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
	scraper := twitterscraper.New()
	tweet, err := scraper.GetTweet("1575080854623174656")
	if err != nil {
		panic(err)
	}
	fmt.Println(tweet.Text)
}
