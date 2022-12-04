<template>
    <div>
        <joke-form
            @submit="onSubmit"
            :loading="loading"
            :joke="joke"></joke-form>
    </div>
</template>

<script lang="ts">
import Config from '@/config/Config';
import type Joke from '@/data/joke';
import { JokeService } from '@/services/joke.service';
import JokeForm from '@/components/jokes/JokeForm.vue';
import { defineComponent } from 'vue';
import type StandardResponse from '@/libs/standard';
import { useAlertStore } from '@/stores/alert';
import type { AlertState } from '@/stores/alert';
import { useUserStore } from '@/stores/user';

interface JokeFormData {
    joke: Joke | undefined;
    loading: boolean;
}

export default defineComponent({
    components: {
        JokeForm,
    },

    data(): JokeFormData {
        return {
            joke: undefined,
            loading: false,
        };
    },

    methods: {
        getJokeByID(jokeID: string) {
            const jokeService = new JokeService(
                Config.apiUrl,
                localStorage.getItem('token'),
            );

            this.loading = true;

            jokeService
                .getJokeByID(jokeID)
                .then((response: Response) => {
                    return response.json();
                })
                .then((response: StandardResponse) => {
                    if (response.status === 200) {
                        this.joke = response.data as Joke;
                    } else {
                        this.alert.showAlert({
                            message: response.message ?? 'Unknown error',
                            messageType: 'error',
                        });
                    }
                })
                .finally(() => {
                    this.loading = false;
                });
        },

        onSubmit(joke: Joke) {
            const jokeService = new JokeService(
                Config.apiUrl,
                localStorage.getItem('token'),
            );

            this.loading = true;

            joke.added_at = null;

            if (this.user.id) {
                joke.author_id = this.user.id;
            }

            jokeService
                .save(joke)
                .then((response) => {
                    return response.json();
                })
                .then((response: StandardResponse) => {
                    const alertState: AlertState = {
                        message: '',
                        messageType: 'info',
                    };

                    if (response.status === 200) {
                        alertState.message = 'Saved successfully';
                    } else {
                        alertState.messageType = 'error';
                        alertState.message =
                            response.message ?? 'Unknown error';
                    }

                    this.alert.showAlert(alertState);
                })
                .finally(() => {
                    this.loading = false;
                });
        },
    },

    mounted() {
        if ('id' in this.$route.params) {
            this.getJokeByID(this.$route.params.id as string);
        }
    },

    setup() {
        const user = useUserStore();
        const alert = useAlertStore();

        return { alert, user };
    },
});
</script>

<style></style>
