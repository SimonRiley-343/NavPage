import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import NavPage from '@/components/NavPage.vue'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
    {
        path: '/',
        name: 'NavPage',
        component: NavPage,
        meta: {
            title: 'WindSpirit IT Doc Lib'
        }
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next()
})

export default router
