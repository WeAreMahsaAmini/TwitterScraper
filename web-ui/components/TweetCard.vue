<script>
import { HandThumbUpIcon, ArrowPathIcon } from "@heroicons/vue/24/solid";
export default {
  props: ["name", "hashtags", "id", "likes", "retweets"],
  components: { HandThumbUpIcon, ArrowPathIcon },
};
</script>
<template>
  <div class="bg-gray-50 hover:bg-white rounded-xl m-1 p-4 w-80 shadow">
    <div class="flex justify-between mb-3">
      <div>
        <div class="font-bold leading-tight">{{ name }}</div>
        <div class="text-sm text-gray-600 leading-tight">@{{ id }}</div>
      </div>
    </div>
    <div class="mb-3 text-gray-700" style="direction: rtl"><slot /></div>
    <div class="mb-3 text-sm leading-tight text-gray-700">
      Hashtags:
      <div
        v-for="hashtag in hashtags"
        :key="hashtag"
        class="inline text-blue-700"
      >
        #{{ hashtag }}
      </div>
    </div>
    <div class="flex justify-center mb-3 text-sm leading-tight text-gray-700">
      <div class="mr-2">
        <HandThumbUpIcon class="inline h-6 w-6 text-black" />: {{ likes }}
      </div>

      <div>
        <ArrowPathIcon class="inline h-6 w-6 text-black" />: {{ retweets }}
      </div>
    </div>
    <div
      class="flex justify-center mb-3 text-sm leading-tight text-gray-700"
    ></div>
  </div>
</template>