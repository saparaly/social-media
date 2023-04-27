import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { name: 'Home', path: '/', component: () => import('../views/HomeView.vue') },
    { name: 'SignUp', path: '/signup', component: () => import('../views/SignupView.vue') },
    { name: 'SignIn', path: '/signin', component: () => import('../views/SigninView.vue') },
  ]
})

export default router
