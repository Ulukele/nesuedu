
import {createRouter, createWebHistory}  from 'vue-router'
// import Auth from "@/modules/auth/Auth.vue"
import TasksPage from "@/modules/tasks/TasksPage.vue"
import MyTasksPage from "@/modules/tasks/UserTasksPage.vue"
import Welcome from "@/modules/welcome/WelcomePage.vue"
import App from "@/modules/index/App.vue"

const routes = [
    {
        path: '/tasks',
        component: TasksPage,
        name: 'tasks'
    },

    {
        path: '/my-tasks',
        component: MyTasksPage,
        name: 'my tasks'
    },

    {
        path: '/',
        component: App,
        name: 'home'
    },

    {
        path: '/welcome-page',
        component: Welcome,
        name: 'welcome page'
    },
]

const history = createWebHistory();

const router = createRouter({
    history,
    routes,
})

export default router