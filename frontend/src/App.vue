<script setup lang="ts">
  import { ref } from 'vue'
  import {EventsOn} from "@/wailsjs/runtime";
  import AboutModal from "@/components/AboutModal.vue";

  let showAboutModal = ref(false);

  EventsOn("toggle-about", (status: boolean) => {
    /*
      This event allows external opening of the AboutModal component from the applications Help menu
    */
    showAboutModal.value = status
  })
</script>

<template>
  <div class="flex">
    <AboutModal :show="showAboutModal" @state-change="(state: Boolean) => this.showAboutModal = state" />
    <!--  Main content  -->
    <div class="flex-none h-screen w-48 mr-1 bg-white dark:bg-zinc-900 shadow-xl">
      <div class="m-2 text-2xl subpixel-antialiased font-bold text-center">
        Papyri
        <span
            class="text-xs font-bold subpixel-antialiased inline-block py-1 px-1 align-middle uppercase rounded text-indigo-600 bg-indigo-200 uppercase last:mr-0 mr-1">
          dev
        </span>
      </div>
      <router-link to="/">
        <div class="p-2 hover:bg-zinc-800 w-full flex items-start" role="link">
          <vue-feather class="mr-1" size="20" type="bookmark" />
          <span class="">Bookmarks</span>
        </div>
      </router-link>
      <router-link to="/sources">
        <div class="p-2 hover:bg-zinc-800 w-full flex items-start" role="link">
          <vue-feather class="mr-1" size="20" type="database" />
          <span>Sources</span>
        </div>
      </router-link>
      <router-link to="/settings">
        <div class="p-2 hover:bg-zinc-800 w-full flex items-start" role="link">
          <vue-feather class="mr-1" size="20" type="settings" />
          <span class="">Settings</span>
        </div>
      </router-link>
    </div>
    <div class="flex-auto m-1">
      <router-view />
    </div>
  </div>
</template>

<style>

</style>
