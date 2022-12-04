<template>
    <v-select :items="languages" v-model="modelValueLocal"></v-select>
</template>

<script lang="ts">
import Config from '@/config/Config';
import type StandardResponse from '@/libs/standard';
import ConfigService from '@/services/config.service';
import { defineComponent } from 'vue';

interface LanguageSelectData {
    languages: string[];
    modelValueLocal: string;
}

export default defineComponent({
    data(): LanguageSelectData {
        return {
            languages: [],
            modelValueLocal: '',
        };
    },
    methods: {
        getLanguages() {
            const configService = new ConfigService(
                Config.apiUrl,
                localStorage.getItem('token'),
            );

            configService
                .getLanguages()
                .then((response: Response) => {
                    return response.json();
                })
                .then((response: StandardResponse) => {
                    if (response.status === 200) {
                        this.languages = response.data as string[];
                    }
                });
        },
    },
    mounted() {
        this.getLanguages();

        this.modelValueLocal = this.modelValue;
    },
    props: {
        modelValue: {
            type: String,
            required: true,
        },
    },
    watch: {
        modelValueLocal(newValue: string) {
            this.$emit('input', newValue);
        },
    },
});
</script>

<style></style>
