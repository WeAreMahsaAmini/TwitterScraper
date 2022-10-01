<template>
  <div v-if="loading">
    <div class="mt-3 text-center text-white"></div>
  </div>
  <div v-else>
    <div class="flex justify-center flex-col px-8 mt-4">
      <h1
        v-if="lang === 'en'"
        class="font-mono text-lg text-center text-gray-100"
      >
        Best of 'for ...' tweet collections (reasons for the protests from
        Iranian)
      </h1>
      <h1 v-else-if="lang === 'fa'" class="text-lg text-center text-gray-100">
        «لیست منتخب توییت های «برای
      </h1>
      <hr class="full-width fill-white mt-5 mx-8" />
      <NuxtLink :href="link" class="fixed top-2 left-3">
        <span class="fi" :class="flag"></span>
      </NuxtLink>
    </div>

    <div class="flex justify-center mt-2" id="pagination">
      <div
        v-for="index in pagesCount"
        :key="index"
        @click="pageChange(index)"
        :class="isActive(index)"
        class="pButton hover:bg-purple-300"
      ></div>
    </div>
    <div class="flex justify-center p-1 mt-2">
      <div v-if="lang === 'en'" class="p-1 content-center text-white">
        <span>Tweet per page:</span>
      </div>
      <input
        :value="tweetPerPage"
        @input="(event) => (tweetPerPage = parseInt(event.target.value || 0))"
        type="number"
        class="
          h-8
          w-12
          rounded
          bg-blue-400
          text-center
          opacity-75
          hover:opacity-100
          text-white
          p-0
          mb-0
        "
        required
      />
      <div
        v-if="lang === 'fa'"
        class="p-1 content-center text-white"
        style="direction: rtl"
      >
        توییت در صفحه:
      </div>
    </div>
    <div class="my-2 px-4">
      <masonry-wall
        :items="pageTweets"
        :column-width="320"
        :gap="15"
        :ssr-columns="1"
      >
        <template #default="{ item, index }">
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
  </div>
</template>

<script>
import json from "~/assets/tweets.json";
import Loading1 from "../components/Loading.vue";
const DEFAULT_TWEETS_PER_PAGE = 25;

export default {
  data() {
    return {
      loading: true,
      tweets: null,
      activeBtnKey: 1,
      pagesCount: null,
      pageTweets: [],
      tweetPerPage: DEFAULT_TWEETS_PER_PAGE,
      lang: "",
    };
  },
  async mounted() {
    this.tweets = json;
    this.pagesCount = Math.ceil(this.tweets.length / DEFAULT_TWEETS_PER_PAGE);
    this.setLang();
    this.loadTweets();
    this.loading = false;
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
    setLang() {
      this.lang = this.$route.hash.substring(1) || "en";
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
    "$route.hash"(newVal) {
      this.setLang();
    },
  },
  components: { Loading1 },
};
</script>
<style>
@import url("https://fonts.googleapis.com/css2?family=Changa:wght@600&family=Noto+Sans+Arabic:wght@500&display=swap");
body {
  font-family: "Noto Sans Arabic", sans-serif;
}
</style>
