<template>
    <div class="h-full md:w-[300px] flex-shrink-0 fixed w-full md:relative md:block z-10">
        <div
            class="p-8 md:border-r h-full border-white border-opacity-20 flex flex-col justify-between bg-black md:bg-transparent ">
            <div>
                <div class="text-5xl font-extrabold mb-12 flex items-center justify-between">
                    <div>
                        <span>Sphinx</span>
                        <span class="text-green-400">.</span>
                    </div>
                    <button @click="emit('close')" class="block md:hidden">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                            class="lucide lucide-x">
                            <path d="M18 6 6 18" />
                            <path d="m6 6 12 12" />
                        </svg>
                    </button>
                </div>
                <div class="space-y-4 text-xl">
                    <NuxtLink :to="route.path" v-for="(route, key) in routes" :key="key"
                        :active-class="'!text-opacity-100 font-bold'"
                        class="flex items center space-x-2 text-white text-opacity-55 hover:text-opacity-100 transition-all ">
                        <div>{{ route.name }}</div>
                    </NuxtLink>
                </div>
            </div>

            <div class="flex justify-between">
                <div class="text-white text-opacity-55">
                    v0.0.1
                </div>
                <div class="text-white text-opacity-55">
                    <a target="_blank" href='https://github.com/MiaoMint/Sphinx'>GitHub</a>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang='ts'>

const emit = defineEmits(['close'])

const routes = [
    {
        name: 'Dashboard',
        path: '/'
    },
    {
        name: 'Domain',
        path: '/domain'
    },
    {
        name: 'Logs',
        path: '/log'
    },
    {
        name: 'Settings',
        path: '/settings'
    },
]


watch(() => useRoute().path, () => {
    emit('close')
})

</script>
<style scoped></style>