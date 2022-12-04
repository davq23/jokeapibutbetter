<template>
    <v-container>
        <div v-if="jokes !== null">
            <joke-list :jokes="jokes" @joke-select="onJokeSelect">
                <v-card
                    style="margin-botton: 2rem"
                    title="Add Joke"
                    append-icon="mdi mdi-plus"
                    @click="$router.push('/jokes/new')"></v-card>
            </joke-list>
        </div>
        <div
            v-else
            style="
                display: flex;
                width: 100%;
                height: 100%;
                justify-content: center;
                align-items: center;
            ">
            <v-progress-circular
                indeterminate
                color="secondary"></v-progress-circular>
        </div>
    </v-container>
</template>

<script lang="ts">
import Config from '@/config/Config';
import type Joke from '@/data/joke';
import { JokeService } from '@/services/joke.service';
import { useAlertStore } from '@/stores/alert';
import { useUserStore } from '@/stores/user';
import { defineComponent } from 'vue';
import JokeList from '@/components/jokes/JokeList.vue';

interface MyJokesData {
    jokes: Joke[] | null;
    offset: string | null;
    language: string | null;
    direction: number;
}

export default defineComponent({
    components: {
        JokeList,
    },

    data(): MyJokesData {
        return {
            jokes: null,
            offset: null,
            language: null,
            direction: 0,
        };
    },

    methods: {
        getMyJokes() {
            const jokeService = new JokeService(
                Config.apiUrl,
                localStorage.getItem('token'),
            );

            jokeService
                .getJokes(
                    this.offset,
                    this.language,
                    this.direction,
                    this.user.id,
                )
                .then((response) => {
                    return response.json();
                })
                .then((response) => {
                    if (response.status == 200) {
                        this.jokes = response.data as Joke[];
                    } else {
                        this.alert.showAlert({
                            message: response.message,
                            messageType: 'error',
                        });
                    }
                });
        },

        onJokeSelect(jokeID: string) {
            this.$router.push(`/jokes/${jokeID}/edit`);
        },
    },

    mounted() {
        this.getMyJokes();
    },

    setup() {
        const user = useUserStore();
        const alert = useAlertStore();

        return {
            user,
            alert,
        };
    },
});
</script>

<style></style>
