package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

type ByLikeAndRetweet []*twitterscraper.TweetResult

func (a ByLikeAndRetweet) Len() int { return len(a) }
func (a ByLikeAndRetweet) Less(i, j int) bool {
	return a[i].Likes+a[i].Retweets > a[j].Likes+a[j].Retweets
}
func (a ByLikeAndRetweet) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

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
	f, err := os.Open("./web-ui/assets/tweets.json")
	if err != nil {
		fmt.Printf("error opening %s: %s", "tweets.json", err)
	}

	defer f.Close()

	byteValue, _ := io.ReadAll(f)
	tweets := []*twitterscraper.TweetResult{}
	json.Unmarshal(byteValue, &tweets)

	f, err = os.Create("./web-ui/assets/tweets.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("total begining: " + strconv.Itoa(len(tweets)))

	total := 0

	currentTime := time.Date(2022, 9, 16, 0, 0, 0, 0, time.Local)
	nextDay := currentTime.Add(24 * time.Hour)
	today := time.Now()

	for today.After(currentTime) {
		fmt.Println("The timeframe is from ", currentTime.Format("2006-1-2"), " until ", nextDay.Format("2006-1-2"))
		totalToday := 0

		for tweet := range scraper.SearchTweets(context.Background(),
			"برای #مهسا_امینی -filter:retweets "+"since:"+currentTime.Format("2006-1-2")+" until:"+nextDay.Format("2006-1-2"),
			10000) {
			total++
			totalToday++
			if tweet.Error != nil {
				panic(tweet.Error)
			}

			if strings.HasPrefix(tweet.Text, "برای") {
				tweets = append(tweets, tweet)
			}
		}
		fmt.Println("Total records received at " + currentTime.Format("2006-1-2") + " is: " + strconv.Itoa(totalToday))

		currentTime = nextDay
		nextDay = currentTime.Add(24 * time.Hour)
	}

	tweets = removeDuplicateValues(tweets)

	fmt.Println("Total records received: " + strconv.Itoa(total))
	fmt.Println("Total records accepted: " + strconv.Itoa(len(tweets)))

	sort.Sort(ByLikeAndRetweet(tweets))

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
