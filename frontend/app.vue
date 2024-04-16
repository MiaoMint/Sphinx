<template>
  <div class="h-screen bg-[#09090b] text-white flex flex-col md:flex-row">
    <Sidebar :class="{
      'hidden': !isOpenSidebar,
    }" @close="isOpenSidebar = false" />
    <hr class="border-white border-opacity-20">
    <div class="w-full overflow-auto ">
      <div class="bg min-h-screen">
        <div class="block md:hidden">
          <div class="h-16 px-10 flex items-center justify-between">
            <h1 class="text-2xl font-bold">Sphinx</h1>
            <button @click="isOpenSidebar = true">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                class="lucide lucide-menu">
                <line x1="4" x2="20" y1="12" y2="12" />
                <line x1="4" x2="20" y1="6" y2="6" />
                <line x1="4" x2="20" y1="18" y2="18" />
              </svg>
            </button>
          </div>
          <hr class="border-white border-opacity-20">
        </div>
        <div class="max-w-[1400px] mx-auto">
          <NuxtPage />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const isOpenSidebar = ref(false)

</script>

<style>
.bg {
  background: linear-gradient(215deg, #5df2581c, transparent 40%);
}
</style>