<template>
    <v-container>
        <v-card>
            <v-card-title>Welcome to the Joke API Explorer</v-card-title>
            <v-card-text>
                <p>
                    This web-app contains a collection of jokes written in
                    Spanish, English and French
                </p>
            </v-card-text>
            <v-card-subtitle
                >Created by David Armando Quintero Granadillo</v-card-subtitle
            >
        </v-card>
    </v-container>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    props: {
        lang: {
            type: String,
            required: false,
            default: 'en',
        },
    },
});
</script>

<style></style>
