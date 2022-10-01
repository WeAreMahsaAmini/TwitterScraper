<template>
  <div>
    <div class="flex justify-center flex-col px-8 mt-4">
      <h1 class="font-mono text-3xl text-center text-gray-100">
        {{
          lang === "en"
            ? "List of reasons for the protests from Iranian"
            : "لیست برای های ایرانیان برای اعتراضات"
        }}
      </h1>
      <hr class="full-width fill-white mt-5 mx-8" />
      <a :href="link" :onclick="changeLang" class="fixed top-2 left-3">
        <span class="fi" :class="flag"></span>
      </a>
    </div>
    <div class="flex justify-center flex-wrap p-5 w-full">
      <masonry-wall
        :items="pageTweets"
        :ssr-columns="1"
        :column-width="300"
        :gap="16"
        :rtl="lang === 'fa'"
        class="w-full"
      >
        <template #default="{ item, index }" class="flex justify-center">
          <tweet-card
            class="bg-opacity-75"
            :id="item.Username"
            :hashtags="item.Hashtags"
            :likes="item.Likes"
            :retweets="item.Retweets"
            :timestamp="item.Timestamp"
            :direction="lang === 'en' ? 'ltr' : 'rtl'"
            :key="index"
          >
            {{ lang === "en" ? item.Translation : item.Text }}
          </tweet-card>
        </template>
      </masonry-wall>
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
import json from "~~/assets/tweets.json";

const DEFAULT_TWEETS_PER_PAGE = 100;
export default {
  data() {
    return {
      tweets: json,
      activeBtnKey: 1,
      pagesCount: Math.ceil(json.length / DEFAULT_TWEETS_PER_PAGE),
      pageTweets: [],
      tweetPerPage: DEFAULT_TWEETS_PER_PAGE,
      lang: document.location.hash.includes("fa") ? "fa" : "en",
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
    window.scrollTo({
      top: 0,
      behavior: "smooth",
    });
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
    changeLang() {
      this.lang = document.location.hash.includes("fa") ? "fa" : "en";
    },
  },
  computed: {
    flag() {
      return this.lang === "en" ? "fi-ir" : "fi-us";
    },
    link() {
      return this.lang === "en" ? "#fa" : "#en";
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
