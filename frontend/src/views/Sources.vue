<script setup>
import firefoxLogo from "../assets/images/Firefox_logo_2019.svg"
import chromeLogo from "../assets/images/Google_Chrome_icon_(September_2014).svg"
import hnLogo from "../assets/images/hacker-news-square-custom.svg"
import githubLogo from "../assets/images/github.svg"
import { BrowserOpenURL } from '../../wailsjs/runtime'
</script>

<script>
export default {
  name: 'Sources',
  components: {},

}
</script>

<template>
  <div class="pr-1">
    <h1 class="text-xl">Sources</h1>
    <p class="text-sm mt-1">Sources are remote repositories where Papyri will look for bookmarks. Sources are safe to import more than once, Papyri will ensure there are no duplicates.</p>

    <div class="mt-2 box-border rounded-lg border p-2 bg-gradient-to-r from-violet-500 to-fuchsia-500">
      <h2 class="text-lg font-black">Manual</h2>
      <p class="text-sm">Some sources prohibit or do not provide a path for programmatic synchronization. Each logo contains a link for exporting data from these sources.</p>
      <h2 class="text mt-1 font-black">Supported</h2>
      <div class="flex items-start">
        <a target="_blank" rel="noreferrer noopener" href="https://support.mozilla.org/en-US/kb/export-firefox-bookmarks-to-backup-or-transfer">
          <img class="mr-2 mt-2" width="25" :src="firefoxLogo" alt="Mozilla Firefox logo" />
        </a>
        <a target="_blank" rel="noreferrer noopener" href="https://support.google.com/chrome/answer/96816?hl=en">
          <img class="mr-2 mt-2" width="25" :src="chromeLogo" alt="Google Chrome logo" />
        </a>
      </div>
    </div>

    <div class="mt-2 box-border rounded-lg border p-2 bg-gradient-to-r from-violet-500 to-fuchsia-500">
      <h2 class="text-lg font-black">Automated</h2>
      <p class="text-sm">Papyri supports automated imports from some sources.</p>
      <h2 class="text mt-1 font-black">Supported</h2>
      <div class="flex items-start">
        <img class="mr-2 mt-2" width="25" :src="hnLogo" alt="HackerNews favorites" />
        <img class="mr-2 mt-2" width="25" :src="githubLogo" alt="GitHub starred" />
      </div>
    </div>
  </div>
</template>

<style>

</style>
