<template>
    <v-container>
        <user-info-form
            v-if="user !== null"
            @submit="onSubmit"
            :user="user.getAsUserData()"></user-info-form>
    </v-container>
</template>

<script lang="ts">
import UserInfoForm from '@/components/users/UserInfoForm.vue';
import { defineComponent } from 'vue';
import type { User } from '@/data/user';
import UserService from '@/services/user.service';
import Config from '@/config/Config';
import type StandardResponse from '@/libs/standard';
import { useAlertStore } from '@/stores/alert';
import { useUserStore } from '@/stores/user';

export default defineComponent({
    components: {
        UserInfoForm,
    },

    methods: {
        onSubmit(user: User) {
            const userService = new UserService(
                Config.apiUrl,
                localStorage.getItem('token'),
            );

            userService
                .save(user)
                .then((response: Response) => {
                    return response.json();
                })
                .then((response: StandardResponse) => {
                    if (response.status === 200) {
                        this.$router.replace('/login');
                    } else {
                        this.alert.showAlert({
                            message: response.message ?? 'Unknown error',
                            messageType: 'error',
                        });
                    }
                });
        },
    },

    setup() {
        const alert = useAlertStore();
        const user = useUserStore();

        return { alert, user };
    },
});
</script>

<style></style>
