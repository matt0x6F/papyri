<script setup lang="ts">
  const props = defineProps({
    // state
    show: Boolean
  })

  const emit = defineEmits<{
    // tells the parent component to close the modal
    (e: 'stateChange', status: Boolean): void
  }>()
</script>

<template>
  <div v-if="show"
       class="modal fade fixed top-0 p-10 left-0 w-full h-full outline-none overflow-x-hidden overflow-y-auto z-20"
       id="aboutModal" tabindex="-1" aria-labelledby="aboutModalLabel" aria-hidden="true">
    <div class="modal-dialog modal-fullscreen relative w-auto pointer-events-none">
      <div
          class="modal-content border-2 dark:border-zinc-800 shadow-lg relative flex flex-col w-full pointer-events-auto bg-white dark:bg-neutral-900 bg-clip-padding rounded-lg outline-none text-current">
        <div
            class="modal-header flex flex-shrink-0 items-center justify-between p-4 border-b dark:border-neutral-800 border-gray-200 rounded-t-md">
          <h5 class="text-xl font-medium leading-normal text-gray-800 dark:text-violet-200 " id="aboutModalLabel">
            About Papyri
          </h5>
          <button type="button"
                  class="btn-close box-content w-4 h-4 p-1 text-black border-none rounded-none opacity-50 focus:shadow-none focus:outline-none focus:opacity-100 hover:text-black hover:opacity-75 hover:no-underline"
                  data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div class="modal-body relative p-4">
          <p>Papyri is a bookmark manager that keeps your data local while syncing bookmarks from remote sources.</p>

          <p class="text-xs pt-2">This is a development release. Please file bugs on our <a
              href="https://github.com/mattouille/papyri">bug tracker</a>.</p>
        </div>
        <div
            class="modal-footer flex flex-shrink-0 flex-wrap items-center justify-end p-4 border-t dark:border-neutral-800 border-gray-200 rounded-b-md">
          <button type="button"
                  class="inline-block px-6 py-2.5
                      bg-purple-600
                      text-white
                      hover:bg-purple-700
                      focus:bg-purple-700
                      active:bg-purple-800

                      dark:bg-violet-300
                      dark:text-violet-900
                      dark:hover:bg-violet-400
                      dark:focus:bg-violet-400
                      dark:active:bg-violet-500

                      font-medium text-xs leading-tight uppercase rounded shadow-md hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0  active:shadow-lg transition duration-150 ease-in-out"
                  @click="$emit('stateChange', false)">
            Close
          </button>
        </div>
      </div>
    </div>
  </div>
</template>