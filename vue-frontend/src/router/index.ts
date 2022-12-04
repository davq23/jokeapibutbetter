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
            path: '/dashboard',
            component: () => import('../views/JokesView.vue'),
        },
        {
            meta: {
                roles: ['ADMIN'],
                authRequired: true,
            },
            name: 'new-joke',
            path: '/jokes/new',
            component: () => import('../views/JokeFormView.vue'),
        },
        {
            meta: {
                roles: ['ADMIN'],
                authRequired: true,
            },
            name: 'edit-joke',
            path: '/jokes/:id/edit',
            component: () => import('../views/JokeFormView.vue'),
        },
        {
            meta: {
                roles: ['ADMIN'],
                authRequired: true,
            },
            name: 'my-jokes',
            path: '/jokes/mine',
            component: () => import('../views/MyJokes.vue'),
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
            component: () => import('../views/LoginView.vue'),
        },
    ],
});

router.beforeResolve(
    async (to: RouteLocationNormalized, from: RouteLocationNormalized) => {
        const user = useUserStore();

        user.setAuthLoaded(false);

        await user.whoIAm();

        user.setAuthLoaded(true);

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
                    name: 'dashboard',
                };
            }
        }
    },
);

export default router;
