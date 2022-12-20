<template>
    <v-container>
        <img :src="profilePicDownloadLink" v-if="profilePicDownloadLink" />
        <image-uploader
            :upload-link="profilePicUploadLink"
            @upload-done="getProfilePicDownloadLink"
            @error="
                alert.showAlert({
                    message: $event,
                    messageType: 'error',
                })
            " />
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
import ImageUploader from '@/components/elements/ImageUploader.vue';

interface UserPreferencesData {
    profilePicUploadLink: string | undefined;
    profilePicDownloadLink: string | undefined;
}

export default defineComponent({
    components: {
        UserInfoForm,
        ImageUploader,
    },

    data(): UserPreferencesData {
        return {
            profilePicUploadLink: undefined,
            profilePicDownloadLink: undefined,
        };
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

        getProfilePicUploadLink() {
            const userService = new UserService(
                Config.apiUrl,
                localStorage.getItem('token'),
            );

            userService
                .getProfilePicUploadLink()
                .then((response: Response) => {
                    return response.json();
                })
                .then((response: StandardResponse) => {
                    if (response.status === 200) {
                        this.profilePicUploadLink = response.data as string;
                    } else {
                        this.alert.showAlert({
                            message: response.message ?? 'Unknown error',
                            messageType: 'error',
                        });
                    }
                });
        },

        getProfilePicDownloadLink() {
            const userService = new UserService(
                Config.apiUrl,
                localStorage.getItem('token'),
            );

            userService
                .getProfilePicDownloadLink()
                .then((response: Response) => {
                    return response.json();
                })
                .then((response: StandardResponse) => {
                    if (response.status === 200) {
                        this.profilePicUploadLink = response.data as string;
                    } else if (response.status === 404) {
                        this.profilePicDownloadLink = undefined;
                    } else {
                        this.alert.showAlert({
                            message: response.message ?? 'Unknown error',
                            messageType: 'error',
                        });
                    }
                });
        },
    },

    mounted() {
        this.getProfilePicDownloadLink();
        this.getProfilePicUploadLink();
    },

    setup() {
        const alert = useAlertStore();
        const user = useUserStore();

        return { alert, user };
    },
});
</script>

<style></style>
