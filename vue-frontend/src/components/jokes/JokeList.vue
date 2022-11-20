<template>
    <div>
        <v-card
            style="white-space: pre"
            v-for="joke in jokes"
            :key="joke.id"
            :text="joke.text"
            :prepend-icon="`fib fi-${getFlagClassByLanguage(joke.lang)}`"
            :subtitle="`Posted  by ${joke.user?.username} at ${formatDate(
                joke.added_at,
            )}`">
            <div style="text-align: end; margin-right: 1rem">
                <rating-input
                    v-if="user.id"
                    v-model="joke.stars"
                    @input="publishRating($event, joke.id)"></rating-input>
            </div>
        </v-card>
    </div>
</template>

<script lang="ts">
import type Joke from '@/data/joke';
import RatingInput from '@/components/ratings/RatingInput.vue';
import { formatDate } from '@/libs/convertDates';

import { VCard } from 'vuetify/components';
import { defineComponent } from 'vue';
import type { PropType } from 'vue';
import { getFlagClassByLanguage } from '@/libs/internationalization';
import { useUserStore } from '@/stores/user';
import Config from '@/config/Config';
import RatingService from '@/services/rating.service';
import type StandardResponse from '@/libs/standard';

export default defineComponent({
    components: {
        RatingInput,
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

        getFlagClassByLanguage(languageCode: string): string {
            return getFlagClassByLanguage(languageCode);
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
                })
                .then((response) => {
                    return response.json();
                })
                .then((jsonResponse: StandardResponse) => {
                    if (jsonResponse.status == 200) {
                        console.log('rated');
                    }
                });
        },
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
