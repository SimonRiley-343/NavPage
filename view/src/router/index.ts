import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import NavPage from '@/components/NavPage.vue';
import Login from '@/components/Login.vue';
import Admin from '@/components/Admin.vue';

Vue.use(VueRouter);

let routes: Array<RouteConfig> = [
    {
        path: '/',
        name: 'NavPage',
        component: NavPage,
        meta: {
            title: 'WindSpirit IT Doc Lib'
        }
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
        meta: {
            title: 'WindSpirit IT Doc Lib - Login'
        }
    },
    {
        path: '/admin',
        name: 'Admin',
        component: Admin,
        meta: {
            title: 'Doc Lib Admin'
        }
    }
];

let router = new VueRouter({
    mode: 'history',
    routes
});

router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title;
    }
    next();
});

export default router;
