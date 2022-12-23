<template>
    <nav>
        <div class="nav-bar">
            <router-link to="/my-tasks" class="nav-bar-btn" style="color: #434343;">
                Мои заказы
            </router-link>

            <hr width="100%" color="#D9D9D9">

            <router-link to="/tasks" class="nav-bar-btn">
                На главную
            </router-link>

            <hr width="100%" color="#D9D9D9">

            <div class="nav-bar-btn" @click="showCreateTaskWind">
                Добавить
            </div>
        </div>
    </nav>
</template>

<script>
export default {
    methods: {
        showCreateTaskWind: function() {
            this.$emit('showCreateTaskWind')
        }
    }
 }
</script>