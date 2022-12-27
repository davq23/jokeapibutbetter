<template>
    <div style="position: relative; background-color: red; height: 10rem">
        <file-drop-area
            :fileTypes="['image/png', 'image/jpeg']"
            @change="handleChangeDropArea"
            @error="$emit('error', $event)">
            <v-progress-circular
                style="position: absolute; top: 50%; left: 50%"
                v-if="uploadProgress !== null"
                v-model="uploadProgress">
            </v-progress-circular>
        </file-drop-area>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import type { PropType } from 'vue';
import FileDropArea from './FileDropArea.vue';
import type StandardResponse from '@/libs/standard';

interface ImageUploaderData {
    uploadProgress: null | string;
}

export default defineComponent({
    components: {
        FileDropArea,
    },
    data(): ImageUploaderData {
        return {
            uploadProgress: null,
        };
    },
    emits: ['upload-done', 'error'],
    methods: {
        handleChangeDropArea(files: FileList) {
            const imageFile = files.item(0);

            if (imageFile) {
                this.getUploadLinkCallback().then((response) => {
                    if (response.status === 200) {
                        this.uploadImage(response.data as string, imageFile);
                    } else {
                        this.$emit(
                            'error',
                            response.message ?? 'Unknown error',
                        );
                    }
                });
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

            xhr.overrideMimeType(file.type);
            xhr.setRequestHeader('Content-Type', file.type);
            xhr.send(file);
        },
    },
    props: {
        getUploadLinkCallback: {
            type: Function as PropType<() => Promise<StandardResponse>>,
            required: true,
        },
    },
});
</script>

<style></style>
