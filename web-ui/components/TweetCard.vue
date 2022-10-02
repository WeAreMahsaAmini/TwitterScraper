<script>
import { HeartIcon, ArrowPathRoundedSquareIcon } from "@heroicons/vue/24/outline";

export default {
  props: [
    "name",
    "hashtags",
    "id",
    "likes",
    "retweets",
    "timestamp",
    "direction",
  ],
  components: { ArrowPathRoundedSquareIcon, HeartIcon },
  computed: {
    humanLikes() {
      if (this.likes > 1000) {
        return `${(this.likes / 1000).toFixed(1)}K`;
      }
      return this.likes;
    },
    humanRetweets() {
      if (this.retweets > 1000) {
        return `${(this.retweets / 1000).toFixed(1)}K`;
      }
      return this.retweets;
    },
    showDate() {
      let date = new Date(this.timestamp * 1000).toDateString();
      let firstSpace = date.indexOf(" ");
      let format = date.slice(firstSpace);
      return format;
    },
    showTime() {
      let time = new Date(this.timestamp * 1000);
      let format = ("0" + time.getHours()).slice(-2) + ":" + ("0" + time.getMinutes()).slice(-2);
      return format;
    }
  },
};
</script>
<template>
  <div class="bg-white border-gray-200 p-4 rounded-xl border w-full md:w-80 shadow">
    <div class="flex justify-between">
      <div class="flex items-center">
        <div class="text-base leading-tight">
          <!-- <span class="text-black font-bold block ">Visualize Value</span> -->

          <span class="text-black font-bold block">@{{ id }}</span>
        </div>
      </div>
      <svg class="text-blue-400 h-6 w-auto inline-block fill-current" viewBox="0 0 24 24">
        <g>
          <path
            d="M23.643 4.937c-.835.37-1.732.62-2.675.733.962-.576 1.7-1.49 2.048-2.578-.9.534-1.897.922-2.958 1.13-.85-.904-2.06-1.47-3.4-1.47-2.572 0-4.658 2.086-4.658 4.66 0 .364.042.718.12 1.06-3.873-.195-7.304-2.05-9.602-4.868-.4.69-.63 1.49-.63 2.342 0 1.616.823 3.043 2.072 3.878-.764-.025-1.482-.234-2.11-.583v.06c0 2.257 1.605 4.14 3.737 4.568-.392.106-.803.162-1.227.162-.3 0-.593-.028-.877-.082.593 1.85 2.313 3.198 4.352 3.234-1.595 1.25-3.604 1.995-5.786 1.995-.376 0-.747-.022-1.112-.065 2.062 1.323 4.51 2.093 7.14 2.093 8.57 0 13.255-7.098 13.255-13.254 0-.2-.005-.402-.014-.602.91-.658 1.7-1.477 2.323-2.41z">
          </path>
        </g>
      </svg>
    </div>
    <p class="text-black block text-base mt-3" :dir="direction">
      <slot />
    </p>
    <p class="text-gray-500 text-base py-1 my-0.5">{{ showTime }} · {{ showDate }}</p>
    <div class="border-gray-200 border border-b-0 my-1"></div>
    <div class="text-gray-500 flex mt-3">
      <div class="flex items-center mr-6">
        <HeartIcon class="inline h-5 w-auto text-red-500" />
        <span class="ml-1 leading-none">{{ humanLikes }}</span>
      </div>
      <div class="flex items-center mr-6">
        <ArrowPathRoundedSquareIcon class="inline h-5 w-auto" />
        <span class="ml-1 leading-none">{{ humanRetweets }}</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.tcard {
  min-width: 100px;
  max-width: 400px;
}
</style>