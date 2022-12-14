<template>
    <joke-card v-if="joke !== null" :joke="joke">
        <v-progress-circular
            v-if="ratings === null"
            indeterminate
            color="secondary"></v-progress-circular>
        <v-container v-else>
            <h5>Ratings</h5>
            <rating-input
                v-for="(rating, index) in ratings"
                :title="rating.user ? rating.user.username : 'User'"
                v-model.number="rating.stars"
                :key="index"></rating-input>
        </v-container>
    </joke-card>
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
import RatingService from '@/services/rating.service';
import type Rating from '@/data/rating';
import RatingInput from '@/components/ratings/RatingInput.vue';

interface SingleJokeView {
    joke: Joke | null;
    ratings: Rating[] | null;
}

export default defineComponent({
    components: {
        JokeCard,
        RatingInput,
    },

    data(): SingleJokeView {
        return {
            joke: null,
            ratings: null,
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

        getRatingsByJoke(id: string) {
            const ratingService = new RatingService(
                Config.apiUrl,
                localStorage.getItem('token'),
            );

            ratingService
                .getAllByJokeID(id)
                .then((response: Response) => {
                    return response.json();
                })
                .then((response: StandardResponse) => {
                    if (response.status === 200) {
                        this.ratings = response.data as Rating[];
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
            this.getRatingsByJoke(this.$route.params.id);
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
