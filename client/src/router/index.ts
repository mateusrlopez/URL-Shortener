import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '@/views/HomeView.vue';
import AccessesView from '@/views/AccessesView.vue';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            name: 'home',
            path: '/',
            component: HomeView,
        },
        {
            name: 'accesses',
            path: '/accesses',
            component: AccessesView,
        },
    ],
});

export default router;
