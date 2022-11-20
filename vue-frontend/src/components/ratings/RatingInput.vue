<template>
    <div class="stars-outer" @click="setClickStars">
        <div
            :key="rerender"
            class="stars-inner"
            :style="`width: ${
                ((stars ?? 0 / maxStars) * 100).toFixed(2) + '%'
            };`"></div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

interface RatingInput {
    stars: number;
    rerender: number;
}

export default defineComponent({
    data(): RatingInput {
        return {
            stars: 0,
            rerender: 0,
        };
    },

    methods: {
        setClickStars(event: MouseEvent) {
            if (event && event.target instanceof HTMLElement) {
                const clientRect = event.target.getBoundingClientRect();
                const xRelativePosition = event.clientX - clientRect.left;

                const score =
                    this.maxStars * (xRelativePosition / clientRect.width);

                this.stars = Math.round(score * 2) / 2;
            }
        },
    },

    mounted() {
        this.stars = this.modelValue ?? 0;

        this.rerender = 1;
    },

    props: {
        maxStars: {
            type: Number,
            default: 5,
        },

        modelValue: Number,
    },

    emits: ['input:modelValue'],

    watch: {
        stars(newValue: number) {
            this.$emit('input:modelValue', newValue);

            this.rerender = this.rerender === 1 ? 0 : 1;
        },
    },
});
</script>

<style scoped>
.stars-outer {
    display: inline-block;
    position: relative;
    font-family: 'FontAwesome';
}

.stars-outer::before {
    content: '\f006 \f006 \f006 \f006 \f006';
}

.stars-inner {
    position: absolute;
    top: 0;
    left: 0;
    white-space: nowrap;
    overflow: hidden;
    width: 0;
}

.stars-inner::before {
    content: '\f005 \f005 \f005 \f005 \f005';
    color: #f8ce0b;
}
</style>
