<template>
    <div
        style="position: relative; background-color: red; height: 10rem"
        @drop="handleDrop"
        @dragover="handleDragOver"
        @dragenter="handleDragOver"
        @dragleave="handleDragOver">
        <v-progress-circular
            style="position: absolute"
            v-if="uploadProgress !== null"
            v-model="uploadProgress">
        </v-progress-circular>
        <div>
            <h3>Drop here to upload image</h3>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

interface ImageUploaderData {
    uploadProgress: null | string;
}

export default defineComponent({
    data(): ImageUploaderData {
        return {
            uploadProgress: null,
        };
    },
    emits: ['upload-done', 'error'],
    methods: {
        handleDragOver(event: DragEvent) {
            event.preventDefault();
        },
        handleDrop(event: DragEvent) {
            event.preventDefault();
            event.stopPropagation();

            if (
                this.uploadLink &&
                event.dataTransfer &&
                event.dataTransfer.files.length > 0
            ) {
                this.uploadImage(this.uploadLink, event.dataTransfer?.files[0]);
            }
        },
        handleProgress(event: ProgressEvent) {
            this.uploadProgress = ((event.loaded / event.total) * 100).toFixed(
                2,
            );
        },
        uploadImage(uploadLink: string, file: File) {
            var xhr = new XMLHttpRequest();

            xhr.open('PUT', uploadLink, true);

            xhr.onload = () => {
                if (xhr.status === 200) {
                    this.$emit('upload-done');
                } else {
                    this.$emit('error', xhr.responseText);
                }
            };

            xhr.onerror = () => {
                this.$emit('error', 'Unknown error');
            };

            xhr.onprogress = this.handleProgress;

            xhr.withCredentials = true;
            xhr.overrideMimeType(file.type);
            xhr.setRequestHeader('Content-Type', file.type);
            xhr.send(file);
        },
    },
    props: {
        uploadLink: {
            type: String,
            required: false,
        },
    },
});
</script>

<style></style>
