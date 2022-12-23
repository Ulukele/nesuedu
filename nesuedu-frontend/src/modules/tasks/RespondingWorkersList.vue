<template>
    <div style="width: 100%; height: fit-content;">
        <ul style="background-color: #F6F6F6; border-radius: 0 0 20px 20px;">
            <worker v-for="worker of workersList" v-bind:worker="worker" @showUserPage="showUserPage"/>
        </ul>
    </div>
</template>

<script>
import Worker from './RespondingWorker.vue';
export default {
    components: {
        Worker
    },

    props: ['workersList'],

    methods: {
        showUserPage: function(uid) {
            this.$emit('showUserPage', uid);
        }
    }
}
</script>