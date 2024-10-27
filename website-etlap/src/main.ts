import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import { createRouter, createWebHistory } from 'vue-router';
import Welcome from './components/Welcome.vue';
import FoodList from './components/FoodList.vue';
import Order from './components/Order.vue';

// https://router.vuejs.org/guide/
const routes = [
    // URL, use only english characters
    // https://github.com/vuejs/vue-router/issues/3110
    { path: '/', component: Welcome },
    { path: '/etelek', component: FoodList },
    { path: '/rendeles', component: Order },
]

const router = createRouter({
    // https://router.vuejs.org/guide/essentials/history-mode.html
    history: createWebHistory(),
    routes,
})
createApp(App)
    .use(router)
    .mount('#app')
