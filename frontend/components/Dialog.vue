<template>
    <div class="fixed bg-black bg-opacity-60 left-0 right-0 bottom-0 top-0 z-30 flex justify-center items-center">
        <div @click="handleClose" class="fixed left-0 right-0 bottom-0 top-0"></div>
        <div class="w-11/12 md:w-auto md:min-w-[500px] zoom-in z-10" ref="dialog">
            <Card class="w-full bg-[#09090b]">
                <div class="flex justify-between mb-4 text-lg">
                    <span class="font-bold">{{ props.title }}</span>
                    <button @click="handleClose">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                            class="lucide lucide-x">
                            <path d="M18 6 6 18" />
                            <path d="m6 6 12 12" />
                        </svg>
                    </button>
                </div>
                <slot></slot>
            </Card>
        </div>
    </div>
</template>
<script setup lang='ts'>

const dialog = ref<HTMLElement | null>(null)

const emit = defineEmits(['close'])

const props = defineProps({
    title: {
        type: String,
        default: ''
    },
})

const handleClose = () => {
    dialog.value?.classList.add('zoom-out')
    dialog.value?.addEventListener('animationend', () => {
        dialog.value?.classList.remove('zoom-out')
        emit('close')
    })
}

</script>
<style scoped></style>