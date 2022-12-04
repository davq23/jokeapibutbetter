<template>
    <form @submit="onSubmit">
        <v-container>
            <v-row>
                <v-col>
                    <v-textarea label="Joke text" v-model="text"></v-textarea>
                </v-col>
            </v-row>
            <v-row>
                <v-col>
                    <v-text-field
                        label="Joke description"
                        v-model="description"></v-text-field>
                </v-col>
            </v-row>
            <v-row>
                <v-col>
                    <select v-model="lang" label="Language">
                        <option value="en_US">English</option>
                        <option value="es_ES">Español</option>
                        <option value="fr_FR">Français</option>
                    </select>
                </v-col>
            </v-row>
            <v-row>
                <v-col>
                    <loading-btn type="submit" :loading="loading" text="Save" />
                </v-col>
            </v-row>
        </v-container>
    </form>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import type { PropType } from 'vue';
import type Joke from '@/data/joke';
import LoadingBtn from '@/components/elements/LoadingBtn.vue';

export default defineComponent({
    components: {
        LoadingBtn,
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
        onSubmit(event: SubmitEvent) {
            event.preventDefault();

            this.$emit('submit', this.$data);
        },
    },

    mounted() {
        if (this.joke) {
            this.id = this.joke.id;
            this.text = this.joke.text;
            this.description = this.joke.description;
            this.lang = this.joke.lang;
        }
    },

    props: {
        joke: {
            type: Object as PropType<Joke | undefined>,
            required: false,
            default: {
                id: '',
                author_id: '',
                description: '',
                text: '',
                added_at: null,
                lang: 'en_US',
            } as Joke,
        },

        loading: {
            type: Boolean,
            required: false,
            default: false,
        },
    },

    watch: {
        joke(newValue: Joke | undefined) {
            if (newValue) {
                this.id = newValue.id;
                this.text = newValue.text;
                this.description = newValue.description;
                this.lang = newValue.lang;
            }
        },
    },
});
</script>

<style></style>
