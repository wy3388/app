import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router'

const routes: RouteRecordRaw[] = [
    {
        path: '/',
        name: 'Home',
        component: () => import('../views/Home.vue')
    },
    {
        path: '/setting',
        name: 'Setting',
        component: () => import('../views/Setting.vue')
    },
    {
        path: '/search',
        name: 'Search',
        component: () => import('../views/Search.vue')
    },
    {
        path: '/read',
        name: "Read",
        component: () => import('../views/Read.vue')
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
