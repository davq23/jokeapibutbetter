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
                    <loading-btn
                        :loading="submitting"
                        type="submit"
                        text="Login"></loading-btn>
                </v-col>
            </v-row>
        </v-container>
    </v-form>
</template>

<script lang="ts">
import { useUserStore } from '@/stores/user';
import { defineComponent } from 'vue';
import { VForm, VContainer, VTextField } from 'vuetify/components';
import LoadingBtn from '@/components/elements/LoadingBtn.vue';
import { useAlertStore } from '@/stores/alert';

interface LoginFormData {
    usernameOrEmail: string;
    password: string;
    showPassword: boolean;
    submitting: boolean;
}

export default defineComponent({
    components: {
        LoadingBtn,
        VContainer,
        VForm,
        VTextField,
    },

    data(): LoginFormData {
        return {
            usernameOrEmail: '',
            password: '',
            showPassword: false,
            submitting: false,
        };
    },

    methods: {
        onSubmit(event: Event) {
            event.preventDefault();

            this.submitting = true;

            this.user
                .login(this.usernameOrEmail, this.password)
                .then((response) => {
                    if (response.status === 200) {
                        this.$emit('redirect', 'jokes');
                    } else {
                        this.alert.showAlert({
                            message: response.message ?? 'ERROR',
                            messageType: 'error',
                        });
                    }
                })
                .finally(() => {
                    this.submitting = false;
                });
        },
    },

    setup() {
        const user = useUserStore();
        const alert = useAlertStore();

        return { user, alert };
    },
});
</script>
