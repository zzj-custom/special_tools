import { createRouter, createWebHistory } from 'vue-router';
import MainRoutes from './MainRoutes';

export const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/:pathMatch(.*)*',
            component: () => import('@/views/pages/Error404.vue')
        },
        MainRoutes,
    ]
});
