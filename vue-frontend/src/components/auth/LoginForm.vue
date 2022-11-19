<template>
    <v-form @submit="onSubmit">
        <v-container>
            <v-row>
                <v-col>
                    <v-text-field
                        label="Username or email"
                        v-model="usernameOrEmail"
                        required></v-text-field>
                </v-col>
            </v-row>
            <v-row>
                <v-col>
                    <v-text-field
                        label="Password"
                        :type="showPassword ? 'text' : 'password'"
                        :append-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
                        @click:append="showPassword = !showPassword"
                        v-model="password"></v-text-field>
                </v-col>
            </v-row>
            <v-row>
                <v-col>
                    <v-btn type="submit">Enter</v-btn>
                </v-col>
            </v-row>
        </v-container>
    </v-form>
</template>

<script lang="ts">
import { useUserStore } from '@/stores/user';
import { defineComponent } from 'vue';
import { VForm, VContainer, VTextField } from 'vuetify/components';

interface LoginFormData {
    usernameOrEmail: string;
    password: string;
    showPassword: boolean;
}

export default defineComponent({
    components: {
        VContainer,
        VForm,
        VTextField,
    },

    data(): LoginFormData {
        return {
            usernameOrEmail: '',
            password: '',
            showPassword: false,
        };
    },

    methods: {
        onSubmit(event: Event) {
            event.preventDefault();
            this.user.login(this.usernameOrEmail, this.password);
        },
    },

    setup() {
        const user = useUserStore();

        return { user };
    },
});
</script>
