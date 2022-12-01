<template>
    <form @submit="onSubmit">
        <v-container>
            <v-row>
                <v-col>
                    <v-text-field
                        label="Joke description"
                        v-model="description"></v-text-field>
                </v-col>
            </v-row>
            <v-row>
                <v-col>
                    <v-textarea label="Joke text" v-model="text"></v-textarea>
                </v-col>
            </v-row>
            <v-row>
                <v-col>
                    <v-select v-model="lang" label="Language">
                        <option value="en_US">English</option>
                        <option value="es_ES">Español</option>
                        <option value="fr_FR">Français</option>
                    </v-select>
                </v-col>
            </v-row>
        </v-container>
    </form>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import type { PropType } from 'vue';
import type Joke from '@/data/joke';

export default defineComponent({
    props: {
        joke: {
            type: () => Object as PropType<Joke>,
            required: false,
            default: {
                id: '',
                author_id: '',
                description: '',
                text: '',
                added_at: '',
                lang: 'en_US',
            },
        },
    },

    data(): Joke {
        return {
            id: '',
            text: '',
            author_id: '',
            description: '',
            lang: 'en_US',
            added_at: '',
            user: undefined,
            stars: undefined,
        };
    },

    emits: ['submit'],

    methods: {
        onSubmit() {
            this.$emit('submit', this.$data);
        },
    },

    mounted() {
        this.$data = this.joke as unknown as Joke;
    },
});
</script>

<style></style>
