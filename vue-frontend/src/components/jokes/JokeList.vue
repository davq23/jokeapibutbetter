<template>
    <div>
        <div v-for="joke in jokes" :key="joke.id">
            <p style="white-space: pre">{{ joke.text }}</p>
        </div>
    </div>
</template>

<script lang="ts">
import type Joke from '@/data/joke';

import { defineComponent } from 'vue';
import type { PropType } from 'vue';

export default defineComponent({
    props: {
        jokes: {
            type: Array as PropType<Joke[]>,
            required: true,
        },
    },

    components: {},
});
</script>

<style></style>
