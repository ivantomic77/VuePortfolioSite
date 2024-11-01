import { createWebHistory, createRouter } from "vue-router";
import MainPage from "../pages/MainPage.vue";
import Post from "../pages/Post.vue";

const routes = [
    { path: '/', component: MainPage},
    { path: '/posts/:slug', component: Post}
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router;