<template>
    <v-btn :type="type" :color="color">
        <v-progress-circular v-if="loading" indeterminate></v-progress-circular>
        {{ text }}
    </v-btn>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    props: {
        color: {
            type: String,
            required: false,
            default: 'primary',
        },
        loading: {
            type: Boolean,
            required: false,
            default: false,
        },
        type: {
            type: String,
            required: false,
            default: 'button',
        },
        text: {
            type: String,
            required: false,
            default: '',
        },
    },
});
</script>

<style></style>
