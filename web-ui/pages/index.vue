<template>
  <div class="leading-loose">
    <div class="flex justify-center flex-wrap py-5 w-full">
      <div v-for="(tweet, index) in pageTweets" :key="index">
        <tweet-card
          :id="tweet.Username"
          :hashtags="tweet.Hashtags"
          :likes="tweet.Likes"
          :retweets="tweet.Retweets"
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
  </div>
</template>

<script>
import json from "./../../tweets.json";
export default {
  data() {
    return {
      tweets: [],
      activeBtnKey: 1,
      pagesCount: 0,
      pageTweets: [],
      tweetPerPage: 10,
    };
  },
  mounted() {
    this.tweets = json;
    this.pagesCount = Math.ceil(this.tweets.length / this.tweetPerPage);
    this.pageTweets = this.tweets.slice(0, this.tweetPerPage);
  },
  computed: {},
  methods: {
    pageChange(key) {
      this.activeBtnKey = key;
      let startIdx = (key - 1) * this.tweetPerPage;
      this.pageTweets = this.tweets.slice(
        startIdx,
        startIdx + this.tweetPerPage
      );
    },
    isActive: function (key) {
      if (this.activeBtnKey === key) {
        return "active";
      }
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