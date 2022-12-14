<template>
    <div>
        <slot></slot>
        <joke-card
            class="grow-on-hover"
            v-for="joke in jokes"
            :key="joke.id"
            :joke="joke"
            @click="$emit('joke-select', joke.id)">
            <div
                style="margin-right: 1rem; text-align: end"
                v-if="joke.author_id !== user.id">
                <rating-input
                    v-if="user.id"
                    title="Your rating"
                    v-model.number="joke.stars"
                    @update:modelValue="
                        publishRating($event, joke.id)
                    "></rating-input>
                <span v-else>Log in to rank this joke</span>
            </div>
        </joke-card>
    </div>
</template>

<script lang="ts">
import type Joke from '@/data/joke';
import RatingInput from '@/components/ratings/RatingInput.vue';
import { formatDate } from '@/libs/convertDates';

import { defineComponent } from 'vue';
import type { PropType } from 'vue';
import { getFlagClassByLanguage } from '@/libs/internationalization';
import { useUserStore } from '@/stores/user';
import Config from '@/config/Config';
import RatingService from '@/services/rating.service';
import type StandardResponse from '@/libs/standard';
import JokeCard from '@/components/jokes/JokeCard.vue';

export default defineComponent({
    components: {
        RatingInput,
        JokeCard,
    },

    emits: ['joke-select'],

    props: {
        fetchMyRatings: {
            type: Boolean,
            required: false,
            default: false,
        },
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
                    user: undefined,
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
