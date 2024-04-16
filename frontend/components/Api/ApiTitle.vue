<template>
    <div class="flex items-center gap-6">
        <button class="text-gray-500 hover:text-white" @click="emit('back')">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                class="lucide lucide-arrow-left">
                <path d="m12 19-7-7 7-7" />
                <path d="M19 12H5" />
            </svg>
        </button>
        <div>
            <div class="text-3xl font-bold mb-2">
                {{ title ?? "Api" }}
            </div>
            <div class="text-sm text-gray-400">
                {{ data?.data.name }} - {{ data?.data.domain || 'loading...' }}
            </div>
        </div>
    </div>
</template>
<script setup lang='ts'>
import type { Domain, Result } from '~/types';
const domainID = useRoute().params.id
const { data, pending, error, refresh } = useFetch<Result<Domain>>(`/api/domain/${domainID}`)
const { title } = defineProps<{
    title?: string
}>()

const emit = defineEmits(['back'])

</script>
<style scoped></style>