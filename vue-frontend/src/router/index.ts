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
                roles: 'admin',
            },
            name: 'new-joke',
            path: '/jokes/new',
            component: () => import('../views/NewJokeView.vue'),
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

        await user.whoIAm();

        if (to.meta.authRequired === true && user.id === null) {
            return {
                name: 'login',
            };
        }
    },
);

export default router;
