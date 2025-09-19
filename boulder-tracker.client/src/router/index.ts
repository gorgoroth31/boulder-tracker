import {createRouter, createWebHashHistory, createWebHistory} from 'vue-router'
import {useAuth0} from '@auth0/auth0-vue';
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue';
import client, {getAccessToken} from "../plugins/auth";
import LogoutView from "../views/LogoutView.vue";

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView,
        },
        {
            path: '/login',
            name: "login",
            component: LoginView,
            beforeEnter: (to, from) => {
                if (client.isAuthenticated.value) {
                    return false
                }
            }
        },
        {
            path: '/logout',
            name: "logout",
            component: LogoutView,
            beforeEnter: (to, from) => {
                if (!client.isAuthenticated.value) {
                    return false
                }
            }
        }
    ],
})

router.beforeEach(async (to, from) => {
    await client.checkSession()
    if (to.name !== 'login' && !client.isAuthenticated.value) {
        return "/login";
    }
})

export default router
