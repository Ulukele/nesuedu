<template>
    <div class="task-page">
        <Header @showUserPage="showUserPage" />
        <div class="tasks-list-filter">
            <va-select class="filter-selector" color="mainBlack" label="Сортировать по" v-model="value" :options="options">
                <template v-slot:option="{ option, index, selectOption }">
                    <div class="d-flex justify-space-between align-center pa-2" @click="selectOption(option)">
                        {{ option }}
                        <va-icon
                            name="sort"
                            size="small"
                            color="#808080"
                            :rotation="(index + 1) % 2 * 180"
                        />
                    </div>
                </template>
            </va-select>
            <TasksList :taskList="sortedTasks"/>
        </div>
        <NavBar></NavBar>
        <UserPage ref="userPage" :user="user" />
    </div>
</template>

<script>
import Header from '../index/Header.vue';
import TasksList from './TasksList.vue';
import NavBar from './TasksNavBar.vue';
import UserPage from '../user/UserPage.vue';
import { axios_utils } from '../../axios_utils'

export default {
    components: {
        Header,
        TasksList,
        NavBar,
        UserPage,
    },

    data: () => ({
        options: ['Цена по возрастанию', 'Цена по убыванию'],
        value: '',
        user: {},
        taskList: []
    }),

    mounted() {
        this.value = this.options[0];
        var id = sessionStorage.getItem('id');
        var sessionID = sessionStorage.getItem("jwt");
        axios_utils.getUser(id, sessionID).then(result => {
            this.user = result.data;
            this.fetchTasks();
        });
    },

    methods: {

        fetchTasks: function() {
            var sessionID = sessionStorage.getItem("jwt");
            axios_utils.getTasks(sessionID).then(result => {
                this.taskList = result.data;
            });
        },

        showUserPage: function(uid) {
            this.$refs.userPage.show = true;
            this.$refs.userPage.isLoaded = true;
            var id = sessionStorage.getItem('id');
            var sessionID = sessionStorage.getItem("jwt");
            axios_utils.getUser(id, sessionID).then(result => {
                this.user = result.data;
            });
        }
    },

    computed: {
        sortedTasks() {
            if (this.value == this.options[0]) {
                let tmp = []
                tmp = JSON.parse(JSON.stringify(this.taskList))
                return tmp.sort((a, b) => parseFloat(a.price) - parseFloat(b.price))
            }
            if (this.value == this.options[1]) {
                let tmp = []
                tmp = JSON.parse(JSON.stringify(this.taskList))
                return tmp.sort((a, b) => parseFloat(b.price) - parseFloat(a.price))
            }
            return this.taskList;
        }
    }
}
</script>