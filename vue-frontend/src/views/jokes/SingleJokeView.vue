<template>
    <joke-card v-if="joke !== null" :joke="joke">
        <v-progress-circular
            v-if="ratings === null"
            indeterminate
            color="secondary"></v-progress-circular>
        <v-container fluid v-else-if="ratings.length > 0">
            <h5>Ratings</h5>
            <rating-input
                v-for="(rating, index) in ratings"
                :title="rating.user ? rating.user.username : 'User'"
                v-model.number="rating.stars"
                :disabled="rating.user && rating.user.id !== user.id"
                @update:modelValue="
                    joke ? publishRating($event, joke.id) : null
                "
                :key="index"></rating-input>
        </v-container>
        <v-container fluid v-else>
            <h5>Be the first to rate this joke</h5>
            <rating-input
                v-if="user.id"
                title="Your rating"
                v-model.number="joke.stars"
                @update:modelValue="
                    joke ? publishRating($event, joke.id) : null
                "></rating-input>
            <span v-else>Log in to rank this joke</span>
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
import { useUserStore } from '@/stores/user';

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

        publishRating(stars: number, jokeID: string) {
            if (this.user.id === null) {
                return;
            }
            const ratingService = new RatingService(
                Config.apiUrl,
                localStorage.getItem('token'),
            );

            ratingService
                .rate({
                    id: '',
                    user_id: this.user.id,
                    joke_id: jokeID,
                    stars,
                    comment: '',
                    user: undefined,
                })
                .then((response) => {
                    return response.json();
                })
                .then((jsonResponse: StandardResponse) => {
                    if (jsonResponse.status == 200) {
                        this.ratings = null;
                        this.getRatingsByJoke(jokeID);
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
        const user = useUserStore();

        return {
            alert,
            user,
        };
    },
});
</script>

<style></style>
