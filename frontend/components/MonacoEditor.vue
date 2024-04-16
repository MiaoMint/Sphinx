<template>
    <div ref="editorContainer"></div>
</template>
<script setup lang='ts'>
import { editor } from 'monaco-editor'
// @ts-ignore
import jsonWorker from 'monaco-editor/esm/vs/language/json/json.worker?worker'
// @ts-ignore
import cssWorker from 'monaco-editor/esm/vs/language/css/css.worker?worker'
// @ts-ignore
import htmlWorker from 'monaco-editor/esm/vs/language/html/html.worker?worker'
// @ts-ignore
import tsWorker from 'monaco-editor/esm/vs/language/typescript/ts.worker?worker'
// @ts-ignore
import EditorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker'


const editorContainer = ref<HTMLDivElement | null>(null)
let monacoEditor: editor.IStandaloneCodeEditor

const props = defineProps<{
    modelValue: string,
    language?: string
}>()

const emit = defineEmits(['update:modelValue'])


// @ts-ignore
self.MonacoEnvironment = {
    getWorker(_: string, label: string) {
        if (label === 'json') {
            return new jsonWorker()
        }
        if (label === 'css' || label === 'scss' || label === 'less') {
            return new cssWorker()
        }
        if (label === 'html' || label === 'handlebars' || label === 'razor') {
            return new htmlWorker()
        }
        if (['typescript', 'javascript'].includes(label)) {
            return new tsWorker()
        }
        return new EditorWorker()
    },
}



onMounted(() => {
    if (!editorContainer.value) {
        return
    }
    monacoEditor = editor.create(editorContainer.value, {
        value: props.modelValue,
        language: props.language,
        theme: 'vs-dark',
    })

    monacoEditor.onDidChangeModelContent(() => {
        console.log(monacoEditor.getValue())
        emit('update:modelValue', monacoEditor.getValue())
    })

})

onUnmounted(() => {
    monacoEditor.dispose()
})

</script>
<style scoped></style>