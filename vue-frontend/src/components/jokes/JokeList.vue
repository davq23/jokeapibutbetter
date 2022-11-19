<template>
    <div>
        <v-card
            style="white-space: pre"
            v-for="joke in jokes"
            :key="joke.id"
            :text="joke.text"
            :subtitle="`Posted  by ${joke.user?.username} at ${formatDate(
                joke.added_at,
            )}`">
        </v-card>
    </div>
</template>

<script lang="ts">
import type Joke from '@/data/joke';

import { formatDate } from '@/libs/convertDates';

import { VCard } from 'vuetify/components';
import { defineComponent } from 'vue';
import type { PropType } from 'vue';

export default defineComponent({
    components: {
        VCard,
    },

    props: {
        jokes: {
            type: Array as PropType<Joke[]>,
            required: true,
        },
    },

    methods: {
        formatDate(datetime: string) {
            return formatDate(datetime);
        },
    },
});
</script>

<style></style>
