import {createRouter, createWebHashHistory, createWebHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LogoutView from "../views/LogoutView.vue";
import {existsUserByClaims} from "../api/api";
import RegistrationView from "../views/RegistrationView.vue";
import AddSessionView from "../views/AddSessionView.vue";
import AboutView from "../views/AboutView.vue";
import mainPageUtils from "../utils/mainPageUtils";
import CallbackView from '../views/CallbackView.vue';
import { useAuthStore } from '@/stores/authStore';

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView,
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
            component: LogoutView
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
        },
        {
            path: '/callback',
            name: 'callback',
            component: CallbackView,
        },
    ],
})

router.beforeEach(async (to, from) => {
    const isRouteLogin = to.name === "about";

    const { isAuthenticated } = useAuthStore()
    
    mainPageUtils.isLoggedIn.value = isAuthenticated();
    
    if (!isAuthenticated() && !isRouteLogin) {
        return "/about";
    }
    
    if (!isAuthenticated()) {
        return;
    }

    const exists = await existsUserByClaims();

    if (to.name !== 'register' && !exists) {
        return "/register";
    }
})

export default router
