<template>
    <v-layout>
        <v-app-bar color="green" title="Jokes App" rounded></v-app-bar>
        <v-navigation-drawer expand-on-hover rail>
            <v-list v-if="user.id">
                <v-list-item>{{ user.username }}</v-list-item>
            </v-list>
            <v-list nav>
                <v-list-item
                    prepend-icon="mdi:mdi-home"
                    title="Home"
                    router
                    :to="'/dashboard'"></v-list-item>
                <v-list-item
                    prepend-icon="mdi:mdi-star"
                    title="About"
                    router
                    :to="'/about'"></v-list-item>
            </v-list>
        </v-navigation-drawer>
        <v-main>
            <RouterView />
        </v-main>
    </v-layout>
</template>

<script lang="ts">
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

export default defineComponent({
    components: {
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

        return { user };
    },

    methods: {},
});
</script>

<style scoped></style>
