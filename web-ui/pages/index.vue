<template>
  <div v-if="loading">
    <div class="mt-3 text-center text-white"></div>
  </div>
  <div class="bg-gray-50" v-else>
    <div class="relative bg-zinc-900 border-b shadow">
      <div
        class="
          mx-auto
          max-w-6xl
          h-16
          flex
          items-center
          justify-between
          px-10
          lg:px-0
        "
      >
        <div v-if="lang === 'en'">
          <h1 class="text-xl text-white font-normal">
            Best of <span class="font-black text-blue-400">for…</span> tweet
            collections
          </h1>
          <h2 class="text-white text-sm leading-none">
            Reasons for the protests from Iranian
          </h2>
        </div>
        <NuxtLink :href="link">
          <span class="fi" :class="flag"></span>
        </NuxtLink>
        <div v-if="lang === 'fa'" dir="rtl">
          <h1 class="text-xl text-white font-normal">
            لیست منتخب توییت‌های
            <span class="font-black text-blue-400">برای…</span>
          </h1>
          <h2 class="text-white text-sm leading-none">
            دلایلی برای اعتراضات در ایران
          </h2>
        </div>
      </div>
    </div>
    <div
      class="
        relative
        gap-2
        items-center
        bg-gray-200
        rounded-b-3xl
        mb-16
        justify-center
        text-center
      "
    >
      <div class="max-w-6xl mx-auto">
        <div class="relative grid grid-cols-1 md:grid-cols-6 lg:grid-cols-5">
          <div
            class="
              col-span-4
              lg:col-span-3
              px-10
              py-10
              lg:px-0 lg:py-10
              text-center
              md:text-start
              z-10
              flex
              items-center
            "
          >
            <div class="space-y-10">
              <div>
                <h1 class="text-4xl font-black">
                  Best of
                  <span class="font-black text-blue-400">for…</span> tweet
                  collections
                </h1>
                <h2 class="text-xl font-light mb-3">
                  Reasons for the protests from Iranian
                </h2>
              </div>
            </div>
          </div>
          <div
            class="
              absolute
              bottom-0
              w-full
              opacity-10
              md:opacity-100
              pt-5
              z-0
              md:relative
              col-span-2
              lg:col-span-2
              flex
              justify-end
              content-end
              items-end
              align-bottom
            "
          >
            <img
              src="../assets/images/poster.png"
              alt="Mahsa Amini Poster"
              class="w-3/4 mx-auto"
            />
          </div>
        </div>
      </div>
    </div>

    <div
      class="
        flex
        justify-between
        items-center
        max-w-6xl
        mx-auto
        md:rounded-xl
        bg-white
        p-3
        shadow
      "
    >
      <div class="flex items-center gap-x-2">
        <div v-if="lang === 'en'">
          <span>Tweet per page:</span>
        </div>
        <input
          :value="tweetPerPage"
          @input="
            (event) => tweetPerPageChange(parseInt(event.target.value || 0))
          "
          type="number"
          class="
            text-sm
            py-2
            w-12
            rounded
            text-center
            border border-gray-300
            mb-0
          "
          required
        />
        <div v-if="lang === 'fa'" dir="rtl">توییت در صفحه:</div>
      </div>
      <div class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
        <div
          @click="pageChange(activeBtnKey - 1)"
          v-if="activeBtnKey > 1"
          class="
            relative
            inline-flex
            cursor-pointer
            items-center
            px-2
            py-2
            rounded-l-md
            border border-gray-300
            text-sm
            font-medium
            text-gray-500
            hover:bg-gray-50
          "
        >
          <svg
            class="h-5 w-5"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
            aria-hidden="true"
          >
            <path
              fill-rule="evenodd"
              d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
              clip-rule="evenodd"
            />
          </svg>
        </div>
        <div
          v-for="index in pagesCount"
          :key="index"
          @click="pageChange(index)"
          :class="isActive(index)"
          class="
            bg-white
            border-gray-300
            text-gray-500
            hover:bg-gray-50
            relative
            inline-flex
            items-center
            px-4
            py-2
            border
            text-sm
            font-medium
            cursor-pointer
          "
        >
          {{ index }}
        </div>
        <div
          @click="pageChange(activeBtnKey + 1)"
          v-if="activeBtnKey < pagesCount"
          class="
            relative
            inline-flex
            cursor-pointer
            items-center
            px-2
            py-2
            rounded-r-md
            border border-gray-300
            bg-white
            text-sm
            font-medium
            text-gray-500
            hover:bg-gray-50
          "
        >
          <svg
            class="h-5 w-5"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
            aria-hidden="true"
          >
            <path
              fill-rule="evenodd"
              d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
              clip-rule="evenodd"
            />
          </svg>
        </div>
      </div>
    </div>
    <div class="flex justify-center flex-wrap p-5 w-full">
      <masonry-wall
        :items="pageTweets"
        :ssr-columns="1"
        :column-width="300"
        :gap="16"
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
  </div>
</template>

<script>
import json from "~/assets/tweets.json";
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
    tweetPerPageChange(key) {
      this.tweetPerPage = key;
      this.pagesCount = Math.ceil(this.tweets.length / this.tweetPerPage);
      this.loadTweets();
    },
    isActive: function (key) {
      if (this.activeBtnKey === key) {
        return "z-10 bg-blue-50 hover:bg-blue-50 border-blue-500 text-blue-500 font-black";
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
};
</script>
<style>
@import url("https://fonts.googleapis.com/css2?family=Vazirmatn:wght@100;200;300;400;500;600;700;800;900&display=swap");

body {
  font-family: "Nunito", sans-serif;
}

[dir="rtl"] {
  font-family: "Vazirmatn";
}
</style>
