package main

import (
	"context"
	"fmt"
	"log"
	"os"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
	scraper := twitterscraper.New()
	f, err := os.Create("tweets.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	fmt.Println("done")
	for tweet := range scraper.SearchTweets(context.Background(),
		"برای #مهسا_امینی", 50) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		_, err2 := f.WriteString(tweet.Text + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
		_, err3 := f.WriteString("-------------------------\n")

		if err3 != nil {
			log.Fatal(err3)
		}
		fmt.Println(tweet.Text)
	}
}
