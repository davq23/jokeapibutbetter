<template>
    <v-card
        style="white-space: pre"
        :key="joke.id"
        :text="joke.text"
        variant="flat"
        @click="!readonly ? $emit('joke-select', joke.id) : null"
        :prepend-icon="`fib fi-${getFlagClassByLanguage(joke.lang)}`"
        :subtitle="`Posted  by ${joke.user?.username}  ${
            joke.added_at ? `${formatDate(joke.added_at)}` : ''
        }`">
        <slot></slot>
    </v-card>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import type { PropType } from 'vue';
import type Joke from '@/data/joke';
import { getFlagClassByLanguage } from '@/libs/internationalization';
import { formatDate } from '@/libs/convertDates';

export default defineComponent({
    components: {},

    data() {
        return {};
    },

    methods: {
        formatDate(datetime: string) {
            return formatDate(datetime);
        },

        getFlagClassByLanguage(languageCode: string): string {
            return getFlagClassByLanguage(languageCode);
        },
    },

    props: {
        joke: {
            type: Object as PropType<Joke>,
            required: true,
        },
        readonly: {
            type: Boolean,
            required: false,
            default: false,
        },
    },
});
</script>

<style></style>
