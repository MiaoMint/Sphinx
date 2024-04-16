<template>
    <div class="h-full">
        <div class="p-8">
            <div class="mb-8 flex justify-between items-center">
                <div class="text-3xl font-bold">
                    Domain
                </div>
                <div class="flex gap-2">
                    <Button @click="handleOpenAddDomainDialog">
                        Add
                    </Button>
                </div>
            </div>
            <div class="flex flex-col gap-4">
                <div v-if="pending">
                    Loading...
                </div>
                <div v-else-if="error">
                    {{ error }}
                </div>
                <Card v-else v-for="(domain, index) in data!.data" :key="index">
                    <div class="text-gray-400 mb-3  text-sm">{{ domain.domain }}</div>
                    <div class="text-lg mb-2">{{ domain.name }}</div>
                    <div class="text-gray-400 mb-1">
                        {{ domain.desc }}
                    </div>
                    <div class="flex justify-between items-center">
                        <div class="flex gap-3">
                            <NuxtLink class="group" :to="{ path: `/domain/${domain.ID}`, }">
                                <span class="text-sm text-gray-400 group-hover:text-white mr-1">API</span>
                                <span class="text-sm text-gray-200 group-hover:text-white">{{ domain.api_count }}</span>
                            </NuxtLink>
                            <div>
                                <span class="text-sm text-gray-400 mr-1">Request</span>
                                <span class="text-sm text-gray-200">4</span>
                            </div>
                        </div>
                        <div class="flex gap-2">
                            <Button @click="handleOpenEditDialog(domain.ID)">Edit</Button>
                            <Button @click="handleDeleteDomain(domain.ID)">Delete</Button>
                        </div>
                    </div>
                </Card>
            </div>
        </div>
        <Dialog v-if="showDialog" @close="handleDialogClose" :title="dialogTitle">
            <form @submit.prevent="dialogHandler()" class="flex flex-col gap-3 ">
                <label for="name">
                    <div class="mb-2 text-sm">Name</div>
                    <input type="text" id="name" placeholder="Name" v-model="name"
                        class="w-full px-3 py-1 text-white bg-transparent border border-white border-opacity-20 rounded-md" />
                </label>
                <label for="domain">
                    <div class="mb-2 text-sm">Doamin</div>
                    <input type="text" id="domain" placeholder="www.example.com" v-model="domain"
                        class="w-full px-3 py-1 text-white bg-transparent border border-white border-opacity-20 rounded-md"
                        required />
                </label>
                <label for="description">
                    <div class="mb-2 text-sm">Description</div>
                    <textarea type="text" id="description" placeholder="Description" v-model="desc"
                        class="w-full px-3 py-1 text-white bg-transparent border border-white border-opacity-20 rounded-md">
                        </textarea>
                </label>
                <div class="flex justify-end gap-2 mt-4">
                    <Button> Save </Button>
                </div>
            </form>
        </Dialog>
    </div>
</template>
<script setup lang='ts'>
import type { Result, Domain } from '@/types'
const { data, pending, error, refresh } = useFetch<Result<Domain[]>>("/api/domain")

const showDialog = ref(false)
const dialogTitle = ref('Add domain')
const name = ref('')
const domain = ref('')
const desc = ref('')
let dialogHandler: Function

const handleOpenAddDomainDialog = () => {
    dialogTitle.value = 'Add domain'
    dialogHandler = handleAddDomain
    showDialog.value = true
}

const handleAddDomain = async () => {
    const res = await $fetch('/api/domain', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: {
            name: name.value,
            domain: domain.value,
            desc: desc.value
        },
    })
    refresh()
    showDialog.value = false
}

const handleDeleteDomain = async (id: number) => {
    const res = await $fetch(`/api/domain/${id}`, {
        method: 'DELETE',
    })
    refresh()
}

const handleOpenEditDialog = async (id: number) => {
    const res = await $fetch<Result<Domain>>(`/api/domain/${id}`)
    const data = res.data
    name.value = data.name
    domain.value = data.domain
    desc.value = data.desc
    dialogHandler = () => handleEditDomain(id)
    dialogTitle.value = 'Edit domain'
    showDialog.value = true
}

const handleEditDomain = async (id: number) => {
    const res = await $fetch(`/api/domain/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: {
            name: name.value,
            domain: domain.value,
            desc: desc.value
        },
    })
    refresh()
    showDialog.value = false
}

const handleDialogClose = () => {
    name.value = ''
    domain.value = ''
    desc.value = ''
    showDialog.value = false
}


</script>
<style scoped></style>