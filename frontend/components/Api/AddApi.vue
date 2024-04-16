<template>
    <div class="h-full">
        <div class="p-8">
            <div class="mb-8 flex justify-between items-center">
                <ApiTitle title="Add Api"@back="emit('page', 'ApiList')" />
                <div class="flex gap-2">
                    <Button @click="handleAddApi">
                        Confirm
                    </Button>
                </div>
            </div>
            <div class="flex flex-col md:flex-row gap-6">
                <form @submit.prevent="handleAddApi" class="flex flex-col gap-3 flex-[2]" ref="form">
                    <label for="name">
                        <div class="mb-2 text-sm">Name</div>
                        <input type="text" id="name" placeholder="Name" v-model="name"
                            class="w-full px-3 py-1 text-white bg-transparent border border-white border-opacity-20 rounded-md" />
                    </label>
                    <label for="path">
                        <div class="mb-2 text-sm">Path</div>
                        <input type="text" id="path" placeholder="/v1/api" v-model="path"
                            class="w-full px-3 py-1 text-white bg-transparent border border-white border-opacity-20 rounded-md"
                            required />
                    </label>
                    <label for="method">
                        <div class="mb-2 text-sm">Method</div>
                        <select id="method" v-model="method"
                            class="w-full px-3 py-1 text-white bg-transparent border border-white border-opacity-20 rounded-md"
                            required>
                            <option class="bg-black" value="GET">GET</option>
                            <option class="bg-black" value="POST">POST</option>
                            <option class="bg-black" value="PUT">PUT</option>
                            <option class="bg-black" value="DELETE">DELETE</option>
                        </select>
                    </label>
                    <label for="handleMode">
                        <div class="mb-2 text-sm">Handle Mode</div>
                        <select id="handleMode" v-model="handleMode"
                            class="w-full px-3 py-1 text-white bg-transparent border border-white border-opacity-20 rounded-md"
                            required>
                            <option class="bg-black" value="ReplaceBody">Replace Body</option>
                            <option class="bg-black" value="ModifyBody">Modify Body</option>
                            <option class="bg-black" value="JavaScript">JavaScript</option>
                        </select>
                    </label>
                </form>
                <div class="w-full  flex-[5]">
                    <div class="mb-2">{{ handleMode }}</div>
                    <MonacoEditor v-model="editorValue" language="json" class="w-full min-h-screen" />
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang='ts'>
import { HandleMode, type Domain, type Result } from '~/types';
const domainID = useRoute().params.id
const emit = defineEmits(['page'])
const form = ref<HTMLFormElement | null>(null)
const name = ref('')
const path = ref('')
const method = ref('GET')
const handleMode = ref<HandleMode>(HandleMode.ReplaceBody)
const editorValue = ref('')


const handleAddApi = async () => {
    if (!form.value?.checkValidity()) {
        form.value?.reportValidity()
        return
    }
    const res = await $fetch(`/api/domain/${domainID}/api`, {
        method: 'POST',
        body: {
            name: name.value,
            path: path.value,
            method: method.value,
            handle_mode: handleMode.value,
            body: editorValue.value,
            domainId: domainID
        },
    })
    console.log(res);
    emit('page', 'ApiList')
}


</script>
<style scoped></style>