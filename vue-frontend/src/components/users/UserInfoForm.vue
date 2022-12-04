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

    mounted() {
        if (this.user) {
            this.id = this.user.id;
            this.username = this.user.username;
            this.email = this.user.email;
            this.roles = this.user.roles;
        }
    },

    props: {
        user: {
            type: Object as PropType<User | null>,
            required: false,
            default: null,
        },
    },
});
</script>
