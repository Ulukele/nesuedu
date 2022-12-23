<template>
    <div style="width: 100%;">
        <ul style="gap: 20px;">
            <Task v-for="task of taskList" v-bind:task="task"/>
        </ul>
    </div>
</template>

<script>
import Task from './Task.vue';
export default {
    components: {
        Task
    },

    props: ['taskList'],
}
</script>
