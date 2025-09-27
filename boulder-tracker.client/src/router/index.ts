import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import {isUserAuthenticated} from "../plugins/auth";
import LogoutView from "../views/LogoutView.vue";
import {existsUserByClaims} from "../api/api";
import RegistrationView from "../views/RegistrationView.vue";
import AddSessionView from "../views/AddSessionView.vue";
import AboutView from "../views/AboutView.vue";
import mainPageUtils from "../utils/mainPageUtils";

const router = createRouter({
    history: createWebHashHistory(),
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
            component: LogoutView,
            beforeEnter: async (to, from) => {
                const isAuthenticated = await isUserAuthenticated();
                if (!isAuthenticated) {
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
    const isRouteLogin = to.name === "about";
    const isAuthenticated = await isUserAuthenticated();

    mainPageUtils.isLoggedIn.value = isAuthenticated;
    
    if (!isAuthenticated && !isRouteLogin) {
        return "about";
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
