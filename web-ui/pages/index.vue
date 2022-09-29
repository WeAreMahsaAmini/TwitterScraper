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
        @input="(event) => (tweetPerPage = parseInt(event.target.value))"
        type="number"
        class="
          h-8
          w-10
          rounded
          bg-blue-400
          text-center
          opacity-75
          hover:opacity-100
          text-white
          p-0
        "
        required
      />
    </div>
  </div>
</template>

<script>
import json from "@/assets/tweets.json";
export default {
  data() {
    return {
      tweets: [],
      activeBtnKey: 1,
      pagesCount: 0,
      pageTweets: [],
      tweetPerPage: 5,
    };
  },
  mounted() {
    this.tweets = json;
    this.pagesCount = Math.ceil(this.tweets.length / this.tweetPerPage);
    this.pageChange(this.activeBtnKey);
  },
  updated() {
    if (this.tweetPerPage > this.tweets.length)
      this.tweetPerPage = this.tweets.length;
    if (this.tweetPerPage < 1) this.tweetPerPage = 1;
    this.pagesCount = Math.ceil(this.tweets.length / this.tweetPerPage);
    this.pageChange(this.activeBtnKey);
  },
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