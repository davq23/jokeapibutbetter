import { useUserStore } from '@/stores/user';
import JokesView from '@/views/JokesView.vue';
import {
    createRouter,
    createWebHistory,
    type RouteLocationNormalized,
} from 'vue-router';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/dashboard',
            name: 'jokes',
            component: JokesView,
        },
        {
            path: '/about',
            name: 'about',
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

        user.whoIAm();
    },
);

export default router;
