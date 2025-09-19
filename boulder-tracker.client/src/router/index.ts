import {createRouter, createWebHashHistory, createWebHistory} from 'vue-router'
import { useAuth0 } from '@auth0/auth0-vue';
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue';

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
    }
  ],
})

router.beforeEach((to, from) => {
  const { isAuthenticated } = useAuth0();

  if (!isAuthenticated.value && to.name !== 'login') {
    return { name: 'login' }
  }
})

export default router
