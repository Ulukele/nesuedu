<template>
    <div class="task-page">
        <Header @showUserPage="showUserPage"/>
            <div class="tasks-list-filter">
                <TasksList v-bind:taskList="taskList" @showUserPage="showUserPage" @updateTasksList="updateTasksList"/>
            </div>
        <NavBar @showCreateTaskWind="showCreateTaskWind"/>
        <UserPage ref="userPage" :user="user"/>
        <CreateTaskPage ref="createTaskPage" @updateTasksList="updateTasksList"/>
    </div>
</template>

<script>
import { axios_utils } from '../../axios_utils'
import Header from '../index/Header.vue';
import TasksList from './UserTasksList.vue';
import NavBar from './UserNavBar.vue';
import UserPage from '../user/UserPage.vue';
import CreateTaskPage from './CreateTaskPage.vue';
export default {
    components: {
        Header,
        TasksList,
        NavBar,
        UserPage,
        CreateTaskPage
    },

    data: () => ({
        options: ['Цена по возрастанию', 'Цена по убыванию', 'Рейтинг заказчика по возрастанию', 'Рейтинг заказчика по убыванию'],
        value: '',
        user: {},
        taskList: []
    }),

    mounted() {
        this.value = this.options[0];
        this.updateTasksList();
    },

    methods: {
        showUserPage: function(uid) {
            var sessionID = sessionStorage.getItem("jwt");
            axios_utils.getUser(uid, sessionID).then(result => {
                this.user = result.data;
            });
            this.$refs.userPage.show = true;
            this.$refs.userPage.isLoaded = true;
        },

        showCreateTaskWind: function() {
            this.$refs.createTaskPage.show = true;
        },

        updateTasksList: function() {
            var userId = sessionStorage.getItem("id");
            var sessionID = sessionStorage.getItem("jwt");
            axios_utils.getTasks(sessionID).then(result => {
                var tasks = result.data;
                this.taskList = tasks.filter(task => {return task.customerId == userId});
            })
        }
    }
}
</script>