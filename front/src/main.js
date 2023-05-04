import { createApp } from 'vue'
import VueCookies from 'vue-cookies'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'

import './assets/main.css'

const app = createApp(App)

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { name: 'Home', path: '/', component: () => import('./views/HomeView.vue') },
        { name: 'Users', path: '/users', component: () => import('./views/UserView.vue')},
        { name: 'SignUp', path: '/signup', component: () => import('./views/SignupView.vue')},
        { name: 'SignIn', path: '/signin', component: () => import('./views/SigninView.vue') },
        { name: 'Logout', path: '/logout', component: () => import('./views/LogoutView.vue') },
        { name: 'CreatePost', path: '/create-post', component: () => import('./views/CreatePostView.vue') },
        { name: 'Post', path: '/post/:id', component: () => import('./views/PostView.vue') },
        { name: 'AppView', path: '/app', component: () => ('./views/AppView.vue') },
    ]
});

app.use(router)
app.use(VueCookies)
app.mount('#app')