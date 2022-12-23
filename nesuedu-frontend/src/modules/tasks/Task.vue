<template>
    <div class="task-item">
        <span class="task-name">{{ task.title }}</span>
        <div class="space-betw" style="width: 30%;">
            <span style="justify-content: flex-start; color:#87C644;">{{ task.audience }}</span>
            <span style="color: #D9D9D9;">{{ taskDeadline }}</span>
        </div>

        <div v-if="!isOpen" class="space-betw" style="margin-bottom: 1.3%; max-width: 95%;">
            <span style="justify-content: flex-start; color: #D9D9D9; cursor: pointer;" @click="isOpen = true">Подробнее...</span>
            <span style="color: #87C644;">{{ task.price }} руб.</span>
        </div>
        <div v-else style="margin-top: 1.3%;">
            <span style="margin-left: 1.3%;">{{ task.description }}</span>
            <div class="space-betw" style="margin-bottom: 1.3%; margin-top: 1.3%; max-width: 97%;">
                <div style="justify-content: flex-start; color: #D9D9D9; cursor: pointer; align-self: flex-end;" @click="isOpen = false">Свернуть</div>
                <div style="justify-content: flex-end;">
                    <span style="color: #87C644;">{{ task.price }} руб.&nbsp;&nbsp;&nbsp;</span>
                    <button v-if="!subscribed" class="respond-btn" @click="subscribe">Откликнуться</button>
                    <button v-else class="respond-btn revert" @click="unsubscribe">Отменить отклик</button>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import { axios_utils } from '../../axios_utils'
import { ref, toHandlerKey } from 'vue'

export default {
    props: ['task'],


    data: () => ({
        subscribed: false,
        isOpen: false,
    }),

    mounted() {
        this.checkSubscribed();
    },

    watch: { 
        task: function(newVal, oldVal) {
            this.checkSubscribed()
        }
    },

    methods: {
        subscribe() {
            var sessionID = sessionStorage.getItem("jwt");
            axios_utils.createSub(this.task.id, sessionID).then(this.checkSubscribed);
        },

        unsubscribe() {
            var sessionID = sessionStorage.getItem("jwt");
            axios_utils.deleteSub(this.task.id, sessionID).then(this.checkSubscribed);
        },

        checkSubscribed() {
            var userId = sessionStorage.getItem("id");
            var sessionID = sessionStorage.getItem("jwt");
            this.subscribed = false;
            axios_utils.getSubs(this.task.id, sessionID).then(result => {
                result.data.forEach(element => {
                    if (element.workerId == userId) {
                        this.subscribed = true;
                        // this.$forceUpdate();
                    }
                });
            });
        }
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
        },
    },
}
</script>