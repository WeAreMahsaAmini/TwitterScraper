package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

type (
	Tweet struct {
		Hashtags     []string
		HTML         string
		ID           string
		Likes        int
		PermanentURL string
		Photos       []string
		Replies      int
		Retweets     int
		Text         string
		Timestamp    int64
		UserID       string
		Username     string
		Videos       []twitterscraper.Video
	}
)

const TWEETS_FILE_PATH = "./web-ui/assets/tweets.json"

// Start date for search for tweets
var START_DATE = time.Date(2022, 9, 16, 0, 0, 0, 0, time.Local)

func main() {
	// Load stored tweets ids from file (fetched before)
	var err error
	var tweetsID []string
	if empty, err := isFileEmpty(TWEETS_FILE_PATH); err != nil {
		log.Fatalf("Error occurred during check file emptiness - %s", err.Error())
	} else if !empty {
		tweetsID, err = loadTweetsIDFromFile(TWEETS_FILE_PATH)
		if err != nil {
			log.Fatalf("Error occurred during load tweets id form file - %s", err.Error())
		}
	}
	fmt.Println("Number of stored tweets is: ", len(tweetsID))

	// Fetch needed tweets
	totalFetched, tweets, err := fetchTweets(tweetsID)
	if err != nil {
		// TODO: there is should be a better error handling
		log.Fatalf("Error occurred during fetch tweets - %s", err.Error())
	}
	// Remove duplicate tweets
	tweets = removeDuplicateValues(tweets)
	fmt.Println("Total records received: " + strconv.Itoa(totalFetched))
	fmt.Println("Total records accepted: " + strconv.Itoa(len(tweets)))

	// Sort tweets by likes and retweets
	sort.Sort(ByLikeAndRetweet(tweets))

	// Prune tweets
	pruneTweets := []*Tweet{}
	for _, tweet := range tweets {
		pruneTweets = append(pruneTweets, &Tweet{
			Hashtags:     tweet.Hashtags,
			HTML:         tweet.HTML,
			ID:           tweet.ID,
			Likes:        tweet.Likes,
			PermanentURL: tweet.PermanentURL,
			Photos:       tweet.Photos,
			Replies:      tweet.Replies,
			Retweets:     tweet.Retweets,
			Text:         tweet.Text,
			Timestamp:    tweet.Timestamp,
			UserID:       tweet.UserID,
			Username:     tweet.Username,
			Videos:       tweet.Videos,
		})
	}

	// Convert data to Json and write it in file
	err = dumpTweetsToFile(TWEETS_FILE_PATH, pruneTweets)
	if err != nil {
		log.Fatalf("Error occurred during dump/write tweets in file - %s", err.Error())
	}
	fmt.Println("Finished")
}
