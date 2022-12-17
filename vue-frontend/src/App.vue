<template>
    <v-layout @change="alert.message = ''" @click="alert.message = ''">
        <template v-if="user.authLoaded">
            <v-app-bar theme="dark">
                <v-toolbar-title
                    @click="$router.push('/')"
                    class="cursor-pointer">
                    <i class="fa-solid fa-masks-theater"></i> Joke API Explorer
                </v-toolbar-title>
                <v-btn
                    variant="text"
                    icon="mdi-dots-vertical"
                    dark
                    @click="drawer = !drawer">
                </v-btn>
            </v-app-bar>
            <v-navigation-drawer
                theme="dark"
                expand-on-hover
                rail
                :location="deviceWidth > 800 ? 'left' : 'bottom'"
                temporary
                v-model="drawer">
                <v-list v-if="user.id && user.email && user.username">
                    <v-list-item
                        prepend-icon="mdi-account"
                        :subtitle="user.email"
                        :title="user.username"
                        :to="'/users/preferences'"></v-list-item>
                </v-list>
                <v-list nav v-if="user.authLoaded">
                    <v-list-item
                        prepend-icon="mdi-home"
                        title="Home"
                        router
                        :to="'/'"></v-list-item>
                    <v-list-item
                        prepend-icon="mdi-script-text"
                        title="All Jokes"
                        router
                        :to="'/jokes'"></v-list-item>
                    <v-list-item
                        v-if="user.id && user.roles.includes('ADMIN')"
                        prepend-icon="mdi-star"
                        title="My Jokes"
                        router
                        :to="'/jokes/mine'"></v-list-item>
                    <v-list-item
                        v-if="user.id === null"
                        prepend-icon="mdi-door"
                        title="Login"
                        router
                        :to="'/login'"></v-list-item>
                    <v-list-item
                        v-if="user.id !== null"
                        prepend-icon="mdi-door-open"
                        title="Logout"
                        @click="logout()"></v-list-item>
                </v-list>
            </v-navigation-drawer>
            <v-main>
                <v-alert
                    prominent
                    v-if="alert.message"
                    :type="alert.messageType"
                    >{{ alert.message }}</v-alert
                >
                <RouterView />
            </v-main>
        </template>
        <loading-view v-else></loading-view>
    </v-layout>
</template>

<script lang="ts">
import LoadingView from './views/auth/LoadingView.vue';
import { defineComponent } from 'vue';
import { RouterView } from 'vue-router';
import {
    VAppBar,
    VBtn,
    VLayout,
    VList,
    VListItem,
    VMain,
    VNavigationDrawer,
} from 'vuetify/components';
import { useUserStore } from './stores/user';
import { useAlertStore } from './stores/alert';

export default defineComponent({
    components: {
        LoadingView,
        RouterView,
        VAppBar,
        VBtn,
        VLayout,
        VList,
        VListItem,
        VMain,
        VNavigationDrawer,
    },

    created() {
        this.deviceWidth = window.innerWidth;
    },

    data() {
        return {
            drawer: false,
            deviceWidth: 0,
        };
    },

    methods: {
        onRedirect(routeName: string) {
            console.log(routeName);
        },

        logout() {
            this.user.logout().then(() => {
                this.$router.replace({ name: 'login' });
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

<style scoped>
.cursor-pointer:hover {
    cursor: pointer;
}
</style>
