<template>
    <div>
        <rating-input
            v-if="!readonly"
            :max-stars="5"
            v-model="rating.stars"></rating-input>
        <rating-input
            v-else
            :max-stars="5"
            :value="rating.stars"
            :readonly="true"></rating-input>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import type { PropType } from 'vue';
import type Rating from '@/data/rating';
import RatingInput from './RatingInput.vue';

export default defineComponent({
    components: {
        RatingInput,
    },
    props: {
        rating: {
            type: Object as PropType<Rating>,
            required: true,
        },

        readonly: {
            type: Boolean,
            required: false,
            default: false,
        },
    },
});
</script>

<style></style>
