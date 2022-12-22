<template>
    <div>
        <p>{{ title }}</p>
        <div class="stars-outer" @click="setClickStars" ref="starsOuter">
            <div
                class="stars-inner fa"
                :style="`width: ${
                    ((modelValue !== undefined ? modelValue : 0) / maxStars) *
                    100
                }%;`"></div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';

export default defineComponent({
    emits: ['update:modelValue'],
    methods: {
        setClickStars(event: MouseEvent) {
            event.stopPropagation();

            if (
                !this.disabled &&
                !this.readonly &&
                this.starsOuter !== null &&
                event &&
                event.target instanceof HTMLElement
            ) {
                const clientRect = this.starsOuter.getBoundingClientRect();
                const xRelativePosition = event.clientX - clientRect.left;

                const score =
                    this.maxStars * (xRelativePosition / clientRect.width);
                this.$emit('update:modelValue', Math.round(score * 2) / 2);
            }
        },
    },

    props: {
        maxStars: {
            type: Number,
            default: 5,
        },

        title: {
            type: String,
            default: '',
        },

        disabled: {
            type: Boolean,
            default: false,
        },

        modelValue: {
            type: [Number, Number as () => number | undefined],
            required: true,
        },

        loading: {
            type: Boolean,
            required: false,
            default: false,
        },

        readonly: {
            type: Boolean,
            required: false,
            default: false,
        },
    },

    setup() {
        const starsOuter = ref<HTMLDivElement | null>(null);

        return {
            starsOuter,
        };
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
    top: 15%;
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
