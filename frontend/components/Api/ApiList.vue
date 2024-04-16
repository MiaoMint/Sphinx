<template>
    <div class="h-full">
        <div class="p-8">
            <div class="mb-8 flex justify-between items-center">
                <ApiTitle @back="useRouter().go(-1)" />
                <div class="flex gap-2">
                    <Button @click="emit('page', 'AddApi')">
                        Add
                    </Button>
                </div>
            </div>
            <div>
                <div v-if="pending">
                    Loading...
                </div>
                <div v-else-if="error">
                    {{ error }}
                </div>
                <div class="flex flex-col gap-4" v-else>
                    <div v-for="(api, index) in data!.data" :key="index">
                        <Card>
                            <div class="flex items-center gap-6">
                                <div class="w-32 flex justify-center items-center flex-col">
                                    <div class="text-3xl  mb-2">
                                        {{ api.method }}
                                    </div>
                                    <div class="text-gray-400 text-sm">
                                        Method
                                    </div>
                                </div>
                                <div class="w-full">
                                    <div class="text-lg mb-2">{{ api.name }}</div>
                                    <div class="text-gray-400 mb-1">
                                        {{ api.path }}
                                    </div>
                                    <div class="flex justify-between items-center">
                                        <div class="flex gap-3">
                                            <div>
                                                <span class="text-sm text-gray-400 mr-1">Handle Mode</span>
                                                <span class="text-sm text-gray-200">{{ api.handle_mode }}</span>
                                            </div>
                                        </div>
                                        <div class="flex gap-2">
                                            <!-- <Button @click="handleOpenEditDialog(api.ID)">Edit</Button> -->
                                            <Button @click="handleDeleteApi(api.ID)">Delete</Button>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </Card>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang='ts'>
import type { API, Result } from '~/types';

const domainID = useRoute().params.id
const { data, pending, error, refresh } = useFetch<Result<API[]>>(`/api/domain/${domainID}/api`)

const emit = defineEmits(['page'])

onMounted(() => {
    refresh()
})

const handleDeleteApi = (id: number) => {
    if (confirm('Are you sure you want to delete this API?')) {
        fetch(`/api/domain/${domainID}/api/${id}`, {
            method: 'DELETE'
        }).then(() => {
            refresh()
        })
    }
}


</script>
<style scoped></style>