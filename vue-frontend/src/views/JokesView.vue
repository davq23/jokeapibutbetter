<template>
    <div>
        <div v-if="jokes !== null">
            <joke-list :jokes="jokes" />
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
    </div>
</template>

<script lang="ts">
import Config from '@/config/Config';
import type Joke from '@/data/joke';
import type StandardResponse from '@/libs/standard';
import JokeList from '@/components/jokes/JokeList.vue';
import { JokeService } from '@/services/joke.service';
import { defineComponent } from 'vue';
import type Rating from '@/data/rating';
import { useUserStore } from '@/stores/user';
import RatingService from '@/services/rating.service';

interface JokeViewData {
    language: string | null;
    offset: string | null;
    direction: number;
    jokes: Joke[] | null;
}

export default defineComponent({
    components: {
        JokeList,
    },

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
                Config.apiUrl,
                localStorage.getItem('token'),
            );

            jokeService
                .getJokes(this.offset, this.language, this.direction, null)
                .then((response: Response) => {
                    return response.json();
                })
                .then((jsonResponse: StandardResponse): void => {
                    if (jsonResponse.status === 200) {
                        this.jokes = jsonResponse.data as Joke[];
                    }
                })
                .finally(() => {
                    if (this.jokes !== null && this.user.id !== null) {
                        this.fetchRatings();
                    }
                });
        },

        async assignRatings(ratings: Rating[]) {
            if (this.jokes === null) {
                return;
            }
            const ratingMap = new Map<string, number>(
                ratings.map((rating: Rating) => [rating.joke_id, rating.stars]),
            );

            for (let index = 0; index < this.jokes.length; index++) {
                if (ratingMap.has(this.jokes[index].id)) {
                    this.jokes[index].stars = ratingMap.get(
                        this.jokes[index].id,
                    );
                }
            }
        },

        fetchRatings() {
            if (this.user.id === null) {
                return;
            }

            const ratingService = new RatingService(
                Config.apiUrl,
                localStorage.getItem('token'),
            );

            ratingService
                .getAllByUserID(this.user.id)
                .then((response) => {
                    return response.json();
                })
                .then((jsonResponse: StandardResponse) => {
                    if (jsonResponse.status == 200) {
                        this.assignRatings(jsonResponse.data as Rating[]);
                    }
                });
        },
    },

    mounted() {
        this.jokes = null;

        if (this.$route.query['lang']) {
            this.language = this.$route.query['lang']?.toString();
        }
        if (this.$route.query['offset']) {
            this.offset = this.$route.query['offset']?.toString();
        }

        this.getJokes();
    },

    setup() {
        const user = useUserStore();

        return {
            user,
        };
    },
});
</script>

<style></style>
