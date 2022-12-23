<template>
    <div>
        <div class="task-item" style="border-radius: 20px 20px 0 0;">
            <div class="space-betw" style="max-width: 97%;">
                <span class="task-name" style="justify-content: flex-start; margin-left: 0;">{{ task.title }}</span>
                <img src="/images/delete-btn.svg" style="cursor: pointer;" @click="deleteTask">
            </div>
            <div class="space-betw" style="width: 30%;">
                <span style="justify-content: flex-start; color:#87C644;">{{ task.audience }}</span>
                <span style="color: #D9D9D9;">{{ taskDeadline }}</span>
            </div>
            <div v-if="!isOpen" class="space-betw" style="margin-bottom: 1.3%; max-width: 97%;">
                <span style="justify-content: flex-start; color: #D9D9D9; cursor: pointer;" @click="isOpen = true">Подробнее...</span>
                <span style="color: #87C644;">{{ task.price }} руб.</span>
            </div>
            <div v-else style="margin-top: 1.3%;">
                <span style="margin-left: 1.3%;">{{ task.description }}</span>
                <div class="space-betw" style="margin-bottom: 1.3%; margin-top: 1.3%; max-width: 97%;">
                    <div style="justify-content: flex-start; color: #D9D9D9; cursor: pointer; align-self: flex-end;" @click="isOpen = false">Свернуть</div>
                    <div style="justify-content: flex-end;">
                        <span style="color: #87C644;">{{ task.price }} руб.</span>
                    </div>
                </div>
            </div>
        </div>
        <RespondingWorkersList v-bind:workersList="workersList" @showUserPage="showUserPage"></RespondingWorkersList>
    </div>
</template>

<script>
import { axios_utils } from '../../axios_utils'
import RespondingWorkersList from './RespondingWorkersList.vue';
export default {
    components: {
        RespondingWorkersList
    },

    props: ['task'],

    data: () => ({
        isOpen: false,
        workersList: [],
    }),

    methods: {
        showUserPage: function(uid) {
            this.$emit('showUserPage', uid);
        }
    },

    mounted() {
        var sessionID = sessionStorage.getItem("jwt");
        axios_utils.getSubs(this.task.id, sessionID).then(result => {
            const promises = result.data.map(async (sub) => {
                return axios_utils.getUser(sub.workerId, sessionID);
            })

            Promise.allSettled(promises).then((results) => {
                this.workersList = results.map(item => {return item.value.data});
            })
        })
    },

    methods: {
        deleteTask: function() {
            var sessionID = sessionStorage.getItem("jwt");
            axios_utils.deleteTask(this.task.id, sessionID).then(result => {
                if (result.status === 200) {
                    this.$emit('doUpdate');
                }
            });
        },
    },

    computed: {
        taskDeadline() {

            return (new Date(this.task.deadline)).toLocaleDateString('ru', {
                weekday:"long",
                year:"numeric",
                month:"short",
                day:"numeric",
                hour:"numeric",
                minute: "numeric"
            });
        }
    },
}
</script>