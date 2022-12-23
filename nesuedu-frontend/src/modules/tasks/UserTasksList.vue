<template>
    <div style="width: 100%;">
        <ul style="gap: 20px;">
            <Task v-for="task of taskList" v-bind:task="task" @showUserPage="showUserPage" @doUpdate="doUpdate"/>
        </ul>
    </div>
</template>

<script>
import Task from './UserTask.vue';
export default {
    components: {
        Task
    },

    props: ['taskList'],

    methods: {
        showUserPage: function(uid) {
            this.$emit('showUserPage', uid);
        },

        doUpdate: function() {
            this.$emit('updateTasksList');
        }
    }
}
</script>