import { useUserStore } from '@/stores/user';

import {
    createRouter,
    createWebHistory,
    type RouteLocationNormalized,
} from 'vue-router';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            name: 'jokes',
            path: '/jokes',
            meta: {
                title: 'All Jokes',
            },
            component: () => import('../views/jokes/JokesView.vue'),
        },
        {
            name: 'home',
            path: '/',
            meta: {
                title: 'Joke API Explorer',
            },
            component: () => import('../views/HomeView.vue'),
        },
        {
            name: 'joke-view',
            path: '/jokes/:id',
            meta: {
                title: 'Joke View',
            },
            component: () => import('../views/jokes/SingleJokeView.vue'),
        },
        {
            meta: {
                roles: ['ADMIN'],
                authRequired: true,
                title: 'New Joke',
            },
            name: 'new-joke',
            path: '/jokes/new',
            component: () => import('../views/jokes/JokeFormView.vue'),
        },
        {
            meta: {
                roles: ['ADMIN'],
                authRequired: true,
                title: 'Edit Joke',
            },
            name: 'edit-joke',
            path: '/jokes/:id/edit',
            component: () => import('../views/jokes/JokeFormView.vue'),
        },
        {
            meta: {
                roles: ['ADMIN'],
                authRequired: true,
                title: 'My Jokes',
            },
            name: 'my-jokes',
            path: '/jokes/mine',
            component: () => import('../views/jokes/MyJokes.vue'),
        },
        {
            meta: {
                authRequired: true,
                self: true,
                title: 'User Preferences',
            },
            name: 'user-preferences',
            path: '/users/preferences',
            component: () => import('../views/users/UserPreferencesView.vue'),
        },
        {
            name: 'about',
            path: '/about',
            // route level code-splitting
            // this generates a separate chunk (About.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () => import('../views/AboutView.vue'),
        },
        {
            path: '/login',
            name: 'login',
            meta: {
                title: 'Login',
            },
            component: () => import('../views/auth/LoginView.vue'),
        },
    ],
});

router.beforeResolve(async (to: RouteLocationNormalized) => {
    const user = useUserStore();

    user.setAuthLoaded(false);

    await user.whoIAm();

    user.setAuthLoaded(true);

    if (to.meta && to.meta.title) {
        document.title = to.meta.title as string;
    }

    if (to.meta.authRequired === true && user.id === null) {
        return {
            name: 'login',
        };
    }

    if (to.meta && to.meta.roles && to.meta.roles instanceof Array) {
        const roleIntersection = to.meta.roles.filter((role) =>
            user.roles.includes(role),
        );

        if (roleIntersection.length === 0) {
            return {
                name: 'home',
            };
        }
    }
});

export default router;
