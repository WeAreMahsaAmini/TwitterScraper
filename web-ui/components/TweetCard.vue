<script>
import {HeartIcon, ArrowPathRoundedSquareIcon } from "@heroicons/vue/24/outline";

export default {
  props: ["name", "hashtags", "id", "likes", "retweets", "timestamp", "direction"],
  components: { ArrowPathRoundedSquareIcon, HeartIcon },
  computed: {
    showDate() {
      let date = new Date(this.timestamp * 1000).toDateString();
      let firstSpace = date.indexOf(" ");
      let format = date.slice(firstSpace);
      return format;
    },
    showTime() {
      let time = new Date(this.timestamp * 1000);
      let format = time.getHours() + ":" + time.getMinutes();
      return format;
    },
    directionStyle() {
      return {
        direction: this.direction
      }
    }
  },
};
</script>
<template>
  <div class="bg-gray-50 hover:bg-white rounded-xl m-auto p-4 w-80 shadow">
    <div class="flex justify-between mb-3">
      <div>
        <div class="font-bold leading-tight">{{ name }}</div>
        <div class="text-sm text-gray-600 leading-tight">
          @{{ id }} · {{ showDate }} · {{ showTime }}
        </div>
      </div>
    </div>
    <div class="mb-3 text-gray-700" :style="directionStyle"><slot /></div>
    <div class="mb-3 text-sm text-gray-700 p-1">
      <div
        v-for="hashtag in hashtags"
        :key="hashtag"
        class="
          inline-block
          bg-blue-100
          text-blue-800 text-xs
          font-semibold
          mr-2
          px-2.5
          py-0.5
          rounded
        "
      >
        {{ "#" + hashtag }}
      </div>
    </div>
    <div class="flex justify-center mb-3 text-sm leading-tight text-gray-700">
      <div class="mr-2">
        <HeartIcon class="inline h-6 w-6 text-black" />: {{ likes }}
      </div>

      <div>
        <ArrowPathRoundedSquareIcon class="inline h-6 w-6 text-black" />: {{ retweets }}
      </div>
    </div>
    <div
      class="flex justify-center mb-3 text-sm leading-tight text-gray-700"
    ></div>
  </div>
</template>