<template>
    <div
        style="background-color: red; height: 100%; width: 100%"
        @drop="handleDrop"
        @dragover="handleDragOver"
        @dragenter="handleDragOver"
        @dragleave="handleDragOver">
        <div>
            <h3>Drop here to upload image</h3>
        </div>
        <slot></slot>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import type { PropType } from 'vue';

export default defineComponent({
    emits: ['error', 'change'],
    methods: {
        handleDragOver(event: DragEvent) {
            event.preventDefault();
        },
        handleDrop(event: DragEvent) {
            event.preventDefault();
            event.stopPropagation();

            let correctFilesUploaded = true;

            if (event.dataTransfer && event.dataTransfer.files.length > 0) {
                for (let i = 0; i < event.dataTransfer.files.length; i++) {
                    console.log(event.dataTransfer.files.item(i)?.type);
                    if (
                        !this.fileTypes.includes(
                            event.dataTransfer.files.item(i)?.type ?? '',
                        )
                    ) {
                        correctFilesUploaded = false;
                        break;
                    }
                }
                if (correctFilesUploaded) {
                    this.$emit('change', event.dataTransfer.files);
                } else {
                    this.$emit('error', 'Invalid file type');
                }
            }
        },
    },
    props: {
        fileTypes: {
            type: Array as PropType<string[]>,
            required: true,
        },
        multiple: {
            type: Boolean,
            required: false,
            default: false,
        },
    },
});
</script>

<style></style>
