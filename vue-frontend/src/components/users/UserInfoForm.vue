<template>
    <v-form @submit="onSubmit">
        <v-container>
            <v-row>
                <v-col>
                    <v-input label="Username" v-model="username" type="text" />
                </v-col>
                <v-col>
                    <v-input label="Email" v-model="email" type="email" />
                </v-col>
            </v-row>
            <v-row>
                <v-col></v-col>
            </v-row>
        </v-container>
    </v-form>
</template>

<script lang="ts">
import type { User } from '@/data/user';
import { defineComponent } from 'vue';
import type { PropType } from 'vue';

export default defineComponent({
    data(): User {
        return {
            id: '',
            username: '',
            email: '',
            roles: ['USER'],
        };
    },

    methods: {
        onSubmit(event: SubmitEvent) {
            event.preventDefault();

            this.$emit('submit', this.$data);
        },
    },

    watch: {
        user(newUser: User | null) {
            console.log(newUser);
            if (newUser) {
                this.id = newUser.id;
                this.username = newUser.username;
                this.email = newUser.email;
                this.roles = newUser.roles;
            }
        },
    },

    props: {
        user: {
            type: Object as PropType<User>,
            required: false,
            default: null,
        },
    },
});
</script>
