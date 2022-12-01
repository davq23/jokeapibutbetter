<template>
    <div>
        <joke-form @submit="onSubmit"></joke-form>
    </div>
</template>

<script lang="ts">
import Config from '@/config/Config';
import type Joke from '@/data/joke';
import { JokeService } from '@/services/joke.service';
import { defineComponent } from 'vue';

export default defineComponent({
    methods: {
        onSubmit(joke: Joke) {
            const jokeService = new JokeService(
                Config.apiUrl,
                localStorage.get('token'),
            );

            jokeService.save(joke);
        },
    },
});
</script>

<style></style>
