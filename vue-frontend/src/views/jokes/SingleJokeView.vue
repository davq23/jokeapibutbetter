<template>
    <joke-card v-if="joke !== null" :joke="joke"></joke-card>
    <v-progress-circular
        v-else
        indeterminate
        color="secondary"></v-progress-circular>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import JokeCard from '@/components/jokes/JokeCard.vue';
import type Joke from '@/data/joke';
import Config from '@/config/Config';
import { JokeService } from '@/services/joke.service';
import type StandardResponse from '@/libs/standard';
import { useAlertStore } from '@/stores/alert';

interface SingleJokeView {
    joke: Joke | null;
}

export default defineComponent({
    components: {
        JokeCard,
    },

    data(): SingleJokeView {
        return {
            joke: null,
        };
    },

    methods: {
        getJoke(id: string) {
            const jokeService = new JokeService(
                Config.apiUrl,
                localStorage.getItem('token'),
            );

            jokeService
                .getJokeByID(id)
                .then((response: Response) => {
                    return response.json();
                })
                .then((response: StandardResponse) => {
                    if (response.status === 200) {
                        this.joke = response.data as Joke;
                    } else {
                        this.alert.showAlert({
                            messageType: 'error',
                            message: response.message ?? 'Unknown error',
                        });
                    }
                });
        },
    },

    mounted() {
        if (
            'id' in this.$route.params &&
            !(this.$route.params.id instanceof Array)
        ) {
            this.getJoke(this.$route.params.id);
        }
    },

    setup() {
        const alert = useAlertStore();

        return {
            alert,
        };
    },
});
</script>

<style></style>
