<template>
    <Card class="w-full flex-[5]">
        <div class="text-2xl mb-8">
            Overview
        </div>
        <VisXYContainer v-if="!pending" height="350px" :margin="{ left: 20, right: 20 }" :data="data?.data">
            <VisLine :x="x" :y="y" color="#41b883" />
            <VisAxis type="x" :grid-line="false" :tick-line="false"
                :tick-format="(index: number) => data?.data[index]?.date" color="#41b883" />
            <VisAxis type="y" :grid-line="false" :tick-line="false" :domain-line="false" color="#41b883" />
        </VisXYContainer>
    </Card>
</template>
<script setup lang='ts'>
import { VisAxis, VisXYContainer, VisLine } from '@unovis/vue'
import type { Result } from '~/types';

type Data = {
    date: string
    count: number
    index: number
}

const { data, pending } = useFetch<Result<Data[]>>("/api/dashboard/overview")


const x = (d: Data) => d.index
const y = (d: Data) => d.count
</script>
<style scoped></style>