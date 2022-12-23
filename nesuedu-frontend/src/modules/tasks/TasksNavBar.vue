<template>
    <nav>
        <div class="nav-bar">
            <router-link to="/tasks" class="nav-bar-btn" style="color: #434343;">
                Заказы
            </router-link>

            <hr width="100%" color="#D9D9D9">

            <router-link to="/my-tasks" class="nav-bar-btn">
                Мои заказы
            </router-link>
        </div>
    </nav>
</template>

<script>
export default { }
</script>
