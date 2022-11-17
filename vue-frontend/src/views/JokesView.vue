<template>
    <div>
        <div v-if="jokes !== null">
            <joke-list :jokes="jokes" />
        </div>
    </div>
</template>

<script lang="ts">
import type Joke from '@/data/joke';
import type StandardResponse from '@/libs/standard';
import { JokeService } from '@/services/joke.service';
import { defineComponent } from 'vue';

interface JokeViewData {
    language: string | null;
    offset: string | null;
    direction: number;
    jokes: Joke[] | null;
}

export default defineComponent({
    components: {},

    data(): JokeViewData {
        return {
            language: null,
            offset: null,
            direction: 0,
            jokes: null,
        };
    },

    methods: {
        getJokes() {
            const jokeService = new JokeService(
                import.meta.env.VITE_JOKEAPI_URL ?? '/',
                localStorage.getItem('token'),
            );

            jokeService
                .getJokes(this.offset, this.language, this.direction)
                .then((response: Response) => {
                    return response.json();
                })
                .then((jsonResponse: StandardResponse): void => {
                    if (jsonResponse.status === 200) {
                        this.jokes = jsonResponse.data as Joke[];
                    }
                });
        },
    },

    mounted() {
        if (this.$route.query['lang']) {
            this.language = this.$route.query['lang']?.toString();
        }
        if (this.$route.query['offset']) {
            this.offset = this.$route.query['offset']?.toString();
        }

        this.getJokes();
    },
});
</script>

<style></style>
