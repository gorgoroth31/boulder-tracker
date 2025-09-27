import {createRouter, createWebHashHistory, createWebHistory} from 'vue-router'
import {useAuth0} from '@auth0/auth0-vue';
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue';
import client, {getAccessToken, isUserAuthenticated} from "../plugins/auth";
import LogoutView from "../views/LogoutView.vue";
import {existsUserByClaims} from "../api/api";
import RegistrationView from "../views/RegistrationView.vue";
import AddSessionView from "../views/AddSessionView.vue";
import AboutView from "../views/AboutView.vue";

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
            path: '/sessions/add',
            name: "addsession",
            component: AddSessionView,
        },
        {
            path: "/about",
            name: "about",
            component: AboutView,
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
        },
        {
            path: '/register',
            name: 'register',
            component: RegistrationView,
            beforeEnter: async (to, from) => {
                const exists = await existsUserByClaims();
                
                // if user exists in db, he isn't allowed to naivgate there
                return !exists;
            }
        }
    ],
})

router.beforeEach(async (to, from) => {
    const isRouteLogin = to.name === "login";
    const isAuthenticated = await isUserAuthenticated();
    
    if (!isAuthenticated && !isRouteLogin) {
        return "login";
    }
    
    if (!isAuthenticated) {
        return;
    }
    
    const exists = await existsUserByClaims();

    if (to.name !== 'register' && !exists) {
        return "/register";
    }
})

export default router
