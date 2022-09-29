package main

import (
	"io"
	"os"
	"log"
	"fmt"
	"context"
	"strconv"
	"strings"
	"encoding/json"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func removeDuplicateValues(intSlice []*twitterscraper.TweetResult) []*twitterscraper.TweetResult {
	keys := make(map[string]bool)
	list := []*twitterscraper.TweetResult{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry.ID]; !value {
			keys[entry.ID] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	scraper := twitterscraper.New()
	f, err := os.Open("tweets.json")
	if err != nil {
		fmt.Printf("error opening %s: %s", "tweets.json", err)
	}

	defer f.Close()

	byteValue, _ := io.ReadAll(f)
	tweets := []*twitterscraper.TweetResult{}
	json.Unmarshal(byteValue, &tweets)

	f, err = os.Create("tweets.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("total begining: " + strconv.Itoa(len(tweets)))

	total := 0
	for tweet := range scraper.SearchTweets(context.Background(), "برای #مهسا_امینی -filter:retweets ", 10000) {
		total++
		if tweet.Error != nil {
			panic(tweet.Error)
		}

		if strings.HasPrefix(tweet.Text, "برای") {
			tweets = append(tweets, tweet)
		}
	}

	tweets = removeDuplicateValues(tweets)

	fmt.Println("total received: " + strconv.Itoa(total))
	fmt.Println("total accepted: " + strconv.Itoa(len(tweets)))

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