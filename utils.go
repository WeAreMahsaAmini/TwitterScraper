package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

// Types
type ByLikeAndRetweet []*twitterscraper.TweetResult

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
func isFileEmpty(path string) (bool, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0444)
	if err != nil {
		return false, err
	}
	b := make([]byte, 1)
	_, err = f.Read(b)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return true, nil
		}
		return false, err
	}
	if bytes.Equal(b, []byte{}) {
		return true, nil
	} else {
		return false, nil
	}
}
func loadTweetsFromFile(path string) ([]*Tweet, error) {
	if empty, err := isFileEmpty(path); err != nil {
		return nil, fmt.Errorf("error occurred during check emptiness of %s file - %w", path, err)
	} else if empty {
		return nil, fmt.Errorf("the %s file is emtpy", path)

	}
	f, err := os.Open(path)

	defer f.Close()

	if err != nil {
		return nil, err
	}
	tweets := []*Tweet{}
	err = json.NewDecoder(f).Decode(&tweets)
	return tweets, err
}

// APPENDS TO FILE
func dumpTweetsToFile(path string, tweets []*Tweet) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	return json.NewEncoder(f).Encode(tweets)
}

// TODO: It should extract only ID fields from json file.(NOT LOAD ALL DATA AND THEN FILTER THE "ID" FIELDS)
func loadTweetsIDFromFile(path string) ([]string, error) {
	data, err := loadTweetsFromFile(path)
	if err != nil {
		return nil, err
	}
	ids := []string{}
	for i := range data {
		ids = append(ids, data[i].ID)
	}
	return ids, nil
}

// UNUSED
// func getTweetsIDs(tweets []*Tweet) []string {
// 	list := []string{}
// 	for i := range tweets {
// 		list = append(list, tweets[i].ID)
// 	}
// 	return list
// }

// We'll do a string binary search
// ID's are number so we can do binary sort on numbers, But then we should convert all string ID's to int (Probably it takes more cost)
func isIDExistInIDs(id string, ids []string) bool {
	return stringBinarySearch(id, ids)
}

func stringBinarySearch(element string, slice []string) bool {
	if !sort.StringsAreSorted(slice) {
		sort.Strings(slice)
	}
	// SearchStrings return where element(input) should be inserted in slice (say returns i)
	// If i == len(ids), means that element not exists and should be inserted in end of slice
	// If i < len(ids), means that element exists or should be inserted in slice[i], So if slice[i] == element, then means that element exists, otherwise it's not exists
	i := sort.SearchStrings(slice, element)
	if i < len(slice) && slice[i] == element {
		return true
	} else {
		return false
	}
}

// Fetch ("برای") tweets
// It's recives stored tweets id for don't fetch tweets that already fetched
// Returns:
// Number of fetched tweets, Tweets, Error
func fetchTweets(stored_tweets_id []string) (int, []*twitterscraper.TweetResult, error) {
	// fetched count (filtered and unfiltered)
	// It's all tweets that program fetched, not only needed ones
	totalCount := 0
	// Fetched tweets
	tweets := []*twitterscraper.TweetResult{}
	// Start time for search
	startDay := START_DATE
	nextDay := startDay.Add(24 * time.Hour)
	today := time.Now()
	// Twitter scrapper
	scraper := twitterscraper.New()

	// Fetch tweets from start time to today
	for today.After(startDay) {
		fmt.Println("The timeframe is from ", startDay.Format("2006-1-2"), " until ", nextDay.Format("2006-1-2")+":")
		// Number of twets that fetched by program, not only needed ones
		totalCountToday := 0
		// Number of tweets that passed the filter and they are was needed ones ("برای" tweets)
		totalPassedToday := 0
		// Get tweets in loop
		for tweet := range scraper.SearchTweets(context.Background(), "برای #مهسا_امینی -filter:retweets "+"since:"+startDay.Format("2006-1-2")+" until:"+nextDay.Format("2006-1-2"), 10000) {
			// Increase all fetched count
			totalCount++
			// Increase today fetched count
			totalCountToday++
			// Return error if error happened
			if tweet.Error != nil {
				return 0, nil, fmt.Errorf("error occurred during get tweets - %w", tweet.Error)
			}
			// If fetched tweet was already exists in the file (already fetched), pass it
			if stored_tweets_id != nil {
				if isIDExistInIDs(tweet.ID, stored_tweets_id) {
					continue
				}
			}
			// If tweet has perfix of "برای" add it to fetched tweets
			if strings.HasPrefix(tweet.Text, "برای") {
				tweets = append(tweets, tweet)
				totalPassedToday++
			}
		}
		// Print fetched tweets in specific date
		fmt.Println("Total tweets fetched at " + startDay.Format("2006-1-2") + " is: " + strconv.Itoa(totalCountToday))
		fmt.Println("Total \"برای\" tweets fetched at " + startDay.Format("2006-1-2") + " is: " + strconv.Itoa(totalPassedToday) + "\n")

		// Change range of days
		startDay = nextDay
		nextDay = startDay.Add(24 * time.Hour)
	}
	// Return tweets
	return totalCount, tweets, nil
}

// ByLikeAndRetweet methods

func (a ByLikeAndRetweet) Len() int { return len(a) }
func (a ByLikeAndRetweet) Less(i, j int) bool {
	return a[i].Likes+a[i].Retweets > a[j].Likes+a[j].Retweets
}
func (a ByLikeAndRetweet) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
