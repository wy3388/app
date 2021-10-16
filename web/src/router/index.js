import {createRouter, createWebHashHistory} from 'vue-router'

const routes = [
    {
        path: '/',
        name: 'Home',
        component: () => import('../views/Home.vue')
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
    },
    {
        path: '/info',
        name: 'Info',
        component: () => import('../views/Info.vue')
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
