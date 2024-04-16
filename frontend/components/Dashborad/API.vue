<template>
    <Card class="w-full flex-[4]">
        <div class="text-2xl mb-8">
            Api
        </div>
        <div class="h-full flex flex-col gap-4">
            <div class="flex justify-between gap-2" v-for="item in data?.data">
                <div class="w-40 flex-shrink-0 text-nowrap overflow-hidden text-ellipsis text-gray-300">{{
                `${item.domain}${item.path}` }}</div>
                <div class="flex justify-end w-full">
                    <div>
                        {{ item.count }}
                    </div>
                    <div class="mx-2 border border-white border-opacity-30"></div>
                    <div class="w-full flex-[6] max-w-44">
                        <div class="bg-[#41b883] h-6 rounded-md"
                            :style="{ width: `calc(100% * ${item.count} / ${data?.data[0].count})` }">
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </Card>
</template>
<script setup lang='ts'>
import type { Result } from '~/types';

type Data = {
    domain: string
    method: string
    path: string
    count: number
    index: number
}

const { data } = useFetch<Result<Data[]>>("/api/dashboard/api")



</script>
<style scoped></style>