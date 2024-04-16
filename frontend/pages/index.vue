<template>
    <div class="h-full">
        <div class="p-8">
            <div class="text-3xl font-bold mb-8">
                Dashboard
            </div>
            <div class="grid grid-cols-1 md:grid-cols-4 gap-5 mb-8">
                <Card>
                    <div class="text-4xl font-bold mb-2">{{ data?.data.domain_count }}</div>
                    <div class="text-sm">Total Domain</div>
                </Card>
                <Card>
                    <div class="text-4xl font-bold mb-2">{{ data?.data.api_count }}</div>
                    <div class="text-sm">Total Api</div>
                </Card>
                <Card>
                    <div class="text-4xl font-bold mb-2">{{ data?.data.request_count }}</div>
                    <div class="text-sm">Total Request</div>
                </Card>
                <Card>
                    <div class="text-4xl font-bold mb-2">{{ data?.data.success_count }}</div>
                    <div class="text-sm">Total Success</div>
                </Card>
            </div>
            <div class="flex gap-8 flex-col md:flex-row mb-8">
                <Overview />
                <Domain />
            </div>
            <API />
        </div>
    </div>
</template>
<script setup lang='ts'>
import { VisAxis, VisStackedBar, VisXYContainer } from '@unovis/vue'
import API from '~/components/Dashborad/API.vue';
import Domain from '~/components/Dashborad/Domain.vue';
import Overview from '~/components/Dashborad/Overview.vue';
import type { Result } from '~/types';

const { data } = useFetch<Result<{
    domain_count: number
    api_count: number
    request_count: number
    success_count: number
}>>("/api/dashboard/total")




</script>
<style scoped></style>