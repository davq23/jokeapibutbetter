<template>
    <v-card
        class="grow-on-hover"
        style="white-space: pre"
        :key="joke.id"
        :text="joke.text"
        @click="$emit('joke-select', joke.id)"
        :prepend-icon="`fib fi-${getFlagClassByLanguage(joke.lang)}`"
        :subtitle="`Posted  by ${joke.user?.username}  ${
            joke.added_at ? `${formatDate(joke.added_at)}` : ''
        }`">
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

    data() {},

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
    },
});
</script>

<style></style>
