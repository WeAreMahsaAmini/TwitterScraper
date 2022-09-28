package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
	scraper := twitterscraper.New()
	f, err := os.Create("tweets.json")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	tweets := []*twitterscraper.TweetResult{}
	for tweet := range scraper.SearchTweets(context.Background(),
		"برای #مهسا_امینی", 100) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}

		if strings.HasPrefix(tweet.Text, "برای") {

			tweets = append(tweets, tweet)
		}

	}

	b, err := json.Marshal(tweets)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err2 := f.WriteString(string(b))

	if err2 != nil {
		log.Fatal(err2)
	}
}
