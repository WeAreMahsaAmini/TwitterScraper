<template>
  <div>
    <div class="flex justify-center flex-wrap py-5 w-full">
      <div v-for="(tweet, index) in pageTweets" :key="index">
        <tweet-card
          class="bg-opacity-75"
          :id="tweet.Username"
          :hashtags="tweet.Hashtags"
          :likes="tweet.Likes"
          :retweets="tweet.Retweets"
          :timestamp="tweet.Timestamp"
        >
          {{ tweet.Text }}
        </tweet-card>
      </div>
    </div>
    <div class="flex justify-center" id="pagination">
      <div
        v-for="index in pagesCount"
        :key="index"
        @click="pageChange(index)"
        :class="isActive(index)"
        class="pButton hover:bg-purple-300"
      ></div>
    </div>
    <div class="flex justify-center mt-2">
      <input
        :value="tweetPerPage"
        @input="(event) => (tweetPerPage = parseInt(event.target.value || 0))"
        type="number"
        class="h-8 w-10 rounded bg-blue-400 text-center opacity-75 hover:opacity-100 text-white p-0"
        required
      />
    </div>
  </div>
</template>

<script>
import json from "@/assets/tweets.json";

const DEFAULT_TWEETS_PER_PAGE = 5;
export default {
  data() {
    return {
      tweets: json,
      activeBtnKey: 1,
      pagesCount: Math.ceil(json.length / DEFAULT_TWEETS_PER_PAGE),
      pageTweets: [],
      tweetPerPage: DEFAULT_TWEETS_PER_PAGE,
    };
  },
  mounted() {
    this.loadTweets();
  },
  updated() {
    if (this.tweetPerPage > this.tweets.length)
      this.tweetPerPage = this.tweets.length;
    if (this.tweetPerPage < 1) this.tweetPerPage = 1;
    this.pagesCount = Math.ceil(this.tweets.length / this.tweetPerPage);
  },
  beforeUpdate() {
    
  },
  methods: {
    pageChange(key) {
      this.activeBtnKey = key;
    },
    isActive: function (key) {
      if (this.activeBtnKey === key) {
        return "active";
      }
    },
    loadTweets() {
      let startIdx = (this.activeBtnKey - 1) * this.tweetPerPage;
      this.pageTweets = this.tweets.slice(
        startIdx,
        startIdx + this.tweetPerPage
      );
    },
  },
  watch: {
    activeBtnKey(val, oldVal) {
      this.loadTweets();
    },
    tweetPerPage(val, oldVal) {
      if (val > this.tweets.length) {
        this.tweetPerPage = this.tweets.length;
      } else if (this.tweetPerPage < 1) {
        this.tweetPerPage = 1;
      }
      this.loadTweets();
    },
  },
};
</script>
<style>
@import url("https://fonts.googleapis.com/css2?family=Changa:wght@600&family=Noto+Sans+Arabic:wght@500&display=swap");
body {
  font-family: "Noto Sans Arabic", sans-serif;
}
</style>
