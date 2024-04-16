<template>
    <component :is="component" @page="handlePage" />
</template>
<script setup lang='ts'>
import AddApi from '~/components/Api/AddApi.vue';
import ApiList from '~/components/Api/ApiList.vue';
const component = shallowRef<Component>(ApiList)

const handlePage = (page: string) => {
    switch (page) {
        case 'AddApi':
            component.value = AddApi
            break
        default:
            component.value = ApiList
            break
    }
}

</script>
<style scoped></style>