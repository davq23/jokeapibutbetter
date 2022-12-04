<template>
    <v-layout @change="alert.message = ''" @click="alert.message = ''">
        <template v-if="user.authLoaded">
            <v-app-bar color="green" title="Jokes App" rounded></v-app-bar>
            <v-navigation-drawer expand-on-hover rail>
                <v-list v-if="user.id && user.email && user.username">
                    <v-list-item
                        prepend-icon="mdi-account"
                        :subtitle="user.email"
                        :title="user.username"></v-list-item>
                </v-list>
                <v-list nav>
                    <v-list-item
                        prepend-icon="mdi-home"
                        title="Home"
                        router
                        :to="'/dashboard'"></v-list-item>
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
import LoadingView from './views/LoadingView.vue';
import { defineComponent } from 'vue';
import { RouterView } from 'vue-router';
import {
    VAppBar,
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
        VLayout,
        VList,
        VListItem,
        VMain,
        VNavigationDrawer,
    },

    setup() {
        const user = useUserStore();
        const alert = useAlertStore();

        return { user, alert };
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
});
</script>

<style scoped></style>
