<template>
    <div class="h-full">
        <div class="p-8">
            <div class="mb-8 flex justify-between items-center">
                <div class="text-3xl font-bold">
                    Logs
                </div>
            </div>
            <Card class="flex flex-col gap-2">
                <!-- 表格 -->
                <div class="overflow-x-auto">
                    <table class="table-auto w-full">
                        <thead class="text-left">
                            <tr>
                                <th class="px-4 py-2">Domain</th>
                                <th class="px-4 py-2">Path</th>
                                <th class="px-4 py-2">Method</th>
                                <th class="px-4 py-2">Status</th>
                                <th class="px-4 py-2">Latency</th>
                                <th class="px-4 py-2">Time</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="item in data?.data" :key="item.ID">
                                <td class="px-4 py-2">{{ item.domain }}</td>
                                <td class="px-4 py-2">{{ item.path }}</td>
                                <td class="px-4 py-2">{{ item.method }}</td>
                                <td class="px-4 py-2">{{ item.status }}</td>
                                <td class="px-4 py-2">{{ item.latency }}</td>
                                <td class="px-4 py-2">{{ new Date(item.CreatedAt) }}</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </Card>
        </div>
    </div>
</template>
<script setup lang='ts'>
import type { Result, Log } from '~/types';



const { data, pending } = useFetch<Result<Log[]>>("/api/log")


</script>
<style scoped></style>